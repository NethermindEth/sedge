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
	"path/filepath"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/spf13/cobra"
)

func RunCmd(sedgeActions actions.SedgeActions, depsMgr dependencies.DependenciesManager) *cobra.Command {
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
			if err := checkDependencies(depsMgr, true, dependencies.Docker); err != nil {
				return err
			}
			return sedgeActions.ValidateDockerComposeFile(filepath.Join(generationPath, configs.DefaultDockerComposeScriptName), services...)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
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
