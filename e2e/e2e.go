package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

type (
	e2eArranger func(t *testing.T, sedgePath string) error
	e2eAct      func(t *testing.T, sedgePath, dataDirPath string)
	e2eAssert   func(t *testing.T, dataDirPath string)
)

type e2eTestCase struct {
	t        *testing.T
	testDir  string
	repoPath string
	arranger e2eArranger
	act      e2eAct
	assert   e2eAssert
}

func newE2ETestCase(t *testing.T, arranger e2eArranger, act e2eAct, assert e2eAssert) *e2eTestCase {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tc := &e2eTestCase{
		t:        t,
		testDir:  t.TempDir(),
		repoPath: filepath.Dir(wd),
		arranger: arranger,
		act:      act,
		assert:   assert,
	}
	t.Logf("Creating new E2E test case (%p). TestDir: %s", tc, tc.testDir)
	checkGoInstalled(t)
	tc.installGoModules()
	tc.build()
	return tc
}

func (e *e2eTestCase) run() {
	// Cleanup environment before and after test
	e.Cleanup()
	defer e.Cleanup()
	if e.arranger != nil {
		err := e.arranger(e.t, e.BinaryPath())
		if err != nil {
			e.t.Fatalf("error in Arrange step: %v", err)
		}
	}
	if e.act != nil {
		e.act(e.t, e.BinaryPath(), e.dataDirPath())
	}
	if e.assert != nil {
		e.assert(e.t, e.dataDirPath())
	}
}

func (e *e2eTestCase) BinaryPath() string {
	binaryName := "sedge"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	return filepath.Join(e.testDir, binaryName)
}

func (e *e2eTestCase) Cleanup() {
	dataDir := e.dataDirPath()
	// Check if data directory exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return
	}
	// Remove docker compose stack
	err := exec.Command("docker", "compose", "-f", filepath.Join(dataDir, "docker-compose.yml"), "down", "--volumes").Run()
	if err != nil {
		e.t.Logf("error removing docker compose stack: %v", err)
	}

	// Remove sedge-data directory
	err = os.RemoveAll(dataDir)
	if err != nil {
		e.t.Logf("error removing data directory: %v", err)
	}
}

func (e *e2eTestCase) installGoModules() {
	e.t.Helper()
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = e.repoPath
	e.t.Logf("Installing Go modules in %s", e.repoPath)
	if err := cmd.Run(); err != nil {
		e.t.Fatalf("error installing Go modules: %v", err)
	} else {
		e.t.Logf("Go modules installed")
	}
}

func (e *e2eTestCase) build() {
	e.t.Helper()
	binaryName := "sedge"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	outPath := filepath.Join(e.testDir, binaryName)
	e.t.Logf("Building binary to %s", outPath)
	err := exec.Command("go", "build", "-o", outPath, filepath.Join(e.repoPath, "cmd", "sedge", "main.go")).Run()
	if err != nil {
		e.t.Fatalf("error building binary: %v", err)
	} else {
		e.t.Logf("binary built")
	}
}

func (e *e2eTestCase) dataDirPath() string {
	return filepath.Join(e.testDir, "sedge-data")
}

func checkGoInstalled(t *testing.T) {
	t.Helper()
	err := exec.Command("go", "version").Run()
	if err != nil {
		t.Fatalf("error checking Go installation: %v", err)
	} else {
		t.Logf("Go installed")
	}
}
