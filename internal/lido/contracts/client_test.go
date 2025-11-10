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
package contracts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectClient(t *testing.T) {
	tcs := []struct {
		name    string
		network string
		wantErr bool
	}{
		{
			name:    "ConnectClient, Holesky",
			network: "holesky",
			wantErr: false,
		},
		{
			name:    "ConnectClient, invalid Network",
			network: "invalid",
			wantErr: true,
		},
		{
			name:    "ConnectClient, Mainnet",
			network: "mainnet",
			wantErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			client, err := ConnectClient(tc.network, false)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})
	}
}

func TestConnectClientWithRPCs(t *testing.T) {
	tcs := []struct {
		name    string
		network string
		RPCs    []string
		wantErr bool
	}{
		{
			name:    "ConnectClientWithRPCs, Holesky",
			network: "holesky",
			RPCs:    []string{"https://holesky.gateway.tenderly.co", "https://eth-holesky-testnet.rpc.grove.city/v1/01fdb492"},
			wantErr: false,
		},
		{
			name:    "ConnectClientWithRPCs, Holesky, invalid RPC",
			network: "holesky",
			RPCs:    []string{"https://www.google.com"},
			wantErr: true,
		},
		{
			name:    "ConnectClientWithRPCs, invalid Network RPCs",
			network: "holesky",
			RPCs:    []string{"https://eth.llamarpc.com"}, // Mainnet RPC
			wantErr: true,
		},
		{
			name:    "ConnectClientWithRPCs, invalid Network RPCs, no HTTPS",
			network: "holesky",
			RPCs:    []string{"wss://holesky.drpc.org"}, // Mainnet RPC
			wantErr: false,
		},
		{
			name:    "ConnectClientWithRPCs, mix Network RPCs",
			network: "holesky",
			RPCs:    []string{"https://holesky.gateway.tenderly.co", "wss://ethereum-rpc.publicnode.com"},
			wantErr: false,
		},
		{
			name:    "ConnectClient, Mainnet",
			network: "mainnet",
			RPCs:    []string{"https://eth.llamarpc.com"},
			wantErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			client, err := ConnectClient(tc.network, false, tc.RPCs...)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})
	}
}

func TestConnectClientWS(t *testing.T) {
	tcs := []struct {
		name    string
		network string
		wantErr bool
	}{
		{
			name:    "ConnectClient, Holesky",
			network: "holesky",
			wantErr: false,
		},
		{
			name:    "ConnectClient, invalid Network",
			network: "invalid",
			wantErr: true,
		},
		{
			name:    "ConnectClient, Mainnet",
			network: "mainnet",
			wantErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			client, err := ConnectClient(tc.network, true)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})
	}
}

func TestConnectClientWSWithRPCs(t *testing.T) {
	tcs := []struct {
		name    string
		network string
		RPCs    []string
		wantErr bool
	}{
		{
			name:    "ConnectClientWithRPCs, Holesky",
			network: "holesky",
			RPCs:    []string{"wss://ethereum-holesky-rpc.publicnode.com", "wss://holesky.drpc.org"},
			wantErr: false,
		},
		{
			name:    "ConnectClientWithRPCs, Holesky, invalid RPC",
			network: "holesky",
			RPCs:    []string{"https://www.google.com"},
			wantErr: false,
		},
		{
			name:    "ConnectClientWithRPCs, invalid Network RPCs",
			network: "holesky",
			RPCs:    []string{"wss://ethereum-rpc.publicnode.com"}, // Mainnet RPC
			wantErr: true,
		},
		{
			name:    "ConnectClientWithRPCs, invalid Network RPCs, no WS",
			network: "holesky",
			RPCs:    []string{"https://ethereum-holesky-rpc.publicnode.com"}, // Mainnet RPC
			wantErr: false,
		},
		{
			name:    "ConnectClientWithRPCs, mix Network RPCs",
			network: "mainnet",
			RPCs:    []string{"https://ethereum-holesky-rpc.publicnode.com", "wss://ethereum-rpc.publicnode.com"},
			wantErr: false,
		},
		{
			name:    "ConnectClient, Mainnet",
			network: "mainnet",
			RPCs:    []string{"wss://ethereum-rpc.publicnode.com"},
			wantErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			client, err := ConnectClient(tc.network, true, tc.RPCs...)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})
	}
}
