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
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/spf13/cobra"
)

func RunCmd(sedgeActions actions.SedgeActions) *cobra.Command {
	// Flags
	var (
		generationPath string
		services       []string
	)

	cmd := &cobra.Command{
		Use:   "run [flags]",
		Short: "Run services",
		Long:  "Run all the generated services",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return fmt.Errorf(configs.ErrCMDArgsNotSupported, "run")
			}
			var found bool

			if err := filepath.Walk(generationPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return fmt.Errorf(configs.ComposeNotFoundErr, path)
				}
				if info.IsDir() {
					return nil
				}
				if info.Name() == configs.DefaultDockerComposeScriptName {
					found = true
					// Check if provided docker-compose script is a valid docker-compose script
					if err := utils.ValidateCompose(path); err != nil {
						return fmt.Errorf(configs.InvalidComposeErr, err)
					}

					// Parse provided docker-compose script
					// cmp, err := parseCompose(path)
					// if err != nil {
					// 	return err
					// }

					// Check if provided services are valid
					if len(services) > 0 {
						actualServices, err := loadServices(path)
						if err != nil {
							return err
						}
						for _, service := range services {
							if !utils.Contains(actualServices, service) {
								return fmt.Errorf(configs.InvalidService, service)
							}
						}
					}
				}
				return nil
			}); err != nil {
				return fmt.Errorf(configs.GenPathErr, err)
			}

			if !found {
				return fmt.Errorf(configs.ComposeNotFoundErr, generationPath)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sedgeActions.ManageDependencies(actions.ManageDependenciesOptions{
				Dependencies: []string{"docker"},
				Install:      true,
			}); err != nil {
				return fmt.Errorf(configs.DependencyErr, err)
			}
			err := sedgeActions.SetupContainers(actions.SetupContainersOptions{
				GenerationPath: generationPath,
				Services:       services,
			})
			if err != nil {
				return fmt.Errorf(configs.SetupContainersErr, err)
			}
			err = sedgeActions.RunContainers(actions.RunContainersOptions{
				GenerationPath: generationPath,
				Services:       services,
			})
			if err != nil {
				return fmt.Errorf(configs.StartingContainersErr, err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "generation path for sedge data")
	cmd.Flags().StringSliceVar(&services, "services", []string{}, "List of services to run. If this flag is not provided, all services will run.")
	return cmd
}

// loadServices get the services from a Sedge valid docker-compose script
func loadServices(path string) ([]string, error) {
	composeBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(configs.ReadingYmlErr, err)
	}
	var composeData map[string]interface{}
	// It was already parsed by parseCompose
	_ = yaml.Unmarshal(composeBytes, &composeData)

	for topLevelK, field := range composeData {
		if topLevelK == "services" {
			services, ok := field.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf(configs.ParsingServicesErr)
			}
			var servicesNames []string
			for service := range services {
				servicesNames = append(servicesNames, service)
			}
			return servicesNames, nil
		}
	}

	return nil, fmt.Errorf(configs.ServicesNotFoundErr)
}
