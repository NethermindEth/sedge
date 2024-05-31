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

type checkContainersTC struct {
	path     string
	runner   commands.CommandRunner
	isErr    bool
	psRunned int
}

func buildCheckContainersTestCase(t *testing.T, caseName string, isErr bool) *checkContainersTC {
	dcPath := t.TempDir()
	test.PrepareTestCaseDir(filepath.Join("testdata", "checks_tests", caseName, configs.DefaultSedgeDataFolderName), dcPath)

	tc := checkContainersTC{}
	tc.path = dcPath
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, int, error) {
			if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "ps") {
				tc.psRunned += 1
				_, err := os.Lstat(filepath.Join(dcPath, configs.DefaultDockerComposeScriptName))
				return "", 1, err
			}
			return "", 0, nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
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

		_, err := CheckContainers(tc.runner, tc.path)
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
