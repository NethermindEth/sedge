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
package contracts

import "github.com/NethermindEth/sedge/configs"

const (
	// Contract names
	CSModule                 = "csmodule"
	CSAccounting             = "sepolia"
	CSFeeDistributor         = "gnosis"
	MEVBoostRelayAllowedList = "mevboostrelayallowedlist"
)

var deployedAddresses = map[string]map[string]string{
	CSModule: {
		configs.NetworkHolesky: "0x4562c3e63c2e586cD1651B958C22F88135aCAd4f",
	},
	CSAccounting: {
		configs.NetworkHolesky: "0xc093e53e8F4b55A223c18A2Da6fA00e60DD5EFE1",
	},
	CSFeeDistributor: {
		configs.NetworkHolesky: "0xD7ba648C8F72669C6aE649648B516ec03D07c8ED",
	},
	MEVBoostRelayAllowedList: {
		configs.NetworkMainnet: "0xF95f069F9AD107938F6ba802a3da87892298610E",
		configs.NetworkHolesky: "0x2d86C5855581194a386941806E38cA119E50aEA3",
	},
}

func DeployedAddresses(contractName string) map[string]string {
	return deployedAddresses[contractName]
}
