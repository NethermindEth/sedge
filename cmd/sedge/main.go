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
package main

import (
	"log"
	"os"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/pkg/slashing"
	"github.com/docker/docker/client"
)

func main() {
	// Prompt used to interact with the user input
	prompt := prompts.NewPromptCli()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()
	serviceManager := services.NewServiceManager(dockerClient)
	slashingManager := slashing.NewSlashingDataManager(dockerClient, serviceManager)
	sedgeActions := actions.NewSedgeActions(serviceManager, slashingManager)
	sedgeCmd := cli.RootCmd()
	sedgeCmd.AddCommand(
		cli.CliCmd(prompt, slashingManager),
		cli.KeysCmd(prompt),
		cli.DownCmd(),
		cli.ClientsCmd(),
		cli.NetworksCmd(),
		cli.LogsCmd(),
		cli.VersionCmd(),
		cli.SlashingExportCmd(sedgeActions),
	)
	if err := sedgeCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
