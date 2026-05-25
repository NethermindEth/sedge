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
package configs

import (
	"errors"
	"fmt"
	"sort"
)

const (
	ArbitrumChainOne     = "arbitrum-one"
	ArbitrumChainSepolia = "arbitrum-sepolia"
)

var (
	ErrInvalidArbitrumChain    = errors.New("invalid arbitrum chain")
	ErrArbitrumChainL1Mismatch = errors.New("arbitrum chain does not match selected L1 network")
	ErrArbitrumNoSelector      = errors.New("must provide --network and/or --chain for arbitrum")
)

type ArbitrumChainConfig struct {
	NethermindChainspec string // --config=<value> for nethermind-arbitrum
	ChainID             uint64 // --chain.id=<value> for nitro
	L1Network           string // sedge network name of the parent chain
	DefaultSnapshotURL  string // optional --init.url=<value>; empty -> --init.empty=true
}

var arbitrumChains = map[string]ArbitrumChainConfig{
	ArbitrumChainOne: {
		NethermindChainspec: "arbitrum-mainnet",
		ChainID:             42161,
		L1Network:           NetworkMainnet,
		DefaultSnapshotURL:  "",
	},
	ArbitrumChainSepolia: {
		NethermindChainspec: "arbitrum-sepolia",
		ChainID:             421614,
		L1Network:           NetworkSepolia,
		DefaultSnapshotURL:  "",
	},
}

func ArbitrumChain(name string) (ArbitrumChainConfig, error) {
	cfg, ok := arbitrumChains[name]
	if !ok {
		return ArbitrumChainConfig{}, fmt.Errorf("%w: %s", ErrInvalidArbitrumChain, name)
	}
	return cfg, nil
}

func ArbitrumSupportedChains() []string {
	out := make([]string, 0, len(arbitrumChains))
	for k := range arbitrumChains {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// ResolveArbitrumChainAndL1 applies the bidirectional inference rules from the design.
// At least one of network/chain must be non-empty.
func ResolveArbitrumChainAndL1(network, chain string) (resolvedNetwork, resolvedChain string, err error) {
	switch {
	case network == "" && chain == "":
		return "", "", ErrArbitrumNoSelector
	case chain == "":
		for name, cfg := range arbitrumChains {
			if cfg.L1Network == network {
				return network, name, nil
			}
		}
		return "", "", fmt.Errorf("%w: no arbitrum chain registered for L1 %q", ErrInvalidArbitrumChain, network)
	case network == "":
		cfg, err := ArbitrumChain(chain)
		if err != nil {
			return "", "", err
		}
		return cfg.L1Network, chain, nil
	default:
		cfg, err := ArbitrumChain(chain)
		if err != nil {
			return "", "", err
		}
		if cfg.L1Network != network {
			return "", "", fmt.Errorf("%w: chain %q expects L1 %q, got %q",
				ErrArbitrumChainL1Mismatch, chain, cfg.L1Network, network)
		}
		return network, chain, nil
	}
}
