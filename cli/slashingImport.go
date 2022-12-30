package cli

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
		Use:   "slashing-import [flags]",
		Short: "Import slashing protection data",
		Long:  "Import slashing protection interchange data (EIP-3076). Validator is stopped if is currently running",
		PreRun: func(cmd *cobra.Command, args []string) {
			if from == "" {
				from = filepath.Join(generationPath, "slashing_export.json")
			}
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

	cmd.Flags().StringVarP(&validatorClient, "validator", "v", "", "validator engine client")
	cmd.Flags().StringVarP(&network, "network", "n", "", "network")
	cmd.Flags().BoolVar(&startValidator, "start-validator", false, "if the validator client is currently stopped, then it is started after slashing import")
	cmd.Flags().BoolVar(&stopValidator, "stop-validator", false, "if the validator client is currently running, then it is not started after slashing import")
	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	cmd.Flags().StringVarP(&from, "from", "o", "", "path to the slashing interchange data file to import")
	return cmd
}
