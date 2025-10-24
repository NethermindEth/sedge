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
	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
)

// Blockstamp represents block information in the new IPFS data structure
type Blockstamp struct {
	BlockHash      string `json:"block_hash"`
	BlockNumber    int64  `json:"block_number"`
	BlockTimestamp int64  `json:"block_timestamp"`
	RefEpoch       int64  `json:"ref_epoch"`
	RefSlot        int64  `json:"ref_slot"`
	SlotNumber     int64  `json:"slot_number"`
	StateRoot      string `json:"state_root"`
}

// Duty represents assigned vs included duties
type Duty struct {
	Assigned int `json:"assigned"`
	Included int `json:"included"`
}

// PerformanceCoefficients represents the weighting for different duties
type PerformanceCoefficients struct {
	AttestationsWeight int `json:"attestations_weight"`
	BlocksWeight       int `json:"blocks_weight"`
	SyncWeight         int `json:"sync_weight"`
}

// Validator represents a single validator's data
type Validator struct {
	AttestationDuty    Duty     `json:"attestation_duty"`
	DistributedRewards *big.Int `json:"distributed_rewards"`
	Performance        float64  `json:"performance"`
	ProposalDuty       Duty     `json:"proposal_duty"`
	RewardsShare       float64  `json:"rewards_share"`
	Slashed            bool     `json:"slashed"`
	Strikes            int      `json:"strikes"`
	SyncDuty           Duty     `json:"sync_duty"`
	Threshold          float64  `json:"threshold"`
}

// Operator represents an operator's data with all validators
type Operator struct {
	DistributedRewards      *big.Int                `json:"distributed_rewards"`
	PerformanceCoefficients PerformanceCoefficients `json:"performance_coefficients"`
	Validators              map[string]Validator    `json:"validators"`
}

// Report : struct that represents a single report in the new IPFS data structure (legacy format)
type Report struct {
	NodeOperatorID      *big.Int               `json:"nodeOperatorId"`
	CumulativeFeeShares *big.Int               `json:"cumulativeFeeShares"`
	PerformanceMetrics  map[string]interface{} `json:"performanceMetrics,omitempty"`
	Strikes             map[string]interface{} `json:"strikes,omitempty"`
	Proof               []string               `json:"proof,omitempty"` // Merkle proof for verification
}

// Tree : struct that represents Merkle Tree data (legacy format)
type Tree struct {
	Format       string   `json:"format"`
	LeafEncoding []string `json:"leafEncoding"`
	Tree         []string `json:"tree"`
	Values       []struct {
		Value     []interface{} `json:"value"`
		TreeIndex int           `json:"treeIndex"`
	} `json:"values"`
}

// NewTreeData : struct that represents the new IPFS data structure with operators and validators
type NewTreeData struct {
	Blockstamp         Blockstamp          `json:"blockstamp"`
	Distributable      *big.Int            `json:"distributable"`
	DistributedRewards *big.Int            `json:"distributed_rewards"`
	Frame              []int64             `json:"frame"`
	Operators          map[string]Operator `json:"operators"`
	RebateToProtocol   *big.Int            `json:"rebate_to_protocol"`
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
		return nil, fmt.Errorf("error getting treeCID: %w", err)
	}

	shares, err := cumulativeFeeShares(treeCID, nodeID)
	if err != nil {
		return nil, fmt.Errorf("error getting Node Operator shares: %w", err)
	}

	bondInfo, err := bond.BondSummary(network, nodeID)
	if err != nil {
		return nil, fmt.Errorf("error getting Node Operator bond: %w", err)
	}

	rewards = new(big.Int).Add(bondInfo.Excess, shares)
	return rewards, nil
}

func cumulativeFeeShares(treeCID string, nodeID *big.Int) (*big.Int, error) {
	// Get tree data from IPFS
	treeData, err := treeData(treeCID)
	if err != nil {
		return nil, fmt.Errorf("error getting tree data: %v", err)
	}

	// Try new format first (list of reports)
	if len(treeData.Values) == 0 {
		// Try to parse as new format
		newTreeData, err := parseNewTreeData(treeCID)
		if err == nil {
			return getSharesFromReports(newTreeData, nodeID)
		}
		log.Debugf("Failed to parse as new format, falling back to legacy format: %v", err)
	}

	// Legacy format: Compare nodeOperatorID in tree with nodeId to get shares
	for _, item := range treeData.Values {
		if len(item.Value) == 2 {
			nodeOperatorId, err1 := convertTreeValuesToBigInt(item.Value[0])
			shares, err2 := convertTreeValuesToBigInt(item.Value[1])
			if err1 != nil || err2 != nil {
				log.Debugf("Error converting values: %v, %v", err1, err2)
				continue
			}
			if nodeOperatorId.Cmp(nodeID) == 0 {
				log.Debugf("shares: %v", shares)
				return shares, nil
			}
		} else {
			log.Debugf("Unexpected value format, expected 2 elements")
		}
	}
	return nil, fmt.Errorf("invalid nodeId")
}

func treeCID(network string) (string, error) {
	var treeCIDString string
	contract, client, err := csFeeDistributorContract(network)
	if err != nil {
		return treeCIDString, fmt.Errorf("failed to call csFeeDistributorContract: %w", err)
	}
	defer client.Close()

	treeCIDString, err = contract.TreeCid(nil)
	if err != nil {
		return treeCIDString, fmt.Errorf("failed to call TreeCid contract method: %w", err)
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

// parseNewTreeData parses the new IPFS data structure with list of reports
func parseNewTreeData(treeCID string) (*NewTreeData, error) {
	var newTreeData NewTreeData
	gatewayURL := fmt.Sprintf("https://ipfs.io/ipfs/%s", treeCID) // Public gateway URL
	resp, err := utils.GetRequest(gatewayURL, time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gatewayURL: %w", err)
	}
	defer resp.Body.Close()

	// Read the data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	// Try to parse as new format first
	if err := json.Unmarshal(data, &newTreeData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON as new format: %w", err)
	}

	// Validate that we have operators
	if len(newTreeData.Operators) == 0 {
		return nil, fmt.Errorf("no operators found in new format")
	}

	return &newTreeData, nil
}

// getSharesFromReports extracts cumulative fee shares from the new operator-based format
func getSharesFromReports(newTreeData *NewTreeData, nodeID *big.Int) (*big.Int, error) {
	nodeIDStr := nodeID.String()
	operator, exists := newTreeData.Operators[nodeIDStr]
	if !exists {
		return nil, fmt.Errorf("nodeID %v not found in operators", nodeID)
	}

	log.Debugf("Found operator %v with distributed rewards: %v", nodeID, operator.DistributedRewards)
	log.Debugf("Performance coefficients: attestations=%d, blocks=%d, sync=%d",
		operator.PerformanceCoefficients.AttestationsWeight,
		operator.PerformanceCoefficients.BlocksWeight,
		operator.PerformanceCoefficients.SyncWeight)
	log.Debugf("Number of validators: %d", len(operator.Validators))

	// Log validator performance data
	for validatorID, validator := range operator.Validators {
		log.Debugf("Validator %s: performance=%.4f, strikes=%d, slashed=%v, rewards=%v",
			validatorID, validator.Performance, validator.Strikes, validator.Slashed, validator.DistributedRewards)
	}

	return operator.DistributedRewards, nil
}

// Converts nodeOperatorId and shares from Tree.Values.Value interface
func convertTreeValuesToBigInt(value interface{}) (*big.Int, error) {
	bigIntValue := new(big.Int)
	bigIntValue.SetString(fmt.Sprintf("%.0f", value), 10)
	return bigIntValue, nil
}

func csFeeDistributorContract(network string) (*Csfeedistributor, *ethclient.Client, error) {
	client, err := contracts.ConnectClient(network, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	contractName := contracts.CSFeeDistributor

	contractAddress, err := contracts.ContractAddressByNetwork(contractName, network)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get deployed contract address: %w", err)
	}

	address := common.HexToAddress(contractAddress)
	contract, err := NewCsfeedistributor(address, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CSFeeDistributor instance: %w", err)
	}
	return contract, client, nil
}
