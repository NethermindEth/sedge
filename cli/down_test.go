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
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/stretchr/testify/assert"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

type downCmdTestCase struct {
	generationPath string
	runner         commands.CommandRunner
	fdOut          *bytes.Buffer
	isErr          bool
}

func buildDownTestCase(t *testing.T, caseName string, isErr bool) *downCmdTestCase {
	tc := downCmdTestCase{}
	configPath := t.TempDir()

	dcPath := filepath.Join(configPath, configs.DefaultSedgeDataFolderName)
	if err := os.Mkdir(dcPath, os.ModePerm); err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "down_tests", caseName, configs.DefaultSedgeDataFolderName), dcPath)
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
	tc.fdOut = new(bytes.Buffer)
	tc.isErr = isErr
	return &tc
}

func TestDownCmd(t *testing.T) {
	tcs := []downCmdTestCase{
		*buildDownTestCase(t, "case_1", false),
	}

	for _, tc := range tcs {
		rootCmd := RootCmd()
		rootCmd.AddCommand(DownCmd(tc.runner))
		rootCmd.SetArgs([]string{"down", "--path", tc.generationPath})
		rootCmd.SetOut(tc.fdOut)
		log.SetOutput(tc.fdOut)

		descr := "sedge down"
		err := rootCmd.Execute()
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

func TestDown_Error(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	// docker compose ps error, PreCheck error
	desc := "docker compose ps error, PreCheck error"
	runner := &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "ps") {
				return "", errors.New("runner error")
			}
			return "", nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}
	tt := buildDownTestCase(t, "case_1", false)

	downCmd := DownCmd(runner)
	downCmd.SetArgs([]string{"--path", tt.generationPath})
	downCmd.SetOutput(io.Discard)
	err := downCmd.Execute()

	if err != nil {
		assert.EqualError(t, err, "it seems docker compose plugin is not installed. Please install it and try again. Error: runner error", desc)
	} else {
		assert.NoError(t, err, desc)
	}

	// docker compose ps --status running error, Check Containers error
	desc = "docker compose ps --status running error, Check Containers error"
	runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "ps") && strings.Contains(c.Cmd, "--filter status=running") {
				return "", errors.New("runner error")
			}
			return "", nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}

	downCmd = DownCmd(runner)
	downCmd.SetArgs([]string{"--path", tt.generationPath})
	downCmd.SetOutput(io.Discard)
	err = downCmd.Execute()

	if err != nil {
		assert.EqualError(t, err, "services of docker-compose script provided are not running. Error: runner error", desc)
	} else {
		assert.NoError(t, err, desc)
	}

	// docker compose down error
	desc = "docker compose down error"
	runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "down") {
				return "", errors.New("runner error")
			}
			return "", nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}

	downCmd = DownCmd(runner)
	downCmd.SetArgs([]string{"--path", tt.generationPath})
	downCmd.SetOutput(io.Discard)
	err = downCmd.Execute()

	if err != nil {
		assert.EqualError(t, err, fmt.Sprintf("command 'docker compose -f %s down' throws error: runner error", filepath.Join(tt.generationPath, "docker-compose.yml")), desc)
	} else {
		assert.NoError(t, err, desc)
	}

	// Generation path error
	desc = "Generation path error"
	runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			return "", nil
		},
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}
	tDir := t.TempDir()

	downCmd = DownCmd(runner)
	downCmd.SetArgs([]string{"--path", tDir})
	downCmd.SetOutput(io.Discard)
	err = downCmd.Execute()

	if err != nil {
		assert.EqualError(t, err, fmt.Sprintf(configs.DockerComposeScriptNotFoundError, tDir, configs.DefaultAbsSedgeDataPath), desc)
	} else {
		assert.NoError(t, err, desc)
	}
}
