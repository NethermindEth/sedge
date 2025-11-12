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

import (
	"fmt"

	"github.com/NethermindEth/sedge/configs"
)

type address = map[string]string

const (
	// Contract names
	CSModule                 = "csmodule"
	CSAccounting             = "csaccounting"
	CSFeeDistributor         = "csfeedistributor"
	MEVBoostRelayAllowedList = "mevboostrelayallowedlist"
	Vebo                     = "vebo"
	CSExitPenalties          = "csexitpenalties"
)

var deployedAddresses = map[string]address{
	CSModule: {
		configs.NetworkMainnet: "0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F",
		configs.NetworkHoodi:   "0x79CEf36D84743222f37765204Bec41E92a93E59d",
	},
	CSAccounting: {
		configs.NetworkMainnet: "0x4d72BFF1BeaC69925F8Bd12526a39BAAb069e5Da",
		configs.NetworkHoodi:   "0xA54b90BA34C5f326BC1485054080994e38FB4C60",
	},
	CSFeeDistributor: {
		configs.NetworkMainnet: "0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0",
		configs.NetworkHoodi:   "0xaCd9820b0A2229a82dc1A0770307ce5522FF3582",
	},
	MEVBoostRelayAllowedList: {
		configs.NetworkMainnet: "0xF95f069F9AD107938F6ba802a3da87892298610E",
		configs.NetworkHoodi:   "0x279d3A456212a1294DaEd0faEE98675a52E8A4Bf",
	},
	Vebo: {
		configs.NetworkMainnet: "0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e",
		configs.NetworkHoodi:   "0x8664d394C2B3278F26A1B44B967aEf99707eeAB2",
	},
	CSExitPenalties: {
		configs.NetworkMainnet: "0x06cd61045f958A209a0f8D746e103eCc625f4193",
		configs.NetworkHoodi:   "0xD259b31083Be841E5C85b2D481Cfc17C14276800",
	},
}

func DeployedAddresses(contractName string) address {
	return deployedAddresses[contractName]
}

func ContractAddressByNetwork(contractName, network string) (string, error) {
	address, found := deployedAddresses[contractName][network]
	if !found {
		return "", fmt.Errorf("no contract code at network %s, please double check a smart contract is deployed on given network", network)
	}
	return address, nil
}
