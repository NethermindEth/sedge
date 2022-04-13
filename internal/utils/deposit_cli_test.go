package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/commands"
)

type generateValidatorKeyTestCase struct {
	runner   commands.CommandRunner
	existing bool
	network  string
	path     string
	isErr    bool
}

func TestGenerateValidatorKey(t *testing.T) {
	tcs := []generateValidatorKeyTestCase{
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "", fmt.Errorf("unexpected error")
					}
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker build") {
						return "", fmt.Errorf("error")
					}
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
		{
			runner: &generateValidatorKeyCMDRunner{
				runCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker build") {
						return "", fmt.Errorf("error")
					}
					return "", nil
				},
				runBash: func(bs commands.BashScript) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
	}

	for _, tc := range tcs {
		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})
		descr := fmt.Sprintf("GenerateValidatorKey(%t, %s, %s)", tc.existing, tc.network, tc.path)

		err := GenerateValidatorKey(tc.existing, tc.network, tc.path)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

type generateValidatorKeyCMDRunner struct {
	runCMD  func(commands.Command) (string, error)
	runBash func(commands.BashScript) (string, error)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerComposeUpCMD(options commands.DockerComposeUpOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeUpCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerPSCMD(options commands.DockerPSOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerPSCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerComposePSCMD(options commands.DockerComposePsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposePSCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerComposeLogsCMD(options commands.DockerComposeLogsOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeLogsCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerBuildCMD(options commands.DockerBuildOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerBuildCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerInspectCMD(options commands.DockerInspectOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerInspectCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) BuildDockerComposeDownCMD(options commands.DockerComposeDownOptions) commands.Command {
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	return r.BuildDockerComposeDownCMD(options)
}

func (cr *generateValidatorKeyCMDRunner) RunCMD(cmd commands.Command) (string, error) {
	return cr.runCMD(cmd)
}

func (cr *generateValidatorKeyCMDRunner) RunBash(script commands.BashScript) (string, error) {
	return cr.runBash(script)
}
