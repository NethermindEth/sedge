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
package package_handler

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/monitoring/env"
	"github.com/NethermindEth/sedge/internal/monitoring/profile"
	"github.com/compose-spec/compose-go/cli"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/afero"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v3"
)

const (
	pkgDirName             = "pkg"
	checksumFileName       = "checksum.txt"
	manifestFileName       = "manifest.yml"
	profileFileName        = "profile.yml"
	manifestSchemaFileName = "schema/manifest_schema.yml"
	profileSchemaFileName  = "schema/profile_schema.yml"
)

// PackageHandler is used to interact with an AVS node software package at the given
// path.
type PackageHandler struct {
	path string
	afs  afero.Fs
}

// NewPackageHandler creates a new PackageHandler instance for the given package path.
func NewPackageHandler(path string) *PackageHandler {
	return &PackageHandler{path: path, afs: afero.NewOsFs()}
}

// NewPackageHandlerOptions is used to provide options to the NewPackageHandlerFromURL
type NewPackageHandlerOptions struct {
	// Path is the path where the package will be cloned
	Path string
	// URL is the URL of the git repository
	URL string
	// GitAuth is used to provide authentication to a private git repository
	GitAuth *GitAuth
}

// GitAuth is used to provide authentication to a private git repository. Two types of
// authentication are supported (tested with GitHub):
//
//  1. Username and password: set both Username and Password fields, and leave the Pat
//     field empty.
//
//  2. Personal access token: set the Username and Pat fields, and leave the Password
//     field empty.
//
// Pat field has more priority than Password field, meaning that if both are set, the
// Pat field will be used.
// TODO: support key authentication
type GitAuth struct {
	Username string
	Password string
	Pat      string
}

func (g *NewPackageHandlerOptions) getAuth() *http.BasicAuth {
	if g.GitAuth == nil {
		return nil
	}
	if g.GitAuth.Pat != "" {
		return &http.BasicAuth{
			Username: g.GitAuth.Username,
			Password: g.GitAuth.Pat,
		}
	}
	return &http.BasicAuth{
		Username: g.GitAuth.Username,
		Password: g.GitAuth.Password,
	}
}

// NewPackageHandlerFromURL clones the package from the given URL and returns. The GitAuth
// field could be used to provide authentication to a private git repository.
func NewPackageHandlerFromURL(opts NewPackageHandlerOptions) (*PackageHandler, error) {
	_, err := git.PlainClone(opts.Path, false, &git.CloneOptions{
		URL:  opts.URL,
		Auth: opts.getAuth(),
	})
	if err != nil {
		if errors.Is(err, transport.ErrAuthenticationRequired) {
			return nil, RepositoryNotFoundOrPrivateError{
				URL: opts.URL,
			}
		}
		if errors.Is(err, transport.ErrRepositoryNotFound) {
			return nil, RepositoryNotFoundError{
				URL: opts.URL,
			}
		}
		return nil, err
	}
	return NewPackageHandler(opts.Path), nil
}

// Check validates a package. It returns an error if the package is invalid.
// It checks the existence of some required files and directories and computes the
// checksums comparing them with the ones listed in the checksum.txt file.
func (p *PackageHandler) Check() error {
	if err := checkPackageDirExist(p.path, pkgDirName, p.afs); err != nil {
		return err
	}
	err := checkPackageFileExist(p.path, checksumFileName, p.afs)
	if err != nil {
		var fileNotFoundErr PackageFileNotFoundError
		if errors.As(err, &fileNotFoundErr) {
			return nil
		}
		return err
	}
	return p.checkSum()
}

// Versions returns the descending sorted list of available versions for the package.
// A version is a git tag that matches the regex `^v\d+\.\d+\.\d+$`.
func (p *PackageHandler) Versions() ([]string, error) {
	pkgRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return nil, err
	}
	tagIter, err := pkgRepo.Tags()
	if err != nil {
		return nil, err
	}
	var versions []string
	tagIter.ForEach(func(ref *plumbing.Reference) error {
		tag := ref.Name().Short()
		if semver.IsValid(tag) {
			versions = append(versions, tag)
		}
		return nil
	})
	if len(versions) == 0 {
		return nil, ErrNoVersionsFound
	}
	semver.Sort(versions)
	return versions, nil
}

// HasVersion returns an error if the given version is not available for the package.
func (p *PackageHandler) HasVersion(version string) error {
	versions, err := p.Versions()
	if err != nil {
		return err
	}
	for _, v := range versions {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("%w: %s", ErrVersionNotFound, version)
}

// CheckoutCommit checkout the cloned repository to the given commit hash. If
// the commit hash is not found, it returns an error.
func (p *PackageHandler) CheckoutCommit(commitHash string) error {
	pkgRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return err
	}
	wt, err := pkgRepo.Worktree()
	if err != nil {
		return err
	}
	return wt.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commitHash),
	})
}

// LatestVersion returns the latest version of the package.
func (p *PackageHandler) LatestVersion() (string, error) {
	versions, err := p.Versions()
	if err != nil {
		return "", err
	}
	return versions[len(versions)-1], nil
}

// CommitPrecedence returns true if the new commit hash is a descendant of the
// old commit hash. It returns an error if the commit hashes are not found.
func (p *PackageHandler) CommitPrecedence(oldCommitHash, newCommitHash string) (bool, error) {
	err := p.CheckoutCommit(newCommitHash)
	if err != nil {
		return false, err
	}
	gitRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return false, err
	}
	commitIter, err := gitRepo.Log(&git.LogOptions{
		From: plumbing.NewHash(newCommitHash),
	})
	if err != nil {
		return false, err
	}
	// Skip first commit, because it is the new commit hash itself
	_, err = commitIter.Next()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}
		return false, err
	}
	for {
		c, err := commitIter.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return false, nil
			}
			return false, err
		}
		if c.Hash.String() == oldCommitHash {
			return true, nil
		}
	}
}

// CheckoutVersion checks out the cloned repository to the given version (tag).
func (p *PackageHandler) CheckoutVersion(version string) error {
	if !semver.IsValid(version) {
		return ErrInvalidVersion
	}
	gitRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return err
	}
	tagIter, err := gitRepo.Tags()
	if err != nil {
		return err
	}
	defer tagIter.Close()
	for {
		tag, err := tagIter.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return ErrNoVersionsFound
			}
			return err
		}
		if tag.Name().Short() == version {
			worktree, err := gitRepo.Worktree()
			if err != nil {
				return fmt.Errorf("error getting worktree: %w", err)
			}
			err = worktree.Checkout(&git.CheckoutOptions{
				Branch: tag.Name(),
			})
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

// CurrentVersion returns the current version of the package, which is tha latest
// tag with version format that points to the current HEAD.
func (p *PackageHandler) CurrentVersion() (string, error) {
	gitRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return "", err
	}
	head, err := gitRepo.Head()
	if err != nil {
		return "", err
	}
	tagIter, err := gitRepo.TagObjects()
	if err != nil {
		return "", err
	}
	var headVersions []string
	for {
		tag, err := tagIter.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		if semver.IsValid(tag.Name) && head.Hash() == tag.Target {
			headVersions = append(headVersions, tag.Name)
		}
	}
	if len(headVersions) == 0 {
		return "", ErrNoVersionsFound
	}
	semver.Sort(headVersions)
	return headVersions[len(headVersions)-1], nil
}

// CurrentHash returns the git hash of the package HEAD
func (p *PackageHandler) CurrentCommitHash() (string, error) {
	gitRepo, err := git.PlainOpen(p.path)
	if err != nil {
		return "", err
	}
	head, err := gitRepo.Head()
	if err != nil {
		return "", err
	}
	return head.Hash().String(), nil
}

// Profiles returns the list of profiles defined in the package for the current version.
func (p *PackageHandler) Profiles() ([]profile.Profile, error) {
	names, err := p.profilesNames()
	if err != nil {
		return nil, err
	}

	profiles := make([]profile.Profile, 0)
	for _, profileName := range names {
		profile, err := p.parseProfile(profileName)
		if err != nil {
			return nil, err
		}
		profile.Name = profileName

		if err := profile.Validate(); err != nil {
			return nil, err
		}

		profiles = append(profiles, *profile)
	}

	return profiles, nil
}

// Profiles returns the list of profiles defined in the package for the current version.
func (p *PackageHandler) Profile(name string) (*profile.Profile, error) {
	names, err := p.profilesNames()
	if err != nil {
		return nil, err
	}

	for _, profileName := range names {
		if profileName == name {
			profile, err := p.parseProfile(profileName)
			if err != nil {
				return nil, err
			}
			profile.Name = profileName
			err = profile.Validate()
			if err != nil {
				return nil, err
			}
			return profile, nil
		}
	}

	return nil, fmt.Errorf("%w: %s", ErrProfileNotFound, name)
}

// CheckComposeProject checks if the compose project for the given profile is valid.
func (p *PackageHandler) CheckComposeProject(profileName string, env map[string]string) error {
	composeFile := filepath.Join(p.path, pkgDirName, profileName, "docker-compose.yml")
	composeExists, err := afero.Exists(p.afs, composeFile)
	if err != nil {
		return err
	}
	if !composeExists {
		return fmt.Errorf("%w: profile %s", ErrProfileComposeFileNotFound, profileName)
	}

	projectOptions, err := cli.NewProjectOptions([]string{composeFile})
	if err != nil {
		return err
	}
	maps.Copy(projectOptions.Environment, env)

	project, err := cli.ProjectFromOptions(projectOptions)
	if err != nil {
		return err
	}

	services := project.AllServices()
	for _, service := range services {
		if service.Build != nil {
			return fmt.Errorf("%w: profile %s, service %s", ErrBuildContextNotAllowed, profileName, service.Name)
		}
	}

	return nil
}

// DotEnv returns the .env file for the given profile.
// Assumes the package has been checked and is valid.
func (p *PackageHandler) DotEnv(profile string) (map[string]string, error) {
	envPath := filepath.Join(p.path, pkgDirName, profile, ".env")
	e, err := env.LoadEnv(p.afs, envPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ReadingDotEnvError{
			pkgPath:     p.path,
			profileName: profile,
		}, err)
	}
	return e, nil
}

// ProfileFS returns the filesystem path for the given profile.
func (p *PackageHandler) ProfilePath(profileName string) string {
	return filepath.Join(p.path, pkgDirName, profileName)
}

// HasPlugin returns true if the package has a plugin.
func (p *PackageHandler) HasPlugin() (bool, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return false, err
	}

	return manifest.Plugin != nil, nil
}

// Plugin returns the plugin for the package.
func (p *PackageHandler) Plugin() (*Plugin, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return nil, err
	}

	if manifest.Plugin == nil {
		return nil, ErrNoPlugin
	}

	return manifest.Plugin, nil
}

func (p *PackageHandler) parseManifest() (*Manifest, error) {
	manifestPath := filepath.Join(p.path, pkgDirName, manifestFileName)
	// Validate YAML Schema
	// TODO: Fix the relative path
	// err := validateYAMLSchema(manifestSchemaFileName, manifestPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("yaml schema validation error: %v", err)
	// }
	// Read the manifest file
	data, err := afero.ReadFile(p.afs, manifestPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ReadingManifestError{
			pkgPath: p.path,
		}, err)
	}

	var manifest Manifest
	err = yaml.Unmarshal(data, &manifest)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ParsingManifestError{
			pkgPath: p.path,
		}, err)
	}

	return &manifest, nil
}

// HardwareRequirements returns the hardware requirements for the specified profile.
// It takes a profile name as input and returns the hardware requirements as a HardwareRequirements struct.
// If the profile does not have any hardware requirements overrides, it returns the hardware requirements
// from the package manifest. If there is an error parsing the manifest or the profile, an error is returned.
func (p *PackageHandler) HardwareRequirements(profileName string) (hardwareRequirements, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return hardwareRequirements{}, err
	}

	hr := manifest.HardwareRequirements
	profile, err := p.parseProfile(profileName)
	if err != nil {
		return hardwareRequirements{}, err
	}
	if profile.HardwareRequirementsOverrides != nil {
		return hardwareRequirements{
			MinCPUCores:                 profile.HardwareRequirementsOverrides.MinCPUCores,
			MinRAM:                      profile.HardwareRequirementsOverrides.MinRAM,
			MinFreeSpace:                profile.HardwareRequirementsOverrides.MinFreeSpace,
			StopIfRequirementsAreNotMet: profile.HardwareRequirementsOverrides.StopIfRequirementsAreNotMet,
		}, nil
	}

	return hr, nil
}

// Path returns the path of the package.
func (p *PackageHandler) Path() string {
	return p.path
}

// ManifestFilePath returns the path of the manifest file.
func (p *PackageHandler) ManifestFilePath() string {
	return filepath.Join(p.path, pkgDirName, manifestFileName)
}

func (p *PackageHandler) profilesNames() ([]string, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return nil, err
	}

	if err := manifest.validate(); err != nil {
		return nil, err
	}

	return manifest.Profiles, nil
}

func (p *PackageHandler) parseProfile(profileName string) (*profile.Profile, error) {
	profilePath := filepath.Join(p.path, pkgDirName, profileName, profileFileName)
	// Validate YAML Schemas
	// TODO: Fix the relative path
	// err := validateYAMLSchema(manifestSchemaFileName, profilePath)
	// if err != nil {
	// 	return nil, fmt.Errorf("yaml schema validation error: %v", err)
	// }
	// Read the profile file
	data, err := afero.ReadFile(p.afs, profilePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ReadingProfileError{
			profileName: profileName,
		}, err)
	}

	var profile profile.Profile
	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ParsingProfileError{
			profileName: profileName,
		}, err)
	}

	return &profile, nil
}

func (p *PackageHandler) checkSum() error {
	currentChecksums, err := parseChecksumFile(filepath.Join(p.path, checksumFileName), p.afs)
	if err != nil {
		return err
	}
	computedChecksums, err := packageHashes(p.path, p.afs)
	if err != nil {
		return err
	}
	if len(currentChecksums) != len(computedChecksums) {
		return fmt.Errorf("%w: expected %d files, got %d", ErrInvalidChecksum, len(currentChecksums), len(computedChecksums))
	}
	for file, hash := range currentChecksums {
		if computedChecksums[file] != hash {
			return fmt.Errorf("%w: checksum mismatch for file %s, expected %s, got %s", ErrInvalidChecksum, file, hash, computedChecksums[file])
		}
	}
	return nil
}

func (p *PackageHandler) SpecVersion() (string, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return "", err
	}

	if err := manifest.validate(); err != nil {
		return "", err
	}

	return manifest.Version, nil
}

func (p *PackageHandler) Name() (string, error) {
	manifest, err := p.parseManifest()
	if err != nil {
		return "", err
	}

	if err := manifest.validate(); err != nil {
		return "", err
	}

	return manifest.Name, nil
}
