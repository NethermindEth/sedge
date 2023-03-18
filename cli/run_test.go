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
package cli_test

import (
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tcs := []struct {
		name  string
		args  []string
		setup func(*testing.T, *sedge_mocks.MockDependenciesManager, *sedge_mocks.MockSedgeActions) string
		err   string
	}{
		{
			name: "Valid flags and valid docker-compose",
			args: []string{"--services", "execution,consensus"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "valid"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml"), []string{"execution", "consensus"}).Return(nil).Times(1),
					a.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution", "consensus"},
					}).Return(nil).Times(1),
					a.EXPECT().RunContainers(actions.RunContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution", "consensus"},
					}).Return(nil).Times(1),
				)
				return
			},
		},
		{
			name: "Valid flags, bad docker-compose",
			args: []string{"--services", "execution,consensus"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "no_services"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml"), []string{"execution", "consensus"}).Return(errors.New("bad docker-compose")).Times(1),
				)
				return
			},
			err: "bad docker-compose",
		},
		{
			name: "With args",
			args: []string{"arg"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				return configs.DefaultAbsSedgeDataPath
			},
			err: "command run does not support arguments. Please use flags instead",
		},
		{
			name: "Without docker",
			args: []string{"--services", "execution,consensus"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "valid"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return(nil, []string{dependencies.Docker}).Times(1),
				)
				return
			},
			err: "missing dependencies: docker",
		},
		{
			name: "Error setting up containers",
			args: []string{"--services", "execution,consensus"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "valid"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml"), []string{"execution", "consensus"}).Return(nil).Times(1),
					a.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution", "consensus"},
					}).Return(errors.New("setup error")).Times(1),
				)
				return
			},
			err: "error setting up service containers: setup error",
		},
		{
			name: "Error running containers",
			args: []string{"--services", "execution,consensus"},
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "valid"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml"), []string{"execution", "consensus"}).Return(nil).Times(1),
					a.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution", "consensus"},
					}).Return(nil).Times(1),
					a.EXPECT().RunContainers(actions.RunContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution", "consensus"},
					}).Return(errors.New("run error")).Times(1),
				)
				return
			},
			err: "error starting service containers: run error",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mActions := sedge_mocks.NewMockSedgeActions(ctrl)
			depsMgr := sedge_mocks.NewMockDependenciesManager(ctrl)
			generationPath := tc.setup(t, depsMgr, mActions)

			args := []string{"--path", generationPath}
			args = append(args, tc.args...)

			t.Logf("Running test with args: %v", args)

			runCmd := cli.RunCmd(mActions, depsMgr)
			runCmd.SetArgs(args)
			runCmd.SetOutput(io.Discard)
			err := runCmd.Execute()

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
