package cli

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ImportKeysCmd(sedgeActions actions.SedgeActions) *cobra.Command {
	// Flags
	var (
		validatorClient   string
		network           string
		stopValidator     bool
		startValidator    bool
		from              string
		customConfigPath  string
		customGenesisPath string
		customDeployBlock string
	)

	cmd := &cobra.Command{
		Use:   "import-key [flags] [validator]",
		Short: "Import validator keys",
		Long: `
Import validator client keys, use the 'from' flag to specify the keys location,
and make sure that follows the EIP-2335: BLS12-381 Keystore standard. This command
assumes that the validator client container exists, stopped or not.

This command stops the validator client during the importing process due to the
validator database being locked while it's running but leaves the validator client
in the same state in which it was found. That means if the validator is running/stopped
before the import, then the validator will be running/stopped after the command
is executed, regardless of whether the export fails or not. To force a different
behavior use --start-validator and --stop-validator flags.

The [validator] is a required argument used to specify which validator client, from
all supported by Sedge (lighthouse, lodestar, prysm or teku), is used to import the
validator keys. This is necessary because each client has its own way to achieve
the importation.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return ErrInvalidNumberOfArguments
			}
			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			validatorClient = args[0]
		},
		Run: func(cmd *cobra.Command, args []string) {
			options := actions.ImportValidatorKeysOptions{
				ValidatorClient: validatorClient,
				Network:         network,
				StopValidator:   stopValidator,
				StartValidator:  startValidator,
				From:            from,
			}
			if customConfigPath != "" {
				options.CustomConfig.NetworkConfigPath = &customConfigPath
			}
			if customGenesisPath != "" {
				options.CustomConfig.GenesisPath = &customGenesisPath
			}
			if customDeployBlock != "" {
				options.CustomConfig.DeployBlockPath = &customDeployBlock
			}
			err := sedgeActions.ImportValidatorKeys(options)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&network, "network", "n", "", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "starts the validator client after import, regardless of the state the validator was in before")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "stops the validator client after import, regardless of the state the validator was in before")
	cmd.Flags().StringVar(&from, "from", filepath.Join(configs.DefaultAbsSedgeDataPath, "keystore"), "path to the validator keys, must follow the EIP-2335: BLS12-381 Keystore standard")
	cmd.Flags().StringVar(&customConfigPath, "custom-config", "", "file path or url to use as custom network config.")
	cmd.Flags().StringVar(&customGenesisPath, "custom-genesis", "", "file path or url to use as custom network genesis.")
	cmd.Flags().StringVar(&customDeployBlock, "custom-deploy-block", "", "custom network deploy block.")
	return cmd
}
