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

// SedgeOptions is an interface for creating Sedge setup options
type SedgeOptions interface {
	Layer1Options
	SupportedNetworks() []string
	OverwriteSettings() OverwriteSettings
	// ValidateSettings can be used to validate the settings for the Sedge Option
	ValidateSettings(settings OptionSettings) error
}

// Layer1Options is an interface for creating Layer 1 setup options
type Layer1Options interface {
	WithdrawalAddress(network string) string
	FeeRecipient(network string) string
	RelayURLs(network string) ([]string, error)
	MEVBoostEnabled(network string) bool
}

// CreateSedgeOptions returns the appropriate SedgeOptions based on the Sedge Node Setup kind.
func CreateSedgeOptions(sedgeSetup string) SedgeOptions {
	switch sedgeSetup {
	case LidoNode:
		return &lidoOptions{}
	default:
		return &ethereumOptions{}
	}
}
