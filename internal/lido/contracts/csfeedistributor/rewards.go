package csfeedistributor

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	bond "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
	"github.com/ethereum/go-ethereum/common"
)

// Tree : struct that reperesents data of Merkle Tree
type Tree struct {
	Format       string   `json:"format"`
	LeafEncoding []string `json:"leafEncoding"`
	Tree         []string `json:"tree"`
	Values       []struct {
		Value     []interface{} `json:"value"`
		TreeIndex int           `json:"treeIndex"`
	} `json:"values"`
}

var deployedContractAddresses = map[string]string{
	configs.NetworkHolesky: "0xD7ba648C8F72669C6aE649648B516ec03D07c8ED",
}

/*
Rewards :
This function is responsible for:
retrieving non-claimed rewards for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
proofStrings ([]string): Merkle Proof of Node Operator Rewards
returns :-
a. *big.Int
Non-claimed rewards
b. error
Error if any
*/
func Rewards(network string, nodeID *big.Int, proofStrings []string) (*big.Int, error) {
	var rewards *big.Int
	contract, err := csFeeDistributorContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csFeeDistributorContract: %w", err)
	}

	treeCID, err := treeCID(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call treeCID: %w", err)
	}

	shares, err := cumulativeFeeShares(treeCID, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call cumulativeFeeShares: %w", err)
	}

	proof, err := convertToByte(proofStrings)
	if err != nil {
		return nil, fmt.Errorf("failed to call convToByte: %w", err)
	}

	fees, err := contract.GetFeesToDistribute(nil, nodeID, shares, proof)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetFeesToDistribute: %w", err)
	}

	bondInfo, err := bond.BondSummary(network, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call BondSummary: %w", err)
	}

	rewards = new(big.Int).Add(bondInfo.Excess, fees)
	return rewards, nil
}

func cumulativeFeeShares(treeCID string, nodeID *big.Int) (*big.Int, error) {
	// Get tree data from IPFS
	treeData, err := treeData(treeCID)
	if err != nil {
		return nil, fmt.Errorf("failed to call treeData: %v", err)
	}

	// Compare nodeOperatorID in tree with nodeId to get shares
	// Binary search for the nodeOperatorId
	low, high := 0, len(treeData.Values)-1
	for low <= high {
		mid := (low + high) / 2
		nodeOperatorId, err := convertToBigInt(treeData.Values[mid].Value[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert nodeOperatorId: %v", err)
		}
		cmp := nodeOperatorId.Cmp(nodeID)
		if cmp == 0 {
			// Node operator ID matches, return the shares
			shares, err := convertToBigInt(treeData.Values[mid].Value[1])
			if err != nil {
				return nil, fmt.Errorf("failed to convert shares: %v", err)
			}
			return shares, nil
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil, fmt.Errorf("invalid nodeId")
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
	resp, err := http.Get(gatewayURL)
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

func convertToBigInt(value interface{}) (*big.Int, error) {
	bigIntValue := new(big.Int)
	bigIntValue.SetString(fmt.Sprintf("%.0f", value), 10)
	return bigIntValue, nil
}

func convertToByte(proofStrings []string) ([][32]byte, error) {
	var proofChunks [][32]byte
	for _, hexStr := range proofStrings {
		// Remove the "0x" prefix
		if len(hexStr) >= 2 && hexStr[:2] == "0x" {
			hexStr = hexStr[2:]
		}

		// Decode the hex string to a byte slice
		bytes, err := hex.DecodeString(hexStr)
		if err != nil {
			return proofChunks, fmt.Errorf("failed to decode hex string: %v", err)
		}

		// Ensure the byte slice is exactly 32 bytes long
		if len(bytes) != 32 {
			return proofChunks, fmt.Errorf("decoded byte slice is not 32 bytes long: %v", bytes)
		}

		// Convert the byte slice to a [32]byte array
		var chunk [32]byte
		copy(chunk[:], bytes)
		proofChunks = append(proofChunks, chunk)
	}
	return proofChunks, nil
}

func csFeeDistributorContract(network string) (*Csfeedistributor, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call ConnectContract: %w", err)
	}
	defer client.Close()

	address := common.HexToAddress(deployedContractAddresses[network])
	contract, err := NewCsfeedistributor(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSAccounting instance: %w", err)
	}
	return contract, nil
}
