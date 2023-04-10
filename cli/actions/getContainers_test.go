package actions_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
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
	mevBoostContainerName          = "sedge-mev-boost"

	executionContainerIp    = "192.168.1.1"
	executionContainerImage = "nethermind/nethermind"
	consensusContainerIp    = "192.168.1.2"
	consensusContainerImage = "sigp/lighthouse"
	validatorContainerIp    = "192.168.1.3"
	validatorContainerImage = "consensys/teku"
	mevBoostContainerIp     = "192.168.1.4"
	mevBoostContainerImage  = "flashbots/mev-boost"

	executionNotFoundErrorMsg = "execution container not found"
	consensusNotFoundErrorMsg = "consensus container not found"
	validatorNotFoundErrorMsg = "validator container not found"

	unexpectedContainerErrorMsg = "unexpected container name"
)

type getContainersTestCase struct {
	name                     string
	getContainersDataOptions actions.GetContainersDataOptions
	expected                 actions.ContainersData
	logsOutput               bytes.Buffer
	isErr                    bool
	expectedErrMsg           string
}

func buildGetContainersDataTestCase(
	t *testing.T,
	name string,
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
	err = os.WriteFile(testCaseFinalDockerComposeFilePath, testCaseDockerComposeFileContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write %s file: %v", testCaseFinalDockerComposeFilePath, err)
	}

	return getContainersTestCase{
		name: name,
		getContainersDataOptions: actions.GetContainersDataOptions{
			DockerComposePath: testCaseFinalDockerComposeFilePath,
		},
		expected:   expected,
		logsOutput: bytes.Buffer{},
		isErr:      isErr,
	}
}

func buildInspectResults(
	containerName,
	containerImage,
	containerIp string,
) types.ContainerJSON {
	return types.ContainerJSON{
		ContainerJSONBase: &types.ContainerJSONBase{
			Name: containerName,
		},
		Config: &container.Config{
			Image: containerImage,
		},
		NetworkSettings: &types.NetworkSettings{
			Networks: map[string]*network.EndpointSettings{
				"sedge-network": { //FIXME: fix in case of network data renaming
					IPAddress: containerIp,
				},
			},
		},
	}
}

func getMockActions(t *testing.T) actions.SedgeActions {
	ctrl := gomock.NewController(t)
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	t.Cleanup(
		func() {
			ctrl.Finish()
		},
	)

	// execution container
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), executionContainerName).Return(buildInspectResults(
		executionContainerName,
		executionContainerImage,
		executionContainerIp,
	), nil).AnyTimes()
	// consensus container
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), consensusContainerName).Return(buildInspectResults(
		consensusContainerName,
		consensusContainerImage,
		consensusContainerIp,
	), nil).AnyTimes()
	// validator container
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), validatorContainerName).Return(buildInspectResults(
		validatorContainerName,
		validatorContainerImage,
		validatorContainerIp,
	), nil).AnyTimes()
	// mev boost container
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), mevBoostContainerName).Return(buildInspectResults(
		mevBoostContainerName,
		mevBoostContainerImage,
		mevBoostContainerIp,
	), nil).AnyTimes()
	// not found containers
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), executionNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(executionNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), consensusNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(consensusNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.Any(), validatorNotFoundContainerName).Return(types.ContainerJSON{}, errors.New(validatorNotFoundErrorMsg)).AnyTimes()
	dockerClient.EXPECT().ContainerInspect(gomock.All(), gomock.All()).Return(types.ContainerJSON{}, errors.New(unexpectedContainerErrorMsg)).AnyTimes()

	return actions.NewSedgeActions(
		actions.SedgeActionsOptions{
			DockerClient:   dockerClient,
			ServiceManager: nil,
			CommandRunner:  nil,
		},
	)
}

func TestGetContainresData(t *testing.T) {
	tcs := []getContainersTestCase{
		buildGetContainersDataTestCase(
			t,
			"Full Node",
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
			"Execution Only",
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
			"No Execution",
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
			"case_executionNF",
			true,
			executionNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{{}}},
		),
		buildGetContainersDataTestCase(
			t,
			"Consensus Not Found",
			"case_consensusNF",
			true,
			consensusNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{{}}},
		),
		buildGetContainersDataTestCase(
			t,
			"Validator Not Found",
			"case_validatorNF",
			true,
			validatorNotFoundErrorMsg,
			actions.ContainersData{Containers: []actions.ContainerData{{}}},
		),
	}

	actions := getMockActions(t)
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
