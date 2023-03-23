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
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStartError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
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
	dockerClient := sedge_mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerStart(context.Background(), "validator-client", gomock.Any()).
		Return(nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	err := serviceManager.Start("validator-client")
	assert.Nil(t, err)
}
