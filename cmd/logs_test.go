package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/test"
	log "github.com/sirupsen/logrus"
)

type logsTestCase struct {
	runner     commands.CommandRunner
	configPath string
	fdOut      *bytes.Buffer
	services   []string
	isErr      bool
	dcPsRuns   int
	dcLogsRuns int
}

func prepareLogsTestCaseConfigDir(name, dest string) (string, error) {
	caseConfigPath := filepath.Join(".", "testdata", "logs_tests", fmt.Sprintf("case_%s", name), "config")
	err := test.PrepareTestCaseDir(caseConfigPath, dest)
	return dest, err
}

func prepareLogsTestCaseDCScripts(t *testing.T, name string) error {
	caseDCScriptsPath := filepath.Join("testdata", "logs_tests", fmt.Sprintf("case_%s", name), "docker-compose-scripts")
	if err := os.Mkdir("docker-compose-scripts", os.ModePerm); err != nil {
		return err
	}
	t.Cleanup(func() {
		os.RemoveAll("docker-compose-scripts")
	})
	err := test.PrepareTestCaseDir(caseDCScriptsPath, "docker-compose-scripts")
	return err
}

func buildLogsTestCase(t *testing.T, testName string, services []string, isErr bool) logsTestCase {
	tcConfigPath, err := prepareLogsTestCaseConfigDir(testName, t.TempDir())
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	err = prepareLogsTestCaseDCScripts(t, testName)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc := logsTestCase{}

	fdOut := new(bytes.Buffer)

	//TODO: allow modification of the simple runner
	runner := test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker-compose") {
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

	tc.runner = &runner
	tc.configPath = tcConfigPath
	tc.fdOut = fdOut
	tc.services = services
	tc.isErr = isErr
	return tc
}

func TestLogsCmd(t *testing.T) {
	tc1 := buildLogsTestCase(
		t,
		"1",
		[]string{"execution", "consensus", "validator"},
		false,
	)

	tcs := []logsTestCase{
		tc1,
	}

	for _, tc := range tcs {
		args := []string{"logs", "--config", tc.configPath, "--tail"}
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
