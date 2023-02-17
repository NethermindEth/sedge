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
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type cmdArgs struct {
	path     string
	services []string
}

func TestRun(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	emptyErr := errors.New("")

	tcs := []struct {
		name      string
		pTestData string
		args      cmdArgs
		preRunErr error
		depsErr   bool
		setupErr  bool
		runErr    bool
	}{
		{
			name:      "Valid args and valid docker-compose",
			pTestData: "valid",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
		},
		{
			name:      "Valid args, bad docker-compose, no services",
			pTestData: "no_services",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			preRunErr: fmt.Errorf(configs.GenPathErr, emptyErr),
		},
		{
			name:      "Valid args, bad docker-compose, empty services, yaml schema error",
			pTestData: "bad_services",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			preRunErr: fmt.Errorf(configs.InvalidComposeErr, emptyErr),
		},
		{
			name:      "Valid args, bad docker-compose, no version",
			pTestData: "no_version",
			args:      cmdArgs{services: []string{"execution", "consensus"}}, // Leave error commented in case we add a version check
			// preRunErr: fmt.Errorf(configs.MissingVersionErr),
		},
		{
			name:      "Bad compose path",
			args:      cmdArgs{path: "bad_path"},
			preRunErr: fmt.Errorf(configs.ComposeNotFoundErr, emptyErr),
		},
		{
			name:      "Valid docker-compose, bad services",
			pTestData: "valid",
			args:      cmdArgs{services: []string{"bad_service"}},
			preRunErr: fmt.Errorf(configs.InvalidService, "bad_service"),
		},
		{
			name:      "Without env file",
			pTestData: "no_env",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			preRunErr: fmt.Errorf(configs.InvalidComposeErr, emptyErr),
		},
		{
			name:      "No compose file in path",
			pTestData: "no_compose",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			preRunErr: fmt.Errorf(configs.ComposeNotFoundErr, emptyErr),
		},
		{
			name:      "Valid docker-compose, no services provided",
			pTestData: "valid",
			args:      cmdArgs{},
		},
		{
			name:      "Dependencies error",
			pTestData: "valid",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			depsErr:   true,
		},
		{
			name:      "Setup error",
			pTestData: "valid",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			setupErr:  true,
		},
		{
			name:      "Run error",
			pTestData: "valid",
			args:      cmdArgs{services: []string{"execution", "consensus"}},
			runErr:    true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			if tc.pTestData != "" {
				tmp := t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", tc.pTestData), tmp)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				if tc.args.path == "" {
					tc.args.path = tmp
				}
			}
			args := []string{"--path", tc.args.path}
			if len(tc.args.services) > 0 {
				args = append(args, "--services")
				args = append(args, strings.Join(tc.args.services, ","))
			} else {
				tc.args.services = []string{}
			}
			t.Logf("Running test with args: %v", args)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mActions := sedge_mocks.NewMockSedgeActions(ctrl)
			if tc.preRunErr == nil {
				if tc.depsErr {
					mActions.EXPECT().ManageDependencies(gomock.Any()).Return(errors.New(configs.DependencyErr)).Times(1)
				} else {
					mActions.EXPECT().ManageDependencies(gomock.Any()).Return(nil).Times(1)

					sopts := actions.SetupContainersOptions{
						GenerationPath: tc.args.path,
						Services:       tc.args.services,
					}
					if tc.setupErr {
						mActions.EXPECT().SetupContainers(sopts).Return(errors.New(configs.SetupContainersErr)).Times(1)
					} else {
						mActions.EXPECT().SetupContainers(sopts).Return(nil).Times(1)

						ropts := actions.RunContainersOptions{
							GenerationPath: tc.args.path,
							Services:       tc.args.services,
						}
						if tc.runErr {
							mActions.EXPECT().RunContainers(ropts).Return(errors.New(configs.StartingContainersErr)).Times(1)
						} else {
							mActions.EXPECT().RunContainers(ropts).Return(nil).Times(1)
						}
					}
				}
			}

			runCmd := cli.RunCmd(mActions)
			runCmd.SetArgs(args)
			runCmd.SetOutput(io.Discard)
			err := runCmd.Execute()

			if tc.preRunErr != nil {
				assert.ErrorContains(t, err, tc.preRunErr.Error())
			} else if !tc.depsErr && !tc.setupErr && !tc.runErr {
				assert.NoError(t, err)
			}

			if tc.depsErr {
				assert.ErrorContains(t, err, fmt.Errorf(configs.DependencyErr, emptyErr).Error())
				//assert.True(t, strings.Contains(err.Error(), configs.DependencyErr))
			}
			if tc.setupErr {
				assert.ErrorContains(t, err, fmt.Errorf(configs.SetupContainersErr, emptyErr).Error())
				//assert.True(t, strings.Contains(err.Error(), configs.SetupContainersErr))
			}
			if tc.runErr {
				assert.ErrorContains(t, err, fmt.Errorf(configs.StartingContainersErr, emptyErr).Error())
				//assert.True(t, strings.Contains(err.Error(), configs.StartingContainersErr))
			}
		})
	}
}
