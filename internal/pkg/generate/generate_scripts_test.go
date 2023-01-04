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
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	wrongDep = "wrong_dep"
)

type genTestData struct {
	Description     string
	GenerationData  *GenData
	ErrorGenCompose error
	ErrorGenEnvFile error
	ErrorCheckFunc  error
	CheckFunc       func(t *testing.T, data *GenData, compose, env io.Reader) error
}

var checkMevServices = func(t *testing.T, data *GenData, compose, env io.Reader) error {
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

	if utils.Contains(data.Services, mevBoost) {

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

var checkExtraFlagsOnExecution = func(t *testing.T, data *GenData, compose, env io.Reader) error {
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

	if composeData.Services.Execution != nil {
		for _, flag := range data.ElExtraFlags {
			assert.True(t, utils.Contains(composeData.Services.Execution.Command, "--"+flag))
		}
	} else {
		return errors.New("execution service is not present")
	}

	return nil
}

var defaultFunc = func(t *testing.T, data *GenData, compose, env io.Reader) error {

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

	if utils.Contains(data.Services, execution) {
		assert.NotNil(t, composeData.Services.Execution)
	}
	if utils.Contains(data.Services, consensus) {
		assert.NotNil(t, composeData.Services.Consensus)
	}
	if utils.Contains(data.Services, validator) {
		assert.NotNil(t, composeData.Services.Validator)
	}
	if utils.Contains(data.Services, mevBoost) {
		assert.NotNil(t, composeData.Services.Mevboost)
	}
	if utils.Contains(data.Services, validatorImport) {
		assert.NotNil(t, composeData.Services.ValidatorImport)
	}
	if utils.Contains(data.Services, configConsensus) {
		assert.NotNil(t, composeData.Services.ConfigConsensus)
	}

	// load .env file
	envData := retriveEnvData(t, env)
	if data.Network == "gnosis" {
		// Check that the right network is set
		assert.Contains(t, envData, "EL_NETWORK")
		assert.Equal(t, envData["EL_NETWORK"], "xdai")

		assert.Contains(t, envData, "CL_NETWORK")
		assert.Equal(t, envData["CL_NETWORK"], "gnosis")

	} else {
		// Check that the right network is set
		assert.Contains(t, envData, "NETWORK")
		assert.Equal(t, envData["NETWORK"], data.Network)
	}

	return nil
}

func generateTestCases(t *testing.T) (tests []genTestData) {
	baseDescription := "Test generation of compose services "
	tests = []genTestData{{Description: baseDescription + " NilData", ErrorGenCompose: EmptyDataError, CheckFunc: defaultFunc}}

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
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, all", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl},
								Network:         network,
								Services:        []string{execution, consensus, validator},
							},
							CheckFunc: defaultFunc,
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, no validator", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{execution, consensus},
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: consensusCl},
								ValidatorClient: &clients.Client{Name: consensusCl, Omitted: true},
								Network:         network,
							},
							CheckFunc: defaultFunc,
						},
						genTestData{
							Description: fmt.Sprintf(baseDescription+"execution: %s, consensus: %s, validator: %s, network: %s, Execution Client not Valid", executionCl, consensusCl, consensusCl, network),
							GenerationData: &GenData{
								Services:        []string{execution, consensus},
								ExecutionClient: &clients.Client{Name: executionCl},
								ConsensusClient: &clients.Client{Name: wrongDep},
								ValidatorClient: &clients.Client{Name: consensusCl, Omitted: true},
								Network:         network,
							},
							ErrorGenCompose: ConsensusClientNotValidError,
							CheckFunc:       defaultFunc,
						})
				}
			}
		}
	}

	return
}

func TestGenerateComposeServices(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []genTestData{
		{
			Description: "Test generation of compose services",
			GenerationData: &GenData{
				Services:        []string{execution, consensus, validator, validatorImport, mevBoost},
				ExecutionClient: &clients.Client{Name: "nethermind"},
				ConsensusClient: &clients.Client{Name: "teku"},
				ValidatorClient: &clients.Client{Name: "teku"},
				Network:         "mainnet",
				Mev:             true,
			},
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
			CheckFunc: checkMevServices,
		},
		{
			Description: "Test EL extra flags",
			GenerationData: &GenData{
				ExecutionClient: &clients.Client{Name: "nethermind"},
				Services:        []string{execution},
				Network:         "mainnet",
				ElExtraFlags:    []string{"extra", "flag"},
			},
			CheckFunc: checkExtraFlagsOnExecution,
		},
	}

	tests = append(tests, generateTestCases(t)...)

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {

			var buffer bytes.Buffer
			err := ComposeFile(tt.GenerationData, io.Writer(&buffer))
			assert.ErrorIs(t, err, tt.ErrorGenCompose)
			if tt.ErrorGenCompose != nil {
				return
			}

			var envBuffer bytes.Buffer
			err = EnvFile(tt.GenerationData, io.Writer(&envBuffer))
			assert.ErrorIs(t, err, tt.ErrorGenEnvFile)
			if tt.ErrorGenEnvFile != nil {
				return
			}

			err = tt.CheckFunc(t, tt.GenerationData, bytes.NewReader(buffer.Bytes()), bytes.NewReader(envBuffer.Bytes()))
			assert.ErrorIs(t, err, tt.ErrorCheckFunc)
		})
	}
}

// TestValidateClients tests the validation of clients
func TestValidateClients(t *testing.T) {
	configs.InitNetworksConfigs()
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

func TestGenComposeFileSimpleExecution(t *testing.T) {
	configs.InitNetworksConfigs()
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
			assert.ErrorIs(t, err, tt.Error)
			if tt.Error != nil {
				return
			}
		})
	}
}

func TestEnvFileAndFlags(t *testing.T) {
	configs.InitNetworksConfigs()
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

func TestCleanGeneratedFiles(t *testing.T) {
	configs.InitNetworksConfigs()
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
			assert.ErrorIs(t, err, tt.Error)
			if tt.Error != nil {
				return
			}
			err = ComposeFile(tt.Data, out)
			assert.ErrorIs(t, err, tt.Error)
			if tt.Error != nil {
				return
			}
			assert.FileExists(t, filepath.Join(path, configs.DefaultDockerComposeScriptName))

			// open output file
			out, err = os.Create(filepath.Join(path, configs.DefaultEnvFileName))
			err = EnvFile(tt.Data, out)
			assert.ErrorIs(t, err, tt.Error)
			if tt.Error != nil {
				return
			}
			assert.FileExists(t, filepath.Join(path, configs.DefaultEnvFileName))

			err = CleanGenerated(path)
			assert.ErrorIs(t, err, tt.Error)
			if tt.Error != nil {
				return
			}
		})
	}
}
