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
package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/docker/docker/api/types/container"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWaitErrCh(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	waitCh := time.After(3 * time.Second)
	wantErr := errors.New("error")
	wantErrCh := make(chan error, 1)
	wantErrCh <- wantErr

	dockerClient.EXPECT().
		ContainerWait(context.Background(), "sedge-validator-client", gomock.Any()).
		Return(make(chan container.WaitResponse), wantErrCh).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	exitCh, errCh := serviceManager.Wait(services.DefaultSedgeValidatorClient, container.WaitConditionNextExit)
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
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	waitCh := time.After(3 * time.Second)
	wantWait := container.WaitResponse{
		StatusCode: 0,
	}
	wantWaitCh := make(chan container.WaitResponse, 1)
	wantWaitCh <- wantWait

	dockerClient.EXPECT().
		ContainerWait(context.Background(), "sedge-validator-client", gomock.Any()).
		Return(wantWaitCh, make(chan error)).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	exitCh, errCh := serviceManager.Wait(services.DefaultSedgeValidatorClient, container.WaitConditionNextExit)
	select {
	case <-waitCh:
		t.Fatal("exit channel timeout")
	case exit := <-exitCh:
		assert.Equal(t, wantWait.StatusCode, exit.StatusCode)
	case <-errCh:
		t.Fatal("unexpected value from error channel")
	}
}
