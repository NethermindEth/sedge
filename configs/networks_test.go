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
			name:    "Valid network, goerli",
			network: "goerli",
			wantErr: false,
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
