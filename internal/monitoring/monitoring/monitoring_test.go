package monitoring

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring/compose"
	"github.com/NethermindEth/sedge/internal/monitoring/data"
	mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/internal/monitoring/monitoring/services/types"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}

	okLocker := func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
		// Create a mock locker
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		gomock.InOrder(
			locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
			locker.EXPECT().Lock().Return(nil),
			locker.EXPECT().Locked().Return(true),
			locker.EXPECT().Unlock().Return(nil),
		)
		return locker
	}

	dotEnvFileWriter := func(t *testing.T, fs afero.Fs, dotenv map[string]string) {
		// Create the dotenv file
		dotenvFile, err := fs.Create(filepath.Join(userDataHome, ".eigen", "monitoring", ".env"))
		require.NoError(t, err)

		// Write the dotenv file
		for key, value := range dotenv {
			_, err := dotenvFile.WriteString(key + "=" + value + "\n")
			require.NoError(t, err)
		}
	}

	tests := []struct {
		name        string
		mocker      func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager)
		setupDotEnv func(t *testing.T, fs afero.Fs, dotenv map[string]string)
		dotenv      map[string]string
		wantErr     bool
	}{
		{
			name: "ok, 1 service",
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
					servicer.EXPECT().SetContainerIP(net.ParseIP("127.0.0.1")).Return(),
				)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("127.0.0.1", nil)

				return []ServiceAPI{
					servicer,
				}, dockerManager
			},
			setupDotEnv: dotEnvFileWriter,
			dotenv: map[string]string{
				"NODE_PORT": "9000",
			},
		},
		{
			name: "ok, 2 services",
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager) {
				service1 := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					service1.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					service1.EXPECT().ContainerName().Return("node1"),
				)

				service2 := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					service2.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					service2.EXPECT().ContainerName().Return("node2"),
				)

				service1.EXPECT().SetContainerIP(net.ParseIP("127.0.0.1")).Return()
				service2.EXPECT().SetContainerIP(net.ParseIP("127.0.0.2")).Return()

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node1").Return("127.0.0.1", nil)
				dockerManager.EXPECT().ContainerIP("node2").Return("127.0.0.2", nil)

				return []ServiceAPI{
					service1, service2,
				}, dockerManager
			},
			setupDotEnv: dotEnvFileWriter,
			dotenv: map[string]string{
				"NODE1_PORT": "9000",
				"NODE2_PORT": "9001",
			},
		},
		{
			name: "error, 1 service, init service error",
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(errors.New("error")),
				)

				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, dockerManager
			},
			setupDotEnv: dotEnvFileWriter,
			dotenv: map[string]string{
				"NODE_PORT": "9000",
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, ContainerIP error",
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
				)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("", errors.New("error"))

				return []ServiceAPI{
					servicer,
				}, dockerManager
			},
			setupDotEnv: dotEnvFileWriter,
			dotenv: map[string]string{
				"NODE_PORT": "9000",
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, ContainerIP gives an invalid IP",
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack, dotenv map[string]string) ([]ServiceAPI, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
				)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("nethermind-loves-eigenlayer", nil)

				return []ServiceAPI{
					servicer,
				}, dockerManager
			},
			setupDotEnv: dotEnvFileWriter,
			dotenv: map[string]string{
				"NODE_PORT": "9000",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			// Create a mock controller
			ctrl := gomock.NewController(t)
			locker := okLocker(t, ctrl)

			afs := afero.NewMemMapFs()
			tt.setupDotEnv(t, afs, tt.dotenv)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{},
				mocks.NewMockComposeManager(ctrl),
				mocks.NewMockDockerManager(ctrl),
				afs,
				locker,
			)

			services, dockerManager := tt.mocker(t, ctrl, manager.stack, tt.dotenv)
			manager.services = services
			manager.dockerManager = dockerManager

			// Init the stack
			err := manager.Init()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInstallStack(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}

	okLocker := func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
		// Create a mock locker
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		gomock.InOrder(
			locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
			locker.EXPECT().Lock().Return(nil),
			locker.EXPECT().Locked().Return(true),
			locker.EXPECT().Unlock().Return(nil),
		)
		// stack.Installed() lock
		gomock.InOrder(
			locker.EXPECT().Lock().Return(nil),
			locker.EXPECT().Locked().Return(true),
			locker.EXPECT().Unlock().Return(nil),
		)
		return locker
	}
	onlyNewLocker := func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
		// Create a mock locker
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)
		return locker
	}

	// Setup mock http server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Listen for the POST request to /-/reload
		if r.URL.Path == "/-/reload" && r.Method == http.MethodPost {
			// All good
			w.WriteHeader(http.StatusOK)
		} else if r.Method != http.MethodGet {
			// Unexpected method
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			// Unexpected path
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	serverPort := strings.Split(server.URL, ":")[2]

	tests := []struct {
		name         string
		mockerLocker func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker
		mocker       func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager)
		wantErr      bool
	}{
		{
			name:         "ok, 1 service, port not occupied",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
					servicer.EXPECT().SetContainerIP(net.ParseIP("127.0.0.1")).Return(),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("127.0.0.1", nil)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
		},
		{
			name:         "ok, 2 services",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE1_PORT": "9000",
					"NODE2_PORT": "9001",
				}
				service1, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)

				// Expect the service to be triggered
				gomock.InOrder(
					service1.EXPECT().DotEnv().Return(map[string]string{
						"NODE1_PORT": "9000",
					}),
					service1.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					service1.EXPECT().Setup(dotenv).Return(nil),
					service1.EXPECT().ContainerName().Return("node1"),
				)
				gomock.InOrder(
					service2.EXPECT().DotEnv().Return(map[string]string{
						"NODE2_PORT": "9000",
					}),
					service2.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					service2.EXPECT().Setup(dotenv).Return(nil),
					service2.EXPECT().ContainerName().Return("node2"),
				)
				service1.EXPECT().SetContainerIP(net.ParseIP("168.0.2.1")).Return()
				service2.EXPECT().SetContainerIP(net.ParseIP("168.0.3.1")).Return()

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node1").Return("168.0.2.1", nil)
				dockerManager.EXPECT().ContainerIP("node2").Return("168.0.3.1", nil)

				return []ServiceAPI{
					service1,
					service2,
				}, composeManager, dockerManager
			},
		},
		{
			name:         "ok, 1 service, port occupied",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				// Convert serverPort to int
				p, err := strconv.Atoi(serverPort)
				require.NoError(t, err)
				sp := strconv.Itoa(p + 1)

				dotenv := map[string]string{
					"NODE_PORT": sp,
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(map[string]string{"NODE_PORT": serverPort}),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
					servicer.EXPECT().SetContainerIP(net.ParseIP("127.1.1.6")).Return(),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("127.1.1.6", nil)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
		},
		{
			name:         "error, 1 service, port not int",
			mockerLocker: onlyNewLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "3RR0R",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				servicer.EXPECT().DotEnv().Return(dotenv)

				composeManager := mocks.NewMockComposeManager(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name:         "error, 1 service, invalid port",
			mockerLocker: onlyNewLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "0",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				servicer.EXPECT().DotEnv().Return(dotenv)

				composeManager := mocks.NewMockComposeManager(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name:         "error, 1 service, init service error",
			mockerLocker: onlyNewLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(errors.New("error")),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, stack setup error",
			mockerLocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(errors.New("error")),
				)
				return locker
			},
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, service setup error",
			mockerLocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				return locker
			},
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(errors.New("error")),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, create error",
			mockerLocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				return locker
			},
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(nil),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(errors.New("error"))

				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, run error",
			mockerLocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				return locker
			},
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(nil),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(errors.New("error"))

				dockerManager := mocks.NewMockDockerManager(ctrl)

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
		{
			name: "error, 1 service, ContainerIP error",
			mockerLocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				return locker
			},
			mocker: func(t *testing.T, ctrl *gomock.Controller, stack *data.MonitoringStack) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				dotenv := map[string]string{
					"NODE_PORT": "9000",
				}
				servicer := mocks.NewMockServiceAPI(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().DotEnv().Return(dotenv),
					servicer.EXPECT().Init(types.ServiceOptions{
						Stack:  stack,
						Dotenv: dotenv,
					}).Return(nil),
					servicer.EXPECT().Setup(dotenv).Return(nil),
					servicer.EXPECT().ContainerName().Return("node"),
				)

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Create(compose.DockerComposeCreateOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: filepath.Join(stack.Path(), "docker-compose.yml")}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("", errors.New("error"))

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			// Create a mock controller
			ctrl := gomock.NewController(t)
			locker := tt.mockerLocker(t, ctrl)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{},
				mocks.NewMockComposeManager(ctrl),
				mocks.NewMockDockerManager(ctrl),
				afero.NewMemMapFs(),
				locker,
			)

			services, composeManager, dockerManager := tt.mocker(t, ctrl, manager.stack)
			manager.services = services
			manager.composeManager = composeManager
			manager.dockerManager = dockerManager

			// Init the stack
			err := manager.InstallStack()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Check the stack is installed
				installed, err := manager.stack.Installed()
				assert.NoError(t, err)
				assert.True(t, installed)
			}
		})
	}
}

func TestAddAndRemoveTarget(t *testing.T) {
	okLocker := func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker {
		// Create a mock locker
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		userDataHome := os.Getenv("XDG_DATA_HOME")
		if userDataHome == "" {
			userHome, err := os.UserHomeDir()
			require.NoError(t, err)
			userDataHome = filepath.Join(userHome, ".local", "share")
		}
		locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)

		return locker
	}

	tests := []struct {
		name          string
		mockerLocker  func(t *testing.T, ctrl *gomock.Controller) *mocks.MockLocker
		mocker        func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager)
		target        types.MonitoringTarget
		labels        map[string]string
		dockerNetwork string
		add           bool
		wantErr       bool
	}{
		{
			name:         "add, ok, prometheus service",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				prometheusService := mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				gomock.InOrder(
					prometheusService.EXPECT().ContainerName().Return(PrometheusContainerName),
					dockerManager.EXPECT().ContainerNetworks(PrometheusContainerName).Return([]string{"eigen_default"}, nil),
					dockerManager.EXPECT().NetworkConnect(PrometheusContainerName, dockerNetwork).Return(nil),
					prometheusService.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--%s++%s", labels[InstanceIDLabel], PrometheusContainerName, dockerNetwork)).Return(nil),
				)

				return []ServiceAPI{
					prometheusService,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
			},
			labels: map[string]string{
				InstanceIDLabel: "becks",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			add: true,
		},
		{
			name:         "add, ok, 1 service",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				service := mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				gomock.InOrder(
					service.EXPECT().ContainerName().Return("service1"),
					service.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--service1++%s", labels[InstanceIDLabel], dockerNetwork)).Return(nil),
				)

				return []ServiceAPI{
					service,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
				Path: "/custom/path",
			},
			labels: map[string]string{
				InstanceIDLabel: "heineken",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			add: true,
		},
		{
			name:         "add, ok, 1 services, prometheus was already added to network",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				prometheusService := mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				gomock.InOrder(
					prometheusService.EXPECT().ContainerName().Return(PrometheusContainerName),
					dockerManager.EXPECT().ContainerNetworks(PrometheusContainerName).Return([]string{"eigen_default", dockerNetwork}, nil),
					prometheusService.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--%s++%s", labels[InstanceIDLabel], PrometheusContainerName, dockerNetwork)).Return(nil),
				)

				return []ServiceAPI{
					prometheusService,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
			},
			labels: map[string]string{
				InstanceIDLabel: "grolsch",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			add: true,
		},
		{
			name:         "add, ok, 2 services, 1 AddTarget error",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				prometheusService, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)

				gomock.InOrder(
					prometheusService.EXPECT().ContainerName().Return(PrometheusContainerName),
					dockerManager.EXPECT().ContainerNetworks(PrometheusContainerName).Return([]string{"eigen_default"}, nil),
					dockerManager.EXPECT().NetworkConnect(PrometheusContainerName, dockerNetwork).Return(nil),
					prometheusService.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--%s++%s", labels[InstanceIDLabel], PrometheusContainerName, dockerNetwork)).Return(nil),
					service2.EXPECT().ContainerName().Return("node2"),
					service2.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--node2++%s", labels[InstanceIDLabel], dockerNetwork)).Return(errors.New("error")),
				)

				return []ServiceAPI{
					prometheusService,
					service2,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
			},
			labels: map[string]string{
				InstanceIDLabel: "carlsberg",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			wantErr: true,
			add:     true,
		},
		{
			name:         "add, ok, 2 services, prometheus ContainerNetworks error",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				prometheusService, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				gomock.InOrder(
					service2.EXPECT().ContainerName().Return("node2"),
					service2.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--node2++%s", labels[InstanceIDLabel], dockerNetwork)).Return(nil),
					prometheusService.EXPECT().ContainerName().Return(PrometheusContainerName),
					dockerManager.EXPECT().ContainerNetworks(PrometheusContainerName).Return(nil, errors.New("error")),
				)

				return []ServiceAPI{
					service2,
					prometheusService,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
			},
			labels: map[string]string{
				InstanceIDLabel: "amstel",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			wantErr: true,
			add:     true,
		},
		{
			name:         "add, ok, 2 services, prometheus NetworkConnect error",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				service1, prometheusService := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				gomock.InOrder(
					service1.EXPECT().ContainerName().Return("node1"),
					service1.EXPECT().AddTarget(target, labels, fmt.Sprintf("%s--node1++%s", labels[InstanceIDLabel], dockerNetwork)).Return(nil),
					prometheusService.EXPECT().ContainerName().Return(PrometheusContainerName),
					dockerManager.EXPECT().ContainerNetworks(PrometheusContainerName).Return([]string{"eigen_default"}, nil),
					dockerManager.EXPECT().NetworkConnect(PrometheusContainerName, dockerNetwork).Return(errors.New("error")),
				)

				return []ServiceAPI{
					service1,
					prometheusService,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
				Port: 9000,
			},
			labels: map[string]string{
				InstanceIDLabel: "budweiser",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			wantErr: true,
			add:     true,
		},
		{
			name:         "remove, ok, 1 service",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					servicer.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					servicer.EXPECT().ContainerName().Return("node"),
					dockerManager.EXPECT().NetworkDisconnect("node", dockerNetwork).Return(nil),
				)

				return []ServiceAPI{
					servicer,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
			},
			labels: map[string]string{
				InstanceIDLabel: "stella",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
		},
		{
			name:         "remove, ok, 2 services, one of them was already removed from network",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				service1, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					service1.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					service1.EXPECT().ContainerName().Return("node1"),
					dockerManager.EXPECT().NetworkDisconnect("node1", dockerNetwork).Return(nil),
					service2.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					service2.EXPECT().ContainerName().Return("node2"),
					dockerManager.EXPECT().NetworkDisconnect("node2", dockerNetwork).Return(assert.AnError),
				)

				return []ServiceAPI{
					service1,
					service2,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "localhost",
			},
			labels: map[string]string{
				InstanceIDLabel: "corona",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
		},
		{
			name:         "remove, ok, 2 services, 1 RemoveTarget error",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				service1, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					service1.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					service1.EXPECT().ContainerName().Return("node1"),
					dockerManager.EXPECT().NetworkDisconnect("node1", dockerNetwork).Return(nil),
					service2.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return("", assert.AnError),
				)

				return []ServiceAPI{
					service1,
					service2,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "remove",
			},
			labels: map[string]string{
				InstanceIDLabel: "perla",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
			wantErr: true,
		},
		{
			name:         "remove, ok, 2 services, 1 NetworkDisconnect error",
			mockerLocker: okLocker,
			mocker: func(t *testing.T, ctrl *gomock.Controller, labels map[string]string, dockerNetwork string, target types.MonitoringTarget) ([]ServiceAPI, *mocks.MockDockerManager) {
				service1, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the service to be triggered
				gomock.InOrder(
					service1.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					service1.EXPECT().ContainerName().Return("node1"),
					dockerManager.EXPECT().NetworkDisconnect("node1", dockerNetwork).Return(nil),
					service2.EXPECT().RemoveTarget(labels[InstanceIDLabel]).Return(dockerNetwork, nil),
					service2.EXPECT().ContainerName().Return("node2"),
					dockerManager.EXPECT().NetworkDisconnect("node2", dockerNetwork).Return(assert.AnError),
				)

				return []ServiceAPI{
					service1,
					service2,
				}, dockerManager
			},
			target: types.MonitoringTarget{
				Host: "remove",
			},
			labels: map[string]string{
				InstanceIDLabel: "yeungling",
				CommitHashLabel: "76973ce6755edb6cce37efd62266e98c838f6968",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock controller
			ctrl := gomock.NewController(t)
			locker := tt.mockerLocker(t, ctrl)

			services, dockerManager := tt.mocker(t, ctrl, tt.labels, tt.dockerNetwork, tt.target)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				services,
				mocks.NewMockComposeManager(ctrl),
				dockerManager,
				afero.NewMemMapFs(),
				locker,
			)

			var err error
			if tt.add {
				// Add the target
				err = manager.AddTarget(tt.target, tt.labels, tt.dockerNetwork)
			} else {
				// Remove the target
				err = manager.RemoveTarget(tt.labels[InstanceIDLabel])
			}
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRun(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}
	composePath := filepath.Join(userDataHome, ".eigen", "monitoring", "docker-compose.yml")

	tests := []struct {
		name    string
		mocker  func(t *testing.T, ctrl *gomock.Controller) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager)
		wantErr bool
	}{
		{
			name: "ok",
			mocker: func(t *testing.T, ctrl *gomock.Controller) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				service1, service2 := mocks.NewMockServiceAPI(ctrl), mocks.NewMockServiceAPI(ctrl)
				service1.EXPECT().ContainerName().Return("node1")
				service2.EXPECT().ContainerName().Return("node2")

				service1.EXPECT().SetContainerIP(net.ParseIP("168.0.2.1")).Return()
				service2.EXPECT().SetContainerIP(net.ParseIP("168.0.3.1")).Return()

				composeManager := mocks.NewMockComposeManager(ctrl)
				// Expect the compose manager to be triggered
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: composePath}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node1").Return("168.0.2.1", nil)
				dockerManager.EXPECT().ContainerIP("node2").Return("168.0.3.1", nil)

				return []ServiceAPI{
					service1, service2,
				}, composeManager, dockerManager
			},
		},
		{
			name: "up error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				composeManager := mocks.NewMockComposeManager(ctrl)
				// Expect the compose manager to be triggered
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: composePath}).Return(errors.New("error"))
				return []ServiceAPI{mocks.NewMockServiceAPI(ctrl)}, composeManager, mocks.NewMockDockerManager(ctrl)
			},
			wantErr: true,
		},
		{
			name: "error, ContainerIP error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) ([]ServiceAPI, *mocks.MockComposeManager, *mocks.MockDockerManager) {
				servicer := mocks.NewMockServiceAPI(ctrl)
				servicer.EXPECT().ContainerName().Return("node")

				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Up(compose.DockerComposeUpOptions{Path: composePath}).Return(nil)

				dockerManager := mocks.NewMockDockerManager(ctrl)
				dockerManager.EXPECT().ContainerIP("node").Return("", errors.New("error"))

				return []ServiceAPI{
					servicer,
				}, composeManager, dockerManager
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an in-memory filesystem
			afs := afero.NewMemMapFs()

			// Create a mock controller
			ctrl := gomock.NewController(t)

			// Create a mock locker
			locker := mocks.NewMockLocker(ctrl)
			// Expect the lock to be acquired
			locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)

			services, composeManager, dockerManager := tt.mocker(t, ctrl)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				services,
				composeManager,
				dockerManager,
				afs,
				locker,
			)

			// Run the stack
			err := manager.Run()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStop(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}
	composePath := filepath.Join(userDataHome, ".eigen", "monitoring", "docker-compose.yml")

	tests := []struct {
		name    string
		mocker  func(t *testing.T, ctrl *gomock.Controller) *mocks.MockComposeManager
		wantErr bool
	}{
		{
			name: "ok",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockComposeManager {
				composeManager := mocks.NewMockComposeManager(ctrl)
				// Expect the compose manager to be triggered
				gomock.InOrder(
					composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath}).Return(nil),
				)
				return composeManager
			},
		},
		{
			name: "down error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockComposeManager {
				composeManager := mocks.NewMockComposeManager(ctrl)
				// Expect the compose manager to be triggered
				composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath}).Return(errors.New("error"))
				return composeManager
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock controller
			ctrl := gomock.NewController(t)

			// Create a mock locker
			locker := mocks.NewMockLocker(ctrl)
			// Expect the lock to be acquired
			locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{mocks.NewMockServiceAPI(ctrl)},
				tt.mocker(t, ctrl),
				mocks.NewMockDockerManager(ctrl),
				afero.NewMemMapFs(),
				locker,
			)

			// Stop the stack
			err := manager.Stop()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStatus(t *testing.T) {
	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}

	tests := []struct {
		name    string
		mocker  func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager
		want    common.Status
		wantErr bool
	}{
		{
			name: "ok",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				gomock.InOrder(
					dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Running, nil),
					dockerManager.EXPECT().ContainerStatus(PrometheusContainerName).Return(common.Running, nil),
					dockerManager.EXPECT().ContainerStatus(NodeExporterContainerName).Return(common.Running, nil),
				)
				return dockerManager
			},
			want: common.Running,
		},
		{
			name: "error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Unknown, errors.New("error"))
				return dockerManager
			},
			want:    common.Unknown,
			wantErr: true,
		},
		{
			name: "Restarting",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				gomock.InOrder(
					dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Restarting, nil),
					dockerManager.EXPECT().ContainerStatus(PrometheusContainerName).Return(common.Restarting, nil),
					dockerManager.EXPECT().ContainerStatus(NodeExporterContainerName).Return(common.Restarting, nil),
				)
				return dockerManager
			},
			want: common.Restarting,
		},
		{
			name: "Paused",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Paused, nil)
				return dockerManager
			},
			want:    common.Broken,
			wantErr: true,
		},
		{
			name: "Exited",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Exited, nil)
				return dockerManager
			},
			want:    common.Broken,
			wantErr: true,
		},
		{
			name: "Dead",
			mocker: func(t *testing.T, ctrl *gomock.Controller) *mocks.MockDockerManager {
				dockerManager := mocks.NewMockDockerManager(ctrl)
				// Expect the docker manager to be triggered
				dockerManager.EXPECT().ContainerStatus(GrafanaContainerName).Return(common.Dead, nil)
				return dockerManager
			},
			want:    common.Broken,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock controller
			ctrl := gomock.NewController(t)

			// Create a mock locker
			locker := mocks.NewMockLocker(ctrl)
			// Expect the lock to be acquired
			locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{mocks.NewMockServiceAPI(ctrl)},
				mocks.NewMockComposeManager(ctrl),
				tt.mocker(t, ctrl),
				afero.NewMemMapFs(),
				locker,
			)

			status, err := manager.Status()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, status)
		})
	}
}

func TestInstallationStatus(t *testing.T) {
	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}

	tests := []struct {
		name    string
		mocker  func(t *testing.T, ctrl *gomock.Controller) (afero.Fs, *mocks.MockLocker)
		want    common.Status
		wantErr bool
	}{
		{
			name: "installed",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (afero.Fs, *mocks.MockLocker) {
				fs := afero.NewMemMapFs()
				// Recreate installed monitoring
				fs.MkdirAll(filepath.Join(userDataHome, ".eigen", "monitoring"), 0o755)
				fs.Create(filepath.Join(userDataHome, ".eigen", "monitoring", "docker-compose.yml"))
				fs.Create(filepath.Join(userDataHome, ".eigen", "monitoring", ".env"))

				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)
				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)

				return fs, locker
			},
			want: common.Installed,
		},
		{
			name: "not installed",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (afero.Fs, *mocks.MockLocker) {
				fs := afero.NewMemMapFs()

				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)
				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)

				return fs, locker
			},
			want: common.NotInstalled,
		},
		{
			name: "lock error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (afero.Fs, *mocks.MockLocker) {
				fs := afero.NewMemMapFs()

				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)
				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(errors.New("error")),
				)

				return fs, locker
			},
			want:    common.Unknown,
			wantErr: true,
		},
		{
			name: "unlock error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (afero.Fs, *mocks.MockLocker) {
				fs := afero.NewMemMapFs()

				// Create a mock locker
				locker := mocks.NewMockLocker(ctrl)
				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(errors.New("error")),
				)

				return fs, locker
			},
			want:    common.Unknown,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock controller
			ctrl := gomock.NewController(t)

			fs, locker := tt.mocker(t, ctrl)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{mocks.NewMockServiceAPI(ctrl)},
				mocks.NewMockComposeManager(ctrl),
				mocks.NewMockDockerManager(ctrl),
				fs,
				locker,
			)

			status, err := manager.InstallationStatus()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, status)
		})
	}
}

func TestCleanup(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		require.NoError(t, err)
		userDataHome = filepath.Join(userHome, ".local", "share")
	}
	composePath := filepath.Join(userDataHome, ".eigen", "monitoring", "docker-compose.yml")

	tests := []struct {
		name      string
		mocker    func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker)
		force     bool
		noInstall bool
		wantErr   bool
	}{
		{
			name: "ok, force false",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker) {
				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath, Volumes: true}).Return(nil)

				locker := mocks.NewMockLocker(ctrl)
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				locker.EXPECT().Lock().Return(nil)

				return composeManager, locker
			},
		},
		{
			name: "ok, force true",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker) {
				composeManager := mocks.NewMockComposeManager(ctrl)

				locker := mocks.NewMockLocker(ctrl)
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)

				return composeManager, locker
			},
			force: true,
		},
		{
			name: "down error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker) {
				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath, Volumes: true}).Return(errors.New("error"))

				locker := mocks.NewMockLocker(ctrl)
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)

				return composeManager, locker
			},
			wantErr: true,
		},
		{
			name: "ok, force false, no install",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker) {
				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath, Volumes: true}).Return(nil)

				locker := mocks.NewMockLocker(ctrl)
				locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker)
				locker.EXPECT().Lock().Return(nil)

				return composeManager, locker
			},
			noInstall: true,
		},
		{
			name: "stack cleanup error, lock error",
			mocker: func(t *testing.T, ctrl *gomock.Controller) (*mocks.MockComposeManager, *mocks.MockLocker) {
				composeManager := mocks.NewMockComposeManager(ctrl)
				composeManager.EXPECT().Down(compose.DockerComposeDownOptions{Path: composePath, Volumes: true}).Return(nil)

				locker := mocks.NewMockLocker(ctrl)
				gomock.InOrder(
					locker.EXPECT().New(filepath.Join(userDataHome, ".eigen", "monitoring", ".lock")).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(true),
					locker.EXPECT().Unlock().Return(nil),
				)
				locker.EXPECT().Lock().Return(errors.New("lock error"))

				return composeManager, locker
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an in-memory filesystem
			afs := afero.NewMemMapFs()

			// Create a mock controller
			ctrl := gomock.NewController(t)

			composeMgr, locker := tt.mocker(t, ctrl)

			// Create a monitoring manager
			manager := NewMonitoringManager(
				[]ServiceAPI{mocks.NewMockServiceAPI(ctrl)},
				composeMgr,
				mocks.NewMockDockerManager(ctrl),
				afs,
				locker,
			)

			if !tt.noInstall {
				err := manager.stack.Setup(map[string]string{"NODE_NAME": "test"}, script)
				require.NoError(t, err)
			}

			err := manager.Cleanup(tt.force)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Check that monitoring stack has been removed
				exists, err := afero.DirExists(afs, filepath.Join(userDataHome, ".eigen", "monitoring"))
				assert.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}
}

func TestServiceEndpoints(t *testing.T) {
	want := map[string]string{
		GrafanaContainerName:      "http://grafana:3005",
		PrometheusContainerName:   "http://prometheus:9095",
		NodeExporterContainerName: "http://node-exporter:9105",
	}

	// Setup mocks
	ctrl := gomock.NewController(t)
	grafanaMock := mocks.NewMockServiceAPI(ctrl)
	promMock := mocks.NewMockServiceAPI(ctrl)
	nodeExporterMock := mocks.NewMockServiceAPI(ctrl)

	// Expect the service to be triggered
	gomock.InOrder(
		grafanaMock.EXPECT().ContainerName().Return(GrafanaContainerName),
		grafanaMock.EXPECT().Endpoint().Return("http://grafana:3005"),
		promMock.EXPECT().ContainerName().Return(PrometheusContainerName),
		promMock.EXPECT().Endpoint().Return("http://prometheus:9095"),
		nodeExporterMock.EXPECT().ContainerName().Return(NodeExporterContainerName),
		nodeExporterMock.EXPECT().Endpoint().Return("http://node-exporter:9105"),
	)

	// Init monitoring manager and services
	manager := MonitoringManager{
		services: []ServiceAPI{grafanaMock, promMock, nodeExporterMock},
	}

	// Check endpoints
	endpoints := manager.ServiceEndpoints()
	assert.Equal(t, want, endpoints)
}
