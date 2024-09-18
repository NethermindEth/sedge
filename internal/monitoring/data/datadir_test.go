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
	"os"
	"path/filepath"
	"testing"

	mocks "github.com/NethermindEth/sedge/internal/monitoring/locker/mocks"
	"github.com/NethermindEth/sedge/internal/monitoring/utils"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDataDir(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name    string
		path    string
		dataDir *DataDir
		locker  *mocks.MockLocker
		err     error
	}
	ts := []testCase{
		func() testCase {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			absPath, err := filepath.Abs(testDir)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name: "path to absolute",
				path: testDir,
				dataDir: &DataDir{
					path:   absPath,
					fs:     fs,
					locker: locker,
				},
				locker: locker,
				err:    nil,
			}
		}(),
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			dataDir, err := NewDataDir(tc.path, fs, tc.locker)
			if tc.err != nil {
				assert.Nil(t, dataDir)
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.dataDir, dataDir)
			}
		})
	}
}

func TestMonitoringStack(t *testing.T) {
	// Create a memory filesystem
	fs := afero.NewMemMapFs()
	userHome, err := os.UserHomeDir()
	require.NoError(t, err)
	basePath := filepath.Join(userHome, ".local", "share", ".sedge")

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "monitoring", ".lock")}).Return(locker).Times(2)

	verify := func(t *testing.T, stack *MonitoringStack) {
		t.Helper()
		assert.Equal(t, filepath.Join(basePath, "/monitoring"), stack.path)
		assert.Equal(t, fs, stack.fs)
		assert.Equal(t, locker, stack.l)

		exists, err := afero.DirExists(fs, filepath.Join(basePath, "/monitoring"))
		assert.NoError(t, err)
		assert.True(t, exists)

		exists, err = afero.Exists(fs, filepath.Join(basePath, "monitoring", ".lock"))
		assert.NoError(t, err)
		assert.True(t, exists)
	}
	// Create a data dir
	dataDir, err := NewDataDirDefault(fs, locker)
	require.NoError(t, err)

	// Create a monitoring stack
	monitoringStack, err := dataDir.MonitoringStack()
	require.NoError(t, err)
	verify(t, monitoringStack)

	// Try to get a monitoring stack while it does exist
	monitoringStack, err = dataDir.MonitoringStack()
	require.NoError(t, err)
	verify(t, monitoringStack)
}

func TestRemoveMonitoringStack(t *testing.T) {
	// Create monitoring stack
	// Create a memory filesystem
	fs := afero.NewMemMapFs()

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join("monitoring", ".lock")}).Return(locker)

	// Create a data dir
	dataDir, err := NewDataDir("/", fs, locker)
	require.NoError(t, err)

	// Create a monitoring stack
	_, err = dataDir.MonitoringStack()
	require.NoError(t, err)

	// Remove monitoring stack
	err = dataDir.RemoveMonitoringStack()
	require.NoError(t, err)

	exists, err := afero.DirExists(fs, filepath.Join("/monitoring"))
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestRemoveMonitoringStackError(t *testing.T) {
	// Create monitoring stack
	// Create a memory filesystem
	fs := afero.NewMemMapFs()

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)

	// Create a data dir
	dataDir, err := NewDataDir("/", fs, locker)
	require.NoError(t, err)

	// Remove monitoring stack
	err = dataDir.RemoveMonitoringStack()
	require.ErrorIs(t, err, ErrMonitoringStackNotFound)
}
