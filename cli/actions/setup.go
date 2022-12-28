package actions

import (
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	log "github.com/sirupsen/logrus"
)

type SetupContainersOptions struct {
	GenerationPath string
}

func (s *sedgeActions) SetupContainers(options SetupContainersOptions) error {
	log.Info("Setting up containers")
	buildCmd := s.commandRunner.BuildDockerComposeBuildCMD(commands.DockerComposeBuildOptions{
		Path: options.GenerationPath,
	})
	if _, err := s.commandRunner.RunCMD(buildCmd); err != nil {
		return err
	}
	pullCmd := s.commandRunner.BuildDockerComposePullCMD(commands.DockerComposePullOptions{
		Path: options.GenerationPath,
	})
	if _, err := s.commandRunner.RunCMD(pullCmd); err != nil {
		return err
	}
	createCmd := s.commandRunner.BuildDockerComposeCreateCMD(commands.DockerComposeCreateOptions{
		Path: options.GenerationPath,
	})
	if _, err := s.commandRunner.RunCMD(createCmd); err != nil {
		return err
	}
	return nil
}
