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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

var (
	inspectExecutionUrl = "http://192.168.128.3"
	inspectConsensusUrl = "http://192.168.128.3"
)

var inspectOut = `
[
	{
		"NetworkSettings": {
			"Bridge": "",
			"SandboxID": "56e2c759c33315c9de009bd70aac0fdeb9367549303433debb71edff8dd4db39",
			"HairpinMode": false,
			"LinkLocalIPv6Address": "",
			"LinkLocalIPv6PrefixLen": 0,
			"Ports": {
				"30303/tcp": [
					{
						"HostIp": "0.0.0.0",
						"HostPort": "30303"
					}
				],
				"30303/udp": [
					{
						"HostIp": "0.0.0.0",
						"HostPort": "30303"
					}
				],
				"8008/tcp": [
					{
						"HostIp": "0.0.0.0",
						"HostPort": "8008"
					}
				],
				"8545/tcp": [
					{
						"HostIp": "0.0.0.0",
						"HostPort": "8560"
					}
				]
			},
			"SandboxKey": "/var/run/docker/netns/56e2c759c333",
			"SecondaryIPAddresses": null,
			"SecondaryIPv6Addresses": null,
			"EndpointID": "",
			"Gateway": "",
			"GlobalIPv6Address": "",
			"GlobalIPv6PrefixLen": 0,
			"IPAddress": "",
			"IPPrefixLen": 0,
			"IPv6Gateway": "",
			"MacAddress": "",
			"Networks": {
				"sedge_network": {
					"IPAMConfig": null,
					"Links": null,
					"Aliases": [
						"execution-client",
						"execution",
						"babf61f2c52a"
					],
					"NetworkID": "b4bb0c21aa1c9495d08309f8f7f4f2fb5a493fd925c880cb146045aafb2f4390",
					"EndpointID": "7832cdd23f1f9f70e38576f8088da61010e057bffb0b98c83bd391065d703ed9",
					"Gateway": "192.168.128.1",
					"IPAddress": "192.168.128.3",
					"IPPrefixLen": 20,
					"IPv6Gateway": "",
					"GlobalIPv6Address": "",
					"GlobalIPv6PrefixLen": 0,
					"MacAddress": "02:42:c0:a8:80:03",
					"DriverOpts": null
				}
			}
		}
	}
]
`

type cliCmdTestCase struct {
	name           string
	configPath     string
	generationPath string
	runner         commands.CommandRunner
	monitor        MonitoringTool
	fdOut          *bytes.Buffer
	args           cliCmdArgs
	isPreErr       bool
	isErr          bool
}

type cliCmdArgs struct {
	yes          bool
	run          bool
	install      bool
	execClient   string
	conClient    string
	valClient    string
	network      string
	feeRecipient string
	services     []string
}

func (args *cliCmdArgs) toString() string {
	s := []string{}
	if args.yes {
		s = append(s, "--yes")
	}
	if args.run {
		s = append(s, "--run")
	}
	if args.install {
		s = append(s, "-i")
	}
	if args.execClient != "" {
		s = append(s, "-e", args.execClient)
	}
	if args.conClient != "" {
		s = append(s, "-c", args.conClient)
	}
	if args.valClient != "" {
		s = append(s, "-v", args.valClient)
	}
	if args.network != "" {
		s = append(s, "-n", args.network)
	}
	if args.feeRecipient != "" {
		s = append(s, "--fee-recipient", args.feeRecipient)
	}
	if len(*services) == 0 {
		s = append(s, "--run-client none")
	} else {
		s = append(s, strings.Join(*services, ", "))
	}
	return strings.Join(s, " ")
}

func resetCliCmd() {
	cfgFile = ""
	executionName = ""
	consensusName = ""
	validatorName = ""
	network = "mainnet"
	feeRecipient = ""
	generationPath = configs.DefaultDockerComposeScriptsPath
	install = false
	run = false
	y = false
	services = &[]string{}
	fallbackEL = &[]string{}
	jwtPath = ""
	checkpointSyncUrl = ""
}

func prepareCliCmd(tc cliCmdTestCase) error {
	// Set output buffers
	rootCmd.SetOut(tc.fdOut)
	log.SetOutput(tc.fdOut)
	// Set config file path
	cfgFile = tc.configPath
	initConfig()
	// Set flags
	generationPath = tc.generationPath
	y = tc.args.yes
	run = tc.args.run
	install = tc.args.install
	services = &tc.args.services
	if tc.args.execClient != "" {
		executionName = tc.args.execClient
	}
	if tc.args.conClient != "" {
		consensusName = tc.args.conClient
	}
	if tc.args.valClient != "" {
		validatorName = tc.args.valClient
	}
	if tc.args.network != "" {
		network = tc.args.network
	}
	if tc.args.feeRecipient != "" {
		feeRecipient = tc.args.feeRecipient
	}
	if err := preRunCliCmd(rootCmd, []string{}); err != nil {
		return err
	}
	commands.InitRunner(func() commands.CommandRunner {
		return tc.runner
	})
	initMonitor(func() MonitoringTool {
		return tc.monitor
	})
	return nil
}

func buildCliTestCase(
	t *testing.T,
	name,
	caseTestDataDir string,
	args cliCmdArgs,
	isPreErr,
	isErr bool,
) *cliCmdTestCase {
	tc := cliCmdTestCase{}
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
		SRunBash: func(bs commands.ScriptFile) (string, error) {
			return "", nil
		},
	}

	// Check for port occupation
	defaultsPorts := map[string]string{
		"ELApi": configs.DefaultApiPortEL,
		"CLApi": configs.DefaultApiPortCL,
	}
	ports, err := utils.AssingPorts("localhost", defaultsPorts)
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
	tc.generationPath = dcPath
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.fdOut = new(bytes.Buffer)
	tc.isPreErr = isPreErr
	tc.isErr = isErr
	return &tc
}

func TestCliCmd(t *testing.T) {
	// TODO: Add more test cases
	tcs := []cliCmdTestCase{
		*buildCliTestCase(
			t,
			"Random clients", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{execution, consensus},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Fixed clients", "case_1",
			cliCmdArgs{
				yes:        true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				valClient:  "lighthouse",
				services:   []string{execution, consensus},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Missing consensus client", "case_1",
			cliCmdArgs{
				yes:        true,
				execClient: "nethermind",
				valClient:  "lighthouse",
				services:   []string{execution, consensus},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Missing validator client", "case_1",
			cliCmdArgs{
				yes:        true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				services:   []string{execution, consensus},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Good network input", "case_1",
			cliCmdArgs{
				yes:        true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				network:    "mainnet",
				services:   []string{execution, consensus},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Bad network input", "case_1",
			cliCmdArgs{
				yes:        true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				network:    "sedge",
				services:   []string{execution, consensus},
			},
			true,
			true,
		),
		*buildCliTestCase(
			t,
			"--run-client all", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{"all"},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"--run-client none", "case_1",
			cliCmdArgs{
				yes: true,
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"--run-client none, execution, ambiguos error", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{execution, "none"},
			},
			true,
			false,
		),
		*buildCliTestCase(
			t,
			"--run-client validator", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{validator},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"--run-client all, validator, ambiguos error", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{validator, "all"},
			},
			true,
			false,
		),
		*buildCliTestCase(
			t,
			"--run-client all, validator, ambiguos error", "case_1",
			cliCmdArgs{
				yes:      true,
				services: []string{validator, "all"},
			},
			true,
			false,
		),
		*buildCliTestCase(
			t,
			"Invalid network", "case_1",
			cliCmdArgs{
				yes:      true,
				network:  "test",
				services: []string{execution, consensus},
			},
			true,
			true,
		),
	}

	t.Cleanup(resetCliCmd)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resetCliCmd()
			descr := fmt.Sprintf("sedge cli %s", tc.args.toString())

			err := prepareCliCmd(tc)
			if tc.isPreErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			} else if !tc.isPreErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}

			errs := runCliCmd(rootCmd, []string{})
			if tc.isErr && (errs == nil || len(errs) < 1) {
				t.Errorf("%s expected to fail", descr)
			} else if !tc.isErr && errs != nil && len(errs) > 0 {
				errsStr := []string{}
				for _, err := range errs {
					errsStr = append(errsStr, fmt.Sprintf("%v", err))
				}
				t.Errorf("%s failed with errors: %v", descr, strings.Join(errsStr, "\n"))
			}
		})
	}
}

// Stub for MonitoringTool interface
type monitorStub struct {
	data  []posmoni.EndpointSyncStatus
	calls int
}

func (ms *monitorStub) TrackSync(done <-chan struct{}, beaconEndpoints, executionEndpoints []string, wait time.Duration) <-chan posmoni.EndpointSyncStatus {
	ms.calls++
	c := make(chan posmoni.EndpointSyncStatus, len(ms.data))
	var w time.Duration

	go func() {
		for {
			select {
			case <-done:
				close(c)
				return
			case <-time.After(w):
				if w == 0 {
					// Don't wait the first time
					w = wait
				}
				for _, d := range ms.data {
					c <- d
				}
			}
		}
	}()

	return c
}

func TestTrackSync(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name    string
		data    []posmoni.EndpointSyncStatus
		isError bool
	}{
		{
			"Test case 1, execution client got an error",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false, Error: errors.New("")},
			},
			true,
		},
		{
			"Test case 2, execution client got an error, consensus client not synced",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false, Error: errors.New("")},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false},
			},
			true,
		},
		{
			"Test case 3, execution client got an error, consensus client synced",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false, Error: errors.New("")},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			true,
		},
		{
			"Test case 4, bad execution client response, good consensus client response",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true, Error: errors.New("")},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			true,
		},
		{
			"Test case 5, consensus client got an error, consensus client not synced",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false, Error: errors.New("")},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
			},
			true,
		},
		{
			"Test case 6, consensus client got an error, consensus client synced",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false, Error: errors.New("")},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
			},
			true,
		},
		{
			"Test case 7, mixed results",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
		},
		{
			"Test case 8, mixed results",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
			},
			false,
		},
		{
			"Test case 9, mixed results, error",
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false},
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: false, Error: errors.New("")},
			},
			true,
		},
	}
	// TODO: Starvation bc a synced status from one of the clients is not tested

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ms := monitorStub{data: tc.data}

			err := trackSync(&ms, "", "", time.Millisecond*100)
			utils.CheckErr("trackSync(...) failed", tc.isError, err)
		})
	}
}
