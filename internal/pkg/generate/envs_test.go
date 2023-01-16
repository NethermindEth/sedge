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

	"github.com/NethermindEth/sedge/configs"
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
	configs.InitNetworksConfigs()
	tests := []struct {
		name          string
		data          *GenData
		Error         error
		fieldsToCheck map[string]string
	}{
		{
			name: "Check ec image",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				// WIll match any image
				"EC_IMAGE_VERSION": "nethermind/nethermind:",
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
		{
			name: "Check RELAY_URL",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "besu"},
				Network:         "mainnet",
			},
			fieldsToCheck: map[string]string{
				"RELAY_URL": "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net",
			},
		},
		{
			name: "Check set of RELAY_URL",
			data: &GenData{
				ExecutionClient: &clients.Client{Name: "geth"},
				Network:         "mainnet",
				RelayURL:        "https://sample.relay",
			},
			fieldsToCheck: map[string]string{
				"RELAY_URL": "https://sample.relay",
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
