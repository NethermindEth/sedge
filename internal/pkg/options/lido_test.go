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
	"testing"

	lido "github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"github.com/stretchr/testify/assert"
)

func TestLidoOptions_SupportedNetworks(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)
	expected := []string{"mainnet", "holesky", "sepolia"}
	assert.Equal(t, expected, options.SupportedNetworks())
}

func TestLidoOptions_OverwriteSettings(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)
	got := options.OverwriteSettings()
	want := OverwriteSettings{
		FeeRecipient:      true,
		RelayURLs:         true,
		MevBoost:          true,
		WithdrawalAddress: true,
	}
	assert.Equal(t, want, got)
}

func TestLidoOptions_WithdrawalAddress(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)
	expected := func(network string) string {
		address, _ := lido.WithdrawalAddress(network)
		return address
	}

	tests := []struct {
		network  string
		expected func(network string) string
	}{
		{
			"mainnet",
			expected,
		},
		{
			"sepolia",
			expected,
		},
		{
			"holesky",
			expected,
		},
		{
			"unsupported",
			func(network string) string {
				return ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			address := options.WithdrawalAddress(tt.network)
			assert.Equal(t, tt.expected(tt.network), address)
		})
	}
}

func TestLidoOptions_FeeRecipient(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)
	expected := func(network string) string {
		address, _ := lido.FeeRecipient(network)
		return address
	}

	tests := []struct {
		network  string
		expected func(network string) string
	}{
		{
			"mainnet",
			expected,
		},
		{
			"sepolia",
			expected,
		},
		{
			"holesky",
			expected,
		},
		{
			"unsupported",
			func(network string) string {
				return ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			recipient := options.FeeRecipient(tt.network)
			assert.Equal(t, tt.expected(tt.network), recipient)
		})
	}
}

func TestLidoOptions_RelayURLs(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)
	expected := func(network string) []string {
		got, _ := mevboostrelaylist.RelaysURI(network)
		return got
	}

	tests := []struct {
		network  string
		expected func(network string) []string
	}{
		{"mainnet", expected},
		{"sepolia", expected},
		{"holesky", expected},
		{"unsupported", func(network string) []string {
			return nil
		}},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			urls, err := options.RelayURLs(tt.network)
			assert.Equal(t, tt.expected(tt.network), urls)
			if tt.expected(tt.network) != nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLidoOptions_MEVBoostEnabled(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)

	tests := []struct {
		network string
		enabled bool
	}{
		{"mainnet", true},
		{"sepolia", false},
		{"holesky", true},
		{"unsupported", false},
	}

	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			enabled := options.MEVBoostEnabled(tt.network)
			assert.Equal(t, tt.enabled, enabled)
		})
	}
}

func TestLidoOptionsValidateSettings(t *testing.T) {
	options := CreateSedgeOptions(LidoNode)

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
			name: "valid settings, sepolia, no mevboost",
			settings: OptionSettings{
				Network:         "sepolia",
				MEVBoostEnabled: false,
			},
			expected: nil,
		},
		{
			name: "invalid settings, sepolia, mevboost",
			settings: OptionSettings{
				Network:         "sepolia",
				MEVBoostEnabled: true,
			},
			expected: ErrInvalidNetworkMevBoost(mevboostrelaylist.LidoSupportedNetworksMevBoost(), "Lido"),
		},
		{
			name: "invalid settings, unsupported network",
			settings: OptionSettings{
				Network:         "unsupported",
				MEVBoostEnabled: false,
			},
			expected: ErrInvalidNetwork(lido.LidoSupportedNetworks(), "Lido"),
		},
		{
			name: "valid settings, holesky",
			settings: OptionSettings{
				Network:         "holesky",
				MEVBoostEnabled: true,
			},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := options.ValidateSettings(tt.settings)
			if tt.expected != nil {
				assert.EqualError(t, err, tt.expected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
