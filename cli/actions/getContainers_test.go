package actions_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	executionContainerName = "sedge-execution-client"
	consensusContainerName = "sedge-consensus-client"
	validatorContainerName = "sedge-validator-client"
	mevBoostContainerName  = "sedge-mev-boost"
)

type getContainersTestCase struct {
	name                     string
	getContainersDataOptions actions.GetContainersDataOptions
	expected                 actions.ContainersData
	isErr                    bool
}

func buildGetContainersDataTestCase(
	t *testing.T,
	name string,
	caseDataDir string,
	isErr bool,
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
		expected: expected,
		isErr:    isErr,
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

	// TODO: add mocks for other methods
	// dockerClient.EXPECT().ContainerInspect(gomock.Any(), executionContainerName).Return(nil, nil).AnyTimes()

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
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  mevBoostContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  consensusContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  validatorContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Execution Only",
			"case_executionOnly",

			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Consensus Only",
			"case_consensusOnly",

			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  consensusContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"Validator Only",
			"case_validatorOnly",
			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  validatorContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Execution",
			"case_noExecution",
			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  mevBoostContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  consensusContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  validatorContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Consensus",
			"case_noConsensus",
			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  mevBoostContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  validatorContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Mev Boost",
			"case_noMev",
			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  consensusContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  validatorContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
		buildGetContainersDataTestCase(
			t,
			"No Validator",
			"case_noValidator",
			false,
			actions.ContainersData{
				Containers: []actions.ContainerData{
					{
						Name:  executionContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  mevBoostContainerName,
						Image: "",
						Ip:    "",
					},
					{
						Name:  consensusContainerName,
						Image: "",
						Ip:    "",
					},
				},
			},
		),
	}

	actions := getMockActions(t)
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			containersData, err := actions.GetContainersData(tc.getContainersDataOptions)
			if tc.isErr {
				assert.Error(t, err)
			} else if !tc.isErr && assert.NoError(t, err) {
				validateContainersData(t, tc.expected, containersData)
			}

		})
	}
}

func validateContainersData(t *testing.T, expected, actual actions.ContainersData) {
	assert.Equal(t, expected, actual)
}
