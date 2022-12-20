package actions_test

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSlashingImport_ValidatorNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
		}).
		Return(make([]types.Container, 0), nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	s := actions.NewSedgeActions(dockerClient, serviceManager)

	err := s.ImportSlashingInterchangeData(actions.SlashingImportOptions{})
	assert.ErrorIs(t, err, services.ErrContainerNotFound)
}

func TestSlashingImport_CheckValidatorFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantError := errors.New("error")

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
		}).
		Return([]types.Container{
			{ID: "validatorctid"},
		}, nil).
		Times(1)
	dockerClient.EXPECT().
		ContainerInspect(gomock.Any(), services.ServiceCtValidator).
		Return(types.ContainerJSON{}, wantError).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	s := actions.NewSedgeActions(dockerClient, serviceManager)

	err := s.ImportSlashingInterchangeData(actions.SlashingImportOptions{})
	assert.ErrorIs(t, err, wantError)
}

func TestSlashingImport_ValidatorStopFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	validatorCtId := "validatorctid"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
		}).
		Return([]types.Container{
			{ID: validatorCtId},
		}, nil).
		Times(1)
	dockerClient.EXPECT().
		ContainerInspect(gomock.Any(), services.ServiceCtValidator).
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				ID: validatorCtId,
				State: &types.ContainerState{
					Running: true,
				},
			},
		}, nil).
		Times(2)
	dockerClient.EXPECT().
		ContainerStop(gomock.Any(), "validatorctid", gomock.Any()).
		Return(errors.New("error")).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	s := actions.NewSedgeActions(dockerClient, serviceManager)

	err := s.ImportSlashingInterchangeData(actions.SlashingImportOptions{})
	assert.ErrorIs(t, err, services.ErrStoppingContainer)
}

func TestSlashingImport_ValidatorRunning(t *testing.T) {
	clients := []string{"prysm", "lighthouse", "lodestar", "teku"}
	for _, validatorClient := range clients {
		t.Run(validatorClient, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mock_client.NewMockAPIClient(ctrl)
			defer ctrl.Finish()

			validatorCtId := "validatorctid"
			slashingCtName := "validator-slashing-data"
			slashingCtId := "slashing-ct-id"

			// Mock ContainerList
			dockerClient.EXPECT().
				ContainerList(gomock.Any(), types.ContainerListOptions{
					All:     true,
					Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
				}).
				Return([]types.Container{
					{ID: validatorCtId},
				}, nil).
				Times(1)
			// Mock ContainerInspect
			dockerClient.EXPECT().
				ContainerInspect(gomock.Any(), services.ServiceCtValidator).
				Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						ID: validatorCtId,
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil).
				Times(3)
			// Mock ContainerStop
			dockerClient.EXPECT().
				ContainerStop(gomock.Any(), validatorCtId, gomock.Any()).
				Return(nil).
				Times(1)
			// Mock ContainerCreate
			dockerClient.EXPECT().
				ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), slashingCtName).
				Return(container.ContainerCreateCreatedBody{ID: slashingCtId}, nil).
				Times(1)
			// Mock ContainerStart
			dockerClient.EXPECT().
				ContainerStart(gomock.Any(), slashingCtId, gomock.Any()).
				Return(nil).
				Times(1)
			dockerClient.EXPECT().
				ContainerStart(gomock.Any(), services.ServiceCtValidator, gomock.Any()).
				Return(nil).
				Times(1)
			// Mock ContainerWait
			exitCh := make(chan container.ContainerWaitOKBody, 1)
			exitCh <- container.ContainerWaitOKBody{
				StatusCode: 0,
			}
			dockerClient.EXPECT().
				ContainerWait(gomock.Any(), slashingCtName, container.WaitConditionNextExit).
				Return(exitCh, make(chan error)).
				Times(1)
			// Mock ContainerRemove
			dockerClient.EXPECT().
				ContainerRemove(gomock.Any(), slashingCtId, types.ContainerRemoveOptions{}).
				Return(nil).
				Times(1)

			serviceManager := services.NewServiceManager(dockerClient)
			s := actions.NewSedgeActions(dockerClient, serviceManager)

			generationPath := t.TempDir()
			from := setupSlashingDataFile(t)
			copiedFile := filepath.Join(generationPath, configs.ValidatorDir, actions.SlashingImportFile)
			err := s.ImportSlashingInterchangeData(actions.SlashingImportOptions{
				ValidatorClient: validatorClient,
				Network:         "sepolia",
				GenerationPath:  generationPath,
				From:            from,
			})
			assert.Nil(t, err)
			assert.FileExists(t, copiedFile)
			copiedData, err := ioutil.ReadFile(copiedFile)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, []byte(slashingFile), copiedData)
		})
	}
}

const (
	slashingFile = ""
)

func setupSlashingDataFile(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "slashing-data.json")
	if err := ioutil.WriteFile(path, []byte(slashingFile), 0o777); err != nil {
		t.Fatal(err.Error())
	}
	return path
}
