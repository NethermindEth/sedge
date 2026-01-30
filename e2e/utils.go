/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func RunSedge(t *testing.T, binaryPath string, args ...string) error {
	dataDir := filepath.Join(filepath.Dir(binaryPath), "sedge-data")
	return RunCommand(t, binaryPath, "sedge", append([]string{"--path", dataDir}, args...)...)
}

func RunSedgeWithOutput(t *testing.T, binaryPath string, args ...string) ([]byte, error) {
	dataDir := filepath.Join(filepath.Dir(binaryPath), "sedge-data")
	out, _, err := runCommandOutput(t, binaryPath, "sedge", append([]string{"--path", dataDir}, args...)...)
	return out, err
}

func RunCommand(t *testing.T, path string, binaryName string, args ...string) error {
	_, _, err := runCommandOutput(t, path, binaryName, args...)
	return err
}

func RunCommandCMD(t *testing.T, path string, binaryName string, args ...string) *exec.Cmd {
	t.Helper()
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start command: %s %s", binaryName, strings.Join(args, " "))
	}
	return cmd
}

func runCommandOutput(t *testing.T, path string, binaryName string, args ...string) ([]byte, *exec.Cmd, error) {
	t.Helper()
	t.Logf("Binary path: %s", path)
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	out, err := cmd.CombinedOutput()
	t.Logf("===== OUTPUT =====\n%s\n==================", out)
	return out, cmd, err
}

func LogAndPipeError(t *testing.T, prefix string, err error) error {
	t.Helper()
	if err != nil {
		t.Log(prefix, err)
	}
	return err
}
