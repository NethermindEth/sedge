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
	"fmt"
	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/test"
	"github.com/NethermindEth/sedge/test/mock_prompts"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type runCmdTestCase struct {
	name       string
	configPath string
	runner     commands.CommandRunner
	monitor    MonitoringTool
	fdOut      *bytes.Buffer
	args       RunCmdFlags
	isPreErr   bool
	isErr      bool
}

func (run *RunCmdFlags) argsList() []string {
	s := make([]string, 0)
	if run.install {
		s = append(s, "-i")
	}
	if run.network != "" {
		s = append(s, "-n", run.network)
	}
	if run.feeRecipient != "" {
		s = append(s, "--fee-recipient", run.feeRecipient)
	}
	if run.generationPath != "" {
		s = append(s, "-p", run.generationPath)
	}
	if run.nodeType != "" {
		s = append(s, "--node-type", run.nodeType)
	}
	if run.nodeValue != "" {
		s = append(s, "--node-value", run.nodeValue)
	}
	return s
}

func (run *RunCmdFlags) toString() string {
	return strings.Join(run.argsList(), " ")
}

func prepareRunCmd(tc runCmdTestCase) {
	// Set output buffers
	log.SetOutput(tc.fdOut)
	// Set config file path
	cfgFile = tc.configPath
	initConfig()
	commands.InitRunner(func() commands.CommandRunner {
		return tc.runner
	})
	initMonitor(func() MonitoringTool {
		return tc.monitor
	})
}

func buildRunTestCase(
	t *testing.T,
	name,
	caseTestDataDir string,
	args RunCmdFlags,
	isPreErr,
	isErr bool,
) *runCmdTestCase {
	tc := runCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "cli_tests", caseTestDataDir, "config"), configPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	dcPath := filepath.Join(configPath, "docker-compose-scripts")
	err = os.Mkdir(dcPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	// TODO: allow runner edition
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			// For getContainerIP logic
			if strings.Contains(c.Cmd, "ps --quiet") {
				return "666", nil
			} else if strings.Contains(c.Cmd, "docker inspect 666") {
				return inspectOut, nil
			}
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}

	// Check for port occupation
	defaultsPorts := map[string]string{
		"ELApi": configs.DefaultApiPortEL,
		"CLApi": configs.DefaultApiPortCL,
	}
	ports, err := utils.AssignPorts("localhost", defaultsPorts)
	if err != nil {
		t.Fatalf(configs.PortOccupationError, err)
	}

	tc.monitor = &monitorStub{
		data: []posmoni.EndpointSyncStatus{
			{Endpoint: inspectExecutionUrl + ":" + ports["ELApi"], Synced: true},
			{Endpoint: inspectConsensusUrl + ":" + ports["CLApi"], Synced: true},
			{Endpoint: inspectExecutionUrl + ":" + ports["ELApi"], Synced: true},
			{Endpoint: inspectConsensusUrl + ":" + ports["CLApi"], Synced: true},
		},
	}

	tc.name = name
	tc.args = args
	tc.args.generationPath = dcPath
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.fdOut = new(bytes.Buffer)
	tc.isPreErr = isPreErr
	tc.isErr = isErr
	return &tc
}

func TestRunCmd(t *testing.T) {
	// TODO: Add more test cases
	tcs := []runCmdTestCase{
		*buildRunTestCase(
			t,
			"Missing node type", "case_1",
			RunCmdFlags{
				CmdFlags: CmdFlags{
					yes: true,
				},
			},
			false,
			true,
		),
		*buildRunTestCase(
			t,
			"Fixed clients", "case_1",
			RunCmdFlags{
				CmdFlags: CmdFlags{
					yes: true,
				},
				nodeValue: "nethermind",
				nodeType:  execution,
			},
			false,
			false,
		),
		*buildRunTestCase(
			t,
			"Wrong value for nodeValue", "case_1",
			RunCmdFlags{
				CmdFlags: CmdFlags{
					yes: true,
				},
				nodeValue: "teku",
				nodeType:  execution,
			},
			false,
			true,
		),
		*buildRunTestCase(
			t,
			"Good network input", "case_1",
			RunCmdFlags{
				CmdFlags: CmdFlags{
					yes:     true,
					network: "mainnet",
				},
				nodeType: consensus,
			},
			false,
			false,
		),
		*buildRunTestCase(
			t,
			"Bad network input", "case_1",
			RunCmdFlags{
				CmdFlags: CmdFlags{
					yes:     true,
					network: "sedge",
				},
				nodeType: execution,
			},
			true,
			true,
		),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			descr := fmt.Sprintf("sedge run %s", tc.args.toString())

			ctrl := gomock.NewController(t)
			prompt := mock_prompts.NewMockPrompt(ctrl)
			defer ctrl.Finish()
			if !tc.isErr {
				prompt.EXPECT().FeeRecipient().Return("0x0000000000000000000000000000000000000000", nil)
			}

			rootCmd := RootCmd()
			rootCmd.AddCommand(RunCmd(prompt))
			argsL := append([]string{"run"}, tc.args.argsList()...)
			rootCmd.SetArgs(argsL)

			prepareRunCmd(tc)

			if err := rootCmd.Execute(); !tc.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if tc.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			}
		})
	}
}
