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

	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIsRunningError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedError := errors.New("error")

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "validator-client").
		Return(types.ContainerJSON{}, expectedError).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	ok, err := serviceManager.IsRunning("validator-client")
	assert.ErrorIs(t, err, expectedError)
	assert.False(t, ok)
}

func TestIsRunningFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "validator-client").
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				State: &types.ContainerState{
					Running: false,
				},
			},
		}, nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	ok, err := serviceManager.IsRunning("validator-client")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestIsRunningTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "validator-client").
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				State: &types.ContainerState{
					Running: true,
				},
			},
		}, nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	ok, err := serviceManager.IsRunning("validator-client")
	assert.Nil(t, err)
	assert.True(t, ok)
}
