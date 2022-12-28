package cli

import (
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RunCmd(sedgeActions actions.SedgeActions) *cobra.Command {
	// Flags
	var (
		generationPath string
		services       *[]string
	)

	cmd := &cobra.Command{
		Use:   "run [flags]",
		Short: "Run services",
		Long:  "Run all the generated services",
		Run: func(cmd *cobra.Command, args []string) {
			err := sedgeActions.SetupContainers(actions.SetupContainersOptions{
				GenerationPath: generationPath,
				Services:       *services,
			})
			if err != nil {
				log.Fatalf("error setting up service containers: %s", err.Error())
			}
			err = sedgeActions.RunContainers(actions.RunContainersOptions{
				GenerationPath: generationPath,
				Services:       *services,
			})
			if err != nil {
				log.Fatalf("error starting service containers: %s", err.Error())
			}
		},
	}

	cmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	services = cmd.Flags().StringArray("services", []string{}, "List of services to run. If this flag is not provided, all services will run.")
	return cmd
}
