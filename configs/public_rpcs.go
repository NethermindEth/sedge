package configs

type RPC struct {
	NetworkName string
	PublicRPCs  []string
}

var networkRPCs = map[string]RPC{
	NetworkMainnet: {
		NetworkName: NetworkMainnet,
		PublicRPCs: []string{
			"https://eth.llamarpc.com",
			"https://eth-pokt.nodies.app",
			"https://rpc.mevblocker.io",
			"https://ethereum-rpc.publicnode.com",
			"https://rpc.flashbots.net",
		},
	},
	NetworkHolesky: {
		NetworkName: NetworkHolesky,
		PublicRPCs: []string{
			"https://1rpc.io/holesky",
			"https://holesky.drpc.org",
			"https://ethereum-holesky-rpc.publicnode.com",
			"https://endpoints.omniatech.io/v1/eth/holesky/public",
			"https://ethereum-holesky.blockpi.network/v1/rpc/public",
		},
	},
}
