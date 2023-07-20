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
package actions_test

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/utils"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newAction(t *testing.T, ctrl *gomock.Controller) actions.SedgeActions {
	t.Helper()
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	serviceManager := services.NewServiceManager(dockerClient)
	return actions.NewSedgeActions(actions.SedgeActionsOptions{
		DockerClient:   dockerClient,
		ServiceManager: serviceManager,
	})
}

func contains(t *testing.T, list []string, str string) bool {
	t.Helper()
	for _, s := range list {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

type genTestData struct {
	name    string
	genData generate.GenData
}

// Test that the generated compose file with dump data is generated correctly
func TestGenerateDockerCompose(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	tests := make([]genTestData, 0)

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

		rNum, err := rand.Int(rand.Reader, big.NewInt(int64(100)))
		if err != nil {
			t.Errorf("rand.Int() failed: %v", err)
		}
		gracePeriod := configs.NetworkEpochTime(network) * time.Duration(int(rNum.Int64()))
		ckptSync := fmt.Sprintf("http://localhost:%d", 40+rNum.Int64())

		for _, executionCl := range executionClients {
			for _, consensusCl := range consensusClients {
				tests = append(tests,
					genTestData{
						name: fmt.Sprintf("execution: %s, network: %s, only execution", executionCl, network),
						genData: generate.GenData{
							ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
							Services:        []string{"execution"},
							Network:         network,
						},
					},
					genTestData{
						name: fmt.Sprintf("execution: %s, network: %s, only execution with tag", executionCl, network),
						genData: generate.GenData{
							ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
							Services:        []string{"execution"},
							Network:         network,
							ContainerTag:    "sampleTag",
						},
					},
					genTestData{
						name: fmt.Sprintf("consensus: %s, network: %s, only consensus", consensusCl, network),
						genData: generate.GenData{
							ConsensusClient:  &clients.Client{Name: consensusCl, Type: "consensus"},
							Services:         []string{"consensus"},
							Network:          network,
							ExecutionApiUrl:  "http://localhost:8545",
							ExecutionAuthUrl: "http://localhost:8551",
						},
					},
					genTestData{
						name: fmt.Sprintf("consensus: %s, network: %s, only consensus with tag, https", consensusCl, network),
						genData: generate.GenData{
							ConsensusClient:  &clients.Client{Name: consensusCl, Type: "consensus"},
							Services:         []string{"consensus"},
							Network:          network,
							ContainerTag:     "sampleTag",
							ExecutionApiUrl:  "https://localhost:8545",
							ExecutionAuthUrl: "https://localhost:8551",
						},
					},
					genTestData{
						name: fmt.Sprintf("consensus: %s, network: %s, only consensus with custom Checkpoint sync URL", consensusCl, network),
						genData: generate.GenData{
							ConsensusClient:   &clients.Client{Name: consensusCl, Type: "consensus"},
							Services:          []string{"consensus"},
							Network:           network,
							CheckpointSyncUrl: ckptSync,
							ExecutionApiUrl:   "http://localhost:8545",
							ExecutionAuthUrl:  "http://localhost:8551",
						},
					},
					genTestData{
						name: fmt.Sprintf("validator: %s, network: %s, only validator", consensusCl, network),
						genData: generate.GenData{
							ValidatorClient: &clients.Client{Name: consensusCl, Type: "validator"},
							Services:        []string{"validator"},
							Network:         network,
							ConsensusApiUrl: "http://localhost:4000",
						},
					},
					genTestData{
						name: fmt.Sprintf("validator: %s, network: %s, only validator, mev-boost on", consensusCl, network),
						genData: generate.GenData{
							ValidatorClient:     &clients.Client{Name: consensusCl, Type: "validator"},
							Services:            []string{"validator"},
							Network:             network,
							ConsensusApiUrl:     "http://localhost:4000",
							MevBoostOnValidator: true,
						},
					},
					genTestData{
						name: fmt.Sprintf("validator: %s, network: %s, only validator with tag, https", consensusCl, network),
						genData: generate.GenData{
							ValidatorClient: &clients.Client{Name: consensusCl, Type: "validator"},
							Services:        []string{"validator"},
							Network:         network,
							ContainerTag:    "sampleTag",
							ConsensusApiUrl: "https://localhost:4000",
						},
					},
				)
				if utils.Contains(validatorClients, consensusCl) {
					tests = append(tests,
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, all", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient: &clients.Client{Name: consensusCl, Type: "consensus"},
								ValidatorClient: &clients.Client{Name: consensusCl, Type: "validator"},
								Services:        []string{"execution", "consensus", "validator"},
								Network:         network,
								Mev:             true,
							},
						},
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, all, no mev-boost", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient: &clients.Client{Name: consensusCl, Type: "consensus"},
								ValidatorClient: &clients.Client{Name: consensusCl, Type: "validator"},
								Services:        []string{"execution", "consensus", "validator"},
								Network:         network,
							},
						},
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, all, with tag", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient: &clients.Client{Name: consensusCl, Type: "consensus"},
								ValidatorClient: &clients.Client{Name: consensusCl, Type: "validator"},
								Services:        []string{"execution", "consensus", "validator"},
								Network:         network,
								ContainerTag:    "sampleTag",
								Mev:             true,
							},
						},
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, all, waitEpoch set and custom Checkpoint Sync URL", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient:    &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient:    &clients.Client{Name: consensusCl, Type: "consensus"},
								ValidatorClient:    &clients.Client{Name: consensusCl, Type: "validator"},
								Services:           []string{"execution", "consensus", "validator"},
								Network:            network,
								ContainerTag:       "sampleTag",
								VLStartGracePeriod: uint(gracePeriod.Abs()),
								CheckpointSyncUrl:  ckptSync,
							},
						},
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, no validator", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient: &clients.Client{Name: consensusCl, Type: "consensus"},
								Services:        []string{"execution", "consensus"},
								Network:         network,
							},
						},
						genTestData{
							name: fmt.Sprintf("execution: %s, consensus: %s, validator: %s, network: %s, no validator, with tag", executionCl, consensusCl, consensusCl, network),
							genData: generate.GenData{
								ExecutionClient: &clients.Client{Name: executionCl, Type: "execution"},
								ConsensusClient: &clients.Client{Name: consensusCl, Type: "consensus"},
								Services:        []string{"execution", "consensus"},
								Network:         network,
								ContainerTag:    "sampleTag",
							},
						},
					)
				}
			}
		}
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			samplePath := t.TempDir()
			sedgeAction := newAction(t, nil)

			// Always set the JWT secret path
			tc.genData.JWTSecretPath = samplePath
			datadirs := make([]datadirsValidation, 0)

			// Setup client images
			if tc.genData.ExecutionClient != nil {
				tc.genData.ExecutionClient.SetImageOrDefault("")
				datadirs = append(datadirs, datadirsValidation{path: configs.ExecutionDir, shouldExist: true})
			} else {
				datadirs = append(datadirs, datadirsValidation{path: configs.ExecutionDir, shouldExist: false})
			}
			if tc.genData.ConsensusClient != nil {
				tc.genData.ConsensusClient.SetImageOrDefault("")
				datadirs = append(datadirs, datadirsValidation{path: configs.ConsensusDir, shouldExist: true})
			} else {
				datadirs = append(datadirs, datadirsValidation{path: configs.ConsensusDir, shouldExist: false})
			}
			if tc.genData.ValidatorClient != nil {
				tc.genData.ValidatorClient.SetImageOrDefault("")
				datadirs = append(datadirs, datadirsValidation{path: configs.ValidatorDir, shouldExist: true})
			} else {
				datadirs = append(datadirs, datadirsValidation{path: configs.ValidatorDir, shouldExist: false})
			}

			_, err := sedgeAction.Generate(actions.GenerateOptions{
				GenerationData: tc.genData,
				GenerationPath: samplePath,
			})
			if err != nil {
				t.Error("GenerateDockerComposeAndEnvFile() failed", err)
				return
			}

			validateGeneration(t, samplePath, datadirs...)
			cmpData, err := generate.ParseCompose(filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
			require.Nil(t, err)
			envData, err := utils.ParseEnv(filepath.Join(samplePath, configs.DefaultEnvFileName))
			require.Nil(t, err)

			// Validate that Execution Client info matches the sample data
			if tc.genData.ExecutionClient != nil {
				// Check that the execution service is set.
				assert.NotNil(t, cmpData.Services.Execution)
				// Check that the execution container image is set.
				assert.NotEmpty(t, cmpData.Services.Execution.Image)
				// Check that the execution container name contains the tag.
				if tc.genData.ContainerTag == "" {
					assert.Equal(t, services.DefaultSedgeExecutionClient, cmpData.Services.Execution.ContainerName)
				} else {
					assert.Equal(t, services.DefaultSedgeExecutionClient+"-sampleTag", cmpData.Services.Execution.ContainerName)
				}

				// Check that mev-boost service is not set when execution only
				if tc.genData.ValidatorClient == nil && tc.genData.ConsensusClient == nil {
					assert.Nil(t, cmpData.Services.Mevboost)
				}
			}

			// Validate that Consensus Client info matches the sample data
			if tc.genData.ConsensusClient != nil {
				// Check that the consensus service is set.
				assert.NotNil(t, cmpData.Services.Consensus)
				// Check that the consensus container image is set.
				assert.NotEmpty(t, cmpData.Services.Consensus.Image)
				// Check that the consensus container name contains the tag.
				if tc.genData.ContainerTag == "" {
					assert.Equal(t, services.DefaultSedgeConsensusClient, cmpData.Services.Consensus.ContainerName)
				} else {
					assert.Equal(t, services.DefaultSedgeConsensusClient+"-sampleTag", cmpData.Services.Consensus.ContainerName)
				}
				// Check that Checkpoint Sync URL is set
				if tc.genData.CheckpointSyncUrl != "" {
					assert.True(t, contains(t, cmpData.Services.Consensus.Command, tc.genData.CheckpointSyncUrl), "Checkpoint Sync URL not found in consensus service command: %s", cmpData.Services.Consensus.Command)
				}

				// Validate Execution API and AUTH URLs
				apiEndpoint, authEndpoint := envData["EC_API_URL"], envData["EC_AUTH_URL"]
				if tc.genData.ExecutionApiUrl != "" {
					assert.Equal(t, tc.genData.ExecutionApiUrl, apiEndpoint, "Execution API URL is not valid %s", apiEndpoint)
				} else {
					re := regexp.MustCompile(`http:\/\/execution:[0-9]+`)
					assert.True(t, re.MatchString(apiEndpoint), "Execution API URL is not valid %s", apiEndpoint)
				}

				if tc.genData.ExecutionAuthUrl != "" {
					assert.Equal(t, tc.genData.ExecutionAuthUrl, authEndpoint, "Execution Auth URL is not valid %s", authEndpoint)
				} else {
					re := regexp.MustCompile(`http:\/\/execution:[0-9]+`)
					assert.True(t, re.MatchString(authEndpoint), "Execution Auth URL is not valid %s", authEndpoint)
				}

				// Check that mev-boost service is not set when consensus only
				if tc.genData.ExecutionClient == nil && tc.genData.ConsensusClient == nil {
					assert.Nil(t, cmpData.Services.Mevboost)
				}
			}

			// Validate that Validator Client info matches the sample data
			if tc.genData.ValidatorClient != nil {
				// Check that the validator service is set.
				assert.NotNil(t, cmpData.Services.Validator)
				// Check that the validator container name contains the tag.
				if tc.genData.ContainerTag == "" {
					assert.Equal(t, services.DefaultSedgeValidatorClient, cmpData.Services.Validator.ContainerName)
				} else {
					assert.Equal(t, services.DefaultSedgeValidatorClient+"-sampleTag", cmpData.Services.Validator.ContainerName)
				}
				// Check that the validator-blocker service is set.
				assert.NotNil(t, cmpData.Services.ValidatorBlocker)
				// Check that the validator grace period is set. Get the number after sleep
				re := regexp.MustCompile("sleep [0-9]+")
				fullSleep := re.FindAllString(cmpData.Services.ValidatorBlocker.Command, -1)
				re = regexp.MustCompile("[0-9]+")
				sleep := re.FindAllString(fullSleep[0], -1)
				sleepTime, err := strconv.Atoi(sleep[0])
				if err != nil {
					t.Error("Failed to parse sleep time", err)
				} else {
					// Check that the sleep time is equal to the grace period
					assert.Equal(t, tc.genData.VLStartGracePeriod, uint(sleepTime))
				}

				// Prysm special case: remove http:// or https:// from the URL
				prysmURL := tc.genData.ConsensusApiUrl
				prysmURL = strings.TrimPrefix(prysmURL, "http://")
				prysmURL = strings.TrimPrefix(prysmURL, "https://")

				// Check Consensus API URL is set and is valid
				uri, err := url.ParseRequestURI(envData["CC_API_URL"])
				assert.Nil(t, err)
				var add_uri *url.URL
				if tc.genData.ValidatorClient.Name == "prysm" {
					add_uri, err = url.ParseRequestURI(envData["CC_ADD_API_URL"])
					assert.Nil(t, err)
				}
				if tc.genData.ConsensusApiUrl != "" && tc.genData.ValidatorClient.Name != "prysm" {
					assert.Equal(t, tc.genData.ConsensusApiUrl, uri.String(), "Consensus API URL is not valid: %s", uri.String())
				} else if tc.genData.ConsensusApiUrl != "" && tc.genData.ValidatorClient.Name == "prysm" {
					assert.Equal(t, prysmURL, add_uri.String(), "Consensus Additional API URL is not valid: %s", uri.String())
				} else {
					var re *regexp.Regexp
					if tc.genData.ConsensusClient.Name == "prysm" {
						re = regexp.MustCompile(`consensus:[0-9]+`)
						assert.True(t, re.MatchString(add_uri.String()), "Consensus Additional API URL is not valid: %s", uri.String())
					} else {
						re = regexp.MustCompile(`http:\/\/consensus:[0-9]+`)
						assert.True(t, re.MatchString(uri.String()), "Consensus API URL is not valid: %s", uri.String())
					}
				}

				// Check that validator-blocker service is set
				assert.NotNil(t, cmpData.Services.ValidatorBlocker)
				// Check that validator-blocker image is set
				assert.Equal(t, "busybox", cmpData.Services.ValidatorBlocker.Image)
				// Check that validator-blocker container name is set properly
				validatorBlockerCtName := "sedge-validator-blocker"
				if tc.genData.ContainerTag != "" {
					validatorBlockerCtName = validatorBlockerCtName + "-" + tc.genData.ContainerTag
				}
				assert.Equal(t, validatorBlockerCtName, cmpData.Services.ValidatorBlocker.ContainerName)
				// Check that validator-blocker command is not empty
				assert.NotEmpty(t, cmpData.Services.ValidatorBlocker.Command)

				// Check that mev-boost service is not set when validator only
				_, mev := envData["MEV"]
				if tc.genData.ExecutionClient == nil && tc.genData.ConsensusClient == nil {
					assert.Nil(t, cmpData.Services.Mevboost)
				} else if mev && tc.genData.Mev { // Check that mev-boost service is set when full-node and mev is enabled
					// Check that mev-boost service is set
					assert.NotNil(t, cmpData.Services.Mevboost)
					// Check that mev-boost image is set
					assert.Equal(t, "flashbots/mev-boost:latest", cmpData.Services.Mevboost.Image)
					// Check that mev-boost entrypoint is set
					assert.NotEmpty(t, cmpData.Services.Mevboost.Entrypoint)
				}
			}

			if tc.genData.ValidatorClient == nil {
				// Check validator blocker is not set if validator is not set
				assert.Nil(t, cmpData.Services.ValidatorBlocker)
			} else {
				// Check that validator-blocker service is set when validator is set
				assert.NotNil(t, cmpData.Services.ValidatorBlocker)
			}
		})
	}
}

func TestFolderCreationOnCompose(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	samplePath := t.TempDir() + "test"
	c := clients.ClientInfo{Network: "mainnet"}
	clientsMap, _ := c.Clients([]clients.ClientType{clients.ExecutionClientType, clients.ConsensusClientType})
	sampleData := generate.GenData{
		ExecutionClient: clientsMap[clients.ExecutionClientType]["nethermind"],
		ConsensusClient: clientsMap[clients.ConsensusClientType]["lighthouse"],
		ValidatorClient: clientsMap[clients.ConsensusClientType]["lighthouse"],
		Services:        []string{"execution", "consensus", "validator"},
		Network:         "mainnet",
		JWTSecretPath:   samplePath,
	}
	sampleData.ExecutionClient.Image = configs.ClientImages.Execution.Nethermind.String()
	sampleData.ConsensusClient.Image = configs.ClientImages.Consensus.Lighthouse.String()
	sampleData.ValidatorClient.Image = configs.ClientImages.Consensus.Lighthouse.String()

	sedgeAction := newAction(t, nil)

	_, err := sedgeAction.Generate(actions.GenerateOptions{
		GenerationData: sampleData,
		GenerationPath: samplePath,
	})
	if !assert.Nilf(t, err, "GenerateDockerComposeAndEnvFile() failed") {
		return
	}

	validateGeneration(t, samplePath)
	// Remove the folder
	err = os.RemoveAll(samplePath)
	require.NoError(t, err, "unable to remove sample folder")
	// Check that the folder was removed
	assert.NoDirExists(t, samplePath)
}

type datadirsValidation struct {
	path        string
	shouldExist bool
}

func validateGeneration(t *testing.T, samplePath string, datadirs ...datadirsValidation) {
	t.Helper()

	// Check that the folder was created
	assert.DirExists(t, samplePath)
	// Check that docker-compose file exists
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	// Check that .env file exist
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultEnvFileName))
	// Check datadirs are correclty generated
	for _, datadir := range datadirs {
		if datadir.shouldExist {
			assert.DirExists(t, filepath.Join(samplePath, datadir.path))
		} else {
			assert.NoDirExists(t, filepath.Join(samplePath, datadir.path))
		}
	}
	// Check compose file correctness
	err := utils.ValidateCompose(filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	require.NoError(t, err, "generated compose file is not valid")
}
