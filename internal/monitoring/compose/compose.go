package compose

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
)

// DockerComposeCmdError represents an error that occurs when running a Docker Compose command.
type DockerComposeCmdError struct {
	cmd string
}

// Error returns a string representation of the DockerComposeCmdError.
func (e DockerComposeCmdError) Error() string {
	return fmt.Sprintf("Docker Compose Manager running 'docker compose %s'", e.cmd)
}

// ComposeManager manages Docker Compose operations.
type ComposeManager struct {
	cmdRunner commands.CommandRunner
}

// NewComposeManager creates a new instance of ComposeManager.
func NewComposeManager(runner commands.CommandRunner) *ComposeManager {
	return &ComposeManager{
		cmdRunner: runner,
	}
}

// Up runs the Docker Compose 'up' command for the specified options.
func (cm *ComposeManager) Up(opts commands.DockerComposeUpOptions) error {
	upCmd := cm.cmdRunner.BuildDockerComposeUpCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(upCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "up"}, err, out)
	}
	return nil
}

// Pull runs the Docker Compose 'pull' command for the specified options.
func (cm *ComposeManager) Pull(opts commands.DockerComposePullOptions) error {
	pullCmd := cm.cmdRunner.BuildDockerComposePullCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(pullCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "pull"}, err, out)
	}
	return nil
}

// Create runs the Docker Compose 'create' command for the specified options.
func (cm *ComposeManager) Create(opts commands.DockerComposeCreateOptions) error {
	createCmd := cm.cmdRunner.BuildDockerComposeCreateCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(createCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "create"}, err, out)
	}
	return nil
}

// Build runs the Docker Compose 'build' command for the specified options.
func (cm *ComposeManager) Build(opts commands.DockerComposeBuildOptions) error {
	buildCmd := cm.cmdRunner.BuildDockerComposeBuildCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(buildCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "build"}, err, out)
	}
	return nil
}

// PS runs the Docker Compose 'ps' command for the specified options and returns
// the list of services.
func (c *ComposeManager) PS(opts commands.DockerComposePsOptions) ([]ComposeService, error) {
	psCmd := c.cmdRunner.BuildDockerComposePSCMD(opts)

	out, exitCode, err := c.cmdRunner.RunCMD(psCmd)
	if err != nil || exitCode != 0 {
		return nil, fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "ps"}, err, out)
	}
	outList := make([]ComposeService, 0)
	if len(out) == 0 {
		return outList, nil
	}
	// Following `if` cases are necessary to handle the different output formats
	// of the `docker compose ps` command. Some times it returns a list of
	// services, other times it returns a single service for the edge case of
	// only one service being present. Depending on the docker compose version.
	if out[0] == '[' {
		// Multiple services
		err = json.Unmarshal([]byte(out), &outList)
		if err != nil {
			return outList, fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "ps"}, err, out)
		}
	} else if out[0] == '{' {
		// Single service
		var s ComposeService
		err = json.Unmarshal([]byte(out), &s)
		if err != nil {
			return outList, fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "ps"}, err, out)
		}
		outList = append(outList, s)
	} else if strings.HasPrefix(out, "null") {
		// No services
		return outList, nil
	} else {
		// Unexpected output
		return outList, fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "ps"}, "unknown output format", out)
	}
	return outList, nil
}

// Logs runs the Docker Compose 'logs' command for the specified options.
func (cm *ComposeManager) Logs(opts commands.DockerComposeLogsOptions) error {
	logsCmd := cm.cmdRunner.BuildDockerComposeLogsCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(logsCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "logs"}, err, out)
	}
	return nil
}

// Stop runs the Docker Compose 'stop' command for the specified options.
func (cm *ComposeManager) Stop(opts DockerComposeStopOptions) error {
	stopCmd := fmt.Sprintf("docker compose -f %s stop", opts.Path)

	if out, exitCode, err := cm.cmdRunner.RunCMD(commands.Command{Cmd: stopCmd, GetOutput: true}); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "stop"}, err, out)
	}
	return nil
}

// Down runs the Docker Compose 'down' command for the specified options.
func (cm *ComposeManager) Down(opts commands.DockerComposeDownOptions) error {
	downCmd := cm.cmdRunner.BuildDockerComposeDownCMD(opts)

	if out, exitCode, err := cm.cmdRunner.RunCMD(downCmd); err != nil || exitCode != 0 {
		return fmt.Errorf("%w: %s. Output: %s", DockerComposeCmdError{cmd: "down"}, err, out)
	}
	return nil
}
