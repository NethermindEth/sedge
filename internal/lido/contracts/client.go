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
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
)

func connectToRPCETH(RPCs []string, network string) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	for _, url := range RPCs {
		client, err = ethclient.Dial(url)
		if err == nil {
			// Try to get the chain ID, which is a basic operation that should work for any Ethereum client
			chainID, err := client.ChainID(context.Background())
			if err == nil {
				expectedChainID := configs.NetworksConfigs()[network].ChainID
				if chainID.Cmp(new(big.Int).SetUint64(expectedChainID)) == 0 {
					// If we successfully got the chain ID and it matches the expected one,
					// we can be reasonably sure this is the correct Ethereum client
					log.Debugf("connected to %s", url)
					return client, nil
				} else {
					log.Errorf("chain ID mismatch: expected %d, got %d", expectedChainID, chainID.Uint64())
				}
			}
			// If there was an error or chain ID mismatch, close the client and continue to the next URL
			client.Close()
		}
	}

	return nil, fmt.Errorf("failed to connect to any RPC URL, either the RPCs are wrong or the network is not supported")
}

// ConnectClient returns a new Ethereum client connected to the given network.
// If websocket is true, it will try to connect to the first available WS RPC.
// If websocket is false, it will try to connect to the first available HTTP RPC.
// If no RPCs are provided, it will use the public RPCs or WSs for the given network.
// The RPCs are shuffled to avoid the same RPC being used by multiple clients.
func ConnectClient(network string, websocket bool, RPCs ...string) (*ethclient.Client, error) {
	var rpcs []string
	var err error

	for _, rpc := range RPCs {
		if websocket {
			if strings.HasPrefix(rpc, "wss://") {
				rpcs = append(rpcs, rpc)
			}
		} else {
			if strings.HasPrefix(rpc, "https://") || strings.HasPrefix(rpc, "http://") {
				rpcs = append(rpcs, rpc)
			}
		}
	}

	if len(rpcs) == 0 {
		if websocket {
			rpcs, err = configs.GetPublicWSs(network)
		} else {
			rpcs, err = configs.GetPublicRPCs(network)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get public RPC: %w", err)
		}
	}

	client, err := connectToRPCETH(rpcs, network)
	if err != nil {
		return nil, err
	}

	return client, nil
}
