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
	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/ui"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listClientsCmd represents the clients command
var listClientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "List supported clients",
	Long:  `List supported clients for execution and consensus engines`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runListClientsCmd(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
}

func runListClientsCmd(cmd *cobra.Command, args []string) error {
	log.Infof("Listing supported clients\n")

	data, err := buildData(clients.GetSupportedClients)
	if err != nil {
		return err
	}

	ui.WriteListClientsTable(cmd.OutOrStdout(), data)

	log.Infof("Listing clients provided in configuration file\n")

	data, err = buildData(configs.GetConfigClients)
	if err != nil {
		return err
	}

	ui.WriteListClientsTable(cmd.OutOrStdout(), data)
	return nil
}

func init() {
	rootCmd.AddCommand(listClientsCmd)
}

/*
buildData :
Builds the data for the supported clients table

params :-
None

returns :-
a. ui.ListClientsTable
Table data
b. error
Error if any
*/
func buildData(getClients func(string) ([]string, error)) (*ui.ListClientsTable, error) {
	executionClients, err := getClients("execution")
	if err != nil {
		return nil, err
	}
	consensusClients, err := getClients("consensus")
	if err != nil {
		return nil, err
	}
	validatorClients, err := getClients("validator")
	if err != nil {
		return nil, err
	}

	return &ui.ListClientsTable{
		ClientTypes: []string{"Execution", "Consensus", "Validator"},
		Clients:     [][]string{executionClients, consensusClients, validatorClients},
	}, nil
}
