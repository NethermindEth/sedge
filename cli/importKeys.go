package cli

import (
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
		Use:   "import-key [flags]",
		Short: "Import validator keys",
		Long:  "Import validator keys. Validator is stopped if is currently running",
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

	cmd.Flags().StringVarP(&validatorClient, "validator", "v", "", "validator engine client")
	cmd.Flags().StringVarP(&network, "network", "n", "", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "if the validator client is currently stopped, then it is started after slashing export")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "if the validator client is currently running, then it is not started after slashing export")
	cmd.Flags().StringVar(&from, "from", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	cmd.Flags().StringVar(&customConfigPath, "custom-config", "", "File path or url to use as custom network config.")
	cmd.Flags().StringVar(&customGenesisPath, "custom-genesis", "", "File path or url to use as custom network genesis.")
	cmd.Flags().StringVar(&customDeployBlock, "custom-deploy-block", "", "Custom network deploy block.")
	return cmd
}
