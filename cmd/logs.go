/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	tail bool
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs [flags] [services]",
	Short: "Get running container logs",
	Long: `Get running container logs using docker-compose CLI. If no services are provided, the logs of all running services will be displayed.

By default will run 'docker-compose -f <script> logs --follow <service>'`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err = utils.PreCheck(generationPath); err != nil {
			log.Fatal(err)
		}

		var rawServices string
		if rawServices, err = utils.CheckContainers(generationPath); err != nil {
			log.Fatal(err)
		}

		file := generationPath + "/" + configs.DefaultDockerComposeScriptName
		// Get logs from docker-compose script services
		services := strings.Split(rawServices, "\n")
		// Remove empty string resulting of spliting the last blank line of rawServices
		services = services[:len(services)-1]
		if len(args) > 0 {
			services = args
		}

		logsCMD := fmt.Sprintf(configs.DockerComposeLogsFollowCMD, file, strings.Join(services, " "))
		if tail {
			logsCMD = fmt.Sprintf(configs.DockerComposeLogsTailCMD, file, strings.Join(services, " "))
		}

		log.Debugf(configs.RunningCommand, logsCMD)
		if _, err := utils.RunCmd(logsCMD, false, false); err != nil {
			log.Fatalf(configs.GettingLogsError, strings.Join(services, " "), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Local flags
	logsCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose script path")

	logsCmd.Flags().BoolVarP(&tail, "tail", "t", false, "Tail the last 20 logs")
}
