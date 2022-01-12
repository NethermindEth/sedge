/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"math"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/ui"
	"github.com/NethermindEth/1Click/internal/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listClientsCmd represents the listClients command
var listClientsCmd = &cobra.Command{
	Use:   "listClients",
	Short: "List supported clients",
	Long:  `List supported clients for execution and consensus engines`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Listing supported clients\n\n")

		data, err := buildData()
		if err != nil {
			log.Fatal(err)
		}

		ui.WriteListClientsTable(data)
	},
}

func init() {
	rootCmd.AddCommand(listClientsCmd)
}

func buildData() ([][]string, error) {
	executionClients, consensusClients := configs.GetClients("executionClients"), configs.GetClients("consensusClients")
	max := int(math.Max(float64(len(executionClients)), float64(len(consensusClients))))

	if max > 0 {
		for _, list := range []*[]string{&executionClients, &consensusClients} {
			for len(*list) < max {
				*list = append(*list, "-")
			}
		}
	} else {
		executionClients, consensusClients = []string{"-"}, []string{"-"}
	}

	data, err := utils.ZipString(executionClients, consensusClients)
	if err != nil {
		return nil, err
	}

	return data, nil
}
