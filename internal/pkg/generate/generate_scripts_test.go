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
	"os"
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

var checkMevServices = func(t *testing.T, data *GenData, services []string, compose, env io.Reader) error {
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

	if utils.Contains(services, mevBoost) {

		if composeData.Services.Mevboost != nil {
			assert.Equal(t, composeData.Services.Mevboost.Image, "flashbots/mev-boost:latest")
			assert.Equal(t, composeData.Services.Mevboost.ContainerName, "mev-boost")
			assert.Equal(t, composeData.Services.Mevboost.Restart, "on-failure")
		} else {
			return errors.New("mevboost service is not present")
		}
	}

	return nil
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
								ValidatorClient: &clients.Client{Name: consensusCl, Omitted: true},
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
								ValidatorClient: &clients.Client{Name: consensusCl, Omitted: true},
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
		{
			Description: "Test generation import",
			GenerationData: &GenData{
				Services:        []string{mevBoost},
				ExecutionClient: &clients.Client{Name: "nethermind", Omitted: true},
				ConsensusClient: &clients.Client{Name: "teku", Omitted: true},
				ValidatorClient: &clients.Client{Name: "teku", Omitted: true},
				Network:         "mainnet",
				Mev:             true,
				MevBoostService: true,
			},
			Services:  []string{mevBoost},
			Error:     nil,
			CheckFunc: checkMevServices,
		},
	}

	tests = append(tests, generateTestCases(t)...)

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {

			var buffer bytes.Buffer
			if tt.GenerationData != nil && tt.GenerationData.Network == "chiado" {
				t.Logf("GenerationData: %+v", tt.GenerationData)
			}
			err := ComposeFile(tt.GenerationData, io.Writer(&buffer))
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

// TestValidateClients tests the validation of clients
func TestValidateClients(t *testing.T) {
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Wrong execution client",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ExecutionClientNotValidError,
		},
		{
			Description: "Wrong consensus client",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ConsensusClientNotValidError,
		},
		{
			Description: "Wrong validator client",
			Data: &GenData{
				ValidatorClient: &clients.Client{Name: "wrong"},
				Network:         "mainnet",
			},
			Error: ValidatorClientNotValidError,
		},
		{
			Description: "Wrong network, empty clients",
			Data: &GenData{
				Network: wrongDep,
			},
			Error: nil,
		},
		{
			Description: "Wrong network, good consensus",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: UnableToGetClientsInfoError,
		},
		{
			Description: "Wrong network, good execution",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Network:         wrongDep,
			},
			Error: UnableToGetClientsInfoError,
		},
		{
			Description: "Wrong network, good validator",
			Data: &GenData{
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: UnableToGetClientsInfoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			err := validateClients(tt.Data)
			assert.ErrorIs(t, err, tt.Error)
		})
	}
}

func TestComposeFileSimple(t *testing.T) {
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Empty data",
			Data: &GenData{
				Network: "mainnet",
			},
			Error: EmptyDataError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			var buffer bytes.Buffer
			err := ComposeFile(tt.Data, io.Writer(&buffer))
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
		})
	}
}

func TestEnvFile(t *testing.T) {
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku", Endpoint: "http://localhost"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Empty data",
			Data: &GenData{
				Network: "mainnet",
			},
			Error: EmptyDataError,
		},
		{
			Description: "Wrong network",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				Network:         wrongDep,
			},
			Error: TemplateNotFoundError,
		},
		{
			Description: "Prysm consensus with ConsensusAdditionalUrl",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "prysm"},
				ValidatorClient: &clients.Client{Name: "prysm"},
				Network:         "mainnet",
				ConsensusApiUrl: "http://localhost:8080",
			},
			Error: nil,
		},
		{
			Description: "Teku consensus with ConsensusAdditionalUrl",
			Data: &GenData{
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				ConsensusApiUrl: "http://localhost:8080",
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			var buffer bytes.Buffer
			err := EnvFile(tt.Data, io.Writer(&buffer))
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
			if tt.Data.ConsensusApiUrl == "" {
				str := buffer.String()
				assert.Contains(t, str, "CC_API_URL="+tt.Data.ConsensusClient.Endpoint+":")
			} else {
				if tt.Data.ConsensusClient.Name == "prysm" && !tt.Data.ValidatorClient.Omitted {
					assert.Contains(t, buffer.String(), "CC_API_URL=consensus:")
				} else {
					assert.Contains(t, buffer.String(), "CC_API_URL="+tt.Data.ConsensusApiUrl)
				}
			}
		})
	}
}

func TestCleanGenerated(t *testing.T) {
	tests := []struct {
		Description string
		Data        *GenData
		Error       error
	}{
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
		{
			Description: "Test generation of compose services",
			Data: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			path := t.TempDir()

			// generate files
			out, err := os.Create(filepath.Join(path, configs.DefaultDockerComposeScriptName))
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
			err = ComposeFile(tt.Data, out)
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
			assert.FileExists(t, filepath.Join(path, configs.DefaultDockerComposeScriptName))

			// open output file
			out, err = os.Create(filepath.Join(path, configs.DefaultEnvFileName))
			err = EnvFile(tt.Data, out)
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
			assert.FileExists(t, filepath.Join(path, configs.DefaultEnvFileName))

			err = CleanGenerated(path)
			if err != nil {
				assert.ErrorIs(t, err, tt.Error)
				return
			}
		})
	}
}
