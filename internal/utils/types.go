package utils

import "text/template"

// DistroInfo : Struct Contains name, architecture and version of the linux distribution of the host machine
type DistroInfo struct {
	Name         string
	Architecture string
	Version      string
}

// Script : Struct Represents a script to be executed
type Script struct {
	// Tmp : script template
	Tmp *template.Template
	// getOutput : True to get output of the script
	GetOutput bool
	// Data: template data object
	Data interface{}
}

// DepositCLI : Struct Data for eth2.0-deposit-cli command template
type DepositCLI struct {
	// Network : Network name, e.g. "mainnet", "prater"
	Network string
	// Path : Path to the keystore folder
	Path string
}
