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
	"fmt"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	"github.com/spf13/afero"
)

const monitoringStackDirName = "monitoring"

// DataDir is the directory where all the data is stored.
type DataDir struct {
	path   string
	fs     afero.Fs
	locker locker.Locker
}

// NewDataDir creates a new DataDir instance with the given path as root.
func NewDataDir(path string, fs afero.Fs, locker locker.Locker) (*DataDir, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &DataDir{path: absPath, fs: fs, locker: locker}, nil
}

// Path returns the path of the data dir.
func (d *DataDir) Path() string {
	return d.path
}

// NewDataDirDefault creates a new DataDir instance with the default path as root.
// Default path is $XDG_DATA_HOME/.sedge or $HOME/.local/share/.sedge if $XDG_DATA_HOME is not set
// as defined in the XDG Base Directory Specification
func NewDataDirDefault(fs afero.Fs, locker locker.Locker) (*DataDir, error) {
	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		userDataHome = filepath.Join(userHome, ".local", "share")
	}
	dataDir := filepath.Join(userDataHome, ".sedge")
	err := fs.MkdirAll(dataDir, 0o755)
	if err != nil {
		return nil, err
	}

	return NewDataDir(dataDir, fs, locker)
}

// MonitoringStack checks if a monitoring stack directory exists in the data directory.
// If the directory does not exist, it creates it and initializes a new MonitoringStack instance.
// If the directory exists, it simply returns a new MonitoringStack instance.
// It returns an error if there is any issue accessing or creating the directory, or initializing the MonitoringStack.
func (d *DataDir) MonitoringStack() (*MonitoringStack, error) {
	monitoringStackPath := filepath.Join(d.path, monitoringStackDirName)
	_, err := d.fs.Stat(monitoringStackPath)
	if os.IsNotExist(err) {
		if err = d.fs.MkdirAll(monitoringStackPath, 0o755); err != nil {
			return nil, err
		}

		monitoringStack := &MonitoringStack{path: monitoringStackPath, fs: d.fs, l: d.locker}
		if err = monitoringStack.Init(); err != nil {
			return nil, err
		}
		return monitoringStack, nil
	} else if err != nil {
		return nil, err
	}

	return newMonitoringStack(monitoringStackPath, d.fs, d.locker), nil
}

// RemoveMonitoringStack removes the monitoring stack directory from the data directory.
// It returns an error if there is any issue accessing or removing the directory.
func (d *DataDir) RemoveMonitoringStack() error {
	monitoringStackPath := filepath.Join(d.path, monitoringStackDirName)
	_, err := d.fs.Stat(monitoringStackPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrMonitoringStackNotFound, monitoringStackPath)
	} else if err != nil {
		return err
	}

	return d.fs.RemoveAll(monitoringStackPath)
}
