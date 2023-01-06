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
	"regexp"
	"runtime"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
)

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
			false,
		},
		{
			"docker",
			false,
		},
		{
			"wR0n9",
			false,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("getScriptPath(%s)", input.dependency)

		path, _, err := getScriptPath(input.dependency)
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

func TestDependencySupported(t *testing.T) {
	if runtime.GOOS != "linux" { // TODO: update when other os's are supported
		t.Skipf("this os is not supported")
	}

	inputs := []struct {
		dependency string
		result     bool
	}{
		{
			"",
			false,
		},
		{
			"docker",
			true,
		},
		{
			"wR0n9",
			false,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("dependencySupported(%s)", input.dependency)
		got := dependencySupported(input.dependency)
		if got != input.result {
			t.Errorf("%s expected %t but got %t", descr, input.result, got)
		}
	}
}

func TestHandleInstructions(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}

	inputs := []struct {
		caseName      string
		dependencies  []string
		handlerResult error
		isErr         bool
	}{
		{
			"No dependencies to handle",
			[]string{},
			nil,
			false,
		},
		{
			"Handle docker and docker-composes",
			[]string{"docker", "docker-compose"},
			nil,
			true,
		},
		{
			"Handle docker and docker-compose with handler failure",
			[]string{"docker", "docker-compose"},
			fmt.Errorf("unexpected error"),
			true,
		},
		{
			"Handle docker and invalid dependency",
			[]string{"docker", "wR0n9"},
			nil,
			true,
		},
	}

	for _, input := range inputs {
		t.Run(input.caseName, func(t *testing.T) {
			descr := fmt.Sprintf("HandlerInstruction(%s,...)", input.dependencies)
			err := HandleInstructions(&test.SimpleCMDRunner{}, input.dependencies, func(_ commands.CommandRunner, dependency string) error {
				contained := false
				for _, expected := range input.dependencies {
					if dependency == expected {
						contained = true
					}
				}
				if !contained {
					t.Errorf("%s handler called on unnexpected dependency %s", descr, dependency)
				}

				t.Logf("%s handler returning error %v", descr, input.handlerResult)
				return input.handlerResult
			})
			if input.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			} else if !input.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}
		})
	}
}

func TestShowInstructions(t *testing.T) {
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
		descr := fmt.Sprintf("ShowInstructions(%s)", input.dependency)

		if err := ShowInstructions(&test.SimpleCMDRunner{}, input.dependency); input.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

func TestInstallDependency(t *testing.T) {
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
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
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
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
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
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
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

		err := InstallDependency(tc.runner, tc.dependency)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
