/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down [flags]",
	Short: "Shutdown sedge running containers",
	Long:  `Shutdown sedge running containers using docker compose CLI. Shortcut for 'docker compose -f <script> down'`,
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
		if _, err := commands.Runner.RunCMD(downCMD); err != nil {
			log.Fatalf(configs.CommandError, downCMD.Cmd, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Local flags
	downCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultSedgeDataPath, "docker-compose script path")
}
