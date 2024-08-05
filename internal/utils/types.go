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
package utils

// DistroInfo : Struct Contains name, architecture and version of the linux distribution of the host machine
type DistroInfo struct {
	Name         string
	Architecture string
	Version      string
}

// DepositCLI : Struct Data for eth2.0-deposit-cli command template
type DepositCLI struct {
	// Network : Network name, e.g. "mainnet", "sepolia"
	Network string
	// Path : Path to the keystore folder
	Path string
	// Password : Keystore password
	Password string
	// LenPass : Slice used to hide password in template
	LenPass []struct{}
	// Image : staking-deposit-cli docker image
	Image string
	// Eth1WithdrawalAddress : Address to be used to create withdrawal credentials
	Eth1WithdrawalAddress string
}

// ValidatorKeyData : Struct Data for Keystore generation
type ValidatorKeyData struct {
	// Existing: True if a existing mnemonic is being used. False to create and use a new one.
	Existing bool
	// Network : Target network
	Network string
	// Path : Path to the keystore folder
	Path string
	// Password : Keystore password
	Password string
	// Eth1WithdrawalAddress : Address to be used to create withdrawal credentials
	Eth1WithdrawalAddress string
}
