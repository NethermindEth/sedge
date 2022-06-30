package cli

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

type logsTestCase struct {
	name          string
	runner        commands.CommandRunner
	configPath    string
	generatedPath string
	fdOut         *bytes.Buffer
	services      []string
	isErr         bool
	dcPsRuns      int
	dcLogsRuns    int
}

func resetLogCmd() {
	cfgFile = ""
	tail = false
	generationPath = configs.DefaultDockerComposeScriptsPath
}

func prepareLogsTestCaseConfigDir(name, dest string) (string, error) {
	caseConfigPath := filepath.Join(".", "testdata", "logs_tests", name, "config")
	configPath := filepath.Join(dest, "config")
	if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
		return "", err
	}
	err := test.PrepareTestCaseDir(caseConfigPath, configPath)
	return filepath.Join(configPath, "config.yaml"), err
}

func prepareLogsTestCaseDCScripts(name, dest string) (string, error) {
	caseDCScriptsPath := filepath.Join("testdata", "logs_tests", name, "docker-compose-scripts")
	dcPath := filepath.Join(dest, "docker-compose-scripts")
	if err := os.MkdirAll(dcPath, os.ModePerm); err != nil {
		return "", err
	}
	err := test.PrepareTestCaseDir(caseDCScriptsPath, dcPath)
	return dcPath, err
}

func prepareFiles(t *testing.T, tc *logsTestCase) {
	tempDir := t.TempDir()
	tcConfigPath, err := prepareLogsTestCaseConfigDir(tc.name, tempDir)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	tcGeneratedPath, err := prepareLogsTestCaseDCScripts(tc.name, tempDir)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc.configPath = tcConfigPath
	tc.generatedPath = tcGeneratedPath
}

func buildLogsTestCase(t *testing.T, testName string, services []string, isErr bool) logsTestCase {
	tc := logsTestCase{}
	tc.name = testName

	fdOut := new(bytes.Buffer)

	//TODO: allow modification of the simple runner
	runner := test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker compose") {
				if strings.Contains(c.Cmd, "logs") {
					tc.dcLogsRuns += 1
					return "some logs", nil
				}
				if strings.Contains(c.Cmd, "ps") {
					tc.dcPsRuns += 1
					return ` Name             Command           State                       Ports
					----------------------------------------------------------------------------
					execution            bash              Up                  0.0.0.0:80->80/tcp
					consensus            bash              Up                  0.0.0.0:80->80/tcp
					validator            bash              Up                  0.0.0.0:80->80/tcp`, nil
				}
			}
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}

	prepareFiles(t, &tc)

	tc.runner = &runner
	tc.fdOut = fdOut
	tc.services = services
	tc.isErr = isErr
	return tc
}

func TestLogsCmd(t *testing.T) {
	tc1 := buildLogsTestCase(
		t,
		"case_1",
		[]string{"execution", "consensus", "validator"},
		false,
	)

	tcs := []logsTestCase{
		tc1,
	}

	t.Cleanup(resetLogCmd)

	for _, tc := range tcs {
		resetLogCmd()
		args := []string{"logs", "--config", tc.configPath, "--path", tc.generatedPath, "--tail"}
		args = append(args, tc.services...)
		rootCmd.SetArgs(args)
		rootCmd.SetOut(tc.fdOut)
		log.SetOutput(tc.fdOut)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		descr := fmt.Sprintf("1click logs --tail %s", strings.Join(tc.services, " "))

		err := rootCmd.Execute()
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
