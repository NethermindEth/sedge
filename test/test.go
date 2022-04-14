package test

import "github.com/NethermindEth/1click/internal/pkg/commands"

type SimpleCMDRunner struct {
	SRunCMD  func(commands.Command) (string, error)
	SRunBash func(commands.BashScript) (string, error)
}

func (cr *SimpleCMDRunner) BuildDockerComposeUpCMD(options commands.DockerComposeUpOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeUpCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerPSCMD(options commands.DockerPSOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerPSCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposePSCMD(options commands.DockerComposePsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposePSCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeLogsCMD(options commands.DockerComposeLogsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeLogsCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerBuildCMD(options commands.DockerBuildOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerBuildCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerInspectCMD(options commands.DockerInspectOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerInspectCMD(options)
}

func (cr *SimpleCMDRunner) BuildDockerComposeDownCMD(options commands.DockerComposeDownOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeDownCMD(options)
}

func (cr *SimpleCMDRunner) RunCMD(cmd commands.Command) (string, error) {
	return cr.SRunCMD(cmd)
}

func (cr *SimpleCMDRunner) RunBash(script commands.BashScript) (string, error) {
	return cr.SRunBash(script)
}
