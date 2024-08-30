package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	mocks "github.com/NethermindEth/sedge/mocks"
)

// Image tests

func TestImageFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedImage := "expected-image"

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				Image: expectedImage,
			},
		}, nil)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	image, err := dockerServicemanager.Image("sedge")
	assert.Nil(t, err)
	assert.Equal(t, expectedImage, image)
}

func TestImageNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedError := errdefs.NotFound(fmt.Errorf("error"))

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{}, expectedError)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	image, err := dockerServicemanager.Image("sedge")
	assert.ErrorIs(t, err, expectedError)
	assert.True(t, errdefs.IsNotFound(err))
	assert.Equal(t, "", image)
}

func ExampleDockerServiceManager_Image() {
	// Create a new Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	// Create a new dockerServicemanager
	dm := NewDockerServiceManager(dockerClient)

	// Get the image name of a running container
	_, err = dm.Image("myContainer")
	if err != nil {
		log.Error(err)
	}
}

// Start tests

func TestStartError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerStart(context.Background(), "sedge", gomock.Any()).
		Return(errors.New("error")).
		Times(1)

	dockerServicemanager := services.NewDockerServiceManager(dockerClient)
	err := dockerServicemanager.Start("sedge")
	assert.ErrorIs(t, err, services.ErrStartingContainer)
}

func TestStartWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerStart(context.Background(), "sedge", gomock.Any()).
		Return(nil).
		Times(1)

	dockerServicemanager := services.NewDockerServiceManager(dockerClient)
	err := dockerServicemanager.Start("sedge")
	assert.Nil(t, err)
}

// Stop tests

func TestStopContainerNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedError := errdefs.NotFound(errors.New("error"))

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{}, expectedError)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	err := dockerServicemanager.Stop("sedge")
	assert.Nil(t, err)
}

func TestStopError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedError := errors.New("error")

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{}, expectedError)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	err := dockerServicemanager.Stop("sedge")
	assert.ErrorIs(t, err, expectedError)
}

func TestStopContainerAlreadyStopped(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				State: &types.ContainerState{
					Running: false,
				},
			},
		}, nil).
		Times(1)

	dockerServiceManager := services.NewDockerServiceManager(dockerClient)
	err := dockerServiceManager.Stop("sedge")
	assert.Nil(t, err)
}

func TestStopContainer(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tests := []struct {
		name           string
		containerState *types.ContainerState
	}{
		{
			name: "Running status success",
			containerState: &types.ContainerState{
				Running: true,
			},
		},
		{
			name: "Restarting status success",
			containerState: &types.ContainerState{
				Restarting: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			defer ctrl.Finish()

			sedgeCtId := "sedgectid"
			dockerClient.EXPECT().
				ContainerInspect(context.Background(), "sedge").
				Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						ID:    sedgeCtId,
						State: tt.containerState,
					},
				}, nil)
			dockerClient.EXPECT().
				ContainerStop(context.Background(), sedgeCtId, gomock.Any()).
				Return(nil)

			dockerServicemanager := services.NewDockerServicemanager(dockerClient)
			err := dockerServicemanager.Stop("sedge")
			assert.Nil(t, err)
		})
	}
}

func TestStopContainerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	expectedError := errors.New("error")
	validatorCtId := "sedgectid"

	dockerClient.EXPECT().
		ContainerInspect(context.Background(), "sedge").
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				ID: validatorCtId,
				State: &types.ContainerState{
					Running: true,
				},
			},
		}, nil).
		Times(1)
	dockerClient.EXPECT().
		ContainerStop(context.Background(), validatorCtId, gomock.Any()).
		Return(expectedError).
		Times(1)

	serviceManager := services.NewServiceManager(dockerClient)
	err := serviceManager.Stop("sedge")
	assert.ErrorIs(t, err, services.ErrStoppingContainer)
}

// ContainerID tests

func TestContainerId(t *testing.T) {
	ctName := "container-name"
	tests := []struct {
		name       string
		containers []types.Container
		wantId     string
		err        error
	}{
		{
			name: "container found",
			containers: []types.Container{
				{
					ID:    "other-id",
					Names: []string{"other-name"},
				},
				{
					ID:    "container-id",
					Names: []string{"/" + ctName},
				},
			},
			wantId: "container-id",
			err:    nil,
		},
		{
			name:       "container not found, no containers",
			containers: []types.Container{},
			wantId:     "",
			err:        services.ErrContainerNotFound,
		},
		{
			name: "container found, no exact match",
			containers: []types.Container{
				{
					ID:    "other-id",
					Names: []string{"other-name"},
				},
				{
					ID:    "container-id",
					Names: []string{ctName + "-2", ctName + "-3"},
				},
			},
			wantId: "",
			err:    services.ErrContainerNotFound,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		dockerClient := mocks.NewMockAPIClient(ctrl)
		defer ctrl.Finish()

		dockerClient.EXPECT().
			ContainerList(gomock.Any(), types.ContainerListOptions{
				All:     true,
				Filters: filters.NewArgs(filters.Arg("name", ctName)),
			}).
			Return(tt.containers, nil).
			Times(1)
		dockerServiceManager := services.NewDockerServiceManager(dockerClient)
		id, err := dockerServiceManager.ContainerId(ctName)
		assert.ErrorIs(t, err, tt.err)
		assert.Equal(t, tt.wantId, id)

	}
}

func TestContainerIdError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantErr := errors.New("error")
	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return(nil, wantErr)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	id, err := dockerServicemanager.ContainerID(containerName)
	assert.ErrorIs(t, err, wantErr)
	assert.Equal(t, "", id)
}

func TestContainerIdNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	containerName := "container-name"

	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", containerName)),
		}).
		Return(make([]types.Container, 0), nil)

	dockerServicemanager := services.NewDockerServicemanager(dockerClient)
	id, err := dockerServicemanager.ContainerID(containerName)
	assert.ErrorIs(t, err, services.ErrContainerNotFound)
	assert.Equal(t, "", id)
}

// Pull tests

func TestPullError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantErr := errors.New("error")
	imageName := "sedge:latest"

	dockerClient.EXPECT().
		ImagePull(gomock.Any(), imageName, gomock.Any()).
		Return(nil, wantErr)

	dockerServicemanager := NewDockerServicemanager(dockerClient)
	err := dockerServicemanager.Pull(imageName)
	assert.ErrorIs(t, err, wantErr)
}

func TestPull(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	imageName := "sedge:latest"

	dockerClient.EXPECT().
		ImagePull(gomock.Any(), imageName, gomock.Any()).
		Return(nil, nil)

	dockerServicemanager := NewDockerServicemanager(dockerClient)
	err := dockerServicemanager.Pull(imageName)
	assert.Nil(t, err)
}

// ContainerLogs tests

func TestContainerLogsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	wantErr := errors.New("error")
	containerName := "sedge"

	dockerClient.EXPECT().
		ContainerLogs(gomock.Any(), containerName, gomock.Any()).
		Return(nil, wantErr)

	dockerServicemanager := NewDockerServicemanager(dockerClient)
	logs, err := dockerServicemanager.ContainerLogs(containerName, containerName)
	assert.ErrorIs(t, err, wantErr)
	assert.Empty(t, logs)
}

type mockReadCloser struct {
	io.Reader
}

func (mockReadCloser) Close() error { return nil }

// Ensure mockReadCloser implements io.ReadCloser
var _ io.ReadCloser = (*mockReadCloser)(nil)

func TestContainerLogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	containerName := "sedge"
	wantLogs := "logs"

	logReader := mockReadCloser{Reader: strings.NewReader(wantLogs)}

	dockerClient.EXPECT().
		ContainerLogs(gomock.Any(), containerName, containerName, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     false,
		}).
		Return(logReader, nil)

	dockerServicemanager := NewDockerServicemanager(dockerClient)
	logs, err := dockerServicemanager.ContainerLogs(containerName, containerName)
	assert.Nil(t, err)
	assert.Equal(t, wantLogs, logs)
}

func TestContainerLogsMerged(t *testing.T) {
	const serviceLogs = `--------INFO:     Started server process [1]
--------INFO:     Waiting for application startup.
--------INFO:     Application startup complete.
--------INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)
--------INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect
--------INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`
	type testCase struct {
		name      string
		services  map[string]string
		opts      ContainerLogsMergedOptions
		mocker    func(t *testing.T, dockerClient *mocks.MockAPIClient)
		wantLines []string
		wantErr   error
	}

	tc := []testCase{
		{
			name: "success with 3 services",
			services: map[string]string{
				"service-1": "id1",
				"service-2": "id2",
				"service-3": "id3",
			},
			mocker: func(t *testing.T, dockerClient *mocks.MockAPIClient) {
				for i := 0; i < 3; i++ {
					serviceReader := new(bytes.Buffer)
					serviceReader.WriteString(serviceLogs)
					dockerClient.EXPECT().ContainerLogs(context.Background(), fmt.Sprintf("id%d", i+1), types.ContainerLogsOptions{
						ShowStdout: true,
						ShowStderr: true,
					}).Return(io.NopCloser(serviceReader), nil)
				}
			},
			wantLines: []string{
				"service-1: INFO:     Started server process [1]",
				"service-1: INFO:     Waiting for application startup.",
				"service-1: INFO:     Application startup complete.",
				"service-1: INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)",
				`service-1: INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect`,
				`service-1: INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`,
				"service-2: INFO:     Started server process [1]",
				"service-2: INFO:     Waiting for application startup.",
				"service-2: INFO:     Application startup complete.",
				"service-2: INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)",
				`service-2: INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect`,
				`service-2: INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`,
				"service-3: INFO:     Started server process [1]",
				"service-3: INFO:     Waiting for application startup.",
				"service-3: INFO:     Application startup complete.",
				"service-3: INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)",
				`service-3: INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect`,
				`service-3: INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`,
			},
		},
		{
			name:    "3 services, 1 error",
			wantErr: fmt.Errorf("error getting logs for service-3: %w", assert.AnError),
			services: map[string]string{
				"service-1": "id1",
				"service-2": "id2",
				"service-3": "id3",
			},
			mocker: func(t *testing.T, dockerClient *mocks.MockAPIClient) {
				for i := 0; i < 2; i++ {
					serviceReader := new(bytes.Buffer)
					serviceReader.WriteString(serviceLogs)
					dockerClient.EXPECT().ContainerLogs(context.Background(), fmt.Sprintf("id%d", i+1), types.ContainerLogsOptions{
						ShowStdout: true,
						ShowStderr: true,
					}).Return(io.NopCloser(serviceReader), nil)
				}
				dockerClient.EXPECT().ContainerLogs(context.Background(), "id3", types.ContainerLogsOptions{
					ShowStdout: true,
					ShowStderr: true,
				}).Return(nil, assert.AnError)
			},
			wantLines: []string{
				"service-1: INFO:     Started server process [1]",
				"service-1: INFO:     Waiting for application startup.",
				"service-1: INFO:     Application startup complete.",
				"service-1: INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)",
				`service-1: INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect`,
				`service-1: INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`,
				"service-2: INFO:     Started server process [1]",
				"service-2: INFO:     Waiting for application startup.",
				"service-2: INFO:     Application startup complete.",
				"service-2: INFO:     Uvicorn running on http://0.0.0.0:8080 (Press CTRL+C to quit)",
				`service-2: INFO:     172.29.0.3:38684 - "GET /metrics HTTP/1.1" 307 Temporary Redirect`,
				`service-2: INFO:     172.29.0.3:38684 - "GET / HTTP/1.1" 200 OK`,
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dockerClient := mocks.NewMockAPIClient(ctrl)
			tt.mocker(t, dockerClient)
			out := new(bytes.Buffer)
			dockerServicemanager := NewDockerServicemanager(dockerClient)
			err := dockerServicemanager.ContainerLogsMerged(context.Background(), out, tt.services, tt.opts)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				allOut, err := io.ReadAll(out)
				assert.NoError(t, err)
				for _, line := range tt.wantLines {
					assert.Contains(t, string(allOut), line)
				}
				assert.Len(t, strings.Split(string(allOut), "\n"), len(tt.wantLines)+1)
			}
		})
	}
}

// Wait tests
func TestWaitErrCh(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
	defer ctrl.Finish()

	waitCh := time.After(3 * time.Second)
	wantErr := errors.New("error")
	wantErrCh := make(chan error, 1)
	wantErrCh <- wantErr

	dockerClient.EXPECT().
		ContainerWait(context.Background(), "sedge-validator-client", gomock.Any()).
		Return(make(chan container.WaitResponse), wantErrCh).
		Times(1)

	dockerServiceManager := services.NewDockerServiceManager(dockerClient)
	exitCh, errCh := dockerServiceManager.Wait(services.DefaultSedgeValidatorClient, container.WaitConditionNextExit)
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
	tests := []struct {
		name     string
		response container.WaitResponse
	}{
		{
			name: "response without error",
			response: container.WaitResponse{
				StatusCode: 0,
			},
		},
		{
			name: "response with error",
			response: container.WaitResponse{
				StatusCode: 0,
				Error: &container.WaitExitError{
					Message: "error",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			defer ctrl.Finish()

			waitCh := time.After(3 * time.Second)
			wantWaitCh := make(chan container.WaitResponse, 1)
			wantWaitCh <- tt.response
			close(wantWaitCh)

			dockerClient.EXPECT().
				ContainerWait(context.Background(), "sedge", gomock.Any()).
				Return(wantWaitCh, make(chan error)).
				Times(1)

			dockerServicemanager := services.NewDockerServicemanager(dockerClient)
			exitCh, errCh := dockerServicemanager.Wait("sedge", container.WaitConditionNextExit)
			select {
			case <-waitCh:
				t.Fatal("exit channel timeout")
			case exit := <-exitCh:
				assert.Equal(t, tt.response.StatusCode, exit.StatusCode)
				if tt.response.Error == nil {
					assert.Nil(t, exit.Error)
				} else {
					assert.Equal(t, tt.response.Error.Message, exit.Error.Message)
				}
			case <-errCh:
				t.Fatal("unexpected value from error channel")
			}
		})
	}
}

func TestContainerStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		want     common.Status
		errorMsg string
		wantErr  bool
	}{
		{
			name:   "container created",
			status: "created",
			want:   common.Created,
		},
		{
			name:   "container running",
			status: "running",
			want:   common.Running,
		},
		{
			name:   "container paused",
			status: "paused",
			want:   common.Paused,
		},
		{
			name:   "container restarting",
			status: "restarting",
			want:   common.Restarting,
		},
		{
			name:   "container removing",
			status: "removing",
			want:   common.Removing,
		},
		{
			name:   "container exited",
			status: "exited",
			want:   common.Exited,
		},
		{
			name:   "container dead",
			status: "dead",
			want:   common.Dead,
		},
		{
			name:     "container unknown",
			status:   "unknown",
			want:     common.Unknown,
			errorMsg: "unknown container status: unknown",
		},
		{
			name:     "bad container",
			want:     common.Unknown,
			wantErr:  true,
			errorMsg: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			defer ctrl.Finish()

			container := types.ContainerJSON{
				ContainerJSONBase: &types.ContainerJSONBase{
					State: &types.ContainerState{
						Status: tt.status,
					},
				},
			}

			if tt.wantErr {
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container.ID).
					Return(container, errors.New("error"))
			} else {
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container.ID).
					Return(container, nil)
			}

			dockerServicemanager := NewDockerServicemanager(dockerClient)
			status, err := dockerServicemanager.ContainerStatus(container.ID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
			} else {
				assert.Equal(t, tt.want, status)
			}
		})
	}
}

func TestPS(t *testing.T) {
	tests := []struct {
		name       string
		containers []types.Container
		want       []ContainerInfo
		wantErr    error
	}{
		{
			name: "one container found",
			containers: []types.Container{
				{
					ID:      "container-id1",
					Names:   []string{"/name1"},
					Image:   "image1",
					ImageID: "image-id1",
					Command: "command1",
					Created: 1234,
					Ports: []types.Port{
						{
							IP:          "127.0.0.1",
							PrivatePort: 3000,
							PublicPort:  3080,
						},
					},
					Status: "running",
					State:  "state1",
				},
			},
			want: []ContainerInfo{
				{
					ID:      "container-id1",
					Names:   []string{"/name1"},
					Image:   "image1",
					Command: "command1",
					Created: 1234,
					Ports: []Port{
						{
							IP:          "127.0.0.1",
							PrivatePort: 3000,
							PublicPort:  3080,
						},
					},
					Status: "running",
				},
			},
			wantErr: nil,
		},
		{
			name: "many containers found",
			containers: []types.Container{
				{
					ID:      "container-id1",
					Names:   []string{"/name1"},
					Image:   "image1",
					ImageID: "image-id1",
					Command: "command1",
					Created: 1234,
					Ports: []types.Port{
						{
							IP:          "127.0.0.1",
							PrivatePort: 3000,
							PublicPort:  3080,
						},
					},
					Status: "running",
					State:  "state1",
				},
				{
					ID:      "container-id2",
					Names:   []string{"/name2"},
					Image:   "image2",
					ImageID: "image-id2",
					Command: "command2",
					Created: 5678,
					Ports: []types.Port{
						{
							IP:          "127.0.0.10",
							PrivatePort: 4000,
							PublicPort:  4080,
						},
					},
					Status: "running",
					State:  "state2",
				},
			},
			want: []ContainerInfo{
				{
					ID:      "container-id1",
					Names:   []string{"/name1"},
					Image:   "image1",
					Command: "command1",
					Created: 1234,
					Ports: []Port{
						{
							IP:          "127.0.0.1",
							PrivatePort: 3000,
							PublicPort:  3080,
						},
					},
					Status: "running",
				},
				{
					ID:      "container-id2",
					Names:   []string{"/name2"},
					Image:   "image2",
					Command: "command2",
					Created: 5678,
					Ports: []Port{
						{
							IP:          "127.0.0.10",
							PrivatePort: 4000,
							PublicPort:  4080,
						},
					},
					Status: "running",
				},
			},
		},
		{
			name:       "none containers",
			containers: []types.Container{},
			want:       []ContainerInfo{},
		},
		{
			name:       "returning error",
			containers: []types.Container{},
			want:       nil,
			wantErr:    errors.New("error listing containers"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			defer ctrl.Finish()

			dockerClient.EXPECT().ContainerList(gomock.Any(), types.ContainerListOptions{}).Return(tt.containers, tt.wantErr)

			dockerServicemanager := NewDockerServicemanager(dockerClient)
			got, err := dockerServicemanager.PS()
			assert.ErrorIs(t, err, tt.wantErr, "Unexpected error returned")
			assert.Equal(t, tt.want, got, "Expected containers does not match with containers obtained.")
		})
	}
}

func TestContainerIP(t *testing.T) {
	tests := []struct {
		name        string
		arg         string
		response    types.ContainerJSON
		want        string
		wantErr     error
		expectedErr error
	}{
		{
			name: "returning error",
			arg:  "sedge",
			response: types.ContainerJSON{
				ContainerJSONBase: &types.ContainerJSONBase{
					ID: "container-Id",
					State: &types.ContainerState{
						Running: true,
					},
				},
				NetworkSettings: &types.NetworkSettings{
					Networks: map[string]*network.EndpointSettings{
						"sedge": {
							IPAddress: "127.0.0.1",
						},
					},
				},
			},
			want:    "",
			wantErr: errors.New("error inspecting container"),
		},
		{
			name: "checking empty networks",
			arg:  "sedge-network",
			response: types.ContainerJSON{
				ContainerJSONBase: &types.ContainerJSONBase{
					ID: "container-Id",
					State: &types.ContainerState{
						Running: true,
					},
				},
				NetworkSettings: &types.NetworkSettings{
					Networks: map[string]*network.EndpointSettings{},
				},
			},
			want:        "",
			expectedErr: ErrNetworksNotFound,
		},
		{
			name: "returning correct IP",
			arg:  "sedge-network",
			response: types.ContainerJSON{
				ContainerJSONBase: &types.ContainerJSONBase{
					ID: "container-Id",
					State: &types.ContainerState{
						Running: true,
					},
				},
				NetworkSettings: &types.NetworkSettings{
					Networks: map[string]*network.EndpointSettings{
						"sedge": {
							IPAddress: "127.0.0.1",
						},
					},
				},
			},
			want: "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			defer ctrl.Finish()
			dockerClient.EXPECT().
				ContainerInspect(context.Background(), tt.arg).
				Return(tt.response, tt.wantErr)

			dockerServicemanager := NewDockerServicemanager(dockerClient)
			got, err := dockerServicemanager.ContainerIP(tt.arg)
			if tt.wantErr == nil {
				assert.ErrorIs(t, err, tt.expectedErr, "Unexpected error returned")
			} else {
				assert.ErrorIs(t, err, tt.wantErr, "Unexpected error returned")
			}
			assert.Equal(t, tt.want, got, "Expected container IP does not match with IP obtained.")
		})
	}
}

func TestContainerNetworks(t *testing.T) {
	tests := []struct {
		name      string
		mocker    func(t *testing.T, container string) *mocks.MockAPIClient
		container string
		want      []string
		wantErr   bool
	}{
		{
			name: "ok, 1 network",
			mocker: func(t *testing.T, container string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container).
					Return(types.ContainerJSON{
						NetworkSettings: &types.NetworkSettings{
							Networks: map[string]*network.EndpointSettings{
								"sedge": {
									IPAddress: "127.0.0.1",
								},
							},
						},
					}, nil)
				return dockerClient
			},
			container: "container-Id1",
			want:      []string{"sedge"},
			wantErr:   false,
		},
		{
			name: "ok, 2 networks",
			mocker: func(t *testing.T, container string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container).
					Return(types.ContainerJSON{
						NetworkSettings: &types.NetworkSettings{
							Networks: map[string]*network.EndpointSettings{
								"sedge": {
									IPAddress: "168.0.0.1",
								},
								"sedge2": {
									IPAddress: "168.0.0.2",
								},
							},
						},
					}, nil)
				return dockerClient
			},
			container: "container-Id2",
			want:      []string{"sedge", "sedge2"},
			wantErr:   false,
		},
		{
			name: "empty networks",
			mocker: func(t *testing.T, container string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container).
					Return(types.ContainerJSON{
						NetworkSettings: &types.NetworkSettings{
							Networks: map[string]*network.EndpointSettings{},
						},
					}, nil)
				return dockerClient
			},
			container: "container-Id3",
			wantErr:   true,
		},
		{
			name: "error inspecting container",
			mocker: func(t *testing.T, container string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					ContainerInspect(context.Background(), container).
					Return(types.ContainerJSON{}, errors.New("error inspecting container"))
				return dockerClient
			},
			container: "container-Id4",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dockerClient := tt.mocker(t, tt.container)
			dockerServicemanager := NewDockerServicemanager(dockerClient)
			got, err := dockerServicemanager.ContainerNetworks(tt.container)
			if tt.wantErr {
				assert.Error(t, err, "Expected error not returned")
			} else {
				assert.NoError(t, err, "Unexpected error returned")
			}
			assert.Len(t, got, len(tt.want), "Expected container networks does not match with networks obtained.")
			for _, network := range tt.want {
				assert.Contains(t, got, network, "Expected container networks does not match with networks obtained.")
			}
		})
	}
}

func TestNetworkConnect(t *testing.T) {
	tests := []struct {
		name      string
		mocker    func(t *testing.T, container, network string) *mocks.MockAPIClient
		container string
		network   string
		wantErr   bool
	}{
		{
			name: "ok",
			mocker: func(t *testing.T, container, network string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					NetworkConnect(context.Background(), network, container, nil).
					Return(nil)
				return dockerClient
			},
			container: "container-Id1",
			network:   "sedge",
		},
		{
			name: "error connecting container to network",
			mocker: func(t *testing.T, container, network string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					NetworkConnect(context.Background(), network, container, nil).
					Return(errors.New("error connecting container to network"))
				return dockerClient
			},
			container: "container-Id2",
			network:   "sedge",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dockerClient := tt.mocker(t, tt.container, tt.network)
			dockerServicemanager := NewDockerServicemanager(dockerClient)
			err := dockerServicemanager.NetworkConnect(tt.container, tt.network)
			if tt.wantErr {
				assert.Error(t, err, "Expected error not returned")
			} else {
				assert.NoError(t, err, "Unexpected error returned")
			}
		})
	}
}

func TestNetworkDisconnect(t *testing.T) {
	tests := []struct {
		name      string
		mocker    func(t *testing.T, container, network string) *mocks.MockAPIClient
		container string
		network   string
		wantErr   bool
	}{
		{
			name: "ok",
			mocker: func(t *testing.T, container, network string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					NetworkDisconnect(context.Background(), network, container, false).
					Return(nil)
				return dockerClient
			},
			container: "container-Id1",
			network:   "sedge",
		},
		{
			name: "error disconnecting container from network",
			mocker: func(t *testing.T, container, network string) *mocks.MockAPIClient {
				ctrl := gomock.NewController(t)
				dockerClient := mocks.NewMockAPIClient(ctrl)
				dockerClient.EXPECT().
					NetworkDisconnect(context.Background(), network, container, false).
					Return(errors.New("error disconnecting container from network"))
				return dockerClient
			},
			container: "container-Id2",
			network:   "sedge",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dockerClient := tt.mocker(t, tt.container, tt.network)
			dockerServicemanager := NewDockerServicemanager(dockerClient)
			err := dockerServicemanager.NetworkDisconnect(tt.container, tt.network)
			if tt.wantErr {
				assert.Error(t, err, "Expected error not returned")
			} else {
				assert.NoError(t, err, "Unexpected error returned")
			}
		})
	}
}

func TestBuildImageFromURI(t *testing.T) {
	type testCase struct {
		name          string
		remote        string
		tag           string
		buildArgs     map[string]*string
		setup         func(*mocks.MockAPIClient)
		expectedError error
	}
	tests := []testCase{
		{
			name:   "success",
			remote: "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
			tag:    "mock-avs-plugin",
			setup: func(dockerClient *mocks.MockAPIClient) {
				buildBody := io.NopCloser(bytes.NewReader([]byte{}))
				buildResponse := types.ImageBuildResponse{
					Body: buildBody,
				}
				loadBody := io.NopCloser(bytes.NewReader([]byte{}))
				loadResponse := types.ImageLoadResponse{
					Body: loadBody,
				}
				dockerClient.EXPECT().ImageBuild(context.Background(), nil, types.ImageBuildOptions{
					RemoteContext: "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
					Tags:          []string{"mock-avs-plugin"},
					Remove:        true,
					ForceRemove:   true,
				}).Return(buildResponse, nil)
				dockerClient.EXPECT().ImageLoad(context.Background(), buildResponse.Body, true).Return(loadResponse, nil)
			},
			expectedError: nil,
		},
		func(t *testing.T) testCase {
			buildArgs := map[string]*string{
				"key1": StringPtr("value1"),
				"key2": StringPtr("value2"),
			}
			return testCase{
				name:      "success, with build args",
				remote:    "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
				tag:       "mock-avs-plugin",
				buildArgs: buildArgs,
				setup: func(dockerClient *mocks.MockAPIClient) {
					buildBody := io.NopCloser(bytes.NewReader([]byte{}))
					buildResponse := types.ImageBuildResponse{
						Body: buildBody,
					}
					loadBody := io.NopCloser(bytes.NewReader([]byte{}))
					loadResponse := types.ImageLoadResponse{
						Body: loadBody,
					}
					dockerClient.EXPECT().ImageBuild(context.Background(), nil, types.ImageBuildOptions{
						RemoteContext: "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
						Tags:          []string{"mock-avs-plugin"},
						Remove:        true,
						ForceRemove:   true,
						BuildArgs:     buildArgs,
					}).Return(buildResponse, nil)
					dockerClient.EXPECT().ImageLoad(context.Background(), buildResponse.Body, true).Return(loadResponse, nil)
				},
				expectedError: nil,
			}
		}(t),
		{
			name:   "build error",
			remote: "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
			tag:    "mock-avs-plugin",
			setup: func(dockerClient *mocks.MockAPIClient) {
				dockerClient.EXPECT().ImageBuild(context.Background(), nil, types.ImageBuildOptions{
					RemoteContext: "https://github.com/NethermindEth/mock-avs-pkg#main:plugin",
					Tags:          []string{"mock-avs-plugin"},
					Remove:        true,
					ForceRemove:   true,
				}).Return(types.ImageBuildResponse{}, errors.New("build error"))
			},
			expectedError: errors.New("build error"),
		},
		{
			name:   "load error",
			remote: "https://github.com/orgname/avs#main:plugin",
			tag:    "orgname-avs-plugin",
			setup: func(dockerClient *mocks.MockAPIClient) {
				buildBody := io.NopCloser(bytes.NewReader([]byte{}))
				buildResponse := types.ImageBuildResponse{
					Body: buildBody,
				}
				dockerClient.EXPECT().ImageBuild(context.Background(), nil, types.ImageBuildOptions{
					RemoteContext: "https://github.com/orgname/avs#main:plugin",
					Tags:          []string{"orgname-avs-plugin"},
					Remove:        true,
					ForceRemove:   true,
				}).Return(buildResponse, nil)
				dockerClient.EXPECT().ImageLoad(context.Background(), buildResponse.Body, true).Return(types.ImageLoadResponse{}, errors.New("load error"))
			},
			expectedError: errors.New("load error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			tt.setup(dockerClient)
			defer ctrl.Finish()

			dockerServicemanager := NewDockerServicemanager(dockerClient)
			err := dockerServicemanager.BuildImageFromURI(tt.remote, tt.tag, tt.buildArgs)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLoadImageContext(t *testing.T) {
	tests := []struct {
		name    string
		path    func(t *testing.T) string
		wantErr error
	}{
		{
			name: "success",
			path: func(t *testing.T) string {
				absPath, err := filepath.Abs(t.TempDir())
				require.NoError(t, err)
				return absPath
			},
			wantErr: nil,
		},
		{
			name: "not absolute path",
			path: func(t *testing.T) string {
				return "not/absolute/path"
			},
			wantErr: errors.New("path must be absolute"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.path(t)
			dockerServicemanager := NewDockerServicemanager(nil)
			_, err := dockerServicemanager.LoadImageContext(path)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDockerServiceManager_Run(t *testing.T) {
	tests := []struct {
		name          string
		setupMock     func(*gomock.Controller) *mocks.MockAPIClient
		image         string
		options       RunOptions
		expectedError error
	}{
		{
			name: "Run successful",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				// Create channels
				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				// Write to one of the channels
				waitCh <- container.WaitResponse{StatusCode: 0}

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerWait(gomock.Any(), "containerID", gomock.Any()).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(gomock.Any(), "containerID", gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerLogs(gomock.Any(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBuffer([]byte{})), nil),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(nil),
				)
				return dockerClient
			},
			image:   "my-image",
			options: RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
		},
		{
			name: "Container create error",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)
				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{}, errors.New("creation error")),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: errors.New("creation error"),
		},
		{
			name: "Container remove error",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				// Create channels
				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				// Write to one of the channels
				waitCh <- container.WaitResponse{StatusCode: 0}

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerWait(gomock.Any(), "containerID", gomock.Any()).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(gomock.Any(), "containerID", gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerLogs(gomock.Any(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBuffer([]byte{})), nil),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(errors.New("remove error")),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: errors.New("remove error"),
		},
		{
			name: "NetworkConnect error",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)
				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(errors.New("network connection error")),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(nil),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: errors.New("network connection error"),
		},
		{
			name: "ContainerStart error",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				// Create channels
				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerWait(gomock.Any(), "containerID", gomock.Any()).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(gomock.Any(), "containerID", gomock.Any()).Return(errors.New("start container error")),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(nil),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: errors.New("start container error"),
		},
		{
			name: "ContainerWait error channel",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				// Create channels
				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				// Write to error channel
				errCh <- errors.New("container wait error")

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerWait(gomock.Any(), "containerID", gomock.Any()).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(gomock.Any(), "containerID", gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerLogs(gomock.Any(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBufferString("container logs")), nil),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(nil),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: errors.New("error waiting for container containerID: container wait error. container logs: container logs"),
		},
		{
			name: "Non-zero exit status",
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				// Create channels
				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				// Write to one of the channels
				waitCh <- container.WaitResponse{StatusCode: 1}

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(gomock.Any(), "my-network", gomock.Any(), gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerWait(gomock.Any(), "containerID", gomock.Any()).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(gomock.Any(), "containerID", gomock.Any()).Return(nil),
					dockerClient.EXPECT().ContainerLogs(gomock.Any(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBufferString("container logs")), nil),
					dockerClient.EXPECT().ContainerRemove(gomock.Any(), "containerID", gomock.Any()).Return(nil),
				)
				return dockerClient
			},
			image:         "my-image",
			options:       RunOptions{Network: "my-network", Args: []string{"arg1", "arg2"}},
			expectedError: fmt.Errorf("unexpected exit code 1 for container containerID. container logs: container logs"),
		},
		{
			name:    "Running on host network",
			image:   "my-image",
			options: RunOptions{Network: "host", Args: []string{}},
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				waitCh <- container.WaitResponse{StatusCode: 0}

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(context.Background(),
						&container.Config{Image: "my-image", Cmd: []string{}},
						&container.HostConfig{},
						gomock.Nil(),
						gomock.Nil(),
						"").
						Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().ContainerWait(context.Background(), "containerID", container.WaitConditionNextExit).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(context.Background(), "containerID", types.ContainerStartOptions{}).Return(nil),
					dockerClient.EXPECT().ContainerLogs(context.Background(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBuffer([]byte{})), nil),
					dockerClient.EXPECT().ContainerRemove(context.Background(), "containerID", types.ContainerRemoveOptions{}).Return(nil),
				)
				return dockerClient
			},
		},
		{
			name:  "with mounts",
			image: "my-image",
			options: RunOptions{
				Network: "my-network",
				Args:    []string{},
				Mounts: []Mount{
					{
						Type:   VolumeTypeBind,
						Source: "/home/user/dir",
						Target: "/container/dir1",
					},
					{
						Type:   VolumeTypeVolume,
						Source: "volume-name",
						Target: "/container/dir2",
					},
				},
			},
			setupMock: func(ctrl *gomock.Controller) *mocks.MockAPIClient {
				dockerClient := mocks.NewMockAPIClient(ctrl)

				waitCh := make(chan container.WaitResponse, 1)
				errCh := make(chan error, 1)

				waitCh <- container.WaitResponse{StatusCode: 0}

				gomock.InOrder(
					dockerClient.EXPECT().ContainerCreate(context.Background(),
						&container.Config{Image: "my-image", Cmd: []string{}},
						&container.HostConfig{
							Mounts: []mount.Mount{
								{
									Type:   mount.TypeBind,
									Source: "/home/user/dir",
									Target: "/container/dir1",
								},
								{
									Type:   mount.TypeVolume,
									Source: "volume-name",
									Target: "/container/dir2",
								},
							},
						},
						gomock.Nil(),
						gomock.Nil(),
						"").
						Return(container.CreateResponse{ID: "containerID"}, nil),
					dockerClient.EXPECT().NetworkConnect(context.Background(), "my-network", "containerID", gomock.Nil()).Return(nil),
					dockerClient.EXPECT().ContainerWait(context.Background(), "containerID", container.WaitConditionNextExit).Return(waitCh, errCh),
					dockerClient.EXPECT().ContainerStart(context.Background(), "containerID", types.ContainerStartOptions{}).Return(nil),
					dockerClient.EXPECT().ContainerLogs(context.Background(), "containerID", types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}).Return(io.NopCloser(bytes.NewBuffer([]byte{})), nil),
					dockerClient.EXPECT().ContainerRemove(context.Background(), "containerID", types.ContainerRemoveOptions{}).Return(nil),
				)
				return dockerClient
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dockerClient := tt.setupMock(ctrl)

			dockerServicemanager := NewDockerServicemanager(dockerClient)

			err := dockerServicemanager.Run(tt.image, tt.options)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

type notFoundError struct{}

func (e notFoundError) NotFound() {}

func (e notFoundError) Error() string {
	return "not found"
}

func TestImageExist(t *testing.T) {
	image := "test-image:v1.0.0"

	tc := []struct {
		name  string
		ok    bool
		err   error
		setup func(*mocks.MockAPIClient)
	}{
		{
			name: "image exists",
			ok:   true,
			err:  nil,
			setup: func(dockerClient *mocks.MockAPIClient) {
				dockerClient.EXPECT().ImageInspectWithRaw(context.Background(), image).Return(types.ImageInspect{}, nil, nil)
			},
		},
		{
			name: "image does not exist",
			ok:   false,
			err:  nil,
			setup: func(dockerClient *mocks.MockAPIClient) {
				dockerClient.EXPECT().ImageInspectWithRaw(context.Background(), image).Return(types.ImageInspect{}, nil, notFoundError{})
			},
		},
		{
			name: "image inspect error",
			ok:   false,
			err:  assert.AnError,
			setup: func(dockerClient *mocks.MockAPIClient) {
				dockerClient.EXPECT().ImageInspectWithRaw(context.Background(), image).Return(types.ImageInspect{}, nil, assert.AnError)
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dockerClient := mocks.NewMockAPIClient(ctrl)
			tt.setup(dockerClient)
			defer ctrl.Finish()
			dockerServicemanager := NewDockerServicemanager(dockerClient)

			ok, err := dockerServicemanager.ImageExist(image)

			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.ok, ok)
		})
	}
}

// IsRunning tests

func TestIsRunningError(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
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
	dockerClient := mocks.NewMockAPIClient(ctrl)
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

	dockerServiceManager := services.NewDockerServiceManager(dockerClient)
	ok, err := dockerServiceManager.IsRunning("validator-client")
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestIsRunningTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	dockerClient := mocks.NewMockAPIClient(ctrl)
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

	dockerServiceManager := services.NewDockerServiceManager(dockerClient)
	ok, err := dockerServiceManager.IsRunning("validator-client")
	assert.Nil(t, err)
	assert.True(t, ok)
}
