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
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	executionContainerName         = "sedge-execution-client"
	executionNotFoundContainerName = "sedge-execution-client-nf"
	consensusContainerName         = "sedge-consensus-client"
	consensusNotFoundContainerName = "sedge-execution-client-nf"
	validatorContainerName         = "sedge-validator-client"
	validatorNotFoundContainerName = "sedge-validator-client-nf"
	starknetContainerName          = "sedge-starknet-client"
	starknetNotFoundContainerName  = "sedge-starknet-client-nf"
	mevBoostContainerName          = "sedge-mev-boost"

	executionContainerIp    = "192.168.1.1"
	executionContainerImage = "nethermind/nethermind"
	consensusContainerIp    = "192.168.1.2"
	consensusContainerImage = "sigp/lighthouse"
	validatorContainerIp    = "192.168.1.3"
	validatorContainerImage = "consensys/teku"
	mevBoostContainerIp     = "192.168.1.4"
	mevBoostContainerImage  = "flashbots/mev-boost"
	starknetContainerIp     = "192.168.1.5"
	starknetContainerImage  = "nethermind/juno"

	executionNotFoundErrorMsg = "execution container not found"
	consensusNotFoundErrorMsg = "consensus container not found"
	validatorNotFoundErrorMsg = "validator container not found"
	starknetNotFoundErrorMsg  = "starknet container not found"

	unexpectedContainerErrorMsg = "unexpected container name"
)

type getContainersTestCase struct {
	name                     string
	getContainersDataOptions actions.GetContainersDataOptions
	expected                 actions.ContainersData
	logsOutput               bytes.Buffer
	isErr                    bool
	expectedErrMsg           string
	containersTag            string
}

func nameWithTag(
	name,
	tag string,
) string {
	return strings.Join([]string{name, tag}, "-")
}

func buildGetContainersDataTestCase(
	t *testing.T,
	name,
	containersTag,
	caseDataDir string,
	isErr bool,
	expectedErrMsg string,
	expected actions.ContainersData,
) getContainersTestCase {
	testCaseDockerComposeFilePath := filepath.Join("testdata", "getContainers_tests", caseDataDir, configs.DefaultDockerComposeScriptName)
	testCaseDockerComposeFileContent, err := os.ReadFile(testCaseDockerComposeFilePath)
	if err != nil {
		t.Fatalf("Failed to read %s file: %v", testCaseDockerComposeFilePath, err)
	}

	testCaseFinalDockerComposeFilePath := filepath.Join(t.TempDir(), configs.DefaultDockerComposeScriptName)
	err = os.WriteFile(testCaseFinalDockerComposeFilePath, testCaseDockerComposeFileContent, 0o644)
	if err != nil {
		t.Fatalf("Failed to write %s file: %v", testCaseFinalDockerComposeFilePath, err)
	}

	return getContainersTestCase{
		name: name,
		getContainersDataOptions: actions.GetContainersDataOptions{
			DockerComposePath: testCaseFinalDockerComposeFilePath,
		},
		expected:      expected,
		logsOutput:    bytes.Buffer{},
		isErr:         isErr,
		containersTag: containersTag,
	}
}

func buildInspectResults(
	containersTag,
	containerName,
	containerImage,
	containerIp string,
) types.ContainerJSON {
	networkName := "sedge-network" // May change in templates
	if containersTag != "" {
		networkName = nameWithTag(networkName, containersTag)
	}
	return types.ContainerJSON{
		ContainerJSONBase: &types.ContainerJSONBase{
			Name: containerName,
		},
		Config: &container.Config{
			Image: containerImage,
		},
		NetworkSettings: &types.NetworkSettings{
			Networks: map[string]*network.EndpointSettings{
				networkName: {
					IPAddress: containerIp,
				},
			},
		},
	}
}

func getMockActions(
	t *testing.T,
	tc getContainersTestCase,
) actions.SedgeActions {
	ctrl := gomock.NewController(t)
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	t.Cleanup(
		func() {
			ctrl.Finish()
		},
	)

	for _, container := range tc.expected.Containers {
		dockerClient.EXPECT().ContainerInspect(gomock.Any(), container.Name).Return(
			buildInspectResults(
				tc.containersTag,
				container.Name,
				container.Image,
				container.Ip,
			),
			nil,
		).Times(1)
	}

	// not found containers
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), executionNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(executionNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), consensusNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(consensusNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), validatorNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(validatorNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), starknetNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(starknetNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.All(), gomock.All()).Return(types.ContainerJSON{}, errors.New(unexpectedContainerErrorMsg)).AnyTimes()

	return actions.NewSedgeActions(
		actions.SedgeActionsOptions{
			DockerClient:   dockerClient,
			ServiceManager: nil,
			CommandRunner:  nil,
		},
	)
}

func TestGetContainersData(t *testing.T) {
	tcs := []getContainersTestCase{
		buildGetContainersDataTestCase(
			t,
			"Full Node",
			"",
			"case_fullNode",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  mevBoostContainerName,
						Image: mevBoostContainerImage,
						Ip:    mevBoostContainerIp,
					},
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
					{
						Name:  validatorContainerName,
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Full Node with Tags",
			"tag",
			"case_fullNodeTag",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  nameWithTag(executionContainerName, "tag"),
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  nameWithTag(mevBoostContainerName, "tag"),
						Image: mevBoostContainerImage,
						Ip:    mevBoostContainerIp,
					},
					{
						Name:  nameWithTag(consensusContainerName, "tag"),
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
					{
						Name:  nameWithTag(validatorContainerName, "tag"),
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Starknet Full Node",
			"",
			"case_fullNodeStarknet",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
					{
						Name:  starknetContainerName,
						Image: starknetContainerImage,
						Ip:    starknetContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Execution Only",
			"",
			"case_executionOnly",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Consensus Only",
			"",
			"case_consensusOnly",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Validator Only",
			"",
			"case_validatorOnly",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  validatorContainerName,
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Starknet Only",
			"",
			"case_starknetOnly",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  starknetContainerName,
						Image: starknetContainerImage,
						Ip:    starknetContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Execution",
			"",
			"case_noExecution",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  mevBoostContainerName,
						Image: mevBoostContainerImage,
						Ip:    mevBoostContainerIp,
					},
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
					{
						Name:  validatorContainerName,
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Consensus",
			"",
			"case_noConsensus",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  mevBoostContainerName,
						Image: mevBoostContainerImage,
						Ip:    mevBoostContainerIp,
					},
					{
						Name:  validatorContainerName,
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Mev Boost",
			"",
			"case_noMev",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
					{
						Name:  validatorContainerName,
						Image: validatorContainerImage,
						Ip:    validatorContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Validator",
			"",
			"case_noValidator",
			false,
			"",
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: executionContainerImage,
						Ip:    executionContainerIp,
					},
					{
						Name:  mevBoostContainerName,
						Image: mevBoostContainerImage,
						Ip:    mevBoostContainerIp,
					},
					{
						Name:  consensusContainerName,
						Image: consensusContainerImage,
						Ip:    consensusContainerIp,
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Execution Not Found",
			"",
			"case_executionNF",
			true,
			executionNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{}},
		),
		buildGetContainersDataTestCase(
			t,
			"Consensus Not Found",
			"",
			"case_consensusNF",
			true,
			consensusNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{}},
		),
		buildGetContainersDataTestCase(
			t,
			"Validator Not Found",
			"",
			"case_validatorNF",
			true,
			validatorNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{}},
		),
		buildGetContainersDataTestCase(
			t,
			"Starknet Not Found",
			"",
			"case_starknetNF",
			true,
			starknetNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{}},
		),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actions := getMockActions(t, tc)
			log.SetOutput(&tc.logsOutput)
			containersData, err := actions.GetContainersData(tc.getContainersDataOptions)
			if tc.isErr {
				assert.Contains(t, tc.logsOutput.String(), tc.expectedErrMsg)
			} else if !tc.isErr && assert.NoError(t, err) {
				validateContainersData(t, tc.expected, containersData)
			}
		})
	}
}

func validateContainersData(t *testing.T, expected, actual actions.ContainersData) {
	assert.Equal(t, expected, actual)
}
