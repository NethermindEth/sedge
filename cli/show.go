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
package cli

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	log "github.com/sirupsen/logrus"
)

func ShowCmd(cmdRunner commands.CommandRunner, sedgeActions actions.SedgeActions, depsMgr dependencies.DependenciesManager) *cobra.Command {
	// Flags
	var generationPath string
	// Build command
	cmd := &cobra.Command{
		Use:   "show [flags]",
		Short: "Show useful information about sedge running containers",
		Long:  `Show useful information about sedge running containers`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := checkDependencies(depsMgr, true, dependencies.Docker); err != nil {
				log.Error("Failed to check dependencies. Run 'sedge deps check' to check dependencies")
				return err
			}
			return sedgeActions.ValidateDockerComposeFile(filepath.Join(generationPath, configs.DefaultDockerComposeScriptName))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := utils.CheckContainers(cmdRunner, generationPath); err != nil {
				return err
			}

			data, err := sedgeActions.GetContainersData(actions.GetContainersDataOptions{
				DockerComposePath: filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
			})
			if err != nil {
				log.Errorf("Failed to get sedge containers data: %v", err)
				return err
			}

			// Remove initial slash from container name
			for i := range data.Containers {
				if len(data.Containers[i].Name) > 0 && data.Containers[i].Name[0] == '/' {
					data.Containers[i].Name = data.Containers[i].Name[1:]
				}
			}

			output, err := yaml.Marshal(data)
			if err != nil {
				log.Errorf("Failed to show sedge containers data: %v", err)
				return err
			}

			log.Info(string(output))
			return err
		},
	}
	// Bind flags
	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "generation path for sedge data")
	return cmd
}
