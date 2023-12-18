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
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ClientsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clients",
		Short: "List supported clients",
		Long:  `List supported clients for execution and consensus engines`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runListClientsCmd(cmd, args); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}

func runListClientsCmd(cmd *cobra.Command, args []string) error {
	networks, err := utils.SupportedNetworks()
	if err != nil {
		return err
	}

	// omit starknet for the following networks
	omitStarknetNetworks := map[string]bool{
		"chiado":  true,
		"goerli":  true,
		"custom":  true,
		"gnosis":  true,
		"holesky": true,
	}

	for _, n := range networks {
		c := clients.ClientInfo{Network: n}

		// Check if 'starknet' should be omitted for this network
		omitStarknet := omitStarknetNetworks[n]

		data, err := buildData(c.SupportedClients, !omitStarknet, nil)
		if err != nil {
			return err
		}

		ui.WriteListClientsTable(cmd.OutOrStdout(), data)
	}

	return nil
}

// /*
// buildData :
// Builds the data for the supported clients table

// params :-
//

// returns :-
// a. ui.ListClientsTable
// Table data
// b. error
// Error if any
// */
func buildData(getClients func(string) ([]string, error), checkStarknet bool, starknetClients []string) (*ui.ListClientsTable, error) {
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

	// Include 'starknet' clients if needed
	if checkStarknet {
		starknetClients, err = getClients("starknet")
		if err != nil {
			return nil, err
		}
	}

	return &ui.ListClientsTable{
		ClientTypes: []string{"Execution", "Consensus", "Validator", "Starknet"},
		Clients:     [][]string{executionClients, consensusClients, validatorClients, starknetClients},
	}, nil
}
