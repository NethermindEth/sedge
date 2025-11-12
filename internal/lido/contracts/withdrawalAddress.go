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
	"sort"

	"github.com/NethermindEth/sedge/configs"
)

type withdrawalAddressConfig struct {
	network           string
	withdrawalAddress string
}

var withdrawalAddress = map[string]withdrawalAddressConfig{
	configs.NetworkMainnet: {
		network:           configs.NetworkMainnet,
		withdrawalAddress: "0xb9d7934878b5fb9610b3fe8a5e441e8fad7e293f",
	},
	configs.NetworkHoodi: {
		network:           configs.NetworkHoodi,
		withdrawalAddress: "0x4473dCDDbf77679A643BdB654dbd86D67F8d32f2",
	},
}

// WithdrawalAddress returns the withdrawal address for the given network
func WithdrawalAddress(network string) (string, bool) {
	config, ok := withdrawalAddress[network]
	return config.withdrawalAddress, ok
}

// LidoWithdrawalSupportedNetworks returns the supported networks for Lido withdrawal
func LidoWithdrawalSupportedNetworks() []string {
	networks := []string{}
	for network := range withdrawalAddress {
		networks = append(networks, network)
	}
	sort.Strings(networks)
	return networks
}

// NetworkSupportedByLidoWithdrawal checks if the given network is supported by Lido
func NetworkSupportedByLidoWithdrawal(network string) bool {
	supportedNetworks := LidoWithdrawalSupportedNetworks()
	for _, supportedNetwork := range supportedNetworks {
		if network == supportedNetwork {
			return true
		}
	}
	return false
}
