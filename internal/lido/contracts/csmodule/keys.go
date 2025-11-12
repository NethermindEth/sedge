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
network (string): The name of the network (e.g."hoodi").
nodeID (*big.Int): Node Operator ID
returns :-
a. Keys
Struct that include keys status
b. error
Error if any
*/
func KeysStatus(network string, nodeID *big.Int) (Keys, error) {
	var keys Keys

	if nodeID.Sign() < 0 {
		return keys, fmt.Errorf("node ID value out-of-bounds: can't be negative")
	}
	contract, client, err := csModuleContract(network)
	if err != nil {
		return keys, fmt.Errorf("failed to call csModuleContract: %w", err)
	}
	defer client.Close()

	nodeOp, err := contract.GetNodeOperatorSummary(nil, nodeID)
	if err != nil {
		return keys, fmt.Errorf("failed to call GetNodeOperatorSummary contract method: %w", err)
	}
	keys.DepositableValidatorsCount = nodeOp.DepositableValidatorsCount
	keys.DepositedValidators = nodeOp.TotalDepositedValidators
	keys.ExitedValidators = nodeOp.TotalExitedValidators
	keys.RefundedValidatorsCount = nodeOp.RefundedValidatorsCount
	keys.StuckValidatorsCount = nodeOp.StuckValidatorsCount

	return keys, nil
}
