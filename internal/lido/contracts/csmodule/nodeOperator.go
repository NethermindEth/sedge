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
	"github.com/ethereum/go-ethereum/common"
)

/*
NodeID :
This function is responsible for:
retrieving NodeOperatorID for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
rewardaddress (string): The reward address of the node operator
returns :-
a. *big.Int
Node Operator ID
b. error
Error if any
*/
func NodeID(network string, rewardAddress string) (*big.Int, error) {
	// Convert the reward address to a common.Address and check if it's zero
	rewardAddr := common.HexToAddress(rewardAddress)
	if rewardAddr == (common.Address{}) {
		return nil, fmt.Errorf("invalid reward address: can't be zero address")
	}

	nodeOperatorIDs, err := nodeOpIDs(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call NodeOpIDs: %w", err)
	}

	for _, nodeID := range nodeOperatorIDs {
		node, err := NodeOperatorInfo(network, nodeID)
		if err != nil {
			return nil, fmt.Errorf("failed to get NodeOperatorInfo: %w", err)
		}
		if node.RewardAddress == rewardAddr || node.ProposedRewardAddress == rewardAddr {
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
	contract, err := csModuleContract(network)
	if err != nil {
		return nodeOperator, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	nodeOperator, err = contract.GetNodeOperator(nil, nodeID)
	if err != nil {
		return nodeOperator, fmt.Errorf("failed to call GetNodeOperator: %w", err)
	}
	return nodeOperator, nil
}

func nodeOpIDs(network string) ([]*big.Int, error) {
	var nodeOperatorIDs []*big.Int
	contract, err := csModuleContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	limit, err := nodeOpsCount(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call nodeOpsCount: %w", err)
	}
	offset := big.NewInt(0)

	nodeOperatorIDs, err = contract.GetNodeOperatorIds(nil, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetNodeOperatorIds: %w", err)
	}
	return nodeOperatorIDs, nil
}

func nodeOpsCount(network string) (*big.Int, error) {
	var nodeOperatorCount *big.Int
	contract, err := csModuleContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	nodeOperatorCount, err = contract.GetNodeOperatorsCount(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetNodeOperatorsCount: %w", err)
	}

	return nodeOperatorCount, nil
}

func csModuleContract(network string) (*Csmodule, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call ConnectContract: %w", err)
	}
	defer client.Close()

	contractName := contracts.CSModule
	address := common.HexToAddress(contracts.DeployedAddresses(contractName)[network])
	contract, err := NewCsmodule(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSModule instance: %w", err)
	}
	return contract, nil
}
