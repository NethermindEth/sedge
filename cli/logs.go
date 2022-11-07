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
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var tail int

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs [flags] [services]",
	Short: "Get running container logs",
	Long: `Get running container logs using docker-compose CLI. If no services are provided, the logs of all running services will be displayed.

By default will run 'docker compose -f <script> logs --follow <service>'`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err = utils.PreCheck(generationPath); err != nil {
			log.Fatal(err)
		}

		var rawServices string
		if rawServices, err = utils.CheckContainers(generationPath); err != nil {
			log.Fatal(err)
		}

		file := filepath.Join(generationPath, configs.DefaultDockerComposeScriptName)
		// Get logs from docker compose script services
		services := strings.Split(rawServices, "\n")
		// Remove empty string resulting of spliting the last blank line of rawServices
		services = services[:len(services)-1]
		if len(args) > 0 {
			services = args
		}

		logsCMD := commands.Runner.BuildDockerComposeLogsCMD(commands.DockerComposeLogsOptions{
			Path:     file,
			Services: services,
			Follow:   tail == 0,
			Tail:     tail,
		})

		log.Debugf(configs.RunningCommand, logsCMD.Cmd)
		if _, err := commands.Runner.RunCMD(logsCMD); err != nil {
			log.Fatalf(configs.GettingLogsError, strings.Join(services, " "), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Local flags
	logsCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultSedgeDataPath, "docker-compose script path")

	logsCmd.Flags().IntVarP(&tail, "tail", "t", 0, "Tail the number of desired logs. If not set, or set to 0, logs are followed.")
}
