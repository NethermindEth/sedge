//go:build linux || darwin
// +build linux darwin

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
	command := fmt.Sprintf("docker compose -f %s up -d", options.Path)
	if len(options.Services) > 0 {
		command += " " + strings.Join(options.Services, " ")
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposePullCMD(options DockerComposePullOptions) Command {
	command := fmt.Sprintf("docker compose -f %s pull", options.Path)
	if len(options.Services) > 0 {
		command += " " + strings.Join(options.Services, " ")
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeCreateCMD(options DockerComposeCreateOptions) Command {
	command := fmt.Sprintf("docker compose -f %s create", options.Path)
	if len(options.Services) > 0 {
		command += " " + strings.Join(options.Services, " ")
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeBuildCMD(options DockerComposeBuildOptions) Command {
	command := fmt.Sprintf("docker compose -f %s build", options.Path)
	if len(options.Services) > 0 {
		command += " " + strings.Join(options.Services, " ")
	}
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
		log.Debug(`Command "docker compose ps" built with "--service" flag.`)
		flags += " --services"
	} else if options.Quiet {
		log.Debug(`Command "docker compose ps" built with "--quiet" flag.`)
		flags += " --quiet"
	}

	if options.FilterRunning {
		flags += " --filter status=running"
	}

	if options.ServiceName != "" {
		name += " " + options.ServiceName
	}

	var command string
	if options.Path != "" {
		command = fmt.Sprintf("docker compose -f %s ps%s%s", options.Path, flags, name)
	} else {
		command = fmt.Sprintf("docker compose ps%s%s", flags, name)
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerComposeLogsCMD(options DockerComposeLogsOptions) Command {
	command := fmt.Sprintf("docker compose -f %s logs", options.Path)
	servs := strings.Join(options.Services, " ")
	if options.Follow {
		log.Debug(`Command "docker compose logs" built with "--follow" flag.`)
		command += fmt.Sprintf(" --follow %s", servs)
	} else if options.Tail > 0 {
		log.Debugf(`Command "docker compose logs" built with "--tail=%d" flag.`, options.Tail)
		command += fmt.Sprintf(" --tail=%d %s", options.Tail, servs)
	} else {
		log.Warn(`Command "docker compose logs" built without "--follow" or "--tail" flags. Add follow argument or make tail argument greater than 0.`)
		command += " " + servs
	}
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildDockerBuildCMD(options DockerBuildOptions) Command {
	command := "docker build --quiet"
	if len(options.Tag) > 0 {
		log.Debug(`Command "docker build" built with "-t" flag.`)
		command += " -t " + options.Tag
	} else {
		log.Debug(`Command "docker build" built without "-t" flag.`)
	}
	if len(options.Args) > 0 {
		for key, value := range options.Args {
			command += fmt.Sprintf(" --build-arg %s=%s", key, value)
		}
	}
	command += " " + options.Path
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
	command := fmt.Sprintf("docker compose -f %s down", options.Path)
	return Command{Cmd: command}
}

func (cr *UnixCMDRunner) BuildCreateFileCMD(options CreateFileOptions) Command {
	return Command{Cmd: fmt.Sprintf("touch %s", options.FileName)}
}

func (cr *UnixCMDRunner) BuildEchoToFileCMD(options EchoToFileOptions) Command {
	return Command{Cmd: fmt.Sprintf("echo %s > %s", options.Content, options.FileName)}
}

func (cr *UnixCMDRunner) BuildOpenTextEditor(options OpenTextEditorOptions) Command {
	return Command{Cmd: fmt.Sprintf("less %s", options.FilePath)}
}

func (cr *UnixCMDRunner) RunCMD(cmd Command) (string, error) {
	if cr.RunWithSudo && !cmd.ForceNoSudo {
		log.Debug(`Running command with sudo.`)
		cmd.Cmd = fmt.Sprintf("sudo %s", cmd.Cmd)
	} else {
		log.Debug(`Running command without sudo.`)
	}
	return runCmd(cmd.Cmd, cmd.GetOutput)
}

func (cr *UnixCMDRunner) RunScript(script ScriptFile) (string, error) {
	return executeBashScript(script)
}
