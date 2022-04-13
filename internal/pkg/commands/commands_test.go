package commands

import (
	"fmt"
	"testing"
	"text/template"
)

func TestRunCmd(t *testing.T) {
	//TODO: improve test and fix pty commands tests
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

		got, err := Runner.RunBash(BashScript{
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

//TODO: add test cases for building and executing docker commands

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
			output: "docker-compose -f ./testdir/docker-compose.yml down",
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
			output: "docker-compose -f ./testdir/docker-compose.yml logs --follow A B",
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
			output: "docker-compose -f ./testdir/docker-compose.yml logs --tail=20 A B",
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
			output: "docker-compose -f ./testdir/docker-compose.yml logs A B",
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
			output: "docker-compose -f ./testdir/docker-compose.yml ps --services --filter status=running",
		},
		{
			descr: `BuildDockerComposePSCMD(DockerComposePsOptions{
				Path:     "./testdir/docker-compose.yml",
				Services: false,
			})`,
			builder: func() string {
				return Runner.BuildDockerComposePSCMD(DockerComposePsOptions{
					Path:     "./testdir/docker-compose.yml",
					Services: false,
				}).Cmd
			},
			output: "docker-compose -f ./testdir/docker-compose.yml ps --filter status=running",
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
			output: "docker-compose -f ./testdir/docker-compose.yml up -d A B",
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
	}

	for _, input := range inputs {
		got := input.builder()
		if got != input.output {
			t.Errorf("%s expected %q but got %q", input.descr, input.output, got)
		}
	}
}
