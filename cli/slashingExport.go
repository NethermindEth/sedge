package cli

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/pkg/slashing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type SlashingExportFlags struct {
	validatorClient string
	network         string
	stopValidator   bool
	startValidator  bool
	generationPath  string
	out             string
}

func SlashingExportCmd(slashingManager slashing.SlashingDataManager, serviceManager services.ServiceManager) *cobra.Command {
	var flags SlashingExportFlags

	cmd := &cobra.Command{
		Use:   "slashing-export [flags]",
		Short: "Export slashing protection data",
		Long:  "Export slashing protection interchange data (EIP-3076). Validator is stopped if is currently running",
		PreRun: func(cmd *cobra.Command, args []string) {
			if flags.out == "" {
				flags.out = filepath.Join(flags.generationPath, "slashing_export.json")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			previouslyRunning, err := serviceManager.IsRunning(services.ServiceValidator)
			if err != nil {
				log.Fatal(err)
			}
			// Stop validator client
			if err := serviceManager.Stop(services.ServiceValidator); err != nil {
				log.Fatal(err)
			}
			// Export slashing data
			if err := slashingManager.Export(flags.validatorClient, flags.network, flags.generationPath, flags.out); err != nil {
				log.Fatal(err)
			}
			// Run validator again
			if (previouslyRunning && !flags.stopValidator) || flags.startValidator {
				if err := serviceManager.Start(services.ServiceValidator); err != nil {
					log.Fatal(err)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&flags.validatorClient, "validator", "v", "", "validator engine client")
	cmd.Flags().StringVarP(&flags.network, "network", "n", "", "network")
	cmd.Flags().BoolVar(&flags.startValidator, "start-validator", false, "if the validator client is currently stopped, then it is started after slashing export")
	cmd.Flags().BoolVar(&flags.stopValidator, "stop-validator", false, "if the validator client is currently running, then it is not started after slashing export")
	cmd.Flags().StringVarP(&flags.generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	cmd.Flags().StringVarP(&flags.out, "out", "o", "", `path to write slashing protection data (default "[GENERATION_PATH]/slashing_export.json")`)
	return cmd
}
