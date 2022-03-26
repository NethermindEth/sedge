/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"math"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/pkg/clients"
	"github.com/NethermindEth/1Click/internal/ui"
	"github.com/NethermindEth/1Click/internal/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listClientsCmd represents the clients command
var listClientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "List supported clients",
	Long:  `List supported clients for execution and consensus engines`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Listing supported clients\n")

		data, err := buildData(clients.GetSupportedClients)
		if err != nil {
			log.Fatal(err)
		}

		ui.WriteListClientsTable(data)

		log.Infof("Listing clients provided in configuration file\n")

		data, err = buildData(configs.GetConfigClients)
		if err != nil {
			log.Fatal(err)
		}

		ui.WriteListClientsTable(data)
	},
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
a. [][]string
Table data
b. error
Error if any
*/
func buildData(getClients func(string) ([]string, error)) ([][]string, error) {
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

	max := int(math.Max(float64(len(executionClients)), float64(len(consensusClients))))
	max = int(math.Max(float64(max), float64(len(validatorClients))))

	if max > 0 {
		for _, list := range [...]*[]string{&executionClients, &consensusClients, &validatorClients} {
			for len(*list) < max {
				*list = append(*list, "-")
			}
		}
	} else {
		executionClients, consensusClients, validatorClients = []string{"-"}, []string{"-"}, []string{"-"}
	}

	data, err := utils.ZipString(executionClients, consensusClients, validatorClients)
	if err != nil {
		return nil, err
	}

	return data, nil
}
