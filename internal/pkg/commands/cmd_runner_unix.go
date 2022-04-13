//go:build !windows
// +build !windows

package commands

import (
	"fmt"
	"strings"
)

type UnixCMDRunner struct {
	RunWithSudo bool
}

func newCMDRunner(options CMDRunnerOptions) CommandRunner {
	return &UnixCMDRunner{
		RunWithSudo: options.RunAsAdmin,
	}
}

func (cr *UnixCMDRunner) BuildDockerComposeUpCMD(options DockerComposeUpOptions) Command {
	servs := strings.Join(options.Services, " ")
	command := fmt.Sprintf("docker-compose -f %s up -d %s", options.Path, servs)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerPSCMD(options DockerPSOptions) Command {
	command := "docker ps"
	if options.All {
		command += " -a"
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposePSCMD(options DockerComposePsOptions) Command {
	servs := ""
	if options.Services {
		servs = " --services"
	}
	command := fmt.Sprintf("docker-compose -f %s ps%s --filter status=running", options.Path, servs)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeLogsCMD(options DockerComposeLogsOptions) Command {
	command := fmt.Sprintf("docker-compose -f %s logs ", options.Path)
	servs := strings.Join(options.Services, " ")
	if options.Follow {
		command += fmt.Sprintf("--follow %s", servs)
	} else if options.Tail > 0 {
		command += fmt.Sprintf("--tail=%d %s", options.Tail, servs)
	} else {
		command += servs
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerBuildCMD(options DockerBuildOptions) Command {
	command := fmt.Sprintf("docker build %s ", options.Path)
	if len(options.Tag) > 0 {
		command += "-t " + options.Tag
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerInspectCMD(options DockerInspectOptions) Command {
	command := "docker inspect " + options.Name
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeDownCMD(options DockerComposeDownOptions) Command {
	command := fmt.Sprintf("docker-compose -f %s down", options.Path)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) RunCMD(cmd Command) (string, error) {
	if cr.RunWithSudo {
		cmd.Cmd = fmt.Sprintf("sudo %s", cmd.Cmd)
	}
	return runCmd(cmd.Cmd, cmd.GetOutput, cmd.RunInPty)
}

func (cr *UnixCMDRunner) RunBash(script BashScript) (string, error) {
	return executeBashScript(script)
}
