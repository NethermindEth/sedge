/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"
	"math"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/alexeyco/simpletable"
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

		writeTable(data)
	},
}

func init() {
	rootCmd.AddCommand(listClientsCmd)
}

func writeTable(data [][]string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Execution Client"},
			{Align: simpletable.AlignCenter, Text: "Consensus Client"},
		},
	}

	for i, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignCenter, Text: interface{}(row[0]).(string)},
			{Align: simpletable.AlignCenter, Text: interface{}(row[1]).(string)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
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
