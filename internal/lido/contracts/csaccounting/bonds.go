package csaccounting

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/ethereum/go-ethereum/common"
)

// BondInfo : Struct represent bond info of Node Operator
type BondInfo struct {
	Current  *big.Int
	Required *big.Int
	Excess   *big.Int
	Missed   *big.Int
}

var deployedContractAddresses = map[string]string{
	configs.NetworkHolesky: "0xc093e53e8F4b55A223c18A2Da6fA00e60DD5EFE1",
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
	contract, err := csAccountingContract(network)
	if err != nil {
		return bondsInfo, fmt.Errorf("failed to call csAccountingContract: %w", err)
	}

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

/*
BondShares :
This function is responsible for:
retrieving bond shares for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. *big.Int
Struct that include keys status
b. error
Error if any
*/
func BondShares(network string, nodeID *big.Int) (*big.Int, error) {
	var shares *big.Int
	contract, err := csAccountingContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csAccountingContract: %w", err)
	}

	shares, err = contract.GetBondShares(nil, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetBondShares: %w", err)
	}
	return shares, nil
}

func csAccountingContract(network string) (*CSAccounting, error) {
	client, err := contracts.ConnectClient(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call ConnectContract: %w", err)
	}
	defer client.Close()

	address := common.HexToAddress(deployedContractAddresses[network])
	contract, err := NewCSAccounting(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSAccounting instance: %w", err)
	}
	return contract, nil
}
