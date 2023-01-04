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
	"runtime"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/client"
)

func main() {
	// Init configs
	configs.InitNetworksConfigs()
	// Commands Runner
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: runtime.GOOS == "linux",
	})
	// Prompt used to interact with the user input
	prompt := prompts.NewPromptCli()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()
	serviceManager := services.NewServiceManager(dockerClient)
	sedgeActions := actions.NewSedgeActions(dockerClient, serviceManager, cmdRunner)
	sedgeCmd := cli.RootCmd()
	sedgeCmd.AddCommand(
		cli.CliCmd(cmdRunner, prompt, serviceManager, sedgeActions),
		cli.KeysCmd(cmdRunner, prompt),
		cli.DownCmd(cmdRunner),
		cli.ClientsCmd(),
		cli.NetworksCmd(),
		cli.LogsCmd(cmdRunner),
		cli.VersionCmd(),
		cli.SlashingExportCmd(sedgeActions),
		cli.RunCmd(cmdRunner, sedgeActions),
		cli.GenerateCmd(prompt, sedgeActions),
	)
	if err := sedgeCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
