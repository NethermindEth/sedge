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
	//"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// Keys : Struct represent keys status of Node Operator
type Keys struct {
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	ExitedValidators           *big.Int
	DepositedValidators        *big.Int
	DepositableValidatorsCount *big.Int
}

/*
KeysStatus :
This function is responsible for:
retrieving keys status for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. Keys
Struct that include keys status
b. error
Error if any
*/
func KeysStatus(network string, nodeID *big.Int) (Keys, error) {
	var keys Keys
	contract, err := csModuleContract(network)
	if err != nil {
		return keys, fmt.Errorf("failed to call csModuleContract: %w", err)
	}

	nodeOp, err := contract.GetNodeOperatorSummary(nil, nodeID)
	if err != nil {
		return keys, fmt.Errorf("failed to call GetNodeOperator: %w", err)
	}
	keys.DepositableValidatorsCount = nodeOp.DepositableValidatorsCount
	keys.DepositedValidators = nodeOp.TotalDepositedValidators
	keys.ExitedValidators = nodeOp.TotalExitedValidators
	keys.RefundedValidatorsCount = nodeOp.RefundedValidatorsCount
	keys.StuckValidatorsCount = nodeOp.StuckValidatorsCount

	return keys, nil
}

/*
SigningKeys :
This function is responsible for:
retrieving signing keys for Lido CSM node
params :-
network (string): The name of the network (e.g."holesky").
nodeID (*big.Int): Node Operator ID
returns :-
a. []string
Slice that includes signing keys
b. error
Error if any
*/
func SigningKeys(network string, nodeID *big.Int) ([]string, error) {
	var keys []string
	contract, err := csModuleContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}
	keysCount, err := nodeOpNonWithdrawnKeys(network, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call  nodeOpNonWithdrawnKeys: %w", err)
	}
	// Retrieve keys one by one
	for i := new(big.Int); i.Cmp(keysCount) < 0; i.Add(i, big.NewInt(1)) {
		keyData, err := contract.GetSigningKeys(nil, nodeID, i, big.NewInt(1))
		if err != nil {
			return nil, fmt.Errorf("failed to call GetSigningKeys: %w", err)
		}

		// Convert the byte array to a hexadecimal string
		key := hex.EncodeToString(keyData)
		log.Print(key)
		keys = append(keys, key)
	}

	return keys, nil
}

func nodeOpNonWithdrawnKeys(network string, nodeID *big.Int) (*big.Int, error) {
	var keysCount *big.Int
	contract, err := csModuleContract(network)
	if err != nil {
		return nil, fmt.Errorf("failed to call csModuleContract: %w", err)
	}
	keysCount, err = contract.GetNodeOperatorNonWithdrawnKeys(nil, nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetNodeOperatorNonWithdrawnKeys: %w", err)
	}
	return keysCount, nil
}
