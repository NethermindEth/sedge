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
			Name:                     NetworkChiado,
			RequireJWT:               true,
			NetworkService:           "merge",
			GenesisForkVersion:       "0x0000006f",
			DefaultCustomConfigSrc:   "https://github.com/gnosischain/configs/raw/main/chiado/config.yaml",
			DefaultCustomGenesisSrc:  "https://github.com/gnosischain/configs/raw/main/chiado/genesis.ssz",
			DefaultCustomDeployBlock: "0",
		},
		{
			Name:               NetworkGnosis,
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
