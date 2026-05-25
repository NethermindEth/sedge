package configs

import (
	"errors"
	"testing"
)

func TestArbitrumChainConfig(t *testing.T) {
	cfg, err := ArbitrumChain("arbitrum-one")
	if err != nil {
		t.Fatalf("arbitrum-one: unexpected error %v", err)
	}
	if cfg.ChainID != 42161 {
		t.Errorf("arbitrum-one ChainID: want 42161, got %d", cfg.ChainID)
	}
	if cfg.NethermindChainspec != "arbitrum-mainnet" {
		t.Errorf("arbitrum-one NethermindChainspec: want arbitrum-mainnet, got %q", cfg.NethermindChainspec)
	}
	if cfg.L1Network != NetworkMainnet {
		t.Errorf("arbitrum-one L1Network: want %q, got %q", NetworkMainnet, cfg.L1Network)
	}

	cfg, err = ArbitrumChain("arbitrum-sepolia")
	if err != nil {
		t.Fatalf("arbitrum-sepolia: unexpected error %v", err)
	}
	if cfg.ChainID != 421614 {
		t.Errorf("arbitrum-sepolia ChainID: want 421614, got %d", cfg.ChainID)
	}
	if cfg.L1Network != NetworkSepolia {
		t.Errorf("arbitrum-sepolia L1Network: want %q, got %q", NetworkSepolia, cfg.L1Network)
	}

	if _, err := ArbitrumChain("does-not-exist"); !errors.Is(err, ErrInvalidArbitrumChain) {
		t.Errorf("unknown chain: want ErrInvalidArbitrumChain, got %v", err)
	}
}

func TestArbitrumSupportedChains(t *testing.T) {
	chains := ArbitrumSupportedChains()
	if len(chains) != 2 {
		t.Fatalf("want 2 chains, got %d (%v)", len(chains), chains)
	}
	wantSet := map[string]bool{"arbitrum-one": false, "arbitrum-sepolia": false}
	for _, c := range chains {
		if _, ok := wantSet[c]; !ok {
			t.Errorf("unexpected chain %q", c)
		}
		wantSet[c] = true
	}
	for c, seen := range wantSet {
		if !seen {
			t.Errorf("missing chain %q in supported list", c)
		}
	}
}

func TestResolveArbitrumChainAndL1(t *testing.T) {
	tests := []struct {
		name      string
		network   string
		chain     string
		wantNet   string
		wantChain string
		wantErr   bool
	}{
		{"L1 mainnet only -> default arbitrum-one", "mainnet", "", "mainnet", "arbitrum-one", false},
		{"L1 sepolia only -> default arbitrum-sepolia", "sepolia", "", "sepolia", "arbitrum-sepolia", false},
		{"chain only -> derive L1", "", "arbitrum-sepolia", "sepolia", "arbitrum-sepolia", false},
		{"both consistent", "mainnet", "arbitrum-one", "mainnet", "arbitrum-one", false},
		{"both inconsistent", "mainnet", "arbitrum-sepolia", "", "", true},
		{"unknown chain", "mainnet", "arbitrum-nova", "", "", true},
		{"unsupported L1 holesky", "holesky", "", "", "", true},
		{"neither given", "", "", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			net, chain, err := ResolveArbitrumChainAndL1(tt.network, tt.chain)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if net != tt.wantNet || chain != tt.wantChain {
				t.Errorf("got (%q,%q), want (%q,%q)", net, chain, tt.wantNet, tt.wantChain)
			}
		})
	}
}
