package configs

var JWTNetworks map[string]bool

func init() {
	JWTNetworks = map[string]bool{
		"mainnet": false,
		"kiln":    true,
	}
}
