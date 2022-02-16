/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	scriptPath string
	tail       bool
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs [flags] [services]",
	Short: "Get running container logs",
	Long: `Get running container logs using docker-compose CLI. If no services are provided, the logs of all running services will be displayed.

By default will run 'docker-compose -f <script> logs --follow <service>'`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check that docker and docker-compose are installed
		pending := utils.CheckDependencies([]string{"docker", "docker-compose"})
		for _, dependency := range pending {
			log.Errorf(configs.DependencyNotInstalledError, dependency)
		}
		if len(pending) > 0 {
			log.Fatal(configs.DependenciesMissingError)
		}

		// Check docker engine is on
		log.Debugf(configs.RunningCommand, configs.DockerPsCMD)
		if _, err := utils.RunCmd(configs.DockerPsCMD, true); err != nil {
			log.Fatalf(configs.DockerEngineOffError, err)
		}

		// Check if docker-compose script was generated
		file := scriptPath + "/" + configs.DefaultDockerComposeScriptName
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Errorf(configs.OpeningFileError, file, err)
			log.Fatalf(configs.DockerComposeScriptNotFoundError, scriptPath, configs.DefaultDockerComposeScriptsPath)
		}

		// Check if docker-compose script is running
		psCMD := fmt.Sprintf(configs.DockerComposePsServicesCMD, file)
		log.Debugf(configs.RunningCommand, configs.DockerPsCMD)
		rawServices, err := utils.RunCmd(psCMD, true)
		if err != nil || rawServices == "\n" {
			if rawServices == "\n" && err == nil {
				err = fmt.Errorf(configs.DockerComposePsReturnedEmptyError)
			}
			log.Fatalf(configs.ScriptIsNotRunningError, err)
		}

		// Get logs from docker-compose script services
		services := strings.Split(rawServices, "\n")
		if len(args) > 0 {
			services = args
		}

		logsCMD := fmt.Sprintf(configs.DockerComposeLogsFollowCMD, file, strings.Join(services, " "))
		if tail {
			logsCMD = fmt.Sprintf(configs.DockerComposeLogsTailCMD, file, strings.Join(services, " "))
		}

		log.Debugf(configs.RunningCommand, logsCMD)
		if _, err := utils.RunCmd(logsCMD, false); err != nil {
			log.Fatal(configs.GettingLogsError)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Local flags
	logsCmd.Flags().StringVarP(&scriptPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose script path")

	logsCmd.Flags().BoolVarP(&tail, "tail", "t", false, "Tail the last 20 logs")
}
