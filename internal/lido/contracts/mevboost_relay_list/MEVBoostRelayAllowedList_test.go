package mevboostrelaylist

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
)

// Network configuration structs
type NetworkConfig struct {
	ContractAddress string
	RPCs            []string
}

// Define configurations for mainnet and Holesky
var mainnetConfig = NetworkConfig{
	ContractAddress: "0xF95f069F9AD107938F6ba802a3da87892298610E",
	RPCs: []string{
		"https://eth.llamarpc.com",
		"https://eth-pokt.nodies.app",
		"https://rpc.mevblocker.io",
		"https://ethereum-rpc.publicnode.com",
		"https://rpc.flashbots.net",
	},
}

var holeskyConfig = NetworkConfig{
	ContractAddress: "0x2d86C5855581194a386941806E38cA119E50aEA3",
	RPCs: []string{
		"https://1rpc.io/holesky",
		"https://ethereum-holesky-rpc.publicnode.com",
		"https://endpoints.omniatech.io/v1/eth/holesky/public",
	},
}

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

func getNetworkConfig(network string) (*NetworkConfig, error) {
	switch network {
	case "mainnet":
		return &mainnetConfig, nil
	case "holesky":
		return &holeskyConfig, nil
	default:
		return nil, fmt.Errorf("Unsupported network: %s", network)
	}
}

func GetRelays(network string) ([]Struct0, error) {
	var relays []Struct0

	config, err := getNetworkConfig(network)
	if err != nil {
		return relays, fmt.Errorf("Failed to get network config: %w", err)
	}
	client, err := connectToRPC(config.RPCs)
	if err != nil {
		return relays, fmt.Errorf("Failed to connect to RPC: %w", err)
	}
	defer client.Close()

	parsedABI, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return relays, fmt.Errorf("Failed to parse ABI: %w", err)
	}
	data, err := parsedABI.Pack("get_relays")
	if err != nil {
		return relays, fmt.Errorf("Failed to pack ABI data: %w", err)
	}

	type CallArgs struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}
	args := CallArgs{
		To:   config.ContractAddress,
		Data: "0x" + hex.EncodeToString(data),
	}

	var result string
	if err := client.Call(&result, "eth_call", args, "latest"); err != nil {
		return relays, fmt.Errorf("Failed to make RPC call: %w", err)
	}

	output, err := hex.DecodeString(result[2:]) // Remove the '0x' prefix
	if err != nil {
		return relays, fmt.Errorf("Failed to decode result hex: %w", err)
	}
	err = parsedABI.UnpackIntoInterface(&relays, "get_relays", output)
	if err != nil {
		return relays, fmt.Errorf("Failed to unpack ABI output: %w", err)
	}

	return relays, nil
}

func TestGetRelays(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tcs := []struct {
		name           string
		network        string
		expectedRelays []Struct0
	}{
		{
			"GetRelays Mainnet", "mainnet", []Struct0{
				{"https://0xa7ab7a996c8584251c8f925da3170bdfd6ebc75d50f5ddc4050a6fdc77f2a3b5fce2cc750d0865e05d7228af97d69561@agnostic-relay.net", "Agnostic", true, "Agnostic Relay"},
				{"https://0xb0b07cd0abef743db4260b0ed50619cf6ad4d82064cb4fbec9d3ec530f7c5e6793d9f286c4e082c0244ffb9f2658fe88@bloxroute.regulated.blxrbdn.com", "bloXroute", true, "bloXroute Regulated Relay"},
				{"https://0xa15b52576bcbf1072f4a011c0f99f9fb6c66f3e1ff321f11f461d15e31b1cb359caa092c71bbded0bae5b5ea401aab7e@aestus.live", "Aestus", true, "Aestus Relay"},
				{"https://0x8b5d2e73e2a3a55c6c87b8b6eb92e0149a125c852751db1422fa951e42a09b82c142c3ea98d0d9930b056a3bc9896b8f@bloxroute.max-profit.blxrbdn.com", "bloXroute", true, "bloXroute Max-Profit Relay"},
				{"https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net", "Flashbots", true, "Flashbots Relay"},
				{"https://0xb3ee7afcf27f1f1259ac1787876318c6584ee353097a50ed84f51a1f21a323b3736f271a895c7ce918c038e4265918be@relay.edennetwork.io", "Eden Network", true, "Eden Network Relay"},
				{"https://0x98650451ba02064f7b000f5768cf0cf4d4e492317d82871bdc87ef841a0743f69f0f1eea11168503240ac35d101c9135@mainnet-relay.securerpc.com/", "Manifold Finance", false, "Manifold SecureRPC Relay"},
				{"https://0xa1559ace749633b997cb3fdacffb890aeebdb0f5a3b6aaa7eeeaf1a38af0a8fe88b9e4b1f61f236d2e64d95733327a62@relay.ultrasound.money", "Ultra Sound", true, "Ultra Sound Relay"},
			},
		},
		{
			"GetRelays Holesky", "holesky", []Struct0{
				{"https://0xb1559beef7b5ba3127485bbbb090362d9f497ba64e177ee2c8e7db74746306efad687f2cf8574e38d70067d40ef136dc@relay-stag.ultrasound.money", "Ultra Sound", true, "Ultra Sound Relay Holesky - no filtering"},
				{"https://0xab78bf8c781c58078c3beb5710c57940874dd96aef2835e7742c866b4c7c0406754376c2c8285a36c630346aa5c5f833@holesky.aestus.live", "Aestus", true, "Aestus Relay Holesky - no filtering"},
				{"http://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@testnet.relay-proxy.blxrbdn.com:18552/", "bloxRoute", false, "bloxRoute Validator Gateway - filtering"},
				{"https://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@bloxroute.holesky.blxrbdn.com", "bloxRoute", true, "bloxRoute Relay - filtering"},
				{"https://0x833b55e20769a8a99549a28588564468423c77724a0ca96cffd58e65f69a39599d877f02dc77a0f6f9cda2a3a4765e56@relay-holesky.beaverbuild.org", "Beaverbuild", false, "Beaverbuild Relay Holesky"},
				{"https://0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12@relay-holesky.edennetwork.io", "Eden", true, "Eden Relay Holesky - Filtering"},
				{"https://0xaa58208899c6105603b74396734a6263cc7d947f444f396a90f7b7d3e65d102aec7e5e5291b27e08d02c50a050825c2f@holesky.titanrelay.xyz/", "Titan", true, "Titan Relay Holesky - nonfiltering"},
				{"https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@boost-relay-holesky.flashbots.net", "Flashbots", true, "Flashbots Boost Holesky - filtering"},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			relays, err := GetRelays(tc.network)
			if err != nil {
				t.Fatalf("Failed to call GetRelays %v", err)
			}

			if !reflect.DeepEqual(relays, tc.expectedRelays) {
				t.Errorf("Relays do not match expected values\nGot: %+v\nExpected: %+v", relays, tc.expectedRelays)
			}
		})
	}
}
