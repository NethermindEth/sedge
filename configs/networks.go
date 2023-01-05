package configs

import (
	"errors"
	"fmt"
	"time"
)

const (
	// Network names
	NetworkMainnet = "mainnet"
	NetworkGoerli  = "goerli"
	NetworkSepolia = "sepolia"
	NetworkGnosis  = "gnosis"
	NetworkChiado  = "chiado"
	NetworkCustom  = "custom"
)

var ErrInvalidNetwork = errors.New("invalid network")

func NetworkCheck(value string) error {
	switch value {
	case NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrInvalidNetwork, value)
	}
}

func NetworkSupported() []string {
	return []string{
		NetworkMainnet,
		NetworkGoerli,
		NetworkSepolia,
		NetworkGnosis,
		NetworkChiado,
		NetworkCustom,
	}
}

func NetworkEpochTime(network string) time.Duration {
	switch network {
	case NetworkMainnet, NetworkGoerli, NetworkSepolia:
		return 7 * time.Minute
	case NetworkGnosis, NetworkChiado:
		return 2 * time.Minute
	default:
		return 7 * time.Minute
	}
}
