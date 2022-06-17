package cli

import (
	"path/filepath"
	"testing"

	"github.com/NethermindEth/1click/test"
)

func resetRootCmd() {
	cfgFile = ""
}

func TestRootCmdExecute(t *testing.T) {
	configPath := t.TempDir()
	t.Cleanup(resetRootCmd)
	err := test.PrepareTestCaseDir(filepath.Join(".", "testdata", "root_test", "config"), configPath)
	if err != nil {
		t.Errorf("Can't create config file: %v", err)
	}

	rootCmd.SetArgs([]string{"--config", configPath})

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("rootCmd.Execute() failed: %v", err)
	}
}
