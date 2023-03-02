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

func TestSupportMEVBoost(t *testing.T) {
	tests := []struct {
		name    string
		network string
		want    bool
	}{
		{
			name:    "Valid network, mainnet",
			network: "mainnet",
			want:    false,
		},
		{
			name:    "Valid network, goerli",
			network: "goerli",
			want:    false,
		},
		{
			name:    "Valid network, sepolia",
			network: "sepolia",
			want:    false,
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
