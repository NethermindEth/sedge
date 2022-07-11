package configs

var JWTNetworks map[string]bool

var NetworksToServices map[string]string

func init() {
	JWTNetworks = map[string]bool{
		"mainnet": false,
		"kiln":    true,
		"ropsten": true,
	}

	NetworksToServices = map[string]string{
		"mainnet": "mainnet",
		"kiln":    "merge",
		"ropsten": "merge",
	}
}
