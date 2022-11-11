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
package commands

import (
	"fmt"
	"testing"
	"text/template"
)

func TestRunCmd(t *testing.T) {
	// TODO: improve test and fix pty commands tests
	inputs := []struct {
		cmd       string
		getOutput bool
		runInPty  bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo hello world",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:       "echo hello world",
			getOutput: true,
			runInPty:  true,
			output:    "hello world\n",
			isErr:     true,
		},
		{
			cmd:   "wr0n6",
			isErr: true,
		},
	}

	InitRunner(func() CommandRunner {
		return NewCMDRunner(CMDRunnerOptions{
			RunAsAdmin: false,
		})
	})

	for _, input := range inputs {
		descr := fmt.Sprintf("RunCmd(%s,%t,%t)", input.cmd, input.getOutput, input.runInPty)

		got, err := Runner.RunCMD(Command{
			Cmd:       input.cmd,
			GetOutput: input.getOutput,
			RunInPty:  input.runInPty,
		})
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if input.getOutput && input.output != got {
				t.Errorf("%s expected %s but got %s", descr, input.output, got)
			}
		}
	}
}

func TestRunBashScript(t *testing.T) {
	inputs := []struct {
		cmd       string
		getOutput bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo hello world",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:   "wr0n6",
			isErr: true,
		},
	}

	InitRunner(func() CommandRunner {
		return NewCMDRunner(CMDRunnerOptions{
			RunAsAdmin: false,
		})
	})

	for _, input := range inputs {
		descr := fmt.Sprintf("RunBashCmd(%s,%t)", input.cmd, input.getOutput)

		tmp, err := template.New("script").Parse(string(input.cmd))
		if err != nil {
			t.Fatalf("Unexpected error at case %q: %v", input.cmd, err)
		}

		got, err := Runner.RunScript(ScriptFile{
			Tmp:       tmp,
			GetOutput: input.getOutput,
		})
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if input.getOutput && input.output != got {
				t.Errorf("%s expected %s but got %s", descr, input.output, got)
			}
		}
	}
}

// TODO: add test cases for building and executing docker commands

func TestBuildCommands(t *testing.T) {
	inputs := [...]struct {
		descr   string
		builder func() string
		output  string
	}{
		{
			descr: `BuildDockerBuildCMD(DockerBuildOptions{
				Path: "./testdir/dockerfile",
				Tag:  "test:latest",
			}`,
			builder: func() string {
				return Runner.BuildDockerBuildCMD(DockerBuildOptions{
					Path: "./testdir/dockerfile",
					Tag:  "test:latest",
				}).Cmd
			},
			output: "docker build ./testdir/dockerfile -t test:latest",
		},
		{
			descr: `BuildDockerBuildCMD(DockerBuildOptions{
				Path: "./testdir/dockerfile",
				Tag:  "",
			}`,
			builder: func() string {
				return Runner.BuildDockerBuildCMD(DockerBuildOptions{
					Path: "./testdir/dockerfile",
					Tag:  "",
				}).Cmd
			},
			output: "docker build ./testdir/dockerfile",
		},
		{
			descr: `BuildDockerComposeDownCMD(DockerComposeDownOptions{
				Path: "./testdir/docker-compose.yml",
			})`,
			builder: func() string {
				return Runner.BuildDockerComposeDownCMD(DockerComposeDownOptions{
					Path: "./testdir/docker-compose.yml",
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml down",
		},
		{
			descr: `BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: []string{"A", "B"},
				Follow:   true,
				Tail:     20,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: []string{"A", "B"},
					Follow:   true,
					Tail:     20,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml logs --follow A B",
		},
		{
			descr: `BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: []string{"A", "B"},
				Follow:   false,
				Tail:     20,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: []string{"A", "B"},
					Follow:   false,
					Tail:     20,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml logs --tail=20 A B",
		},
		{
			descr: `BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: []string{"A", "B"},
				Follow:   false,
				Tail:     -1,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposeLogsCMD(DockerComposeLogsOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: []string{"A", "B"},
					Follow:   false,
					Tail:     -1,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml logs A B",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: true,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps --services",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: true,
				FilterRunning: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:          "./testdir/docker-compose.yml",
					Services:      true,
					FilterRunning: true,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps --services --filter status=running",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Quiet: true,
				FilterRunning: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:          "./testdir/docker-compose.yml",
					Quiet:         true,
					FilterRunning: true,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps --quiet --filter status=running",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Quiet: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:  "./testdir/docker-compose.yml",
					Quiet: true,
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps --quiet",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path: "./testdir/docker-compose.yml",
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				ServiceName: "service",
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:        "./testdir/docker-compose.yml",
					ServiceName: "service",
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps service",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Quiet:       true,
				ServiceName: "service",
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:        "./testdir/docker-compose.yml",
					Quiet:       true,
					ServiceName: "service",
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml ps --quiet service",
		},
		{
			descr: `BuildDockerComposeUpCMD(DockerComposeUpOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: []string{"A", "B"},
			})`,
			builder: func() string {
				return Runner.BuildDockerComposeUpCMD(DockerComposeUpOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: []string{"A", "B"},
				}).Cmd
			},
			output: "docker compose -f ./testdir/docker-compose.yml up -d A B",
		},
		{
			descr: `BuildDockerInspectCMD(DockerInspectOptions{
				Name: "test:latest",
			})`,
			builder: func() string {
				return Runner.BuildDockerInspectCMD(DockerInspectOptions{
					Name: "test:latest",
				}).Cmd
			},
			output: "docker inspect test:latest",
		},
		{
			descr: `BuildDockerInspectCMD(DockerInspectOptions{
				Name: "test",
				Format: "'{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'",
			})`,
			builder: func() string {
				return Runner.BuildDockerInspectCMD(DockerInspectOptions{
					Name:   "test",
					Format: "'{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'",
				}).Cmd
			},
			output: "docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' test",
		},
		{
			descr: `BuildDockerPSCMD(DockerPSOptions{
				All: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerPSCMD(DockerPSOptions{
					All: true,
				}).Cmd
			},
			output: "docker ps -a",
		},
		{
			descr: `BuildDockerPSCMD(DockerPSOptions{
				All: true,
			})`,
			builder: func() string {
				return Runner.BuildDockerPSCMD(DockerPSOptions{
					All: false,
				}).Cmd
			},
			output: "docker ps",
		},
		{
			descr: `BuildCreateFileCMD(CreateFileOptions{
				FileName: "./testdir/testfile",
			})`,
			builder: func() string {
				return Runner.BuildCreateFileCMD(CreateFileOptions{
					FileName: "./testdir/testfile",
				}).Cmd
			},
			output: "touch ./testdir/testfile",
		},
		{
			descr: `BuildEchoToFileCMD(EchoToFileOptions{
				FileName: "./testdir/testfile",
				Content: "test",
			})`,
			builder: func() string {
				return Runner.BuildEchoToFileCMD(EchoToFileOptions{
					FileName: "./testdir/testfile",
					Content:  "test",
				}).Cmd
			},
			output: "echo test > ./testdir/testfile",
		},
	}

	for _, input := range inputs {
		got := input.builder()
		if got != input.output {
			t.Errorf("%s expected %q but got %q", input.descr, input.output, got)
		}
	}
}
