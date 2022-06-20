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

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/internal/utils"
	"github.com/NethermindEth/1click/test"
	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	log "github.com/sirupsen/logrus"
)

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
	random       bool
	run          bool
	install      bool
	execClient   string
	conClient    string
	valClient    string
	network      string
	feeRecipient string
}

func (args *cliCmdArgs) toString() string {
	s := []string{}
	if args.random {
		s = append(s, "-r")
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
	run = tc.args.run
	install = tc.args.install
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

func buildCliTestCase(t *testing.T, name, caseTestDataDir string, args cliCmdArgs, monitorData []posmoni.EndpointSyncStatus, isPreErr, isErr bool) *cliCmdTestCase {
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
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}

	tc.monitor = &monitorStub{data: monitorData}

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
			"Random clients",
			"case_1",
			cliCmdArgs{
				random:  true,
				run:     true,
				install: true,
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Fixed clients",
			"case_1",
			cliCmdArgs{
				run:        true,
				install:    true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				valClient:  "lighthouse",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Missing consensus client",
			"case_1",
			cliCmdArgs{
				run:        true,
				install:    true,
				execClient: "nethermind",
				valClient:  "lighthouse",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Missing validator client", "case_1",
			cliCmdArgs{
				run:        true,
				install:    true,
				execClient: "nethermind",
				conClient:  "lighthouse",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Good network input", "case_1",
			cliCmdArgs{
				run:        true,
				install:    true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				network:    "mainnet",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
		*buildCliTestCase(
			t,
			"Bad network input", "case_1",
			cliCmdArgs{
				run:        true,
				install:    true,
				execClient: "nethermind",
				conClient:  "lighthouse",
				network:    "1click",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			true,
			true,
		),
		*buildCliTestCase(
			t,
			"Bad fee recipient input", "case_1",
			cliCmdArgs{
				run:          true,
				install:      true,
				execClient:   "nethermind",
				conClient:    "lighthouse",
				network:      "kiln",
				feeRecipient: "666",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			true,
			false,
		),
		*buildCliTestCase(
			t,
			"Good fee recipient input", "case_1",
			cliCmdArgs{
				run:          true,
				install:      true,
				execClient:   "nethermind",
				conClient:    "lighthouse",
				network:      "kiln",
				feeRecipient: "0x5c00ABEf07604C59Ac72E859E5F93D5ab8546F83",
			},
			[]posmoni.EndpointSyncStatus{
				{Endpoint: configs.OnPremiseExecutionURL, Synced: true},
				{Endpoint: configs.OnPremiseConsensusURL, Synced: true},
			},
			false,
			false,
		),
	}

	t.Cleanup(resetCliCmd)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resetCliCmd()
			descr := fmt.Sprintf("1click cli %s", tc.args.toString())

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

			err := trackSync(&ms, time.Millisecond*100)
			utils.CheckErr("trackSync(...) failed", tc.isError, err)
		})
	}
}
