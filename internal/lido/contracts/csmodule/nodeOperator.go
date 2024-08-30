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
package csmodule

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
NodeID :
This function is responsible for:
retrieving NodeOperatorID for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
rewardAddress (string): The reward address of the node operator
returns :-
a. *big.Int
Node Operator ID
b. error
Error if any
*/
func NodeID(network string, rewardAddress string) (*big.Int, error) {
	err := ui.EthAddressValidator(rewardAddress, false)
	if err != nil {
		return nil, fmt.Errorf("invalid reward address: %w", err)
	}
	// Convert the reward address to a common.Address and check if it's zero
	rewardAddr := common.HexToAddress(rewardAddress)
	if rewardAddr == (common.Address{}) {
		return nil, fmt.Errorf("invalid reward address: can't be zero address")
	}

	nodeOperatorIDs, err := nodeOpIDs(network)
	if err != nil {
		return nil, fmt.Errorf("error getting Node Operators ID: %w", err)
	}

	for _, nodeID := range nodeOperatorIDs {
		node, err := NodeOperatorInfo(network, nodeID)
		if err != nil {
			return nil, fmt.Errorf("error getting Node Operator Information: %w", err)
		}
		if node.RewardAddress == rewardAddr {
			return nodeID, nil
		}
	}
	return nil, fmt.Errorf("invalid reward address: %s", rewardAddress)
}

/*
NodeOperatorInfo :
This function is responsible for:
retrieving NodeOperator info for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. NodeOperator
Struct that includes Node Operator info
b. error
Error if any
*/
func NodeOperatorInfo(network string, nodeID *big.Int) (NodeOperator, error) {
	var nodeOperator NodeOperator
	contract, client, err := csModuleContract(network)
	defer client.Close()
	if err != nil {
		return nodeOperator, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	nodeOperator, err = contract.GetNodeOperator(nil, nodeID)
	if err != nil {
		return nodeOperator, fmt.Errorf("failed to call GetNodeOperator contract method: %w", err)
	}
	return nodeOperator, nil
}

func nodeOpIDs(network string) ([]*big.Int, error) {
	var nodeOperatorIDs []*big.Int
	contract, client, err := csModuleContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}
	defer client.Close()

	limit, err := nodeOpsCount(network)
	if err != nil {
		return nil, fmt.Errorf("error getting total number of Node Operators: %w", err)
	}
	offset := big.NewInt(0)

	nodeOperatorIDs, err = contract.GetNodeOperatorIds(nil, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetNodeOperatorIds contract method: %w", err)
	}
	return nodeOperatorIDs, nil
}

func nodeOpsCount(network string) (*big.Int, error) {
	var nodeOperatorCount *big.Int
	contract, client, err := csModuleContract(network)
	defer client.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	nodeOperatorCount, err = contract.GetNodeOperatorsCount(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetNodeOperatorsCount contract method: %w", err)
	}

	return nodeOperatorCount, nil
}

func csModuleContract(network string) (*Csmodule, *ethclient.Client, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	contractName := contracts.CSModule

	contractAddress, err := contracts.ContractAddressByNetwork(contractName, network)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get deployed contract address: %w", err)
	}

	address := common.HexToAddress(contractAddress)
	contract, err := NewCsmodule(address, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CSModule instance: %w", err)
	}
	return contract, client, nil
}
