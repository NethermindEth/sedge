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
package csaccounting

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BondInfo : Struct represent bond info of Node Operator
type BondInfo struct {
	Current  *big.Int
	Required *big.Int
	Excess   *big.Int
	Missed   *big.Int
}

/*
BondSummary :
This function is responsible for:
retrieving bond info for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. BondInfo
Struct that include bond info
b. error
Error if any
*/
func BondSummary(network string, nodeID *big.Int) (BondInfo, error) {
	var bondsInfo BondInfo
	if nodeID.Sign() < 0 {
		return bondsInfo, fmt.Errorf("node ID value out-of-bounds: can't be negative")
	}

	contract, client, err := csAccountingContract(network)
	if err != nil {
		return bondsInfo, fmt.Errorf("failed to call csAccountingContract: %w", err)
	}
	defer client.Close()

	result, err := contract.GetBondSummary(nil, nodeID)
	if err != nil {
		return bondsInfo, fmt.Errorf("failed to call GetBondSummary: %w", err)
	}
	bondsInfo.Current = result.Current
	bondsInfo.Required = result.Required

	// Calculate excess and missed bond amounts
	excess := new(big.Int).Sub(bondsInfo.Current, bondsInfo.Required)
	missed := new(big.Int).Sub(bondsInfo.Required, bondsInfo.Current)

	// Set to zero if negative
	if excess.Sign() < 0 {
		excess.SetInt64(0)
	}
	if missed.Sign() < 0 {
		missed.SetInt64(0)
	}

	bondsInfo.Excess = excess
	bondsInfo.Missed = missed
	return bondsInfo, nil
}

func csAccountingContract(network string) (*Csaccounting, *ethclient.Client, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	contractName := contracts.CSAccounting
	address := common.HexToAddress(contracts.DeployedAddresses(contractName)[network])
	contract, err := NewCsaccounting(address, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CSAccounting instance: %w", err)
	}
	return contract, client, nil
}
