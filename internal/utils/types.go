package utils

// DistroInfo : Struct Contains name, architecture and version of the linux distribution of the host machine
type DistroInfo struct {
	Name         string
	Architecture string
	Version      string
}

// DepositCLI : Struct Data for eth2.0-deposit-cli command template
type DepositCLI struct {
	// Network : Network name, e.g. "mainnet", "prater"
	Network string
	// Path : Path to the keystore folder
	Path string
}
