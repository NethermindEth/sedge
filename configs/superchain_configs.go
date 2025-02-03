package configs

import _ "embed"

//go:embed superchain/mainnet.yaml
var superchainMainnetConfig string

//go:embed superchain/sepolia.yaml
var superchainSepoliaConfig string

// GetSuperchainConfig returns the appropriate superchain config based on the network
func GetSuperchainConfig(network string) string {
	switch network {
	case "mainnet":
		return superchainMainnetConfig
	case "sepolia":
		return superchainSepoliaConfig
	default:
		return ""
	}
}
