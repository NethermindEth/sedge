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
	configPath     string
	generationPath string
	runner         commands.CommandRunner
	fdOut          *bytes.Buffer
	args           []string
	isErr          bool
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

func buildCliTestCase(t *testing.T, caseName string, args []string, isErr bool) *cliCmdTestCase {
	tc := cliCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "cli_tests", caseName, "config"), configPath)
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

	tc.args = args
	tc.generationPath = dcPath
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.fdOut = new(bytes.Buffer)
	tc.isErr = isErr
	return &tc
}

func TestCliCmd(t *testing.T) {
	//TODO: allow to test error programs
	tcs := []cliCmdTestCase{
		*buildCliTestCase(t, "case_1", []string{"-r"}, false),
		*buildCliTestCase(t, "case_1", []string{"-e", "nethermind", "-c", "lighthouse", "-v", "lighthouse"}, false),
		// *buildCliTestCase(t, "case_1", []string{"-e", "nethermind", "-v", "lighthouse"}, true),
		// *buildCliTestCase(t, "case_1", []string{"-e", "nethermind", "-c", "lighthouse"}, true),
	}

	t.Cleanup(resetCliCmd)

	for _, tc := range tcs {
		resetCliCmd()
		rootCmd.SetArgs(append([]string{"cli", "--config", tc.configPath, "--path", tc.generationPath, "-i", "--run"}, tc.args...))
		rootCmd.SetOut(tc.fdOut)
		log.SetOutput(tc.fdOut)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		descr := fmt.Sprintf("1click cli -i --run %s", strings.Join(tc.args, " "))
		err := rootCmd.Execute()
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
