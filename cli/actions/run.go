package actions

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
)

type RunContainersOptions struct {
	GenerationPath string
	Services       []string
}

func (s *sedgeActions) RunContainers(options RunContainersOptions) error {
	upCmd := s.commandRunner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	_, err := s.commandRunner.RunCMD(upCmd)
	return err
}
