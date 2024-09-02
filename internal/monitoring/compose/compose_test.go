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
package compose

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUp(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposeUpOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts: commands.DockerComposeUpOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposeUpOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "up"},
		},
		{
			name: "it runs the correct command when no services are specified",
			opts: commands.DockerComposeUpOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{},
			},
			runCMDError: nil,
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			var expectedCmd string
			if len(tt.opts.Services) > 0 {
				expectedCmd = "docker compose -f " + tt.opts.Path + " up -d " + strings.Join(tt.opts.Services, " ")
			} else {
				expectedCmd = "docker compose -f " + tt.opts.Path + " up -d"
			}
			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposeUpCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Up(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Up() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Up command
	opts := commands.DockerComposeUpOptions{
		Path:     "/path/to/docker-compose.yml",
		Services: []string{"service1", "service2"},
	}

	// Run the Docker Compose Up command
	manager.Up(opts)
}

func TestPull(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposePullOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts: commands.DockerComposePullOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposePullOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "pull"},
		},
		{
			name: "it runs the correct command when no services are specified",
			opts: commands.DockerComposePullOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{},
			},
			runCMDError: nil,
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			var expectedCmd string
			if len(tt.opts.Services) > 0 {
				expectedCmd = "docker compose -f " + tt.opts.Path + " pull " + strings.Join(tt.opts.Services, " ")
			} else {
				expectedCmd = "docker compose -f " + tt.opts.Path + " pull"
			}
			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposePullCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Pull(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Pull() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Pull command
	opts := commands.DockerComposePullOptions{
		Path:     "/path/to/docker-compose.yml",
		Services: []string{"service1", "service2"},
	}

	// Run the Docker Compose Pull command
	manager.Pull(opts)
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposeCreateOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts: commands.DockerComposeCreateOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposeCreateOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "create"},
		},
		{
			name: "it runs the correct command when no services are specified",
			opts: commands.DockerComposeCreateOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{},
			},
			runCMDError: nil,
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			expectedCmd := "docker compose -f " + tt.opts.Path + " create"

			if len(tt.opts.Services) > 0 {
				expectedCmd += " " + strings.Join(tt.opts.Services, " ")
			}

			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposeCreateCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Create(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Create() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Create command
	opts := commands.DockerComposeCreateOptions{
		Path:     "/path/to/docker-compose.yml",
		Services: []string{"service1", "service2"},
	}

	// Run the Docker Compose Create command
	manager.Create(opts)
}

func TestBuild(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposeBuildOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts: commands.DockerComposeBuildOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposeBuildOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "build"},
		},
		{
			name: "it runs the correct command when no services are specified",
			opts: commands.DockerComposeBuildOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{},
			},
			runCMDError: nil,
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			var expectedCmd string
			if len(tt.opts.Services) > 0 {
				expectedCmd = "docker compose -f " + tt.opts.Path + " build " + strings.Join(tt.opts.Services, " ")
			} else {
				expectedCmd = "docker compose -f " + tt.opts.Path + " build"
			}

			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposeBuildCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Build(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Build() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Build command
	opts := commands.DockerComposeBuildOptions{
		Path:     "/path/to/docker-compose.yml",
		Services: []string{"service1", "service2"},
	}

	// Run the Docker Compose Build command
	manager.Build(opts)
}

func TestPS(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposePsOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command with all options set",
			opts: commands.DockerComposePsOptions{
				Path:          "/path/to/docker-compose.yml",
				Services:      true,
				Quiet:         true,
				FilterRunning: true,
				ServiceName:   "service1",
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only Services option set",
			opts: commands.DockerComposePsOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: true,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only Quiet option set",
			opts: commands.DockerComposePsOptions{
				Path:  "/path/to/docker-compose.yml",
				Quiet: true,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only FilterRunning option set",
			opts: commands.DockerComposePsOptions{
				Path:          "/path/to/docker-compose.yml",
				FilterRunning: true,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only ServiceName option set",
			opts: commands.DockerComposePsOptions{
				Path:        "/path/to/docker-compose.yml",
				ServiceName: "service1",
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name:        "it runs the correct command with no options or path set",
			opts:        commands.DockerComposePsOptions{},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposePsOptions{
				Path:          "/path/to/docker-compose.yml",
				Services:      true,
				Quiet:         true,
				FilterRunning: true,
				ServiceName:   "service1",
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "ps"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			var expectedCmd string
			if tt.opts.Path != "" {
				expectedCmd = "docker compose -f " + tt.opts.Path + " ps"
			} else {
				expectedCmd = "docker compose ps"
			}

			if tt.opts.Services {
				expectedCmd += " --services"
			}
			if tt.opts.Quiet {
				expectedCmd += " --quiet"
			}
			if tt.opts.FilterRunning {
				expectedCmd += " --filter status=running"
			}
			if tt.opts.ServiceName != "" {
				expectedCmd += " " + tt.opts.ServiceName
			}

			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposePSCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			_, err := manager.PS(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_PS() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose PS command
	opts := commands.DockerComposePsOptions{
		Path:          "/path/to/docker-compose.yml",
		Services:      true,
		Quiet:         true,
		FilterRunning: true,
		ServiceName:   "service1",
	}

	// Run the Docker Compose PS command
	manager.PS(opts)
}

func TestLogs(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposeLogsOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command with all options set",
			opts: commands.DockerComposeLogsOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
				Follow:   true,
				Tail:     10,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only Follow option set",
			opts: commands.DockerComposeLogsOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
				Follow:   true,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it runs the correct command with only Tail option set",
			opts: commands.DockerComposeLogsOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
				Tail:     10,
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposeLogsOptions{
				Path:     "/path/to/docker-compose.yml",
				Services: []string{"service1", "service2"},
				Follow:   true,
				Tail:     10,
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "logs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			expectedCmd := "docker compose -f " + tt.opts.Path + " logs"
			if tt.opts.Follow {
				expectedCmd += " --follow"
			}
			if tt.opts.Tail > 0 {
				expectedCmd += " --tail=" + strconv.Itoa(tt.opts.Tail)
			}
			expectedCmd += " " + strings.Join(tt.opts.Services, " ")

			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposeLogsCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Logs(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Logs() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Logs command
	opts := commands.DockerComposeLogsOptions{
		Path:     "/path/to/docker-compose.yml",
		Services: []string{"service1", "service2"},
		Follow:   true,
		Tail:     10,
	}

	// Run the Docker Compose Logs command
	manager.Logs(opts)
}

func TestStop(t *testing.T) {
	tests := []struct {
		name        string
		opts         DockerComposeStopOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts:  DockerComposeStopOptions{
				Path: "/path/to/docker-compose.yml",
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts:  DockerComposeStopOptions{
				Path: "/path/to/docker-compose.yml",
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "stop"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			expectedCmd := "docker compose -f " + tt.opts.Path + " stop"

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(commands.Command{Cmd: expectedCmd, GetOutput: true}).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(commands.Command{Cmd: expectedCmd, GetOutput: true}).Return("", 0, nil)
			}

			err := manager.Stop(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDown(t *testing.T) {
	tests := []struct {
		name        string
		opts        commands.DockerComposeDownOptions
		runCMDError error
		wantError   error
	}{
		{
			name: "it runs the correct command",
			opts: commands.DockerComposeDownOptions{
				Path: "/path/to/docker-compose.yml",
			},
			runCMDError: nil,
			wantError:   nil,
		},
		{
			name: "it returns an error if RunCMD fails",
			opts: commands.DockerComposeDownOptions{
				Path: "/path/to/docker-compose.yml",
			},
			runCMDError: errors.New("command failed"),
			wantError:   DockerComposeCmdError{cmd: "down"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRunner := sedge_mocks.NewMockCommandRunner(ctrl)

			manager := NewComposeManager(mockRunner)

			expectedCmd := "docker compose -f " + tt.opts.Path + " down"

			expectedCommand := commands.Command{
				Cmd:      expectedCmd,
				GetOutput: false,
			}
			mockRunner.EXPECT().BuildDockerComposeDownCMD(tt.opts).Return(expectedCommand)

			if tt.runCMDError != nil {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 1, tt.runCMDError)
			} else {
				mockRunner.EXPECT().RunCMD(expectedCommand).Return("", 0, nil)
			}

			err := manager.Down(tt.opts)

			if tt.wantError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleComposeManager_Down() {
	// Create a new CMDRunner with admin privileges
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{RunAsAdmin: true})

	// Create a new ComposeManager with the CMDRunner
	manager := NewComposeManager(cmdRunner)

	// Define the options for the Docker Compose Down command
	opts := commands.DockerComposeDownOptions{
		Path: "/path/to/docker-compose.yml",
	}

	// Run the Docker Compose Down command
	manager.Down(opts)
}
