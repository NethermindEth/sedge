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
	"github.com/NethermindEth/1click/internal/ui"
	"github.com/NethermindEth/1click/internal/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listNetworksCmd represents the networks command
var listNetworksCmd = &cobra.Command{
	Use:   "networks",
	Short: "List supported networks",
	Long:  `List supported networks`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runListNetworksCmd(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
}

func runListNetworksCmd(cmd *cobra.Command, args []string) error {
	// Get supported networks and print table of networks
	networks, err := utils.SupportedNetworks()
	if err != nil {
		return err
	}

	log.Infof("Listing supported networks\n")

	ui.WriteListNetworksTable(cmd.OutOrStdout(), networks)

	return nil
}

func init() {
	rootCmd.AddCommand(listNetworksCmd)
}
