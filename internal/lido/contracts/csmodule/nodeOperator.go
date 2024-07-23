package csmodule

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/ethereum/go-ethereum/common"
)

var deployedContractAddresses = map[string]string{
	configs.NetworkHolesky: "0x4562c3e63c2e586cD1651B958C22F88135aCAd4f",
}

func csModuleContract(network string) (*CSModule, error) {
	client, err := contracts.ConnectContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call ConnectContract: %w", err)
	}
	defer client.Close()
	address := common.HexToAddress(deployedContractAddresses[network])
	contract, err := NewCSModule(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSModule instance: %w", err)
	}
	return contract, nil
}

func NodeID(network string, rewardAddress string) (*big.Int, error) {
	nodeOperatorIDs, err := nodeOpIDs(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call NodeOpIDs: %w", err)
	}
	rewardAddr := common.HexToAddress(rewardAddress)

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
