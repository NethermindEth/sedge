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
	"fmt"

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

func ConnectClient(network string) (*ethclient.Client, error) {
	rpcs, err := configs.GetPublicRPCs(network)
	if err != nil {
		return nil, fmt.Errorf("failed to get public RPC: %w", err)
	}

	// Connect to the RPC endpoint
	client, err := connectToRPCETH(rpcs)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	return client, nil
}
