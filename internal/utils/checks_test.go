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
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
)

func validatePending(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for _, dep1 := range got {
		found := false
		for _, dep2 := range want {
			if dep1 == dep2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestCheckDependencies(t *testing.T) {
	inputs := [...]struct {
		dependencies []string
		pending      []string
	}{
		{},
		{
			[]string{"wr0n9"},
			[]string{"wr0n9"},
		},
		{
			[]string{""},
			[]string{""},
		},
		{
			[]string{"curl"},
			[]string{},
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("CheckDependencies(%s)", input.dependencies)
		got := CheckDependencies(input.dependencies)
		if !validatePending(got, input.pending) {
			t.Errorf("%s expected %s but got %s", descr, input.pending, got)
		}
	}
}

func TestPreCheck(t *testing.T) {
	tcs := [...]struct {
		name            string
		caseTestDataDir string
		path            string
		runner          commands.CommandRunner
		isErr           bool
		noDocker        bool
		noCompose       bool
	}{
		{
			name:            "Ok",
			caseTestDataDir: "case_1",
			path:            t.TempDir(),
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
				},
				SRunBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			isErr: false,
		},
		{
			name:            "Missing docker compose yaml",
			caseTestDataDir: "case_2",
			path:            t.TempDir(),
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
				},
				SRunBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			isErr: true,
		},
		{
			name:            "Command failure",
			caseTestDataDir: "case_1",
			path:            t.TempDir(),
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					return "", fmt.Errorf("test unknown error")
				},
				SRunBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			isErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			descr := fmt.Sprintf("PreCheck(%s)", tc.path)
			dPath, dcPath := "", ""

			if !tc.noDocker {
				dPath = test.CreateFakeDep(t, "docker")
			}
			if !tc.noCompose {
				dcPath = test.CreateFakeDep(t, "docker-compose")
			}

			commands.InitRunner(func() commands.CommandRunner {
				return tc.runner
			})

			test.PrepareTestCaseDir(filepath.Join("testdata", "checks_tests", tc.caseTestDataDir, "docker-compose-scripts"), tc.path)

			err := PreCheck(tc.path)
			if tc.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			} else if !tc.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}

			if dPath != "" {
				test.DeleteFakeDep(dPath)
			}
			if dcPath != "" {
				test.DeleteFakeDep(dcPath)
			}
		})
	}
}

type checkContainersTC struct {
	path     string
	runner   commands.CommandRunner
	isErr    bool
	psRunned int
}

func buildCheckContainersTestCase(t *testing.T, caseName string, isErr bool) *checkContainersTC {
	dcPath := t.TempDir()
	test.PrepareTestCaseDir(filepath.Join("testdata", "checks_tests", caseName, "docker-compose-scripts"), dcPath)

	tc := checkContainersTC{}
	tc.path = dcPath
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "ps") {
				tc.psRunned += 1
				_, err := os.Lstat(filepath.Join(dcPath, configs.DefaultDockerComposeScriptName))
				return "", err
			}
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}
	tc.isErr = isErr
	return &tc
}

func TestCheckContainers(t *testing.T) {
	tcs := [...]*checkContainersTC{
		buildCheckContainersTestCase(
			t,
			"case_1",
			false,
		),
		buildCheckContainersTestCase(
			t,
			"case_2",
			true,
		),
	}

	for _, tc := range tcs {
		descr := fmt.Sprintf("CheckContainers(%s)", tc.path)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		_, err := CheckContainers(tc.path)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if tc.psRunned < 1 {
				t.Errorf("%s didn't run docker compose ps", descr)
			}
		}
	}
}
