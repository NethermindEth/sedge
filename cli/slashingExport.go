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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func SlashingExportCmd(sedgeActions actions.SedgeActions) *cobra.Command {
	// Flags
	var (
		validatorClient string
		network         string
		stopValidator   bool
		startValidator  bool
		generationPath  string
		out             string
	)

	cmd := &cobra.Command{
		Use:   "slashing-export [flags]",
		Short: "Export slashing protection data",
		Long:  "Export slashing protection interchange data (EIP-3076). Validator is stopped if is currently running",
		PreRun: func(cmd *cobra.Command, args []string) {
			if out == "" {
				out = filepath.Join(generationPath, "slashing_export.json")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := sedgeActions.ExportSlashingInterchangeData(actions.SlashingExportOptions{
				ValidatorClient: validatorClient,
				Network:         network,
				StopValidator:   stopValidator,
				StartValidator:  startValidator,
				GenerationPath:  generationPath,
				Out:             out,
			})
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&validatorClient, "validator", "v", "", "validator engine client")
	cmd.Flags().StringVarP(&network, "network", "n", "", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "if the validator client is currently stopped, then it is started after slashing export")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "if the validator client is currently running, then it is not started after slashing export")
	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	cmd.Flags().StringVarP(&out, "out", "o", "", `path to write slashing protection data (default "[GENERATION_PATH]/slashing_export.json")`)
	return cmd
}
