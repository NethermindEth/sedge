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
	"archive/tar"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	mocks "github.com/NethermindEth/sedge/mocks"
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
			wd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			relativePth, err := filepath.Rel(wd, testDir)
			if err != nil {
				t.Fatal(err)
			}
			absPath, err := filepath.Abs(relativePth)
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

func TestDataDir_Instance(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name       string
		locker     locker.Locker
		instanceId string
		path       string
		instance   *Instance
		err        error
		mockCtrl   *gomock.Controller
	}
	ts := []testCase{
		func() testCase {
			path := t.TempDir()
			// Create instance dir path
			err := fs.MkdirAll(filepath.Join(path, nodesDirName, "mock-avs-default"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			// Create state.json file
			stateFile, err := fs.Create(filepath.Join(path, nodesDirName, "mock-avs-default", "state.json"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = stateFile.WriteString(`{"name":"mock-avs","url":"` + common.MockAvsPkg.Repo() + `","version":"` + common.MockAvsPkg.Version() + `","profile":"option-returner","tag":"default"}`)
			if err != nil {
				t.Fatal(err)
			}
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)
			locker.EXPECT().New(filepath.Join(path, nodesDirName, "mock-avs-default", ".lock")).Return(locker)
			return testCase{
				name:       "valid instance",
				locker:     locker,
				instanceId: "mock-avs-default",
				path:       path,
				instance: &Instance{
					Name:    "mock-avs",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Tag:     "default",
					Profile: "option-returner",
					path:    filepath.Join(path, nodesDirName, "mock-avs-default"),
					fs:      fs,
					locker:  locker,
				},
				err:      nil,
				mockCtrl: ctrl,
			}
		}(),
		func() testCase {
			path := t.TempDir()
			// Create instance dir path
			err := fs.MkdirAll(filepath.Join(path, nodesDirName, "mock-avs-default"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			// Create state.json file
			stateFile, err := fs.Create(filepath.Join(path, nodesDirName, "mock-avs-default", "state.json"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = stateFile.WriteString(`{"url":"` + common.MockAvsPkg.Repo() + `","version":"` + common.MockAvsPkg.Version() + `","profile":"option-returner","tag":"default"}`)
			if err != nil {
				t.Fatal(err)
			}
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)
			return testCase{
				name:       "invalid instance, state without name field",
				locker:     locker,
				instanceId: "mock-avs-default",
				path:       path,
				instance:   nil,
				err:        ErrInvalidInstance,
				mockCtrl:   ctrl,
			}
		}(),
		func() testCase {
			path := t.TempDir()
			// Create nodes dir
			err := fs.MkdirAll(filepath.Join(path, nodesDirName), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)
			return testCase{
				name:       "instance not found",
				locker:     locker,
				instanceId: "mock-avs-default",
				path:       path,
				instance:   nil,
				err:        ErrInvalidInstanceDir,
				mockCtrl:   ctrl,
			}
		}(),
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.mockCtrl.Finish()
			dataDir, err := NewDataDir(tc.path, fs, tc.locker)
			assert.NoError(t, err)
			instance, err := dataDir.Instance(tc.instanceId)
			if tc.err != nil {
				assert.Nil(t, instance)
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.instance, instance)
			}
		})
	}
}

func TestDataDir_InitInstance(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name       string
		path       string
		instance   *Instance
		err        error
		locker     *mocks.MockLocker
		afterCheck func(t *testing.T)
	}
	ts := []testCase{
		func() testCase {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, nodesDirName, "mock-avs-default"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			return testCase{
				name: "instance already exists",
				path: path,
				instance: &Instance{
					Name:    "mock-avs",
					Tag:     "default",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Profile: "option-returner",
					fs:      fs,
					locker:  locker,
				},
				err:    ErrInstanceAlreadyExists,
				locker: locker,
			}
		}(),
		func() testCase {
			path := t.TempDir()
			return testCase{
				name: "invalid instance, state without url field",
				path: path,
				instance: &Instance{
					Name:    "mock-avs",
					Tag:     "default",
					Version: common.MockAvsPkg.Version(),
					Profile: "option-returner",
					fs:      fs,
				},
				err: ErrInvalidInstance,
				afterCheck: func(t *testing.T) {
					assert.NoDirExists(t, filepath.Join(path, nodesDirName, "mock-avs-default"))
				},
			}
		}(),
		func() testCase {
			path := t.TempDir()
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)
			locker.EXPECT().New(filepath.Join(path, nodesDirName, "mock-avs-default", ".lock")).Return(locker)

			return testCase{
				name: "valid instance",
				path: path,
				instance: &Instance{
					Name:    "mock-avs",
					Tag:     "default",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Profile: "option-returner",
					fs:      fs,
					locker:  locker,
				},
				err:    nil,
				locker: locker,
				afterCheck: func(t *testing.T) {
					assert.DirExists(t, filepath.Join(path, nodesDirName, "mock-avs-default"))
					assert.FileExists(t, filepath.Join(path, nodesDirName, "mock-avs-default", "state.json"))
					assert.FileExists(t, filepath.Join(path, nodesDirName, "mock-avs-default", ".lock"))
				},
			}
		}(),
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			dataDir, err := NewDataDir(tc.path, fs, tc.locker)
			assert.NoError(t, err)
			err = dataDir.InitInstance(tc.instance)
			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}
			if tc.afterCheck != nil {
				tc.afterCheck(t)
			}
		})
	}
}

func TestDataDir_HasInstance(t *testing.T) {
	type testCase struct {
		name       string
		dataDir    *DataDir
		instanceId string
		has        bool
	}
	ts := []testCase{
		func() testCase {
			fs := afero.NewMemMapFs()
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			dataDir, err := NewDataDir(testDir, fs, locker)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "instance dir not found",
				dataDir:    dataDir,
				instanceId: "mock_avs-latest",
				has:        false,
			}
		}(),
		func() testCase {
			fs := afero.NewMemMapFs()
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			dataDir, err := NewDataDir(testDir, fs, locker)
			if err != nil {
				t.Fatal(err)
			}
			err = fs.MkdirAll(filepath.Join(testDir, "nodes", "mock_avs-latest"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "instance dir found",
				dataDir:    dataDir,
				instanceId: "mock_avs-latest",
				has:        true,
			}
		}(),
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			has := tc.dataDir.HasInstance(tc.instanceId)
			assert.Equal(t, tc.has, has)
		})
	}
}

func TestDataDir_RemoveInstance(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name       string
		dataDir    *DataDir
		instanceId string
		err        error
	}
	ts := []testCase{
		func() testCase {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			dataDir, err := NewDataDir(testDir, fs, locker)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "instance dir not found",
				dataDir:    dataDir,
				instanceId: "mock_avs-latest",
				err:        fmt.Errorf("%w: mock_avs-latest", ErrInstanceNotFound),
			}
		}(),
		func() testCase {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			err := os.MkdirAll(filepath.Join(testDir, nodesDirName, "mock_avs-latest"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			dataDir, err := NewDataDir(testDir, fs, locker)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "success",
				dataDir:    dataDir,
				instanceId: "mock_avs-latest",
				err:        nil,
			}
		}(),
		func() testCase {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			testDir := t.TempDir()
			err := os.MkdirAll(filepath.Join(testDir, nodesDirName), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			_, err = os.Create(filepath.Join(testDir, nodesDirName, "mock_avs-test"))
			if err != nil {
				t.Fatal(err)
			}
			dataDir, err := NewDataDir(testDir, fs, locker)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "file instead of directory",
				dataDir:    dataDir,
				instanceId: "mock_avs-test",
				err:        errors.New("mock_avs-test is not a directory"),
			}
		}(),
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dataDir.RemoveInstance(tc.instanceId)
			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func TestDataDir_InstancePath(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name       string
		path       string
		instanceId string
		want       string
		wantErr    error
	}
	tests := []testCase{
		func() testCase {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, nodesDirName, "mock-avs-default"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:       "instance dir exists",
				path:       path,
				instanceId: "mock-avs-default",
				want:       filepath.Join(path, nodesDirName, "mock-avs-default"),
				wantErr:    nil,
			}
		}(),
		{
			name:       "instance not found",
			path:       t.TempDir(),
			instanceId: "mock-avs-default",
			want:       "",
			wantErr:    ErrInstanceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			d, err := NewDataDir(tt.path, fs, locker)
			assert.NoError(t, err)
			got, err := d.InstancePath(tt.instanceId)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDir_InitTemp(t *testing.T) {
	fs := afero.NewOsFs()

	type tc struct {
		name    string
		path    string
		id      string
		want    string
		wantErr error
		check   func(t *testing.T)
	}
	tests := []tc{
		func() tc {
			path := t.TempDir()
			return tc{
				name: "empty data dir",
				path: path,
				id:   "temp-dir-id",
				want: filepath.Join(path, tempDir, "temp-dir-id"),
				check: func(t *testing.T) {
					assert.DirExists(t, filepath.Join(path, tempDir, "temp-dir-id"))
				},
			}
		}(),
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name: "empty temp dir",
				path: path,
				id:   "temp-dir-id",
				want: filepath.Join(path, tempDir, "temp-dir-id"),
				check: func(t *testing.T) {
					assert.DirExists(t, filepath.Join(path, tempDir, "temp-dir-id"))
				},
			}
		}(),
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir, "temp-dir-id"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name:    "already exists",
				path:    path,
				id:      "temp-dir-id",
				want:    filepath.Join(path, tempDir, "temp-dir-id"),
				wantErr: nil,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			dataDir, err := NewDataDir(tt.path, fs, locker)
			assert.NoError(t, err)
			got, err := dataDir.InitTemp(tt.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			if tt.check != nil {
				tt.check(t)
			}
		})
	}
}

func TestDataDir_RemoveTemp(t *testing.T) {
	fs := afero.NewOsFs()

	type tc struct {
		name  string
		path  string
		id    string
		check func(t *testing.T)
	}
	tests := []tc{
		{
			name: "empty data dir",
			path: t.TempDir(),
			id:   "mock-avs-default",
		},
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name: "empty temp dir",
				path: path,
				id:   "mock-avs-default",
			}
		}(),
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir, "temp-dir-id"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name: "temp dir exists",
				path: path,
				id:   "temp-dir-id",
				check: func(t *testing.T) {
					assert.NoDirExists(t, filepath.Join(path, tempDir, "temp-dir-id"))
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			dataDir, err := NewDataDir(tt.path, fs, locker)
			assert.NoError(t, err)
			gotErr := dataDir.RemoveTemp(tt.id)
			assert.NoError(t, gotErr)
			if tt.check != nil {
				tt.check(t)
			}
		})
	}
}

func TestDataDir_TempPath(t *testing.T) {
	fs := afero.NewOsFs()

	type tc struct {
		name    string
		path    string
		id      string
		want    string
		wantErr error
	}
	tests := []tc{
		{
			name:    "empty data dir",
			path:    t.TempDir(),
			id:      "temp-dir-id",
			want:    "",
			wantErr: ErrTempDirDoesNotExist,
		},
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name:    "empty temp dir",
				path:    path,
				id:      "temp-dir-id",
				want:    "",
				wantErr: ErrTempDirDoesNotExist,
			}
		}(),
		func() tc {
			path := t.TempDir()
			err := fs.MkdirAll(filepath.Join(path, tempDir, "temp-dir-id"), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name:    "temp dir exists",
				path:    path,
				id:      "temp-dir-id",
				want:    filepath.Join(path, tempDir, "temp-dir-id"),
				wantErr: nil,
			}
		}(),
		func() tc {
			path := t.TempDir()
			err := os.MkdirAll(filepath.Join(path, tempDir), 0o755)
			if err != nil {
				t.Fatal(err)
			}
			_, err = os.Create(filepath.Join(path, tempDir, "temp-dir-id"))
			if err != nil {
				t.Fatal(err)
			}
			return tc{
				name:    "not a directory",
				path:    path,
				id:      "temp-dir-id",
				want:    "",
				wantErr: ErrTempIsNotDir,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			dataDir, err := NewDataDir(tt.path, fs, locker)
			assert.NoError(t, err)
			gotPath, gotErr := dataDir.TempPath(tt.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, gotErr, tt.wantErr)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, tt.want, gotPath)
			}
		})
	}
}

func TestDataDir_initBackupDir(t *testing.T) {
	tc := []struct {
		name  string
		err   error
		setup func(*testing.T) *DataDir
	}{
		{
			name: "backup dir exists",
			err:  nil,
			setup: func(t *testing.T) *DataDir {
				fs := afero.NewMemMapFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(testDir, 0o755)
				require.NoError(t, err)
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
		{
			name: "backup dir does not exist",
			err:  nil,
			setup: func(t *testing.T) *DataDir {
				fs := afero.NewMemMapFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(filepath.Join(testDir, backupDir), 0o755)
				require.NoError(t, err)
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup(t)
			err := d.initBackupDir()
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				exists, err := afero.DirExists(d.fs, filepath.Join(d.path, backupDir))
				require.NoError(t, err)
				assert.True(t, exists)
			}
		})
	}
}

func TestDataDir_HasBackup(t *testing.T) {
	backup := Backup{
		InstanceId: "mock-avs-default",
		Timestamp:  time.Unix(1696340865, 0),
		Version:    common.MockAvsPkg.Version(),
		Commit:     common.MockAvsPkg.CommitHash(),
		Url:        common.MockAvsPkg.Repo(),
	}
	tc := []struct {
		name  string
		setup func() *DataDir
		ok    bool
		err   error
	}{
		{
			name: "backup dir does not exist",
			ok:   false,
			err:  nil,
			setup: func() *DataDir {
				fs := afero.NewMemMapFs()
				testDir := t.TempDir()
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
		{
			name: "backup file does not exist",
			ok:   false,
			err:  nil,
			setup: func() *DataDir {
				fs := afero.NewMemMapFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(filepath.Join(testDir, backupDir), 0o755)
				require.NoError(t, err)
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
		{
			name: "backup file exists",
			ok:   true,
			err:  nil,
			setup: func() *DataDir {
				fs := afero.NewMemMapFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(filepath.Join(testDir, backupDir), 0o755)
				require.NoError(t, err)
				file, err := fs.Create(filepath.Join(testDir, backupDir, backup.Id()+".tar"))
				require.NoError(t, err)
				require.NoError(t, file.Close())
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			ok, err := d.HasBackup(backup.Id())
			if tt.err != nil {
				require.Error(t, err)
				assert.False(t, ok)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.ok, ok)
			}
		})
	}
}

func TestDataDir_InitBackup(t *testing.T) {
	backup := Backup{
		InstanceId: "mock-avs-default",
		Timestamp:  time.Unix(1696340865, 0),
		Version:    common.MockAvsPkg.Version(),
		Commit:     common.MockAvsPkg.CommitHash(),
		Url:        common.MockAvsPkg.Repo(),
	}
	tc := []struct {
		name  string
		err   error
		setup func() *DataDir
	}{
		{
			name: "success, backup dir does not exist",
			err:  nil,
			setup: func() *DataDir {
				fs := afero.NewOsFs()
				testDir := t.TempDir()
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
		{
			name: "success, backup file does not exist",
			err:  nil,
			setup: func() *DataDir {
				fs := afero.NewOsFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(filepath.Join(testDir, backupDir), 0o755)
				require.NoError(t, err)
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
		{
			name: "error, backup file exists",
			err:  ErrBackupAlreadyExists,
			setup: func() *DataDir {
				fs := afero.NewOsFs()
				testDir := t.TempDir()
				err := fs.MkdirAll(filepath.Join(testDir, backupDir), 0o755)
				require.NoError(t, err)
				file, err := fs.Create(filepath.Join(testDir, backupDir, backup.Id()+".tar"))
				require.NoError(t, err)
				require.NoError(t, file.Close())
				return &DataDir{
					path: testDir,
					fs:   fs,
				}
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			err := d.InitBackup(&backup)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				bStat, err := d.fs.Stat(d.BackupPath(backup.Id()))
				require.NoError(t, err)
				require.Equal(t, bStat.Mode(), os.FileMode(0o644))
				require.Equal(t, bStat.Size(), int64(1024))
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
	locker.EXPECT().New(filepath.Join(basePath, "/monitoring", ".lock")).Return(locker).Times(2)

	verify := func(t *testing.T, stack *MonitoringStack) {
		t.Helper()
		assert.Equal(t, filepath.Join(basePath, "/monitoring"), stack.path)
		assert.Equal(t, fs, stack.fs)
		assert.Equal(t, locker, stack.l)

		exists, err := afero.DirExists(fs, filepath.Join(basePath, "/monitoring"))
		assert.NoError(t, err)
		assert.True(t, exists)

		exists, err = afero.Exists(fs, filepath.Join(basePath, "/monitoring", ".lock"))
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

func TestDataDir_BackupList(t *testing.T) {
	type testData struct {
		backup    Backup
		state     []byte
		timestamp time.Time
	}
	tc := []struct {
		name string
		data []testData
		err  error
	}{
		{
			name: "success",
			data: []testData{
				{
					backup: Backup{
						InstanceId: "mock-avs-default",
						Timestamp:  time.Unix(1696420902, 0),
						Version:    "v5.5.1",
						Commit:     "d5af645fffb93e8263b099082a4f512e1917d0af",
						Url:        "https://github.com/NethermindEth/mock-avs-pkg",
					},
					state: []byte(`
					{
						"name": "mock-avs",
						"url": "https://github.com/NethermindEth/mock-avs-pkg",
						"version": "v5.5.1",
						"spec_version": "v0.1.0",
						"commit": "d5af645fffb93e8263b099082a4f512e1917d0af",
						"profile": "option-returner",
						"tag": "default",
						"monitoring": {
						  "targets": [
							{
							  "service": "main-service",
							  "port": "8080",
							  "path": "/metrics"
							}
						  ]
						},
						"api": {
						  "service": "main-service",
						  "port": "8080"
						},
						"plugin": {
						  "image": "mock-avs-plugin:v0.2.0"
						}
					  }
					`),
					timestamp: time.Unix(1696420902, 0),
				},
				{
					backup: Backup{
						InstanceId: "mock-avs-second",
						Timestamp:  time.Unix(1696421646, 0),
						Version:    "v5.5.0",
						Commit:     "a3406616b848164358fdd24465b8eecda5f5ae34",
						Url:        "https://github.com/NethermindEth/mock-avs-pkg",
					},
					state: []byte(`
					{
						"name": "mock-avs",
						"url": "https://github.com/NethermindEth/mock-avs-pkg",
						"version": "v5.5.0",
						"spec_version": "v0.1.0",
						"commit": "a3406616b848164358fdd24465b8eecda5f5ae34",
						"profile": "option-returner",
						"tag": "second",
						"monitoring": {
						  "targets": [
							{
							  "service": "main-service",
							  "port": "8080",
							  "path": "/metrics"
							}
						  ]
						},
						"api": {
						  "service": "main-service",
						  "port": "8080"
						},
						"plugin": {
						  "image": "mock-avs-plugin:v0.1.0"
						}
					  }
					`),
					timestamp: time.Unix(1696421646, 0),
				},
			},
			err: nil,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewOsFs()
			dataDirPath, err := afero.TempDir(fs, "", "datadir")
			require.NoError(t, err)
			dataDir, err := NewDataDir(dataDirPath, fs, nil)
			require.NoError(t, err)

			backups := make([]Backup, 0, len(tt.data))

			for _, d := range tt.data {
				err = dataDir.InitBackup(&d.backup)
				require.NoError(t, err)
				backupTarPath := dataDir.BackupPath(d.backup.Id())
				backupTarFile, err := fs.OpenFile(backupTarPath, os.O_WRONLY, 0o644)
				require.NoError(t, err)
				tarWriter := tar.NewWriter(backupTarFile)
				tarAddStateJson(t, tarWriter, d.state)
				tarAddTimestamp(t, tarWriter, d.timestamp)
				err = tarWriter.Close()
				require.NoError(t, err)
				backups = append(backups, d.backup)
			}

			got, err := dataDir.BackupList()

			for i := 0; i < len(got); i++ {
				got[i].Id()
			}

			if tt.err != nil {
				require.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, backups, got)
			}
		})
	}
}

func TestRemoveMonitoringStack(t *testing.T) {
	// Create monitoring stack
	// Create a memory filesystem
	fs := afero.NewMemMapFs()

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	locker.EXPECT().New(filepath.Join("/monitoring", ".lock")).Return(locker)

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

func tarAddStateJson(t *testing.T, tarWriter *tar.Writer, state []byte) {
	t.Helper()
	header := &tar.Header{
		Name:    "data/state.json",
		Size:    int64(len(state)),
		Mode:    0o644,
		ModTime: time.Now(),
	}
	err := tarWriter.WriteHeader(header)
	require.NoError(t, err)
	_, err = tarWriter.Write(state)
	require.NoError(t, err)
}

func tarAddTimestamp(t *testing.T, tarWriter *tar.Writer, timestamp time.Time) {
	t.Helper()
	data := strconv.FormatInt(timestamp.Unix(), 10)
	header := &tar.Header{
		Name:    "timestamp",
		Size:    int64(len(data)),
		Mode:    0o644,
		ModTime: time.Now(),
	}
	err := tarWriter.WriteHeader(header)
	require.NoError(t, err)
	_, err = tarWriter.Write([]byte(data))
	require.NoError(t, err)
}
