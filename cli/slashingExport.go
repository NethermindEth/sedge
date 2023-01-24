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
		Use:   "slashing-export [flags] [validator]",
		Short: "Export slashing protection data",
		Long: `
Export Slashing Protection Interchange Format (EIP-3076) data. This command assumes
that the validator client container exists, stopped or not and that its database
is already initialized. Take in mind that the validator client generates slashing
protection data after some time running, so for some clients export slashing protection
data just after start the client could produce some errors.

This command stops the validator client during the exporting process due to the
validator database being locked while it's running but leaves the validator client
in the same state in which it was found. That means if the validator is running/stopped
before the export, then the validator will be running/stopped after the command
is executed, regardless of whether the export fails or not. To force a different
behavior use --start-validator and --stop-validator flags.

The [validator] is a required argument used to specify which validator client, from
all supported by Sedge (lighthouse, lodestar, prysm or teku), is used to exporting
the Slashing Protection data. This is necessary because each client has its own way
to achieve the exportation.`,
		Example: `
sedge slashing-export --out slashing-data.json prysm
sedge slashing-export --out slashing-data.json --stop-validator lodestar 
sedge slashing-export --out slashing-data.json --start-validator lighthouse`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ErrInvalidNumberOfArguments
			}
			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if out == "" {
				out = filepath.Join(generationPath, "slashing_export.json")
			}
			if err := configs.NetworkCheck(network); err != nil {
				log.Fatal(err)
			}
			validatorClient = args[0]
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

	cmd.Flags().StringVarP(&network, "network", "n", "mainnet", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "starts the validator client after export, regardless of the state the validator was in before")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "stops the validator client after export, regardless of the state the validator was in before")
	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "path to the generation directory")
	cmd.Flags().StringVarP(&out, "out", "o", "", `path to write slashing protection data (default "[GENERATION_PATH]/slashing_export.json")`)
	return cmd
}
