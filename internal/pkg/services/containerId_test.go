package services_test

import (
	"errors"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestContainerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantId := "container-id"
	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return([]types.Container{
			{
				ID: wantId,
			},
		}, nil).
		Times(1)
	serviceManager := services.NewServiceManager(dockerClient)
	id, err := serviceManager.ContainerId(containerName)
	assert.Nil(t, err)
	assert.Equal(t, wantId, id)
}

func TestContainerIdError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantErr := errors.New("error")
	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return(nil, wantErr).
		Times(1)
	serviceManager := services.NewServiceManager(dockerClient)
	id, err := serviceManager.ContainerId(containerName)
	assert.ErrorIs(t, err, wantErr)
	assert.Equal(t, "", id)
}

func TestContainerIdNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return(make([]types.Container, 0), nil).
		Times(1)
	serviceManager := services.NewServiceManager(dockerClient)
	id, err := serviceManager.ContainerId(containerName)
	assert.ErrorIs(t, err, services.ErrContainerNotFound)
	assert.Equal(t, "", id)
}

func TestContainerIdMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return(make([]types.Container, 2), nil).
		Times(1)
	serviceManager := services.NewServiceManager(dockerClient)
	id, err := serviceManager.ContainerId(containerName)
	assert.ErrorIs(t, err, services.ErrMultipleContainers)
	assert.Equal(t, "", id)
}
