package actions_test

import (
	"errors"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/golang/mock/gomock"
)

func validatorNotFoundHelper(t *testing.T, ctrl *gomock.Controller) actions.SedgeActions {
	t.Helper()
	dockerClient := mock_client.NewMockAPIClient(ctrl)

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
		}).
		Return(make([]types.Container, 0), nil).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	return actions.NewSedgeActions(dockerClient, serviceManager, nil)
}

func checkValidatorFailureHelper(t *testing.T, ctrl *gomock.Controller, wantError error) actions.SedgeActions {
	dockerClient := mock_client.NewMockAPIClient(ctrl)

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
	return actions.NewSedgeActions(dockerClient, serviceManager, nil)
}

func validatorStopFailureHelper(t *testing.T, ctrl *gomock.Controller) actions.SedgeActions {
	dockerClient := mock_client.NewMockAPIClient(ctrl)

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
	return actions.NewSedgeActions(dockerClient, serviceManager, nil)
}
