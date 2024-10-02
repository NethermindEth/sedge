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
package contracts

import (
	"slices"

	"github.com/NethermindEth/sedge/configs"
)

type feeRecipientConfig struct {
	network             string
	feeRecipientAddress string
	weight              int // Weight of the network for sorting purposes
}

var feeRecipient = map[string]feeRecipientConfig{
	configs.NetworkMainnet: {
		network:             configs.NetworkMainnet,
		feeRecipientAddress: "0x388C818CA8B9251b393131C08a736A67ccB19297",
		weight:              1,
	},
	configs.NetworkHolesky: {
		network:             configs.NetworkHolesky,
		feeRecipientAddress: "0xE73a3602b99f1f913e72F8bdcBC235e206794Ac8",
		weight:              2,
	},
	configs.NetworkSepolia: {
		network:             configs.NetworkSepolia,
		feeRecipientAddress: "0x94B1B8e2680882f8652882e7F196169dE3d9a3B2",
		weight:              3,
	},
}

// FeeRecipient returns the fee recipient address for the given network
func FeeRecipient(network string) (string, bool) {
	config, ok := feeRecipient[network]
	return config.feeRecipientAddress, ok
}

// LidoSupportedNetworks returns the supported networks for Lido
func LidoSupportedNetworks() []string {
	networks := []string{}
	for network := range feeRecipient {
		networks = append(networks, network)
	}
	slices.SortFunc(networks, func(a, b string) int {
		return feeRecipient[a].weight - feeRecipient[b].weight
	})
	return networks
}
