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
package dependencies

import (
	"fmt"
	"regexp"
	"runtime"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	"github.com/stretchr/testify/assert"
)

func TestDependenciesManager_Supported(t *testing.T) {
	tests := []struct {
		name         string
		os           []string
		dependencies []string
		supported    []string
		unsupported  []string
	}{
		{
			name:         "No dependencies",
			os:           []string{"linux", "darwin", "windows"},
			dependencies: []string{},
			supported:    []string{},
			unsupported:  []string{},
		},
		{
			name:         "Supported docker in linux",
			os:           []string{"linux"},
			dependencies: []string{"docker", "wR0n9"},
			supported:    []string{"docker"},
			unsupported:  []string{"wR0n9"},
		},
		{
			name:         "Nothing supported in darwin and windows",
			os:           []string{"darwin", "windows"},
			dependencies: []string{"docker", "wR0n9"},
			supported:    []string{},
			unsupported:  []string{"docker", "wR0n9"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if !contains(tc.os, runtime.GOOS) {
				t.Skipf("Test not supported in %s", runtime.GOOS)
			}
			depsMgr := NewDependenciesManager(nil)
			gotSupported, gotUnsupported, gotErr := depsMgr.Supported(tc.dependencies)
			assert.NoError(t, gotErr)
			assert.Len(t, gotSupported, len(tc.supported))
			for _, s := range gotSupported {
				assert.Contains(t, tc.supported, s)
			}
			assert.Len(t, gotUnsupported, len(tc.unsupported))
			for _, s := range gotUnsupported {
				assert.Contains(t, tc.unsupported, s)
			}
		})
	}
}

func TestDependenciesManager_ShowInstructions(t *testing.T) {
	tests := []struct {
		name       string
		os         []string
		dependency string
		isErr      bool
	}{
		{
			"empty",
			[]string{"linux", "darwin", "windows"},
			"",
			true,
		},
		{
			"docker",
			[]string{"linux"},
			"docker",
			false,
		},
		{
			"wrong",
			[]string{"linux"},
			"wR0n9",
			true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if !contains(tc.os, runtime.GOOS) {
				t.Skipf("Test not supported in %s", runtime.GOOS)
			}
			depsMgr := NewDependenciesManager(nil)
			err := depsMgr.ShowInstructions(tc.dependency)
			if tc.isErr && err == nil {
				t.Errorf("ShowInstructions(%s) expected to fail.", tc.dependency)
			} else if !tc.isErr && err != nil {
				t.Errorf("ShowInstructions(%s) failed: %v", tc.dependency, err)
			}
		})
	}
}

func TestDependenciesManager_Install(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}

	tcs := []struct {
		dependency string
		runner     commands.CommandRunner
		isErr      bool
	}{
		{
			dependency: "docker",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, int, error) {
					return "", 0, nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			isErr: false,
		},
		{
			dependency: "docker-compose",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, int, error) {
					return "", 0, nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", fmt.Errorf("test unexpected error")
				},
			},
			isErr: true,
		},
		{
			dependency: "docker-comp0se",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, int, error) {
					return "", 0, nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			isErr: true,
		},
	}

	for _, tc := range tcs {
		descr := fmt.Sprintf("InstallDependency(%s)", tc.dependency)
		cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{
			RunAsAdmin: false,
		})
		depsMgr := NewDependenciesManager(cmdRunner)
		err := depsMgr.Install(tc.dependency)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

func TestGetScriptPath(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}

	inputs := []struct {
		dependency string
		isErr      bool
	}{
		{
			"",
			true,
		},
		{
			"docker",
			false,
		},
		{
			"wR0n9",
			true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("getScriptPath(%s)", input.dependency)

		path, err := getScriptPath(input.dependency)
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else {
				// TODO: improve results tests

				// Check if path have correct format for dependency
				match, err := regexp.MatchString(`^setup/.+/`+input.dependency+`/.*\.sh$`, path)
				if err != nil || !match {
					t.Errorf("returned path %s is invalid", path)
				}

				// TODO: validate if installation path distro is correct
			}
		}

	}
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
