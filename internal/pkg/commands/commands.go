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

type CommandRunner interface {
	BuildDockerComposeUpCMD(options DockerComposeUpOptions) Command

	BuildDockerComposePullCMD(options DockerComposePullOptions) Command

	BuildDockerPSCMD(options DockerPSOptions) Command

	BuildDockerComposePSCMD(options DockerComposePsOptions) Command

	BuildDockerComposeLogsCMD(options DockerComposeLogsOptions) Command

	BuildDockerBuildCMD(options DockerBuildOptions) Command

	BuildDockerPullCMD(options DockerBuildOptions) Command

	BuildDockerInspectCMD(options DockerInspectOptions) Command

	BuildDockerComposeDownCMD(options DockerComposeDownOptions) Command

	BuildCreateFileCMD(options CreateFileOptions) Command

	BuildEchoToFileCMD(options EchoToFileOptions) Command

	RunCMD(cmd Command) (string, error)

	RunScript(script ScriptFile) (string, error)
}
