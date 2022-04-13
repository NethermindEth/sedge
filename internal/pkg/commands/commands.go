package commands

type CommandRunner interface {
	BuildDockerComposeUpCMD(options DockerComposeUpOptions) Command

	BuildDockerPSCMD(options DockerPSOptions) Command

	BuildDockerComposePSCMD(options DockerComposePsOptions) Command

	BuildDockerComposeLogsCMD(options DockerComposeLogsOptions) Command

	BuildDockerBuildCMD(options DockerBuildOptions) Command

	BuildDockerInspectCMD(options DockerInspectOptions) Command

	BuildDockerComposeDownCMD(options DockerComposeDownOptions) Command

	RunCMD(cmd Command) (string, error)

	RunBash(script BashScript) (string, error)
}

var Runner CommandRunner

func init() {
	Runner = newCMDRunner(CMDRunnerOptions{
		RunAsAdmin: true,
	})
}
