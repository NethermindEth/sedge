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
	"errors"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/stretchr/testify/assert"
)

func TestEthereumOptions_SupportedNetworks(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)
	got := options.SupportedNetworks()
	want := []string{"mainnet", "holesky", "sepolia", "gnosis", "chiado"}
	assert.Equal(t, want, got)
}

func TestEthereumOptions_OverwriteSettings(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)
	got := options.OverwriteSettings()
	want := OverwriteSettings{}
	assert.Equal(t, want, got)
}

func TestEthereumOptions_WithdrawalAddress(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)

	tests := []struct {
		network  string
		expected string
	}{
		{"mainnet", ""},
		{"unsupported", ""},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			address := options.WithdrawalAddress(tt.network)
			assert.Equal(t, tt.expected, address)
		})
	}
}

func TestEthereumOptions_FeeRecipient(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)

	tests := []struct {
		network  string
		expected string
	}{
		{"mainnet", ""},
		{"unsupported", ""},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			recipient := options.FeeRecipient(tt.network)
			assert.Equal(t, tt.expected, recipient)
		})
	}
}

func TestEthereumOptions_RelayURLs(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)
	expected := func(network string) []string {
		return configs.NetworksConfigs()[network].RelayURLs
	}
	empty := func(network string) []string {
		return nil
	}

	tests := []struct {
		network  string
		expected func(string) []string
		err      error
	}{
		{"mainnet", expected, nil},
		{"holesky", expected, nil},
		{"sepolia", expected, nil},
		{"gnosis", expected, nil},
		{"chiado", expected, nil},
		{"unsupported", empty, ErrInvalidNetwork(options.SupportedNetworks(), "Ethereum")},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			urls, err := options.RelayURLs(tt.network)
			assert.Equal(t, tt.expected(tt.network), urls)
			if tt.err != nil {
				assert.ErrorContains(t, err, "invalid network: Choose valid network for Ethereum")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEthereumOptions_MEVBoostEnabled(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)
	enabled := func(network string) bool {
		return configs.NetworksConfigs()[network].SupportsMEVBoost
	}
	empty := func(network string) bool {
		return false
	}

	tests := []struct {
		network  string
		expected func(string) bool
	}{
		{"mainnet", enabled},
		{"holesky", enabled},
		{"sepolia", enabled},
		{"gnosis", enabled},
		{"chiado", enabled},
		{"unsupported", empty},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			enabled := options.MEVBoostEnabled(tt.network)
			assert.Equal(t, tt.expected(tt.network), enabled)
		})
	}
}

func TestEthereumOptions_ValidateSettings(t *testing.T) {
	options := CreateSedgeOptions(EthereumNode)

	tests := []struct {
		name     string
		settings OptionSettings
		expected error
	}{
		{
			name: "valid settings, mainnet",
			settings: OptionSettings{
				Network:         "mainnet",
				MEVBoostEnabled: true,
			},
			expected: nil,
		},
		{
			name: "valid settings, holesky",
			settings: OptionSettings{
				Network:         "holesky",
				MEVBoostEnabled: true,
			},
			expected: nil,
		},
		{
			name: "valid settings, sepolia",
			settings: OptionSettings{
				Network:         "sepolia",
				MEVBoostEnabled: true,
			},
			expected: nil,
		},
		{
			name: "valid settings, gnosis",
			settings: OptionSettings{
				Network:         "gnosis",
				MEVBoostEnabled: false,
			},
			expected: nil,
		},
		{
			name: "invalid settings, unsupported network",
			settings: OptionSettings{
				Network:         "unsupported",
				MEVBoostEnabled: false,
			},
			expected: errors.New("invalid network: Choose valid network for Ethereum"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := options.ValidateSettings(tt.settings)
			if tt.expected != nil {
				assert.ErrorContains(t, err, tt.expected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
