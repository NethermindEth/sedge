package cmd

import (
	"testing"
)

func TestRootCmdExecute(t *testing.T) {
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("rootCmd.Execute() failed: %v", err)
	}
}
