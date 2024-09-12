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
package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/NethermindEth/sedge/configs"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connectToRPCETH(RPCs []string) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	for _, url := range RPCs {
		client, err = ethclient.Dial(url)
		if err == nil {
			return client, nil
		}
	}

	return nil, fmt.Errorf("failed to connect to any RPC URL")
}

func ConnectClient(network string, RPCs ...string) (*ethclient.Client, error) {
	var rpcs []string
	var err error

	if len(RPCs) == 0 {
		rpcs, err = configs.GetPublicRPCs(network)
		if err != nil {
			return nil, fmt.Errorf("failed to get public RPC: %w", err)
		}
	} else {
		rpcs = RPCs
	}

	client, err := connectToRPCETH(rpcs)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	// Verify that the client is indeed an Ethereum RPC client
	if client != nil {
		// Try to get the chain ID, which is a basic operation that should work for any Ethereum client
		chainID, err := client.ChainID(context.Background())
		if err == nil {
			expectedChainID := configs.NetworksConfigs()[network].ChainID
			if chainID.Cmp(new(big.Int).SetUint64(expectedChainID)) == 0 {
				// If we successfully got the chain ID and it matches the expected one,
				// we can be reasonably sure this is the correct Ethereum client
				return client, nil
			}
		}
		// If there was an error or chain ID mismatch, close the client and continue to the next URL
		client.Close()
	}

	return nil, fmt.Errorf("failed to connect to RPC: %w", err)
}
