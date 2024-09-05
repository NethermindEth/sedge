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
package grafana

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/NethermindEth/sedge/internal/monitoring"
	"github.com/NethermindEth/sedge/internal/monitoring/data"
	mocks "github.com/NethermindEth/sedge/internal/monitoring/locker/mocks"
	"github.com/NethermindEth/sedge/internal/monitoring/services/types"
	"github.com/NethermindEth/sedge/internal/monitoring/utils"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type Config struct {
	APIVersion  int          `yaml:"apiVersion"`
	Datasources []Datasource `yaml:"datasources"`
}

type Datasource struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Access   string   `yaml:"access"`
	URL      string   `yaml:"url"`
	UID      string   `yaml:"uid"`
	JsonData JsonData `yaml:"jsonData"`
}

type JsonData struct {
	HTTPMethod                    string                       `yaml:"httpMethod"`
	ManageAlerts                  bool                         `yaml:"manageAlerts"`
	PrometheusType                string                       `yaml:"prometheusType"`
	PrometheusVersion             string                       `yaml:"prometheusVersion"`
	IncrementalQuerying           bool                         `yaml:"incrementalQuerying"`
	IncrementalQueryOverlapWindow string                       `yaml:"incrementalQueryOverlapWindow"`
	CacheLevel                    string                       `yaml:"cacheLevel"`
	ExemplarTraceIdDestinations   []ExemplarTraceIdDestination `yaml:"exemplarTraceIdDestinations"`
}

type ExemplarTraceIdDestination struct {
	DatasourceUid string `yaml:"datasourceUid"`
	Name          string `yaml:"name"`
}

func TestInit(t *testing.T) {
	// Create an in-memory filesystem
	afs := afero.NewMemMapFs()
	userHome, err := os.UserHomeDir()
	require.NoError(t, err)
	basePath := filepath.Join(userHome, ".sedge")

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)

	// Expect the lock to be acquired
	locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "/monitoring", ".lock")}).Return(locker)

	// Create a new DataDir with the in-memory filesystem
	dataDir, err := data.NewDataDir(basePath, afs, locker)
	require.NoError(t, err)
	stack, err := dataDir.MonitoringStack()
	require.NoError(t, err)

	tests := []struct {
		name    string
		options types.ServiceOptions
		wantErr bool
	}{
		{
			name: "ok",
			options: types.ServiceOptions{
				Dotenv: map[string]string{
					"GRAFANA_PORT": "3000",
				},
				Stack: stack,
			},
		},
		{
			name: "missing grafana port",
			options: types.ServiceOptions{
				Dotenv: map[string]string{},
				Stack:  stack,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grafana := NewGrafana()
			err := grafana.Init(tt.options)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, stack, grafana.stack)
				assert.Equal(t, tt.options.Dotenv["GRAFANA_PORT"], strconv.Itoa(int(grafana.port)))
			}
		})
	}
}

func TestSetup(t *testing.T) {
	userHome, err := os.UserHomeDir()
	require.NoError(t, err)
	basePath := filepath.Join(userHome, ".sedge")

	okLocker := func(t *testing.T) *mocks.MockLocker {
		// Create a mock locker
		ctrl := gomock.NewController(t)
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		gomock.InOrder(
			locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "/monitoring", ".lock")}).Return(locker),
			locker.EXPECT().Lock().Return(nil),
			locker.EXPECT().Locked().Return(true),
			locker.EXPECT().Unlock().Return(nil),
		)
		for i := 0; i < 9; i++ {
			gomock.InOrder(
				locker.EXPECT().Lock().Return(nil),
				locker.EXPECT().Locked().Return(true),
				locker.EXPECT().Unlock().Return(nil),
			)
		}
		return locker
	}
	onlyNewLocker := func(t *testing.T) *mocks.MockLocker {
		// Create a mock locker
		ctrl := gomock.NewController(t)
		locker := mocks.NewMockLocker(ctrl)

		// Expect the lock to be acquired
		locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "/monitoring", ".lock")}).Return(locker)
		return locker
	}

	tests := []struct {
		name    string
		mocker  func(t *testing.T) *mocks.MockLocker
		options map[string]string
		wantErr bool
	}{
		{
			name:   "ok",
			mocker: okLocker,
			options: map[string]string{
				"PROM_PORT":    "9090",
				"GRAFANA_PORT": "3000",
			},
		},
		{
			name:   "missing prometheus port",
			mocker: onlyNewLocker,
			options: map[string]string{
				"GRAFANA_PORT": "3000",
			},
			wantErr: true,
		},
		{
			name:   "empty prometheus port",
			mocker: onlyNewLocker,
			options: map[string]string{
				"PROM_PORT":    "",
				"GRAFANA_PORT": "3000",
			},
			wantErr: true,
		},
		{
			name: "lock error",
			mocker: func(t *testing.T) *mocks.MockLocker {
				// Create a mock locker
				ctrl := gomock.NewController(t)
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "/monitoring", ".lock")}).Return(locker),
					locker.EXPECT().Lock().Return(fmt.Errorf("error")),
				)
				return locker
			},
			options: map[string]string{
				"PROM_PORT":    "9090",
				"GRAFANA_PORT": "3000",
			},
			wantErr: true,
		},
		{
			name: "unlock error",
			mocker: func(t *testing.T) *mocks.MockLocker {
				// Create a mock locker
				ctrl := gomock.NewController(t)
				locker := mocks.NewMockLocker(ctrl)

				// Expect the lock to be acquired
				gomock.InOrder(
					locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join(basePath, "/monitoring", ".lock")}).Return(locker),
					locker.EXPECT().Lock().Return(nil),
					locker.EXPECT().Locked().Return(false),
				)
				return locker
			},
			options: map[string]string{
				"PROM_PORT":    "9090",
				"GRAFANA_PORT": "3000",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an in-memory filesystem
			afs := afero.NewMemMapFs()

			// Create a new DataDir with the in-memory filesystem
			dataDir, err := data.NewDataDir(basePath, afs, tt.mocker(t))
			require.NoError(t, err)
			stack, err := dataDir.MonitoringStack()
			require.NoError(t, err)

			// Create a new Grafana service
			grafana := NewGrafana()
			grafana.Init(types.ServiceOptions{
				Stack:  stack,
				Dotenv: tt.options,
			})

			// Setup the Grafana service
			err = grafana.Setup(tt.options)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				assert.NoError(t, err)

				// Check the Grafana config file
				promYmlPath := filepath.Join(basePath, "monitoring", "grafana", "provisioning", "datasources", "prom.yml")
				ok, err := afero.Exists(afs, promYmlPath)
				assert.True(t, ok)
				assert.NoError(t, err)

				// Read the prom.yml file
				var prom Config
				promYml, err := afero.ReadFile(afs, promYmlPath)
				assert.NoError(t, err)
				err = yaml.Unmarshal(promYml, &prom)
				assert.NoError(t, err)

				// Check the Prometheus port
				promEndpoint := fmt.Sprintf("http://%s:%s", monitoring.PrometheusServiceName, tt.options["PROM_PORT"])
				assert.Equal(t, promEndpoint, prom.Datasources[0].URL)

				// Check the Dashboards config file
				dashboardsYmlPath := filepath.Join(basePath, "monitoring", "grafana", "provisioning", "dashboards", "dashboards.yml")
				ok, err = afero.Exists(afs, dashboardsYmlPath)
				assert.True(t, ok)
				assert.NoError(t, err)

				// Check the provisioned dashboards
				foldersToCheck := []string{
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards"),
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards", "common-metrics"),
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards", "node-exporter"),
				}
				filesToCheck := []string{
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards", "common-metrics", "common-metrics.json"),
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards", "common-metrics", "common-metrics-global.json"),
					filepath.Join(basePath, "monitoring", "grafana", "data", "dashboards", "node-exporter", "node-exporter.json"),
				}
				for _, folder := range foldersToCheck {
					ok, err = afero.DirExists(afs, folder)
					assert.True(t, ok)
					assert.NoError(t, err)
				}
				for _, file := range filesToCheck {
					ok, err = afero.Exists(afs, file)
					assert.True(t, ok)
					assert.NoError(t, err)
				}
			}
		})
	}
}

func TestDotEnv(t *testing.T) {
	// Create a new Grafana service
	grafana := NewGrafana()
	// Verify the dotEnv
	assert.EqualValues(t, dotEnv, grafana.DotEnv())
}

func TestSetContainerIP(t *testing.T) {
	tests := []struct {
		name string
		ip   net.IP
	}{
		{
			name: "ok",
			ip:   net.ParseIP("127.0.0.1"),
		},
		{
			name: "empty",
			ip:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Grafana service
			grafana := NewGrafana()
			grafana.SetContainerIP(tt.ip)
			assert.Equal(t, tt.ip, grafana.containerIP)
		})
	}
}

func TestContainerName(t *testing.T) {
	want := monitoring.GrafanaContainerName

	// Create a new Grafana service
	grafana := NewGrafana()
	assert.Equal(t, want, grafana.ContainerName())
}

func TestEndpoint(t *testing.T) {
	dotenv := map[string]string{
		"GRAFANA_PORT": "3333",
	}
	want := "http://168.66.77.88:3333"

	// Create a new Grafana service
	grafana := NewGrafana()
	err := grafana.Init(types.ServiceOptions{
		Dotenv: dotenv,
	})
	require.NoError(t, err)
	grafana.SetContainerIP(net.ParseIP("168.66.77.88"))

	endpoint := grafana.Endpoint()
	assert.Equal(t, want, endpoint)
}
