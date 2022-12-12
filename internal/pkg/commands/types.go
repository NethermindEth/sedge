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

import "text/template"

// Command : Represents a command
type Command struct {
	// Cmd : command string
	Cmd string
	// GetOutput : get output of command
	GetOutput bool
	// RunInPty : run command in a pty
	RunInPty bool
}

// BashScript : Represents a script to be executed
type BashScript struct {
	// Tmp : script template
	Tmp *template.Template
	// getOutput : get output of the script
	GetOutput bool
	// Data: template data object
	Data interface{}
}

// CMDRunnerOptions : Options for configuring a commands runner
type CMDRunnerOptions struct {
	// RunAsAdmin : True to run commands as admin
	RunAsAdmin bool
}

// DockerComposeUpOptions : Represent docker compose up command options
type DockerComposeUpOptions struct {
	// Path : path to docker-compose.yaml
	Path string
	// Services : services names
	Services []string
}

// DockerComposePullOptions represents 'docker compose pull' command options
type DockerComposePullOptions struct {
	// Path to the docker-compose.yaml
	Path string
	// Services names
	Services []string
}

type DockerComposeCreateOptions struct {
	Path string
}

// DockerPSOptions : Represent docker ps command options
type DockerPSOptions struct {
	// All : use with --all
	All bool
}

// DockerComposePsOptions : Represents docker compose ps command options
type DockerComposePsOptions struct {
	// Path : path to docker-compose.yaml
	Path string
	// Services : use with --services to display services
	Services bool
	// Quiet : use with --quietto display only IDs
	Quiet bool
	// ServiceName: Service argument
	ServiceName string
	// FilterRunning : use with --filter status=running
	FilterRunning bool
}

// DockerComposeLogsOptions : Represents docker compose log command options
type DockerComposeLogsOptions struct {
	// Path : path to docker-compose.yaml
	Path string
	// Services : services names
	Services []string
	// Follow : use with --follow
	Follow bool
	// Tail : if greater than 0 and Follow is False used for --tail
	Tail int
}

// DockerBuildOptions : Represents docker build command options
type DockerBuildOptions struct {
	// Path : path to dockerfile
	Path string
	// Tag : docker build --tag tag
	Tag string
}

// DockerInspectOptions : Represents docker inspect command options
type DockerInspectOptions struct {
	// Name : docker object name
	Name string
	// Format : Go template for --format flag
	Format string
}

// DockerComposeDownOptions : Represents docker compose down command options
type DockerComposeDownOptions struct {
	// Path : path to docker-compose.yaml
	Path string
}

// CreateFileOptions : Represents create file command options
type CreateFileOptions struct {
	// FileName : path to file
	FileName string
}

// EchoToFileOptions : Represents echo to file command options
type EchoToFileOptions struct {
	// FileName : path to file
	FileName string
	// Content : content to be written to file
	Content string
}
