package actions

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	log "github.com/sirupsen/logrus"
)

type SetupContainersOptions struct {
	GenerationPath string
	Services       []string
}

func (s *sedgeActions) SetupContainers(options SetupContainersOptions) error {
	log.Info("Setting up containers")
	buildCmd := s.commandRunner.BuildDockerComposeBuildCMD(commands.DockerComposeBuildOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	log.Debugf(configs.RunningCommand, buildCmd.Cmd)
	if _, err := s.commandRunner.RunCMD(buildCmd); err != nil {
		return err
	}
	pullCmd := s.commandRunner.BuildDockerComposePullCMD(commands.DockerComposePullOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	log.Debugf(configs.RunningCommand, pullCmd.Cmd)
	if _, err := s.commandRunner.RunCMD(pullCmd); err != nil {
		return err
	}
	createCmd := s.commandRunner.BuildDockerComposeCreateCMD(commands.DockerComposeCreateOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	log.Debugf(configs.RunningCommand, pullCmd.Cmd)
	if _, err := s.commandRunner.RunCMD(createCmd); err != nil {
		return err
	}
	return nil
}
