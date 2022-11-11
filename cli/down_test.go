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
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

type downCmdTestCase struct {
	configPath     string
	generationPath string
	runner         commands.CommandRunner
	fdOut          *bytes.Buffer
	isErr          bool
}

func resetDownCmd() {
	cfgFile = ""
}

func buildDownTestCase(t *testing.T, caseName string, isErr bool) *downCmdTestCase {
	tc := downCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "down_tests", caseName, "config"), configPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	dcPath := filepath.Join(configPath, "docker-compose-scripts")
	if err = os.Mkdir(dcPath, os.ModePerm); err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	err = test.PrepareTestCaseDir(filepath.Join("testdata", "down_tests", caseName, "docker-compose-scripts"), dcPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	// TODO: allow runner edition
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			return "", nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}

	tc.generationPath = dcPath
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.fdOut = new(bytes.Buffer)
	tc.isErr = isErr
	return &tc
}

func TestDownCmd(t *testing.T) {
	// TODO: allow to test error programs
	tcs := []downCmdTestCase{
		*buildDownTestCase(t, "case_1", false),
	}

	t.Cleanup(resetDownCmd)

	for _, tc := range tcs {
		resetDownCmd()
		rootCmd := RootCmd()
		rootCmd.AddCommand(DownCmd())
		rootCmd.SetArgs([]string{"down", "--config", tc.configPath, "--path", tc.generationPath})
		rootCmd.SetOut(tc.fdOut)
		log.SetOutput(tc.fdOut)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		descr := "sedge down"
		err := rootCmd.Execute()
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
