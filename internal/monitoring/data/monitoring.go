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
package data

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	"github.com/spf13/afero"
)

// MonitoringStack represents the data stored about the monitoring stack
type MonitoringStack struct {
	path string
	fs   afero.Fs
	l    locker.Locker
}

// newMonitoringStack creates a new monitoring stack with the given path as root.
func newMonitoringStack(path string, fs afero.Fs, locker locker.Locker) *MonitoringStack {
	lock := locker.New(filepath.Join(path, ".lock"))
	return &MonitoringStack{path: path, fs: fs, l: lock}
}

// Init initializes a new monitoring stack with the given path as root.
func (m *MonitoringStack) Init() error {
	// Create the lock file
	_, err := m.fs.Create(filepath.Join(m.path, ".lock"))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInitializingMonitoringStack, err)
	}
	m.l = m.l.New(filepath.Join(m.path, ".lock"))
	return nil
}

// Lock locks the monitoring stack
func (m *MonitoringStack) lock() error {
	if m.l == nil {
		return ErrStackNotInitialized
	}
	return m.l.Lock()
}

// Unlock unlocks the monitoring stack
func (m *MonitoringStack) unlock() error {
	if m.l == nil || !m.l.Locked() {
		return errors.New("monitoring stack is not locked")
	}
	return m.l.Unlock()
}

// Setup sets up the monitoring stack with the given environment variables and
// docker-compose.yml file.
func (m *MonitoringStack) Setup(env map[string]string, monitoringFs fs.FS) (err error) {
	err = m.lock()
	if err != nil {
		return err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	// Create .env file
	envFile, err := m.fs.Create(filepath.Join(m.path, ".env"))
	if err != nil {
		return err
	}
	for k, v := range env {
		_, err = envFile.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return fmt.Errorf("failed to write to .env file: %w", err)
		}
	}
	defer envFile.Close()

	// Copy docker-compose.yml
	mComposeFile, err := monitoringFs.Open("script/docker-compose.yml")
	if err != nil {
		return err
	}
	defer mComposeFile.Close()
	composeFile, err := m.fs.Create(filepath.Join(m.path, "docker-compose.yml"))
	if err != nil {
		return err
	}
	defer composeFile.Close()
	if _, err := io.Copy(composeFile, mComposeFile); err != nil {
		return err
	}

	return nil
}

// CreateDir creates a new directory in the monitoring stack at the given path.
// It creates all the parent directories if they don't exist.
// It does nothing if the directory already exists.
func (m *MonitoringStack) CreateDir(path string) (err error) {
	err = m.lock()
	if err != nil {
		return err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	return m.fs.MkdirAll(filepath.Join(m.path, path), 0o755)
}

// Create creates a new file in the monitoring stack at the given path.
func (m *MonitoringStack) Create(path string) (f afero.File, err error) {
	err = m.lock()
	if err != nil {
		return nil, err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	return m.fs.Create(filepath.Join(m.path, path))
}

// ReadFile reads the file at the given path in the monitoring stack.
func (m *MonitoringStack) ReadFile(path string) (data []byte, err error) {
	err = m.lock()
	if err != nil {
		return nil, err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	data, err = afero.ReadFile(m.fs, filepath.Join(m.path, path))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadingFile, err)
	}
	return
}

// WriteFile writes the given data to the file at the given path in the monitoring stack.
// It creates the file if it doesn't exist.
// It overwrites the file if it already exists.
func (m *MonitoringStack) WriteFile(path string, data []byte) (err error) {
	err = m.lock()
	if err != nil {
		return err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	err = afero.WriteFile(m.fs, filepath.Join(m.path, path), data, 0o644)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWritingFile, err)
	}
	return nil
}

// Installed checks if the monitoring stack is installed.
func (m *MonitoringStack) Installed() (installed bool, err error) {
	err = m.lock()
	if err != nil {
		if errors.Is(err, ErrStackNotInitialized) {
			return false, nil
		}
		return false, err
	}
	defer func() {
		unlockErr := m.unlock()
		if err == nil {
			err = unlockErr
		}
	}()

	toCheck := []string{
		".env",
		"docker-compose.yml",
	}
	for _, path := range toCheck {
		_, err = m.fs.Stat(filepath.Join(m.path, path))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return false, nil
			}
			return false, err
		}
	}
	return true, nil
}

// Path returns the path to the monitoring stack datadir.
func (m *MonitoringStack) Path() string {
	return m.path
}

// Cleanup removes the monitoring stack datadir. If force is true, it doesn't
// lock the monitoring stack.
func (m *MonitoringStack) Cleanup(force bool) (err error) {
	if !force {
		err = m.lock()
		if err != nil {
			return err
		}
		if !m.l.Locked() {
			m.l = nil
			return nil // Already unlocked
		}
		err := m.l.Unlock()
		if err == nil {
			m.l = nil
		}
	}
	return m.fs.RemoveAll(m.path)
}
