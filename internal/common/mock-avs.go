/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	dataFile                = "/tmp/mock-avs-versions.yml"
	cacheFile               = "/tmp/mock-avs-versions-cache.json"
	mockAVSRepo             = "https://github.com/NethermindEth/mock-avs"
	mockAVSPkgRepo          = "https://github.com/NethermindEth/mock-avs-pkg"
	optionReturnerImageName = "mock-avs-option-returner"
	healthCheckerImageName  = "mock-avs-health-checker"
	pluginImageName         = "mock-avs-plugin"
)

// Global variables to store the latest versions of the mock-avs and
// mock-avs-pkg repositories, and the docker image names for the
// option-returner, health-checker profiles, and the mock-avs plugin.
var (
	MockAvsSrc          MockAVS
	MockAvsPkg          MockAVS
	OptionReturnerImage MockAVSImage
	HealthCheckerImage  MockAVSImage
	PluginImage         MockAVSImage
)

type Repos struct {
	Timestamp time.Time     `yml:"timestamp"`
	Repos     []MockAVSData `yml:"repos"`
}

type MockAVSData struct {
	Repo       string `yml:"repo"`
	Version    string `yml:"version"`
	CommitHash string `yml:"commitHash"`
}

type MockAVS struct {
	repo       string
	version    string
	commitHash string
}

func NewMockAVS(repo string, version string, commitHash string) *MockAVS {
	return &MockAVS{
		repo:       repo,
		version:    version,
		commitHash: commitHash,
	}
}

func (m *MockAVS) Repo() string {
	return m.repo
}

func (m *MockAVS) Version() string {
	return m.version
}

func (m *MockAVS) CommitHash() string {
	return m.commitHash
}

type MockAVSImage struct {
	image string
	tag   string
}

func NewMockAVSImage(image, tag string) *MockAVSImage {
	return &MockAVSImage{
		image: image,
		tag:   tag,
	}
}

func (m *MockAVSImage) Image() string {
	return m.image
}

func (m *MockAVSImage) Tag() string {
	return m.tag
}

func (m *MockAVSImage) FullImage() string {
	return fmt.Sprintf("%s:%s", m.image, m.tag)
}

// SetMockAVSs set up the MockAVS and MockAVSPkg data structures with
// the latest versions of the mock-avs and mock-avs-pkg repositories.
// It reads the data from the mock-avs-versions.yml file, which is
// generated if it doesn't exist or if it is older than one hour.
// It also sets up the OptionReturnerImage, HealthCheckerImage and
// PluginImage data structures using as tag the latest version of the
// mock-avs repository.
func SetMockAVSs() error {
	if err := checkCache(); err != nil {
		return fmt.Errorf("error checking cache for mock-avs data: %w", err)
	}

	// Read the data from the mock-avs-versions.yml file
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return fmt.Errorf("error reading mock-avs data file: %w", err)
	}

	// Unmarshal the data into the Repos struct
	var reposData Repos
	err = yaml.Unmarshal(file, &reposData)
	if err != nil {
		return fmt.Errorf("error unmarshalling mock-avs data: %w", err)
	}

	// Set the MockAvsSrc and MockAvsPkg global variables
	for _, repo := range reposData.Repos {
		if repo.Repo == mockAVSRepo {
			MockAvsSrc = *NewMockAVS(repo.Repo, repo.Version, repo.CommitHash)
		} else {
			MockAvsPkg = *NewMockAVS(repo.Repo, repo.Version, repo.CommitHash)
		}
	}

	OptionReturnerImage = *NewMockAVSImage(optionReturnerImageName, MockAvsSrc.Version())
	HealthCheckerImage = *NewMockAVSImage(healthCheckerImageName, MockAvsSrc.Version())
	PluginImage = *NewMockAVSImage(pluginImageName, MockAvsSrc.Version())

	return nil
}

func checkCache() error {
	repos := []string{
		mockAVSRepo,
		mockAVSPkgRepo,
	}

	data := make([]MockAVSData, 0)
	if shouldUpdateFile(dataFile) {
		for _, repo := range repos {
			tag, commitHash, err := latestGitTagAndCommitHash(repo)
			if err != nil {
				return fmt.Errorf("error fetching latest git tag and commit: %w", err)
			}

			data = append(data, MockAVSData{
				Repo:       repo,
				Version:    tag,
				CommitHash: commitHash,
			})
		}

		err := writeYMLFile(dataFile, Repos{Timestamp: time.Now(), Repos: data})
		if err != nil {
			return fmt.Errorf("error writing to yml file: %w", err)
		}
	}

	return nil
}

func shouldUpdateFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return true
	}

	cache, err := readYMLFile(filePath)
	if err != nil {
		return true
	}

	return time.Since(cache.Timestamp).Hours() < 1
}

func readYMLFile(filePath string) (Repos, error) {
	var cache Repos

	file, err := os.ReadFile(cacheFile)
	if err != nil {
		return cache, err
	}

	err = yaml.Unmarshal(file, &cache)
	return cache, err
}

func writeYMLFile(filePath string, data Repos) error {
	ymlContent, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, ymlContent, 0o644)
	if err != nil {
		return err
	}

	return nil
}

type Tag struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
	} `json:"commit"`
}

func latestGitTagAndCommitHash(repoURL string) (string, string, error) {
	// Extract the repo owner and name from the URL
	parts := strings.Split(strings.TrimRight(strings.TrimPrefix(repoURL, "https://github.com/"), "/"), "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid repo URL")
	}
	owner, repo := parts[0], parts[1]

	// GitHub API endpoint to get tags
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo)

	var tag, commitHash string

	operation := func() error {
		resp, err := http.Get(apiURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to fetch data from GitHub API, status code: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var tags []Tag
		if err := json.Unmarshal(body, &tags); err != nil {
			return err
		}

		if len(tags) == 0 {
			return fmt.Errorf("no tags found in the repository")
		}

		log.Debugf("Latest Tag: %s\nCommit Hash: %s\n", tags[0].Name, tags[0].Commit.Sha)
		tag = tags[0].Name
		commitHash = tags[0].Commit.Sha
		return nil
	}

	// Using exponential backoff for retries
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 5 * time.Second
	if err := backoff.Retry(operation, bo); err != nil {
		return "", "", err
	}

	return tag, commitHash, nil
}
