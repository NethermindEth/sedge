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
	services   []string
	scriptPath string
	tail       bool
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get running container logs",
	Long: `Get running container logs using docker-compose CLI.

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
		if _, err := utils.RunCmd(configs.DockerPsCMD, true); err != nil {
			log.Fatal(configs.DockerEngineOffError, err)
		}

		// Check if docker-compose script was generated
		file := scriptPath + "/" + configs.DefaultDockerComposeScriptName
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Errorf(configs.OpeningFileError, file, err)
			log.Fatal(configs.DockerComposeScriptNotFoundError, scriptPath, configs.DefaultDockerComposeScriptsPath)
		}

		// Check if docker-compose script is running
		rawServices, err := utils.RunCmd(configs.DockerComposePsServicesCMD, true, file)
		if err != nil || rawServices == "" {
			if rawServices == "" && err == nil {
				err = fmt.Errorf(configs.DockerComposePsReturnedEmptyError)
			}
			log.Fatal(configs.ScriptIsNotRunningError, err)
		}

		// TODO: Get logs from docker-compose script services
		services := strings.Split(rawServices, "\n")
		params := append([]string{file}, services...)

		logsCMD := configs.DockerComposeLogsFollowCMD
		if tail {
			logsCMD = configs.DockerComposeLogsTailCMD
		}

		if _, err := utils.RunCmd(logsCMD, false, params...); err != nil {
			log.Fatal(configs.GettingLogsError)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Local flags
	logsCmd.Flags().StringArrayVarP(&services, "services", "s", []string{"execution", "consensus", "validator"}, "List of services to get the logs from")

	logsCmd.Flags().StringVarP(&scriptPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose script path")

	logsCmd.Flags().BoolVarP(&tail, "tail", "t", false, "Tail the last 20 logs")
}
