package e2e

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func runSedge(t *testing.T, binaryPath string, args ...string) error {
	dataDir := filepath.Join(filepath.Dir(binaryPath), "sedge-data")
	return runCommand(t, binaryPath, append([]string{"--path", dataDir}, args...)...)
}

func runCommand(t *testing.T, path string, args ...string) error {
	_, err := runCommandOutput(t, path, args...)
	return err
}

func runCommandOutput(t *testing.T, path string, args ...string) ([]byte, error) {
	t.Helper()
	t.Logf("Binary path: %s", path)
	t.Logf("Running command: sedge %s", strings.Join(args, " "))
	out, err := exec.Command(path, args...).CombinedOutput()
	t.Logf("===== OUTPUT =====\n%s\n==================", out)
	return out, err
}
