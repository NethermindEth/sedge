/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"strings"

	"github.com/NethermindEth/1Click/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listClientsCmd represents the listClients command
var listClientsCmd = &cobra.Command{
	Use:   "listClients",
	Short: "List supported clients",
	Long:  `List supported clients for execution and consensus engines`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Listing supported clients")
		log.Info("Eth1 Clients: \n%s\n\n", strings.Join(configs.Eth1ClientsSupported, ", "))
		log.Info("Consensus Clients: \n%s\n", strings.Join(configs.ConsensusClientsSupported, ", "))
	},
}

func init() {
	rootCmd.AddCommand(listClientsCmd)
}
