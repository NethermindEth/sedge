package cli

import (
	"errors"
	"path"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidNumberOfArguments = errors.New("invalid number of arguments")
	ErrInvalidValidatorArgument = errors.New("invalid validator argument")
)

func SlashingImportCmd(sedgeActions actions.SedgeActions) *cobra.Command {
	// Flags
	var (
		validatorClient string
		network         string
		stopValidator   bool
		startValidator  bool
		generationPath  string
		from            string
	)

	cmd := &cobra.Command{
		Use:   "slashing-import [validator]",
		Short: "Import slashing protection data",
		Long: `
Import Slashing Protection Interchange Format (EIP-3076) data. This command assumes
that the validator client container exists, stopped or not and that its database
is already initialized. The validator database is initialized if the validator is
running or has already run but is stopped, and also after importing the validator keys.

This command stops the validator client during the importing process due to the
validator database being locked while it's running but leaves the validator client
in the same state in which it was found. That means if the validator is running/stopped
before the import, then the validator will be running/stopped after the command
is executed, regardless of whether the export fails or not. To force a different
behavior use --start-validator and --stop-validator flags.`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if from == "" {
				from = path.Join(configs.DefaultDockerComposeScriptsPath, "slashing-export.json")
			}
			if err := configs.NetworkCheck(network); err != nil {
				log.Fatal(err)
			}
			validatorClient = args[0]
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := sedgeActions.ImportSlashingInterchangeData(actions.SlashingImportOptions{
				ValidatorClient: validatorClient,
				Network:         network,
				StopValidator:   stopValidator,
				StartValidator:  startValidator,
				GenerationPath:  generationPath,
				From:            from,
			})
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&network, "network", "n", "mainnet", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "starts the validator client after import, regardless of the state the validator was in before")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "stops the validator client after import, regardless of the state the validator was in before")
	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "path to the generation directory")
	cmd.Flags().StringVarP(&from, "from", "f", "", "path to the JSON file in the EIP-3076 format with the slashing protection data to import (default: <generation-dir>/slashing_export.json)")
	return cmd
}
