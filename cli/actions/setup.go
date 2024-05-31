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
	SkipPull       bool
}

func (s *sedgeActions) SetupContainers(options SetupContainersOptions) error {
	log.Info("Setting up containers")
	buildCmd := s.commandRunner.BuildDockerComposeBuildCMD(commands.DockerComposeBuildOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	log.Infof(configs.RunningCommand, buildCmd.Cmd)
	if _, _, err := s.commandRunner.RunCMD(buildCmd); err != nil {
		return err
	}
	if !options.SkipPull {
		pullCmd := s.commandRunner.BuildDockerComposePullCMD(commands.DockerComposePullOptions{
			Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
			Services: options.Services,
		})
		log.Infof(configs.RunningCommand, pullCmd.Cmd)
		if _, _, err := s.commandRunner.RunCMD(pullCmd); err != nil {
			return err
		}
	} else {
		log.Warn("Skipping 'docker compose pull' step")
	}
	createCmd := s.commandRunner.BuildDockerComposeCreateCMD(commands.DockerComposeCreateOptions{
		Path:     filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName),
		Services: options.Services,
	})
	log.Infof(configs.RunningCommand, createCmd.Cmd)
	if _, _, err := s.commandRunner.RunCMD(createCmd); err != nil {
		return err
	}
	return nil
}
