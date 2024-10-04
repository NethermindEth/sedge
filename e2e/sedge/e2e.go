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
	"runtime"
	"testing"

	base "github.com/NethermindEth/sedge/e2e"
)

type (
	e2eArranger func(t *testing.T, sedgePath string) error
	e2eAct      func(t *testing.T, sedgePath, dataDirPath string)
	e2eAssert   func(t *testing.T, dataDirPath string)
)

type e2eSedgeTestCase struct {
	base.E2ETestCase
	arranger e2eArranger
	act      e2eAct
	assert   e2eAssert
}

func newE2ESedgeTestCase(t *testing.T, arranger e2eArranger, act e2eAct, assert e2eAssert) *e2eSedgeTestCase {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tc := &e2eSedgeTestCase{
		E2ETestCase: base.E2ETestCase{
			T:          t,
			TestDir:    t.TempDir(),
			RepoPath:   filepath.Dir(filepath.Dir(wd)),
			BinaryName: "sedge",
		},
		arranger: arranger,
		act:      act,
		assert:   assert,
	}
	t.Logf("Creating new E2E test case (%p). TestDir: %s", tc, tc.TestDir)
	base.CheckGoInstalled(t)
	tc.E2ETestCase.InstallGoModules()
	tc.E2ETestCase.Build()
	return tc
}

func (e *e2eSedgeTestCase) run() {
	// Cleanup environment before and after test
	e.Cleanup()
	defer e.Cleanup()
	if e.arranger != nil {
		err := e.arranger(e.T, e.BinaryPath())
		if err != nil {
			e.T.Fatalf("error in Arrange step: %v", err)
		}
	}
	if e.act != nil {
		e.act(e.T, e.BinaryPath(), e.dataDirPath())
	}
	if e.assert != nil {
		e.assert(e.T, e.dataDirPath())
	}
}

func (e *e2eSedgeTestCase) BinaryPath() string {
	binaryName := "sedge"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	return filepath.Join(e.TestDir, binaryName)
}

func (e *e2eSedgeTestCase) Cleanup() {
	dataDir := e.dataDirPath()
	// Check if data directory exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return
	}
	// Remove docker compose stack
	err := exec.Command("docker", "compose", "-f", filepath.Join(dataDir, "docker-compose.yml"), "down", "--volumes").Run()
	if err != nil {
		e.T.Logf("error removing docker compose stack: %v", err)
	}

	// Remove sedge-data directory
	err = os.RemoveAll(dataDir)
	if err != nil {
		e.T.Logf("error removing data directory: %v", err)
	}
}

func (e *e2eSedgeTestCase) dataDirPath() string {
	return filepath.Join(e.TestDir, "sedge-data")
}
