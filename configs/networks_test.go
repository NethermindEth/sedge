package configs

import "testing"

func TestNetworkCheck(t *testing.T) {
	tests := []struct {
		name    string
		network string
		wantErr bool
	}{
		{
			name:    "Valid network, mainnet",
			network: "mainnet",
			wantErr: false,
		},
		{
			name:    "Invalid network, goerli",
			network: "goerli",
			wantErr: true,
		},
		{
			name:    "Valid network, sepolia",
			network: "sepolia",
			wantErr: false,
		},
		{
			name:    "Valid network, gnosis",
			network: "gnosis",
			wantErr: false,
		},
		{
			name:    "Valid network, chiado",
			network: "chiado",
			wantErr: false,
		},
		{
			name:    "Valid network, custom",
			network: "custom",
			wantErr: false,
		},
		{
			name:    "Invalid network",
			network: "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NetworkCheck(tt.network); (err != nil) != tt.wantErr {
				t.Errorf("NetworkCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSupportMEVBoost(t *testing.T) {
	tests := []struct {
		name    string
		network string
		want    bool
	}{
		{
			name:    "Valid network, mainnet",
			network: "mainnet",
			want:    true,
		},
		{
			name:    "Invalid network, goerli",
			network: "goerli",
			want:    false,
		},
		{
			name:    "Valid network, sepolia",
			network: "sepolia",
			want:    true,
		},
		{
			name:    "Valid network, holesky",
			network: "holesky",
			want:    true,
		},
		{
			name:    "Valid network, gnosis",
			network: "gnosis",
			want:    false,
		},
		{
			name:    "Valid network, chiado",
			network: "chiado",
			want:    false,
		},
		{
			name:    "Valid network, custom",
			network: "custom",
			want:    false,
		},
		{
			name:    "Invalid network",
			network: "invalid",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SupportsMEVBoost(tt.network); got != tt.want {
				t.Errorf("SupportsMEVBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPublicRPCs_Randomness(t *testing.T) {
	networkRPCs = map[string]RPC{
		"testnet": {PublicRPCs: []string{"rpc1", "rpc2", "rpc3", "rpc4", "rpc5"}},
	}

	// Run multiple times to check randomness
	iterations := 100
	results := make([]string, iterations)

	for i := 0; i < iterations; i++ {
		got, err := GetPublicRPCs("testnet")
		if err != nil {
			t.Fatalf("GetPublicRPCs() error = %v", err)
		}
		results[i] = got[0] // Store the first RPC of each result
	}

	// Check if we have different first elements (indicating randomness)
	uniqueFirstElements := make(map[string]bool)
	for _, r := range results {
		uniqueFirstElements[r] = true
	}

	if len(uniqueFirstElements) < 2 {
		t.Errorf("GetPublicRPCs() doesn't seem to randomize the order. Got %d unique first elements out of %d iterations", len(uniqueFirstElements), iterations)
	}
}
