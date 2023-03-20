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
	"runtime"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Commands Runner
	cmdRunner := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: runtime.GOOS == "linux",
	})

	// Prompt used to interact with the user input
	prompt := ui.NewPrompter()

	// Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	// Docker service
	serviceManager := services.NewServiceManager(dockerClient)

	// Init dependencies manager
	depsMgr := dependencies.NewDependenciesManager(cmdRunner)

	// Init Sedge Actions
	sdgOpts := actions.SedgeActionsOptions{
		DockerClient:   dockerClient,
		ServiceManager: serviceManager,
		CommandRunner:  cmdRunner,
	}
	sedgeActions := actions.NewSedgeActions(sdgOpts)

	sedgeCmd := cli.RootCmd()
	sedgeCmd.AddCommand(
		cli.CliCmd(prompt, sedgeActions, depsMgr),
		cli.KeysCmd(cmdRunner, prompt),
		cli.DownCmd(cmdRunner, sedgeActions, depsMgr),
		cli.ClientsCmd(),
		cli.NetworksCmd(),
		cli.LogsCmd(cmdRunner, sedgeActions, depsMgr),
		cli.VersionCmd(),
		cli.SlashingExportCmd(sedgeActions, depsMgr),
		cli.SlashingImportCmd(sedgeActions, depsMgr),
		cli.RunCmd(sedgeActions, depsMgr),
		cli.ImportKeysCmd(sedgeActions, depsMgr),
		cli.GenerateCmd(sedgeActions),
		cli.DependenciesCommand(depsMgr),
	)
	sedgeCmd.SilenceErrors = true
	sedgeCmd.SilenceUsage = true
	if err := sedgeCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
