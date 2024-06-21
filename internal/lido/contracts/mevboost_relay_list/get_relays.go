package mevboostrelaylist

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
)

// Relay : Struct represent data of MEV-Boost Relay
type Relay struct {
	Uri         string `yaml:"Uri"`
	Operator    string `yaml:"Operator"`
	IsMandatory bool   `yaml:"IsMandatory"`
	Description string `yaml:"Description"`
}

// Define deployed contract addresses for Mainnet and Holesky
var deployedContractAddresses = map[string]string{
	configs.NetworkMainnet: "0xF95f069F9AD107938F6ba802a3da87892298610E",
	configs.NetworkHolesky: "0x2d86C5855581194a386941806E38cA119E50aEA3",
}

/*
connectToRPC :
This function is responsible for:
attempting to connect to the first available RPC endpoint from the provided list.
params :-
RPCs ([]string): A list of RPC endpoint URLs.
returns :-
a. *rpc.Client
The connected RPC client
b. error
Error if no connections could be established
*/
func connectToRPC(RPCs []string) (*rpc.Client, error) {
	var client *rpc.Client
	var err error

	for _, url := range RPCs {
		client, err = rpc.DialHTTP(url)
		if err == nil {
			return client, nil
		}
	}

	return nil, fmt.Errorf("failed to connect to any RPC URL")
}

/*
GetRelays :
This function is responsible for :-
retrieving a list of relays from the MEV-Boost Allowed List contract for a given network.
params :-
network (string): The name of the network (e.g., "mainnet", "holesky").
returns :-
a. []Relay
List of relays
b. error
Error if any
*/
func GetRelays(network string) ([]Relay, error) {
	var relays []Relay
	rpcs, err := configs.GetPublicRPCs(network)
	if err != nil {
		return relays, fmt.Errorf("failed to get public RPC: %w", err)
	}

	// Connect to the RPC endpoint
	client, err := connectToRPC(rpcs)
	if err != nil {
		return relays, fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	// Parse the ABI of the contract
	parsedABI, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return relays, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Pack the data for the "get_relays" method
	data, err := parsedABI.Pack("get_relays")
	if err != nil {
		return relays, fmt.Errorf("failed to pack ABI data: %w", err)
	}

	// Prepare the RPC call arguments
	type CallArgs struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}
	args := CallArgs{
		To:   deployedContractAddresses[network],
		Data: "0x" + hex.EncodeToString(data),
	}

	// Execute the RPC call
	var result string
	if err := client.Call(&result, "eth_call", args, "latest"); err != nil {
		return relays, fmt.Errorf("failed to make RPC call: %w", err)
	}

	// Decode the result from the RPC call
	output, err := hex.DecodeString(result[2:]) // Remove the '0x' prefix
	if err != nil {
		return relays, fmt.Errorf("failed to decode result hex: %w", err)
	}

	// Unpack the result into the relays slice
	err = parsedABI.UnpackIntoInterface(&relays, "get_relays", output)
	if err != nil {
		return relays, fmt.Errorf("failed to unpack ABI output: %w", err)
	}

	return relays, nil
}
