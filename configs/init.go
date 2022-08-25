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
package configs

var JWTNetworks map[string]bool

var NetworksToServices map[string]string

var NetworksToCheckpointUrl map[string]string

func init() {
	JWTNetworks = map[string]bool{
		"mainnet": true,
		"kiln":    true,
		"ropsten": true,
		"prater":  true,
		"sepolia": true,
		"denver":  true,
		"chiado":  true,
	}

	NetworksToServices = map[string]string{
		"mainnet": "merge",
		"kiln":    "merge",
		"ropsten": "merge",
		"prater":  "merge",
		"sepolia": "merge",
		"denver":  "merge",
		"chiado":  "merge",
	}

	NetworksToCheckpointUrl = map[string]string{
		"mainnet": "",
		"kiln":    "",
		"ropsten": "https://ropsten.checkpoint-sync.ethdevops.io",
		"prater":  "https://goerli.checkpoint-sync.ethdevops.io",
		"sepolia": "https://sepolia.checkpoint-sync.ethdevops.io",
		"denver":  "",
		"chiado":  "",
	}
}
