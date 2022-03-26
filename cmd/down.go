/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down [flags]",
	Short: "Shutdown 1click running containers",
	Long:  `Shutdown 1click running containers using docker-compose CLI. Shortcut for 'docker-compose -f <script> down'`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.PreCheck(generationPath); err != nil {
			log.Fatal(err)
		}

		if _, err := utils.CheckContainers(generationPath); err != nil {
			log.Fatal(err)
		}

		downCMD := fmt.Sprintf(configs.DockerComposeDownCMD, generationPath+"/"+configs.DefaultDockerComposeScriptName)

		log.Debugf(configs.RunningCommand, downCMD)
		if _, err := utils.RunCmd(downCMD, false, false); err != nil {
			log.Fatalf(configs.CommandError, downCMD, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Local flags
	downCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose script path")
}
