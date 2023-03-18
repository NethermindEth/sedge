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
package test

// notest

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
	err = os.Chmod(filepath.Join(depPath, dependency), 0o777)
	if err != nil {
		t.Fatalf("Can't fake dependency %s", dependency)
	}

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
	SRunCMD  func(commands.Command) (string, int, error)
	SRunBash func(commands.ScriptFile) (string, error)
}

func (cr *SimpleCMDRunner) BuildDockerComposeUpCMD(options commands.DockerComposeUpOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeUpCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposePullCMD(options commands.DockerComposePullOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposePullCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeCreateCMD(options commands.DockerComposeCreateOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeCreateCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeBuildCMD(options commands.DockerComposeBuildOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeBuildCMD(options)
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

func (cr *SimpleCMDRunner) BuildDockerComposeVersionCMD() commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeVersionCMD()
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

func (cr *SimpleCMDRunner) BuildCreateFileCMD(options commands.CreateFileOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildCreateFileCMD(options)
}

func (cr *SimpleCMDRunner) BuildEchoToFileCMD(options commands.EchoToFileOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildEchoToFileCMD(options)
}

func (cr *SimpleCMDRunner) BuildOpenTextEditor(options commands.OpenTextEditorOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildOpenTextEditor(options)
}

func (cr *SimpleCMDRunner) RunCMD(cmd commands.Command) (string, int, error) {
	if cr.SRunCMD == nil {
		return "", 0, nil
	}
	return cr.SRunCMD(cmd)
}

func (cr *SimpleCMDRunner) RunScript(script commands.ScriptFile) (string, error) {
	if cr.SRunBash == nil {
		return "", nil
	}
	return cr.SRunBash(script)
}
