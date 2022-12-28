package actions

import "github.com/NethermindEth/sedge/internal/pkg/commands"

type RunContainersOptions struct {
	GenerationPath string
	Services       []string
}

func (s *sedgeActions) RunContainers(options RunContainersOptions) error {
	upCmd := s.commandRunner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     options.GenerationPath,
		Services: options.Services,
	})
	_, err := s.commandRunner.RunCMD(upCmd)
	return err
}
