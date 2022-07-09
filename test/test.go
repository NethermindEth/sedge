package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
)

// Creates an empty file in a temp dir with "dependency" as name and add
// temp dir to PATH. Return the path for the new dependency.
func CreateFakeDep(t *testing.T, dependency string) (depPath string) {
	depPath = t.TempDir()
	file, err := os.Create(filepath.Join(depPath, dependency))
	if err != nil {
		t.Fatalf("Can't fake dependency %s", dependency)
	}
	file.Close()

	PATH := os.Getenv("PATH")
	err = os.Setenv("PATH", fmt.Sprintf("%s:%s", PATH, depPath))
	if err != nil {
		t.Fatalf("Can't fake dependency %s", dependency)
	}
	return
}

// Remove the "depPath" from PATH.
func DeleteFakeDep(depPath string) {
	PATH := os.Getenv("PATH")
	PATH = strings.ReplaceAll(PATH, ":"+depPath, "")
	os.Setenv("PATH", PATH)
}

// Struct for creating a commands.CommandRunner mocks
type SimpleCMDRunner struct {
	SRunCMD  func(commands.Command) (string, error)
	SRunBash func(commands.BashScript) (string, error)
}

func (cr *SimpleCMDRunner) BuildDockerComposeUpCMD(options commands.DockerComposeUpOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeUpCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerPSCMD(options commands.DockerPSOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerPSCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposePSCMD(options commands.DockerComposePsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposePSCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeLogsCMD(options commands.DockerComposeLogsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeLogsCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerBuildCMD(options commands.DockerBuildOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerBuildCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerPullCMD(options commands.DockerBuildOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerPullCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerInspectCMD(options commands.DockerInspectOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerInspectCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeDownCMD(options commands.DockerComposeDownOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeDownCMD(options)
}

func (cr *SimpleCMDRunner) RunCMD(cmd commands.Command) (string, error) {
	return cr.SRunCMD(cmd)
}

func (cr *SimpleCMDRunner) RunBash(script commands.BashScript) (string, error) {
	return cr.SRunBash(script)
}
