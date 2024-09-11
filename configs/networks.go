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
	"math/rand"
	"time"
)

const (
	// Network names
	NetworkMainnet = "mainnet"
	NetworkSepolia = "sepolia"
	NetworkGnosis  = "gnosis"
	NetworkChiado  = "chiado"
	NetworkHolesky = "holesky"
	NetworkCustom  = "custom"
)

var ErrInvalidNetwork = errors.New("invalid network")

func NetworkCheck(value string) error {
	switch value {
	case NetworkMainnet, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky, NetworkCustom:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrInvalidNetwork, value)
	}
}

func NetworkSupported() []string {
	// notest
	return []string{
		NetworkMainnet,
		NetworkSepolia,
		NetworkGnosis,
		NetworkChiado,
		NetworkHolesky,
		NetworkCustom,
	}
}

func NetworkEpochTime(network string) time.Duration {
	switch network {
	case NetworkMainnet, NetworkSepolia:
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

func GetPublicRPCs(network string) ([]string, error) {
	rpcs, exists := networkRPCs[network]
	if !exists {
		return nil, fmt.Errorf("invalid network")
	}
	// Create a copy of the slice to avoid modifying the original
	shuffledRPCs := make([]string, len(rpcs.PublicRPCs))
	copy(shuffledRPCs, rpcs.PublicRPCs)

	// Shuffle the slice to randomize the order
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(shuffledRPCs), func(i, j int) {
		shuffledRPCs[i], shuffledRPCs[j] = shuffledRPCs[j], shuffledRPCs[i]
	})

	return shuffledRPCs, nil
}
