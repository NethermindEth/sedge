/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"path/filepath"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/internal/utils"
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

		downCMD := commands.Runner.BuildDockerComposeDownCMD(commands.DockerComposeDownOptions{
			Path: filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		})

		log.Debugf(configs.RunningCommand, downCMD.Cmd)
		_, err := commands.Runner.RunCMD(downCMD)
		if err != nil {
			log.Fatalf(configs.CommandError, downCMD.Cmd, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Local flags
	downCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose script path")
}
