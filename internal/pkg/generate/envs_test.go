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
package generate

import (
	"bytes"
	"io"
	"strings"
	"testing"

	clientsimages "github.com/NethermindEth/sedge/configs/images"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/stretchr/testify/assert"
)

// retrieveEnvData is a helper function that retrieves the environment data from the given reader
func retrieveEnvData(t *testing.T, reader io.Reader) map[string]string {
	envFile, err := io.ReadAll(reader)
	if err != nil {
		t.Error("unable to read .env file")
	}
	data := make(map[string]string)
	for _, line := range strings.Split(string(envFile), "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, "=")
		if len(split) != 2 {
			// ignore lines that doesn't contain "=" because are comments
			continue
		}
		data[split[0]] = split[1]
	}
	return data
}

// TestGenerateEnvFile tests the generation of the .env file
func TestGenerateEnvFile(t *testing.T) {
	// default clients images
	clientsImages, err := clientsimages.NewDefaultClientsImages()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name          string
		data          *GenData
		Error         error
		fieldsToCheck map[string]string
	}{
		{
			name: "Check ec image",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind", Image: clientsImages.Execution().Nethermind().String()},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				"EC_IMAGE_VERSION": clientsImages.Execution().Nethermind().String(),
			},
		},
		{
			name: "Check set of ec image",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "besu", Image: "custom"},
				Network:         "mainnet",
				JWTSecretPath:   "/tmp/jwt",
			},
			fieldsToCheck: map[string]string{
				"EC_IMAGE_VERSION":   "custom",
				"EC_DATA_DIR":        "./execution-data",
				"EC_JWT_SECRET_PATH": "/tmp/jwt",
			},
		},
		// { // TODO: Uncomment when the refactor CLI PR is merged and default relays can be get from the config, this way we can test the default relays
		// 	name: "Check RELAY_URLS",
		// 	data: &GenData{
		// 		ValidatorClient: &clients.Client{Name: "prysm"},
		// 		Network:         "mainnet",
		// 		Mev:             true,
		// 	},
		// 	fieldsToCheck: map[string]string{
		// 		"RELAY_URLS": "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net",
		// 	},
		// },
		{
			name: "Check set of RELAY_URLS, mainnet",
			data: &GenData{
				Services:        []string{consensus, validator, execution, mevBoost},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
				RelayURLs:       []string{"https://sample.relay", "https://sample.relay2"},
			},
			fieldsToCheck: map[string]string{
				"RELAY_URLS": "https://sample.relay,https://sample.relay2",
			},
		},
		{
			name: "Check set of RELAY_URLS, goerli",
			data: &GenData{
				Services:        []string{consensus, validator, execution, mevBoost},
				ValidatorClient: &clients.Client{Name: "prysm"},
				Network:         "goerli",
				Mev:             true,
				RelayURLs:       []string{"https://sample.relay", "https://sample.relay2"},
			},
			fieldsToCheck: map[string]string{
				"RELAY_URLS": "https://sample.relay,https://sample.relay2",
			},
		},
		{
			name: "Check set of RELAY_URLS, sepolia",
			data: &GenData{
				Services:        []string{consensus, validator, execution, mevBoost},
				ValidatorClient: &clients.Client{Name: "lodestar"},
				Network:         "sepolia",
				Mev:             true,
				RelayURLs:       []string{"https://sample.relay", "https://sample.relay2"},
			},
			fieldsToCheck: map[string]string{
				"RELAY_URLS": "https://sample.relay,https://sample.relay2",
			},
		},
		{
			name: "Check Graffiti, no graffiti set, name from execution-validator",
			data: &GenData{
				ValidatorClient: &clients.Client{Name: "prysm"},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				"GRAFFITI":     "nethermind-prysm",
				"KEYSTORE_DIR": "./keystore",
			},
		},
		{
			name: "Check Graffiti, no graffiti set, name from execution-consensus, consensus=validator",
			data: &GenData{
				ValidatorClient: &clients.Client{Name: "prysm"},
				ConsensusClient: &clients.Client{Name: "prysm"},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				"GRAFFITI":     "nethermind-prysm",
				"KEYSTORE_DIR": "./keystore",
			},
		},
		{
			name: "Check Graffiti, no graffiti set, name from execution-consensus, consensus!=validator",
			data: &GenData{
				ValidatorClient: &clients.Client{Name: "lodestar"},
				ConsensusClient: &clients.Client{Name: "prysm"},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				"GRAFFITI":     "nethermind-prysm-lodestar",
				"KEYSTORE_DIR": "./keystore",
			},
		},
		{
			name: "Check validator Graffiti and keystore",
			data: &GenData{
				ValidatorClient: &clients.Client{Name: "prysm"},
				Network:         "mainnet",
				Graffiti:        "graffiti",
			},
			fieldsToCheck: map[string]string{
				"GRAFFITI":     "graffiti",
				"KEYSTORE_DIR": "./keystore",
			},
		},
		{
			name: "Check wrong network",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "erigon"},
				Network:         "wrong",
			},
			Error: ErrTemplateNotFound,
		},
		{
			name: "Check network",
			data: &GenData{
				Network:         "sepolia",
				ConsensusClient: &clients.Client{Name: "teku"},
			},
			fieldsToCheck: map[string]string{
				"NETWORK": "sepolia",
			},
		},
		{
			name: "Check gnosis network",
			data: &GenData{
				Network:         "gnosis",
				ConsensusClient: &clients.Client{Name: "teku"},
			},
			fieldsToCheck: map[string]string{
				"EL_NETWORK": "xdai",
				"CL_NETWORK": "gnosis",
			},
		},
		{
			name: "Check RELAY_URLS is set if mev is set",
			data: &GenData{
				Network:         "mainnet",
				Services:        []string{consensus, validator, execution, mevBoost},
				ConsensusClient: &clients.Client{Name: "teku"},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ValidatorClient: &clients.Client{Name: "prysm"},
				Mev:             true,
			},
			fieldsToCheck: map[string]string{
				"RELAY_URLS": "https://0xa7ab7a996c8584251c8f925da3170bdfd6ebc75d50f5ddc4050a6fdc77f2a3b5fce2cc750d0865e05d7228af97d69561@agnostic-relay.net,https://0x9000009807ed12c1f08bf4e81c6da3ba8e3fc3d953898ce0102433094e5f22f21102ec057841fcb81978ed1ea0fa8246@builder-relay-mainnet.blocknative.com,https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com,https://0x8b5d2e73e2a3a55c6c87b8b6eb92e0149a125c852751db1422fa951e42a09b82c142c3ea98d0d9930b056a3bc9896b8f@bloxroute.max-profit.blxrbdn.com,https://0xb0b07cd0abef743db4260b0ed50619cf6ad4d82064cb4fbec9d3ec530f7c5e6793d9f286c4e082c0244ffb9f2658fe88@bloxroute.regulated.blxrbdn.com,https://0xb3ee7afcf27f1f1259ac1787876318c6584ee353097a50ed84f51a1f21a323b3736f271a895c7ce918c038e4265918be@relay.edennetwork.io,https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net,https://0xa1559ace749633b997cb3fdacffb890aeebdb0f5a3b6aaa7eeeaf1a38af0a8fe88b9e4b1f61f236d2e64d95733327a62@relay.ultrasound.money",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// generate buffer to write the .env file
			var buffer bytes.Buffer
			if err := EnvFile(tt.data, &buffer); err != nil {
				if err != tt.Error {
					t.Error("unable to generate .env file")
				}
				return
			}
			// read the .env file
			data := retrieveEnvData(t, &buffer)
			for key, value := range tt.fieldsToCheck {
				assert.Contains(t, data, key)
				assert.True(t, strings.Contains(strings.ReplaceAll(data[key], "\r", ""), value))
			}
		})
	}
}

// Test some env vars doesn't exist
// TODO: add more tests cases
func TestMissingEnvVars(t *testing.T) {
	tests := []struct {
		name          string
		data          *GenData
		Error         error
		fieldsToCheck []string
	}{
		{
			name: "Check RELAY_URLS",
			data: &GenData{
				Services:        []string{consensus, validator, execution},
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			fieldsToCheck: []string{
				"RELAY_URLS",
			},
		},
		{
			name: "Check RELAY_URLS",
			data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			fieldsToCheck: []string{
				"RELAY_URLS",
			},
		},
		{
			name: "Check set of RELAY_URLS",
			data: &GenData{
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				RelayURLs:       []string{"https://sample.relay", "https://sample.relay2"},
			},
			fieldsToCheck: []string{
				"RELAY_URLS",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// generate buffer to write the .env file
			var buffer bytes.Buffer
			if err := EnvFile(tt.data, &buffer); err != nil {
				if err != tt.Error {
					t.Error("unable to generate .env file")
				}
				return
			}
			// read the .env file
			data := retrieveEnvData(t, &buffer)
			for key := range tt.fieldsToCheck {
				assert.NotContains(t, data, key)
			}
		})
	}
}
