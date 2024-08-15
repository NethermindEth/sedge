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
package csfeedistributor

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	bond "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Tree : struct that reperesents Merkle Tree data
type Tree struct {
	Format       string   `json:"format"`
	LeafEncoding []string `json:"leafEncoding"`
	Tree         []string `json:"tree"`
	Values       []struct {
		Value     []interface{} `json:"value"`
		TreeIndex int           `json:"treeIndex"`
	} `json:"values"`
}

/*
Rewards :
This function is responsible for:
retrieving non-claimed rewards for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. *big.Int
Non-claimed rewards
b. error
Error if any
*/
func Rewards(network string, nodeID *big.Int) (*big.Int, error) {
	var rewards *big.Int
	if nodeID.Sign() < 0 {
		return nil, fmt.Errorf("node ID value out-of-bounds: can't be negative")
	}

	treeCID, err := treeCID(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call treeCID: %w", err)
	}

	shares, err := cumulativeFeeShares(treeCID, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call cumulativeFeeShares: %w", err)
	}

	bondInfo, err := bond.BondSummary(network, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call BondSummary: %w", err)
	}

	rewards = new(big.Int).Add(bondInfo.Excess, shares)
	return rewards, nil
}

func cumulativeFeeShares(treeCID string, nodeID *big.Int) (*big.Int, error) {
	// Get tree data from IPFS
	treeData, err := treeData(treeCID)
	if err != nil {
		return nil, fmt.Errorf("failed to call treeData: %v", err)
	}

	index, err := binarySearchNodeID(nodeID, treeData)
	if err != nil {
		return nil, fmt.Errorf("failed to find node ID: %v", err)
	}

	shares, err := convertTreeValuesToBigInt(treeData.Values[index].Value[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert shares: %v", err)
	}
	return shares, nil
}

func treeCID(network string) (string, error) {
	var treeCIDString string
	contract, err := csFeeDistributorContract(network)
	if err != nil {
		return treeCIDString, fmt.Errorf("failed to call csFeeDistributorContract: %w", err)
	}

	treeCIDString, err = contract.TreeCid(nil)
	if err != nil {
		return treeCIDString, fmt.Errorf("failed to call TreeCid: %w", err)
	}
	return treeCIDString, nil
}

func treeData(treeCID string) (Tree, error) {
	var treeData Tree
	gatewayURL := fmt.Sprintf("https://ipfs.io/ipfs/%s", treeCID) // Public gateway URL
	resp, err := utils.GetRequest(gatewayURL, time.Second)
	if err != nil {
		return treeData, fmt.Errorf("failed to connect to gatewayURL: %w", err)
	}
	defer resp.Body.Close()

	// Read and print the data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return treeData, fmt.Errorf("failed to read data: %w", err)
	}

	// Collect data in treeData struct
	if err := json.Unmarshal(data, &treeData); err != nil {
		return treeData, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return treeData, nil
}

// Converts nodeOperatorId and shares from Tree.Values.Value interface
func convertTreeValuesToBigInt(value interface{}) (*big.Int, error) {
	bigIntValue := new(big.Int)
	bigIntValue.SetString(fmt.Sprintf("%.0f", value), 10)
	return bigIntValue, nil
}

func binarySearchNodeID(nodeID *big.Int, treeData Tree) (int, error) {
	// Compare nodeOperatorID in tree with nodeId to get shares
	low, high := 0, len(treeData.Values)-1
	for low <= high {
		mid := (low + high) / 2
		nodeOperatorId, err := convertTreeValuesToBigInt(treeData.Values[mid].Value[0])
		if err != nil {
			return 0, fmt.Errorf("failed to convert nodeOperatorId: %v", err)
		}
		cmp := nodeOperatorId.Cmp(nodeID)
		if cmp == 0 {
			return mid, nil
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, fmt.Errorf("invalid nodeId")
}

func csFeeDistributorContract(network string) (*Csfeedistributor, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to client: %w", err)
	}
	defer client.Close()

	contractName := contracts.CSFeeDistributor
	address := common.HexToAddress(contracts.DeployedAddresses(contractName)[network])
	contract, err := NewCsfeedistributor(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSFeeDistributor instance: %w", err)
	}
	return contract, nil
}
