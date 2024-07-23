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

func ConnectContract(network string) (*ethclient.Client, error) {
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
