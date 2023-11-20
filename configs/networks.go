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
package configs

import (
	"errors"
	"fmt"
	"time"
)

const (
	// Network names
	NetworkMainnet   = "mainnet"
	NetworkGoerli    = "goerli"
	NetworkSepolia   = "sepolia"
	NetworkGnosis    = "gnosis"
	NetworkChiado    = "chiado"
	NetworkHolesky   = "holesky"
	NetworkVolta     = "volta"
	NetworkEnergyWeb = "energyweb"
	NetworkCustom    = "custom"
)

var ErrInvalidNetwork = errors.New("invalid network")

// added volta and EnergyWeb
func NetworkCheck(value string) error {
	switch value {
	case NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky, NetworkCustom, NetworkVolta, NetworkEnergyWeb:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrInvalidNetwork, value)
	}
}

func NetworkSupported() []string {
	// notest
	return []string{
		NetworkMainnet,
		NetworkGoerli,
		NetworkSepolia,
		NetworkGnosis,
		NetworkChiado,
		NetworkHolesky,
		NetworkCustom,
		NetworkVolta,
		NetworkEnergyWeb,
	}
}

func NetworkEpochTime(network string) time.Duration {
	switch network {
	case NetworkMainnet, NetworkGoerli, NetworkSepolia:
		return 7 * time.Minute
	case NetworkGnosis, NetworkChiado:
		return 2 * time.Minute
	default:
		return 7 * time.Minute
	}
}

func SupportsMEVBoost(network string) bool {
	out, ok := networksConfigs[network]
	return ok && out.SupportsMEVBoost
}
