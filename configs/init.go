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

var networkConfigs map[string]NetworkConfig

func NetworkConfigs() map[string]NetworkConfig {
	return networkConfigs
}

func InitNetworkConfigs() {
	networkConfigs = make(map[string]NetworkConfig)
}

func AddNetwork(network NetworkConfig) {
	if networkConfigs == nil {
		networkConfigs = make(map[string]NetworkConfig)
	}
	networkConfigs[network.Name] = network
}

func init() {
	InitNetworkConfigs()

	configs := []NetworkConfig{
		{
			Name:               NetworkMainnet,
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00000000",
		},
		{
			Name:               NetworkGoerli,
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00001020",
		},
		{
			Name:               NetworkSepolia,
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x90000069",
		},
		{
			Name:               NetworkChiado,
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x0000006f",
		},
		{
			Name:               NetworkGnosis,
			RequireJWT:         true,
			NetworkService:     "merge",
			GenesisForkVersion: "0x00000064",
		},
	}
	for _, config := range configs {
		AddNetwork(config)
	}
}
