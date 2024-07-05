package contracts

import (
	"sort"

	"github.com/NethermindEth/sedge/configs"
)

type WithdrawalAddressConfig struct {
	Network           string
	WithdrawalAddress string
}

var WithdrawalAddress = map[string]WithdrawalAddressConfig{
	configs.NetworkMainnet: {
		Network:           configs.NetworkMainnet,
		WithdrawalAddress: "0xb9d7934878b5fb9610b3fe8a5e441e8fad7e293f",
	},
	configs.NetworkHolesky: {
		Network:           configs.NetworkHolesky,
		WithdrawalAddress: "0xF0179dEC45a37423EAD4FaD5fCb136197872EAd9",
	},
}

func GetLidoKeysSupportedNetworks() []string {
	networks := []string{}
	for network := range WithdrawalAddress {
		networks = append(networks, network)
	}
	sort.Strings(networks)
	return networks
}

func NetworkSupportedByLidoKeys(network string) ([]string, bool) {
	supportedNetworks := GetLidoKeysSupportedNetworks()
	var supported bool
	for _, supportedNetwork := range supportedNetworks {
		if network == supportedNetwork {
			supported = true
		}
	}
	if !supported {
		return supportedNetworks, supported
	}
	return nil, supported
}
