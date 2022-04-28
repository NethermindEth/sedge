package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/test"
	log "github.com/sirupsen/logrus"
)

type cliCmdTestCase struct {
	name           string
	configPath     string
	generationPath string
	runner         commands.CommandRunner
	fdOut          *bytes.Buffer
	args           cliCmdArgs
	isPreErr       bool
	isErr          bool
}

type cliCmdArgs struct {
	random     bool
	run        bool
	install    bool
	execClient string
	conClient  string
	valClient  string
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
	return strings.Join(s, " ")
}

func resetCliCmd() {
	cfgFile = ""
	executionName = ""
	consensusName = ""
	validatorName = ""
	generationPath = configs.DefaultDockerComposeScriptsPath
	randomize = false
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
	randomize = tc.args.random
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
	if err := preRunCliCmd(rootCmd, []string{}); err != nil {
		return err
	}
	commands.InitRunner(func() commands.CommandRunner {
		return tc.runner
	})
	return nil
}

func buildCliTestCase(t *testing.T, name, caseTestDataDir string, args cliCmdArgs, isPreErr, isErr bool) *cliCmdTestCase {
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
			} else if !tc.isErr && errs != nil && len(errs) > 1 {
				errsStr := []string{}
				for _, err := range errs {
					errsStr = append(errsStr, fmt.Sprintf("%v", err))
				}
				t.Errorf("%s failed with errors: %v", descr, strings.Join(errsStr, "\n"))
			}
		})
	}
}
