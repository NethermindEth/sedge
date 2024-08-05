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
package options

import (
	"golang.org/x/exp/slices"

	"github.com/NethermindEth/sedge/configs"
)

// ethereumOptions for Ethereum Node
type ethereumOptions struct{}

var _ SedgeOptions = (*ethereumOptions)(nil)

func (e *ethereumOptions) SupportedNetworks() []string {
	networks := []string{}
	for _, network := range configs.NetworksConfigs() {
		if network.Name == configs.NetworkCustom {
			continue
		}
		networks = append(networks, network.Name)
	}
	slices.SortFunc(networks, func(a, b string) int {
		return configs.NetworksConfigs()[a].Weight - configs.NetworksConfigs()[b].Weight
	})
	return networks
}

func (e *ethereumOptions) OverwriteSettings() OverwriteSettings {
	return OverwriteSettings{
		FeeRecipient:      false,
		RelayURLs:         false,
		MevBoost:          false,
		WithdrawalAddress: false,
	}
}

func (e *ethereumOptions) WithdrawalAddress(network string) string {
	return ""
}

func (e *ethereumOptions) FeeRecipient(network string) string {
	return ""
}

func (e *ethereumOptions) RelayURLs(network string) ([]string, error) {
	networkConfig, ok := configs.NetworksConfigs()[network]
	if !ok {
		return nil, ErrInvalidNetwork(e.SupportedNetworks(), "Ethereum")
	}
	return networkConfig.RelayURLs, nil
}

func (e *ethereumOptions) MEVBoostEnabled(network string) bool {
	return configs.NetworksConfigs()[network].SupportsMEVBoost
}

func (e *ethereumOptions) ValidateSettings(settings OptionSettings) error {
	// HACK: Ethereum manages MEV Boost at a template level, thus we don't need to check if MEV Boost is enabled and supported by the network
	// It might be a good idea to remove this logic from the template and manage MEV Boost at the network level
	// Check if MEVBoost is enabled and supported by the network
	// if settings.MEVBoostEnabled && !e.MEVBoostEnabled(settings.Network) {
	// 	mevNetworks := make([]string, 0)
	// 	for _, network := range e.SupportedNetworks() {
	// 		if e.MEVBoostEnabled(network) {
	// 			mevNetworks = append(mevNetworks, network)
	// 		}
	// 	}
	// 	return ErrInvalidNetworkMevBoost(mevNetworks, "Ethereum")
	// }

	// Check if the network is supported
	supportedNetworks := e.SupportedNetworks()
	if !slices.Contains(supportedNetworks, settings.Network) {
		return ErrInvalidNetwork(supportedNetworks, "Ethereum")
	}
	return nil
}
