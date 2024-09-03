package main

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

var (
	fileName  string
	cacheFile string
)

type MockAVS struct {
	Repo       string `yml:"repo"`
	Version    string `yml:"version"`
	CommitHash string `yml:"commitHash"`
}

type Repos struct {
	Timestamp time.Time `yml:"timestamp"`
	Repos     []MockAVS `yml:"repos"`
}

func init() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	// Construct the file paths
	fileName = fmt.Sprintf("%s/package_handler/tmp/mock-avs-versions.yml", wd)
	cacheFile = fmt.Sprintf("%s/package_handler/tmp/mock-avs-versions-cache.json", wd)
}

func main() {
	repos := []string{
		"https://github.com/NethermindEth/mock-avs",
		"https://github.com/NethermindEth/mock-avs-pkg",
	}

	data := make([]MockAVS, 0)
	if shouldUpdateFile(fileName) {
		for _, repo := range repos {
			tag, commitHash, err := latestGitTagAndCommitHash(repo)
			if err != nil {
				fmt.Println("Error fetching latest git tag:", err)
				return
			}

			data = append(data, MockAVS{
				Repo:       repo,
				Version:    tag,
				CommitHash: commitHash,
			})
		}

		err := writeYMLFile(fileName, Repos{Timestamp: time.Now(), Repos: data})
		if err != nil {
			fmt.Println("Error writing to yml file:", err)
			return
		}
	}
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
