package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/monitoring/env"
	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	"github.com/NethermindEth/eigenlayer/internal/profile"
	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// InstanceId returns the instance ID for the given name and tag
func InstanceId(name, tag string) string {
	return fmt.Sprintf("%s-%s", name, tag)
}

// Instance represents the data stored about a node software instance
type Instance struct {
	Name              string            `json:"name"`
	URL               string            `json:"url"`
	Version           string            `json:"version"`
	SpecVersion       string            `json:"spec_version"`
	Commit            string            `json:"commit,omitempty"`
	Profile           string            `json:"profile"`
	Tag               string            `json:"tag"`
	MonitoringTargets MonitoringTargets `json:"monitoring"`
	APITarget         *APITarget        `json:"api,omitempty"`
	Plugin            *Plugin           `json:"plugin,omitempty"`
	path              string
	fs                afero.Fs
	locker            locker.Locker
}

func (i *Instance) ID() string {
	return InstanceId(i.Name, i.Tag)
}

type MonitoringTargets struct {
	Targets []MonitoringTarget `json:"targets"`
}

type MonitoringTarget struct {
	Service string `json:"service"`
	Port    string `json:"port"`
	Path    string `json:"path"`
}

type APITarget struct {
	Service string `json:"service"`
	Port    string `json:"port"`
}

type Plugin struct {
	Image string `json:"image"`
}

func (p *Plugin) validate() error {
	if p.Image == "" {
		return fmt.Errorf("%w: plugin image is empty", ErrInvalidInstance)
	}
	return nil
}

// newInstance creates a new instance with the given path as root. It loads the
// state.json file and validates it.
func newInstance(path string, fs afero.Fs, locker locker.Locker) (*Instance, error) {
	i := Instance{
		path: path,
		fs:   fs,
	}
	stateFile, err := i.fs.Open(filepath.Join(i.path, "state.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("%w %s: state.json not found", ErrInvalidInstanceDir, path)
		}
		return nil, err
	}
	defer func() {
		closeErr := stateFile.Close()
		if err == nil {
			err = closeErr
		}
	}()

	stateData, err := io.ReadAll(stateFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(stateData, &i)
	if err != nil {
		return nil, fmt.Errorf("%w %s: invalid state.json file: %s", ErrInvalidInstance, path, err)
	}
	err = i.validate()
	if err != nil {
		return nil, err
	}
	i.locker = locker.New(filepath.Join(path, ".lock"))
	return &i, nil
}

// init initializes a new instance with the given path as root. It creates the
// .lock and state.json files. If the instance is invalid, an error is returned.
func (i *Instance) init(instancePath string, fs afero.Fs, locker locker.Locker) error {
	i.fs = fs
	i.locker = locker
	i.path = instancePath
	err := i.validate()
	if err != nil {
		return err
	}
	err = i.fs.MkdirAll(instancePath, 0o755)
	if err != nil {
		return err
	}

	// Create the lock file
	_, err = i.fs.Create(filepath.Join(i.path, ".lock"))
	if err != nil {
		return err
	}
	// Set lock
	i.locker = i.locker.New(filepath.Join(i.path, ".lock"))

	// Create state file
	stateFile, err := i.fs.Create(filepath.Join(i.path, "state.json"))
	if err != nil {
		return err
	}
	defer func() {
		closeErr := stateFile.Close()
		if err == nil {
			err = closeErr
		}
	}()

	stateData, err := json.Marshal(i)
	if err != nil {
		return err
	}
	_, err = stateFile.Write(stateData)
	return err
}

// Setup creates the instance directory and copies the profile files into it from
// the given fs.FS. It also creates the .env file with the given environment variables
// on the env map.
func (i *Instance) Setup(env map[string]string, profilePath string) (err error) {
	err = i.lock()
	if err != nil {
		return err
	}
	defer func() {
		unlockErr := i.unlock()
		if err == nil {
			err = unlockErr
		}
	}()
	// Create .env file
	envFile, err := i.fs.Create(filepath.Join(i.path, ".env"))
	if err != nil {
		return err
	}
	for k, v := range env {
		envFile.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
	defer envFile.Close()

	// Copy src directory
	err = afero.Walk(i.fs, profilePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(profilePath, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(i.path, relPath)
		if info.IsDir() {
			if err := i.fs.MkdirAll(targetPath, 0o755); err != nil {
				return err
			}
		} else {
			// Skip .env file
			if info.Name() == ".env" {
				return nil
			}
			pkgFile, err := i.fs.Open(path)
			if err != nil {
				return err
			}
			defer pkgFile.Close()
			targetFile, err := i.fs.Create(targetPath)
			if err != nil {
				return err
			}
			if _, err := io.Copy(targetFile, pkgFile); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Check if docker-compose.yml exists
	exists, err := afero.Exists(i.fs, i.ComposePath())
	if err != nil {
		return fmt.Errorf("could not check if docker-compose.yml exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("%w: docker-compose.yml not found", ErrInvalidInstance)
	}
	return nil
}

// ComposePath returns the path to the docker-compose.yml file of the instance.
func (i *Instance) ComposePath() string {
	return filepath.Join(i.path, "docker-compose.yml")
}

// ComposeProject returns the compose project of the instance.
func (i *Instance) ComposeProject() (*types.Project, error) {
	// Load instance environment variables
	instanceEnv, err := i.Env()
	if err != nil {
		return nil, err
	}
	// Build project options with the instance environment
	projectOptions, err := cli.NewProjectOptions([]string{i.ComposePath()})
	if err != nil {
		return nil, err
	}
	maps.Copy(projectOptions.Environment, instanceEnv)
	// Load project from options
	return cli.ProjectFromOptions(projectOptions)
}

// ProfileFile returns the data from the profile.yml file of the instance.
func (i *Instance) ProfileFile() (*profile.Profile, error) {
	if err := i.lock(); err != nil {
		return nil, err
	}
	defer i.unlock()

	var p profile.Profile
	profileFilePath := filepath.Join(i.path, "profile.yml")
	exists, err := afero.Exists(i.fs, profileFilePath)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("profile file not found for instance %s", i.ID())
	}
	profileFile, err := i.fs.Open(profileFilePath)
	if err != nil {
		return nil, err
	}
	defer profileFile.Close()
	profileFileData, err := io.ReadAll(profileFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(profileFileData, &p)
	if err != nil {
		return nil, err
	}
	return &p, p.Validate()
}

// Env returns the environment variables from the .env file of the instance.
func (i *Instance) Env() (map[string]string, error) {
	if err := i.lock(); err != nil {
		return nil, err
	}
	defer i.unlock()
	envPath := filepath.Join(i.path, ".env")
	return env.LoadEnv(i.fs, envPath)
}

// lock locks the .lock file of the instance.
func (i *Instance) lock() error {
	return i.locker.Lock()
}

// unlock unlocks the .lock file of the instance.
func (i *Instance) unlock() error {
	if i.locker == nil || !i.locker.Locked() {
		return errors.New("instance is not locked")
	}
	return i.locker.Unlock()
}

func (i *Instance) validate() error {
	if i.Name == "" {
		return fmt.Errorf("%w: name is empty", ErrInvalidInstance)
	}
	if i.URL == "" {
		return fmt.Errorf("%w: url is empty", ErrInvalidInstance)
	}
	if i.Version == "" && i.Commit == "" {
		return fmt.Errorf("%w: version and commit are empty", ErrInvalidInstance)
	}
	if i.Profile == "" {
		return fmt.Errorf("%w: profile is empty", ErrInvalidInstance)
	}
	if i.Tag == "" {
		return fmt.Errorf("%w: tag is empty", ErrInvalidInstance)
	}

	if i.Plugin != nil {
		if err := i.Plugin.validate(); err != nil {
			return err
		}
	}
	return nil
}
