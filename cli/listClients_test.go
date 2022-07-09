package cli

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

type listClientsTestCase struct {
	name       string
	logsOut    *bytes.Buffer
	tableOut   *bytes.Buffer
	configPath string
	isErr      bool
}

func resetListClientsCmd() {
	cfgFile = ""
}

func buildListClientTestCase(t *testing.T, name, caseTestDataDir string, isErr bool) listClientsTestCase {
	tc := listClientsTestCase{}
	tmpDir := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "list_clients_tests", caseTestDataDir, "config"), tmpDir)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc.name = name
	tc.configPath = filepath.Join(tmpDir, "config.yaml")
	tc.logsOut = new(bytes.Buffer)
	tc.tableOut = new(bytes.Buffer)
	tc.isErr = isErr
	return tc
}

func TestListClientsCmd(t *testing.T) {
	tcs := [...]listClientsTestCase{
		buildListClientTestCase(t, "Ok", "case_1", false),
		buildListClientTestCase(t, "Missing validator clients", "case_2", true),
		buildListClientTestCase(t, "Using json format", "case_3", false),
	}

	t.Cleanup(resetListClientsCmd)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resetListClientsCmd()
			rootCmd.SetOut(tc.tableOut)
			log.SetOutput(tc.logsOut)
			args := []string{}
			// Configure config path
			cfgFile = tc.configPath
			initConfig()

			err := runListClientsCmd(rootCmd, args)
			if tc.isErr && err == nil {
				t.Error("sedge clients expected to fail")
			} else if !tc.isErr && err != nil {
				t.Errorf("sedge clients failed: %v", err)
			}
		})
	}

}
