package cli

import (
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/pkg/slashing"
	"github.com/spf13/cobra"
)

type SlashingExportFlags struct {
	validatorClient string
	network         string
	stopValidator   bool
	startValidator  bool
}

func SlashingExportCmd(slashingManager slashing.SlashingDataManager) *cobra.Command {
	var flags SlashingExportFlags

	cmd := &cobra.Command{
		Use:   "slashing-export [flags]",
		Short: "Export slashing protection data",
		Long:  "Export slashing protection interchange data (EIP-3076). Validator is stopped if is currently running",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			previouslyRunning, err := services.IsRunning(services.ServiceValidator)
			if err != nil {
				return err
			}
			// Stop validator client
			if err := services.Stop(services.ServiceValidator); err != nil {
				return err
			}
			// Export slashing data
			if err := slashingManager.Export(flags.validatorClient, flags.network); err != nil {
				return err
			}
			// Run validator again
			if (previouslyRunning && !flags.stopValidator) || flags.startValidator {
				if err := services.Start(services.ServiceValidator); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.validatorClient, "validator", "v", "", "Validator engine client")
	cmd.Flags().StringVarP(&flags.network, "network", "n", "", "Network")
	cmd.Flags().BoolVar(&flags.startValidator, "start-validator", false, "If the validator client is currently stopped, then it is started after slashing import")
	cmd.Flags().BoolVar(&flags.stopValidator, "stop-validator", false, "If the validator client is currently running then it is not started after slashing import")
	return cmd
}
