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
	lido "github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"golang.org/x/exp/slices"
)

// lidoOptions for Lido Node
type lidoOptions struct{}

var _ SedgeOptions = (*lidoOptions)(nil)

func (l *lidoOptions) SupportedNetworks() []string {
	options := lido.LidoSupportedNetworks()
	return options
}

func (l *lidoOptions) OverwriteSettings() OverwriteSettings {
	return OverwriteSettings{
		FeeRecipient:      true,
		RelayURLs:         true,
		MevBoost:          true,
		WithdrawalAddress: true,
	}
}

func (l *lidoOptions) WithdrawalAddress(network string) string {
	supported := lido.NetworkSupportedByLidoWithdrawal(network)
	if supported {
		wa, _ := lido.WithdrawalAddress(network)
		return wa
	}
	return ""
}

func (l *lidoOptions) FeeRecipient(network string) string {
	supported := lido.LidoSupportedNetworks()
	if slices.Contains(supported, network) {
		fr, _ := lido.FeeRecipient(network)
		return fr
	}
	return ""
}

func (l *lidoOptions) RelayURLs(network string) ([]string, error) {
	relayURLs, err := mevboostrelaylist.RelaysURI(network)
	if err != nil {
		return nil, err
	}
	return relayURLs, nil
}

func (l *lidoOptions) MEVBoostEnabled(network string) bool {
	_, supported := mevboostrelaylist.NetworkSupportedByLidoMevBoost(network)
	return supported
}

func (l *lidoOptions) ValidateSettings(settings OptionSettings) error {
	// Check if MEVBoost is enabled and supported by the network
	if settings.MEVBoostEnabled {
		options, supported := mevboostrelaylist.NetworkSupportedByLidoMevBoost(settings.Network)
		if !supported {
			return ErrInvalidNetworkMevBoost(options, "Lido")
		}
	}

	// Check if the network is supported by Lido
	supportedNetworks := lido.LidoSupportedNetworks()
	if !slices.Contains(supportedNetworks, settings.Network) {
		return ErrInvalidNetwork(supportedNetworks, "Lido")
	}
	return nil
}
