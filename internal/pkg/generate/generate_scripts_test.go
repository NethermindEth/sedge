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
	"errors"
	"fmt"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const (
	wrongDep = "wrong_dep"
)

type genTestData struct {
	Description    string
	GenerationData *GenData
	Services       []string
	Error          error
	CheckFunc      func(t *testing.T, data *GenData, services []string, compose, env io.Reader) error
}

var defaultFunc = func(t *testing.T, data *GenData, services []string, compose, env io.Reader) error {

	// load compose file
	composeBytes, err := ioutil.ReadAll(compose)
	if err != nil {
		return err
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeBytes, &composeData)
	if err != nil {
		return err
	}

	if utils.Contains(services, execution) && composeData.Services.Execution == nil {
		return errors.New("execution service should not be omitted")
	}
	if utils.Contains(services, consensus) && composeData.Services.Consensus == nil {
		return errors.New("consensus service should not be omitted")
	}
	if utils.Contains(services, validator) && composeData.Services.Validator == nil {
		return errors.New("validator service should not be omitted")
	}
	if utils.Contains(services, mevBoost) && composeData.Services.Mevboost == nil {
		return errors.New("mev boost service should not be omitted")
	}
	if utils.Contains(services, validatorImport) && composeData.Services.ValidatorImport == nil {
		return errors.New("validator import service should not be omitted")
	}
	if utils.Contains(services, configConsensus) && composeData.Services.ConfigConsensus == nil {
		return errors.New("validator export service should not be omitted")
	}
	return nil
}

func generateTestCases(t *testing.T) (tests []genTestData) {
	baseDescription := "Test generation of compose services "
	tests = []genTestData{{Description: baseDescription + " NilData", Error: EmptyDataError, CheckFunc: defaultFunc}}

	networks, err := utils.SupportedNetworks()
	if err != nil {
		t.Error("SupportedNetworks() failed", err)
	}

	for _, network := range networks {
		c := clients.ClientInfo{Network: network}

		executionClients, err := c.SupportedClients("execution")
		if err != nil {
			t.Errorf("SupportedClients(\"execution\") failed: %v", err)
		}
		consensusClients, err := c.SupportedClients("consensus")
		if err != nil {
			t.Errorf("SupportedClients(\"consensus\") failed: %v", err)
		}
		validatorClients, err := c.SupportedClients("validator")
		if err != nil {
			t.Errorf("SupportedClients(\"validator\") failed: %v", err)
		}

		// TODO: Add CheckpointSyncUrl, FallbackELUrls and FeeRecipient to test data
		for _, executionCl := range executionClients {
			for _, consensusCl := range consensusClients {
				if utils.Contains(validatorClients, consensusCl) {
					tests = append(tests,
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl},
								Network:         "sepolia",
							},
							Services:  []string{execution, consensus, validator},
							Error:     nil,
							CheckFunc: defaultFunc,
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl, Omited: true},
								Network:         "sepolia",
							},
							Services:  []string{execution, consensus},
							Error:     nil,
							CheckFunc: defaultFunc,
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: wrongDep},
								ValidatorClient: &clients.Client{Name: consensusCl, Omited: true},
								Network:         "sepolia",
							},
							Services:  []string{execution, consensus},
							Error:     ConsensusClientNotValidError,
							CheckFunc: defaultFunc,
						})
				}
			}
		}
	}

	return
}

func TestGenerateComposeServices(t *testing.T) {
	tests := []genTestData{
		{
			Description: "Test generation of compose services",
			GenerationData: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Services:  []string{execution, consensus, validator, validatorImport, mevBoost},
			Error:     nil,
			CheckFunc: defaultFunc,
		},
	}

	tests = append(tests, generateTestCases(t)...)

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {

			var buffer bytes.Buffer
			if tt.GenerationData != nil && tt.GenerationData.Network == "chiado" {
				t.Logf("GenerationData: %+v", tt.GenerationData)
			}
			err := genComposeFile(tt.GenerationData, io.Writer(&buffer))
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}

			err = tt.CheckFunc(t, tt.GenerationData, tt.Services, bytes.NewReader(buffer.Bytes()), nil)
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
			}
		})
	}
}

// Test that the generated compose file with dump data is generated correctly
func TestGenerateDockerCompose(t *testing.T) {
	samplePath := t.TempDir()
	sampleData := &GenData{
		ExecutionClient: &clients.Client{Name: "nethermind", Omited: false},
		Network:         "mainnet",
	}

	err := GenerateDockerComposeAndEnvFile(sampleData, samplePath)
	if err != nil {
		t.Error("GenerateDockerComposeAndEnvFile() failed", err)
		return
	}

	// Check that docker-compose file exists
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	// Check that .env file doesn't exist
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultEnvFileName))

	// Validate that Execution Client info matches the sample data
	// load the docker-compose file
	composeFile, err := ioutil.ReadFile(filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	if err != nil {
		t.Error("unable to read docker-compose.yml")
	}
	var composeData ComposeData
	err = yaml.Unmarshal(composeFile, &composeData)
	if err != nil {
		t.Error("unable to parse docker-compose.yml")
	}

	// Check that the execution client is nethermind
	if composeData.Services.Execution.ContainerName != "execution-client" {
		t.Error("execution client image does not match")
	}

	// Check other services are nil
	if composeData.Services.Consensus != nil {
		t.Error("consensus client should be nil")
	}

}
