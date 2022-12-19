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

// TODO: Remove public level to this variable (NetworksConfigs), use getters to access instead
var NetworksConfigs map[string]NetworkConfig = map[string]NetworkConfig{
	EthereumMainnet.Name:        EthereumMainnet,
	GoerliEthereumTestnet.Name:  GoerliEthereumTestnet,
	SepoliaEthereumTestnet.Name: SepoliaEthereumTestnet,
	ChiadoGnosisTestnet.Name:    ChiadoGnosisTestnet,
	GnosisMainnet.Name:          GnosisMainnet,
	CustomNetwork.Name:          CustomNetwork,
}

var EthereumMainnet NetworkConfig = NetworkConfig{
	Name:               "mainnet",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x00000000",
}

var GoerliEthereumTestnet NetworkConfig = NetworkConfig{
	Name:               "goerli",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x00001020",
}

var SepoliaEthereumTestnet NetworkConfig = NetworkConfig{
	Name:               "sepolia",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x90000069",
}

var ChiadoGnosisTestnet NetworkConfig = NetworkConfig{
	Name:               "chiado",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x0000006f",
}

var GnosisMainnet NetworkConfig = NetworkConfig{
	Name:               "gnosis",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x00000064",
}

var CustomNetwork NetworkConfig = NetworkConfig{
	Name:               "custom",
	RequireJWT:         true,
	NetworkService:     "merge",
	GenesisForkVersion: "0x00000000", // TODO: only affects keystores generation, ensure the deposit method does not conflict over this.
}

// TODO: add doc
// TODO: test
func CheckNetwork(networkName string) error {
	if _, ok := NetworksConfigs[strings.ToLower(networkName)]; !ok {
		return fmt.Errorf(UnknownNetworkError, networkName)
	}
	return nil
}
