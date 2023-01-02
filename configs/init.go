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
	"fmt"
	"strings"
)

var networksConfigs map[string]NetworkConfig

var CustomNetwork NetworkConfig = NetworkConfig{
	Name:               "custom",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x00000000", // TODO: only affects keystores generation, ensure the deposit method does not conflict over this.
}

func NetworksConfigs() map[string]NetworkConfig {
	return networksConfigs
}

func InitNetworksConfigs() {
	networksConfigs = make(map[string]NetworkConfig)

	configs := []NetworkConfig{
		{
			Name:               "mainnet",
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00000000",
		},
		{
			Name:               "goerli",
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00001020",
		},
		{
			Name:               "sepolia",
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x90000069",
		},
		{
			Name:               "chiado",
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x0000006f",
		},
		{
			Name:               "gnosis",
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00000064",
		},
		CustomNetwork,
	}
	for _, config := range configs {
		AddNetwork(config)
	}
}

func AddNetwork(network NetworkConfig) {
	if networksConfigs == nil {
		networksConfigs = make(map[string]NetworkConfig)
	}
	_, found := networksConfigs[network.Name]
	if !found {
		networksConfigs[network.Name] = network
	}
}

func CheckNetwork(networkName string) error {
	if _, ok := networksConfigs[strings.ToLower(networkName)]; !ok {
		return fmt.Errorf(UnknownNetworkError, networkName)
	}
	return nil
}
