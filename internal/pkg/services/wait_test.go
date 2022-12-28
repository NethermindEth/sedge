package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types/container"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWaitErrCh(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	waitCh := time.After(3 * time.Second)
	wantErr := errors.New("error")
	wantErrCh := make(chan error, 1)
	wantErrCh <- wantErr

	dockerClient.EXPECT().
		ContainerWait(context.Background(), "validator-client", gomock.Any()).
		Return(make(chan container.ContainerWaitOKBody), wantErrCh).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	exitCh, errCh := serviceManager.Wait(services.ServiceCtValidator, container.WaitConditionNextExit)
	select {
	case <-waitCh:
		t.Fatal("err channel timeout")
	case <-exitCh:
		t.Fatal("unexpected value from exit channel")
	case err := <-errCh:
		assert.ErrorIs(t, err, wantErr)
	}
}

func TestWaitExitCh(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	waitCh := time.After(3 * time.Second)
	wantWait := container.ContainerWaitOKBody{
		StatusCode: 0,
	}
	wantWaitCh := make(chan container.ContainerWaitOKBody, 1)
	wantWaitCh <- wantWait

	dockerClient.EXPECT().
		ContainerWait(context.Background(), "validator-client", gomock.Any()).
		Return(wantWaitCh, make(chan error)).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	exitCh, errCh := serviceManager.Wait(services.ServiceCtValidator, container.WaitConditionNextExit)
	select {
	case <-waitCh:
		t.Fatal("exit channel timeout")
	case exit := <-exitCh:
		assert.Equal(t, wantWait.StatusCode, exit.StatusCode)
	case <-errCh:
		t.Fatal("unexpected value from error channel")
	}
}
