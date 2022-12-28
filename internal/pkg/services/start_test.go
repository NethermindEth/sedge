package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStartError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerStart(context.Background(), "validator-client", gomock.Any()).
		Return(errors.New("error")).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	err := serviceManager.Start("validator-client")
	assert.ErrorIs(t, err, services.ErrStartingContainer)
}

func TestStartWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerStart(context.Background(), "validator-client", gomock.Any()).
		Return(nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	err := serviceManager.Start("validator-client")
	assert.Nil(t, err)
}
