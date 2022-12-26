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
	"os"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/prompts"
)

func main() {
	// Prompt used to interact with the user input
	prompt := prompts.NewPromptCli()
	sedgeCmd := cli.RootCmd()
	sedgeCmd.AddCommand(
		cli.CliCmd(prompt),
		cli.KeysCmd(prompt),
		cli.DownCmd(),
		cli.ClientsCmd(),
		cli.NetworksCmd(),
		cli.LogsCmd(),
		cli.VersionCmd(),
		cli.GenerateCmd(prompt),
	)
	if err := sedgeCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
