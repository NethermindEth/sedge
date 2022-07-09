package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/configs"
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
	generationPath = configs.DefaultDockerComposeScriptsPath
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
		SRunBash: func(bs commands.BashScript) (string, error) {
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
	//TODO: allow to test error programs
	tcs := []downCmdTestCase{
		*buildDownTestCase(t, "case_1", false),
	}

	t.Cleanup(resetDownCmd)

	for _, tc := range tcs {
		resetDownCmd()
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
