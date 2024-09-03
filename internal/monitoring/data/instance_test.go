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
package data_test

import (
	"fmt"
	"io"
	"maps"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring/data/testdata"
	mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInstance(t *testing.T) {
	fs := afero.NewOsFs()

	type testCase struct {
		name     string
		path     string
		instance *Instance
		mocker   func(*mocks.MockLocker)
		err      error
	}
	ts := []testCase{
		func() testCase {
			testDir := t.TempDir()
			return testCase{
				name:     "empty directory",
				path:     testDir,
				instance: nil,
				err:      ErrInvalidInstanceDir,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			_, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:     "empty state file",
				path:     testDir,
				instance: &Instance{path: testDir},
				err:      ErrInvalidInstance,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			stateFile, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			defer stateFile.Close()
			_, err = io.WriteString(stateFile, "{}")
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:     "valid state file (empty state)",
				path:     testDir,
				instance: &Instance{path: testDir},
				err:      ErrInvalidInstance,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			stateFile, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			defer stateFile.Close()
			_, err = io.WriteString(stateFile, `{"name":"test_name","url":"`+common.MockAvsPkg.Repo()+`","version":"`+common.MockAvsPkg.Version()+`","commit":"`+common.MockAvsPkg.CommitHash()+`","profile":"mainnet","tag":"test_tag"}`)
			if err != nil {
				t.Fatal(err)
			}

			return testCase{
				name: "valid state file",
				path: testDir,
				instance: &Instance{
					Name:    "test_name",
					Tag:     "test_tag",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Commit:  common.MockAvsPkg.CommitHash(),
					Profile: "mainnet",
					path:    testDir,
					fs:      fs,
				},
				mocker: func(locker *mocks.MockLocker) {
					locker.EXPECT().New(filepath.Join(testDir, ".lock")).Return(locker)
				},
				err: nil,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			stateFile, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			defer stateFile.Close()
			_, err = io.WriteString(stateFile, `{"name":"test_name","url":"`+common.MockAvsPkg.Repo()+`","version":"`+common.SpecVersion+`"}`)
			if err != nil {
				t.Fatal(err)
			}
			return testCase{
				name:     "invalid state file, missing fields",
				path:     testDir,
				instance: nil,
				err:      ErrInvalidInstance,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			stateFile, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			defer stateFile.Close()
			_, err = io.WriteString(stateFile, `{
				"name":"test_name",
				"url":"`+common.MockAvsPkg.Repo()+`",
				"version":"`+common.MockAvsPkg.Version()+`",
				"commit":"`+common.MockAvsPkg.CommitHash()+`",
				"profile":"mainnet",
				"tag":"test_tag",
				"plugin":{
					"image":"`+common.PluginImage.FullImage()+`"
					}
				}`)
			if err != nil {
				t.Fatal(err)
			}

			return testCase{
				name: "with plugin, remote image",
				path: testDir,
				instance: &Instance{
					Name:    "test_name",
					Tag:     "test_tag",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Commit:  common.MockAvsPkg.CommitHash(),
					Profile: "mainnet",
					Plugin: &Plugin{
						Image: common.PluginImage.FullImage(),
					},
					fs:   fs,
					path: testDir,
				},
				mocker: func(locker *mocks.MockLocker) {
					locker.EXPECT().New(filepath.Join(testDir, ".lock")).Return(locker)
				},
				err: nil,
			}
		}(),
		func() testCase {
			testDir := t.TempDir()
			stateFile, err := fs.Create(testDir + "/state.json")
			if err != nil {
				t.Fatal(err)
			}
			defer stateFile.Close()
			_, err = io.WriteString(stateFile, `{"name":"test_name","url":"`+common.MockAvsPkg.Repo()+`","version":"`+common.MockAvsPkg.Version()+`","commit":"`+common.MockAvsPkg.CommitHash()+`","profile":"mainnet","tag":"test_tag","plugin":{}}`)
			if err != nil {
				t.Fatal(err)
			}

			return testCase{
				name: "error, empty plugin",
				path: testDir,
				instance: &Instance{
					Name:    "test_name",
					Tag:     "test_tag",
					URL:     common.MockAvsPkg.Repo(),
					Version: common.MockAvsPkg.Version(),
					Commit:  common.MockAvsPkg.CommitHash(),
					Profile: "mainnet",
					Plugin:  &Plugin{},
					path:    testDir,
				},
				err: ErrInvalidInstance,
			}
		}(),
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock locker
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			locker := mocks.NewMockLocker(ctrl)

			if tc.mocker != nil {
				tc.mocker(locker)
				tc.instance.locker = locker
			}

			instance, err := newInstance(tc.path, fs, locker)
			if tc.err != nil {
				assert.Nil(t, instance)
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.Equal(t, *tc.instance, *instance)
				assert.NoError(t, err)
			}
		})
	}
}

func TestInstance_Init(t *testing.T) {
	// TODO: Use always the latest version of mock-avs
	ts := []struct {
		name      string
		instance  *Instance
		stateJSON []byte
		err       error
		mocker    func(path string, locker *mocks.MockLocker)
	}{
		{
			name:      "invalid instance",
			instance:  &Instance{},
			stateJSON: nil,
			err:       ErrInvalidInstance,
		},
		{
			name: "valid instance",
			instance: &Instance{
				Name:        "test_name",
				Tag:         "test_tag",
				URL:         common.MockAvsPkg.Repo(),
				Version:     common.MockAvsPkg.Version(),
				SpecVersion: common.SpecVersion,
				Commit:      common.MockAvsPkg.CommitHash(),
				Profile:     "option-returner",
				MonitoringTargets: MonitoringTargets{
					Targets: []MonitoringTarget{
						{
							Service: "main-service",
							Port:    "8080",
							Path:    "/metrics",
						},
					},
				},
			},
			stateJSON: []byte(`{"name":"test_name","url":"` + common.MockAvsPkg.Repo() + `","version":"` + common.MockAvsPkg.Version() + `","spec_version":"` + common.SpecVersion + `","commit":"` + common.MockAvsPkg.CommitHash() + `","profile":"option-returner","tag":"test_tag","monitoring":{"targets":[{"service":"main-service","port":"8080","path":"/metrics"}]}}`),
			mocker: func(path string, locker *mocks.MockLocker) {
				locker.EXPECT().New(filepath.Join(path, ".lock")).Return(locker)
			},
		},
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()

			// Create a mock locker
			ctrl := gomock.NewController(t)
			locker := mocks.NewMockLocker(ctrl)

			path := t.TempDir()

			if tc.mocker != nil {
				tc.mocker(path, locker)
			}

			err := tc.instance.init(path, fs, locker)
			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				stateFile, err := fs.Open(filepath.Join(path, "state.json"))
				assert.NoError(t, err)
				stateData, err := io.ReadAll(stateFile)
				assert.NoError(t, err)
				assert.Equal(t, tc.stateJSON, stateData)
			}
		})
	}
}

func TestInstance_Setup(t *testing.T) {
	fs := afero.NewMemMapFs()
	instancePath, err := afero.TempDir(fs, "", "instance")
	require.NoError(t, err)

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	gomock.InOrder(
		locker.EXPECT().New(filepath.Join(instancePath, ".lock")).Return(locker),
		locker.EXPECT().Lock().Return(nil),
		locker.EXPECT().Locked().Return(true),
		locker.EXPECT().Unlock().Return(nil),
	)

	i := Instance{
		Name:    "mock-avs",
		URL:     common.MockAvsPkg.Repo(),
		Version: common.MockAvsPkg.Version(),
		Commit:  common.MockAvsPkg.CommitHash(),
		Profile: "option-returner",
		Tag:     "test-tag",
	}
	err = i.init(instancePath, fs, locker)
	if err != nil {
		t.Fatal(err)
	}
	env := map[string]string{
		"VAR_1": "value-1",
	}
	profilePath := testdata.SetupProfileFS(t, "option-returner", fs)

	err = i.Setup(env, profilePath)
	assert.NoError(t, err)

	exists, err := afero.Exists(fs, filepath.Join(instancePath, ".env"))
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = afero.Exists(fs, filepath.Join(instancePath, "docker-compose.yml"))
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = afero.Exists(fs, filepath.Join(instancePath, "src"))
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = afero.Exists(fs, filepath.Join(instancePath, "profile.yml"))
	assert.NoError(t, err)
	assert.True(t, exists)

	envFile, err := fs.Open(filepath.Join(instancePath, ".env"))
	assert.NoError(t, err)
	envData, err := io.ReadAll(envFile)
	assert.NoError(t, err)
	assert.Equal(t, []byte("VAR_1=value-1\n"), envData)
}

func TestInstance_Env(t *testing.T) {
	fs := afero.NewMemMapFs()
	tc := []struct {
		name    string
		env     string
		wantEnv map[string]string
		wantErr bool
	}{
		{
			name:    "empty env",
			env:     "empty-env",
			wantEnv: map[string]string{},
			wantErr: false,
		},
		{
			name: "with values",
			env:  "with-values",
			wantEnv: map[string]string{
				"MAIN_SERVICE_NAME": "main-service",
				"MAIN_PORT":         "8080",
				"NETWORK_NAME":      "sedge",
			},
			wantErr: false,
		},
	}
	for i, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			instancePath, err := afero.TempDir(fs, "", fmt.Sprintf("instance-%d", i))
			require.NoError(t, err, "failed to create instance directory")

			envFile, err := fs.Create(filepath.Join(instancePath, ".env"))
			require.NoError(t, err, "failed to create .env file")

			envData := testdata.GetEnv(t, tt.env)
			_, err = io.Copy(envFile, envData)
			require.NoError(t, err, "failed to copy env data")

			err = envFile.Close()
			require.NoError(t, err, "failed to close env file")

			ctrl := gomock.NewController(t)
			l := mocks.NewMockLocker(ctrl)
			defer ctrl.Finish()
			gomock.InOrder(
				l.EXPECT().Lock().Return(nil),
				l.EXPECT().Locked().Return(true),
				l.EXPECT().Unlock().Return(nil),
			)

			i := Instance{
				path:   instancePath,
				fs:     fs,
				locker: l,
			}

			e, err := i.Env()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, maps.Equal(tt.wantEnv, e), "envs are not equal")
			}
		})
	}
}

func TestInstance_ComposeProject(t *testing.T) {
	fs := afero.NewOsFs()
	dir := testdata.SetupProfileFS(t, "option-returner", fs)

	ctrl := gomock.NewController(t)
	l := mocks.NewMockLocker(ctrl)
	defer ctrl.Finish()
	gomock.InOrder(
		l.EXPECT().Lock().Return(nil),
		l.EXPECT().Locked().Return(true),
		l.EXPECT().Unlock().Return(nil),
	)

	i := Instance{
		path:   dir,
		locker: l,
		fs:     fs,
	}
	p, err := i.ComposeProject()
	require.NoError(t, err)

	// Check services
	require.Len(t, p.Services, 1)
	require.Equal(t, "main-service", p.Services[0].Name)
	// Check main-service ports
	mainService := p.Services[0]
	require.Len(t, mainService.Ports, 1)
	require.Equal(t, uint32(8080), mainService.Ports[0].Target)
	require.Equal(t, "8080", mainService.Ports[0].Published)
	// Check main-service container name
	require.Equal(t, "main-service", mainService.ContainerName)
}
