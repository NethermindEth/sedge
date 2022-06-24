//go:build linux || darwin
// +build linux darwin

package commands

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type UnixCMDRunner struct {
	RunWithSudo bool
}

func NewCMDRunner(options CMDRunnerOptions) CommandRunner {
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
	} else {
		log.Debug(`Command "docker ps" built without the "--all" flag.`)
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposePSCMD(options DockerComposePsOptions) Command {
	flags := ""
	name := ""
	if options.Services {
		log.Debug(`Command "docker-compose ps" built with "--service" flag.`)
		flags += " --services"
	} else if options.Quiet {
		log.Debug(`Command "docker-compose ps" built with "--quiet" flag.`)
		flags += " --quiet"
	}

	if options.FilterRunning {
		flags += " --filter status=running"
	}

	if options.ServiceName != "" {
		name += " " + options.ServiceName
	}
	command := fmt.Sprintf("docker-compose -f %s ps%s%s", options.Path, flags, name)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeLogsCMD(options DockerComposeLogsOptions) Command {
	command := fmt.Sprintf("docker-compose -f %s logs", options.Path)
	servs := strings.Join(options.Services, " ")
	if options.Follow {
		log.Debug(`Command "docker-compose logs" built with "--follow" flag.`)
		command += fmt.Sprintf(" --follow %s", servs)
	} else if options.Tail > 0 {
		log.Debugf(`Command "docker-compose logs" built with "--tail=%d" flag.`, options.Tail)
		command += fmt.Sprintf(" --tail=%d %s", options.Tail, servs)
	} else {
		log.Warn(`Command "docker-compose logs" built without "--follow" or "--tail" flags. Add follow argument or make tail argument greater than 0.`)
		command += " " + servs
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerBuildCMD(options DockerBuildOptions) Command {
	command := fmt.Sprintf("docker build %s", options.Path)
	if len(options.Tag) > 0 {
		log.Debug(`Command "docker build" built with "-t" flag.`)
		command += " -t " + options.Tag
	} else {
		log.Debug(`Command "docker build" built withot "-t" flag.`)
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerPullCMD(options DockerBuildOptions) Command {
	command := fmt.Sprintf("docker pull %s", options.Tag)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerInspectCMD(options DockerInspectOptions) Command {
	flags := ""
	if options.Format != "" {
		flags += " --format " + options.Format
	}
	command := fmt.Sprintf("docker inspect%s %s", flags, options.Name)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeDownCMD(options DockerComposeDownOptions) Command {
	command := fmt.Sprintf("docker-compose -f %s down", options.Path)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) RunCMD(cmd Command) (string, error) {
	if cr.RunWithSudo {
		log.Debug(`Running command with sudo.`)
		cmd.Cmd = fmt.Sprintf("sudo %s", cmd.Cmd)
	} else {
		log.Debug(`Running command without sudo.`)
	}
	return runCmd(cmd.Cmd, cmd.GetOutput, cmd.RunInPty)
}

func (cr *UnixCMDRunner) RunBash(script BashScript) (string, error) {
	return executeBashScript(script)
}
