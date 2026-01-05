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
package aztec_exporter

import (
	"net"
	"path/filepath"
	"strings"
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
)

func TestInit(t *testing.T) {
	afs := afero.NewMemMapFs()

	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join("monitoring", ".lock")}).Return(locker)
	locker.EXPECT().Lock().Return(nil).AnyTimes()
	locker.EXPECT().Locked().Return(true).AnyTimes()
	locker.EXPECT().Unlock().Return(nil).AnyTimes()

	dataDir, err := data.NewDataDir("/", afs, locker)
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
					"AZTEC_EXPORTER_PORT":           "9464",
					"AZTEC_EXPORTER_OTLP_GRPC_PORT": "4317",
					"AZTEC_EXPORTER_OTLP_HTTP_PORT": "4318",
					"AZTEC_EXPORTER_LOG_LEVEL":      "info",
					"AZTEC_EXPORTER_METRIC_EXPIRY":  "5m",
					"AZTEC_EXPORTER_CONF":           "./aztec-exporter/config.yml",
				},
				Stack: stack,
			},
		},
		{
			name: "missing prometheus port",
			options: types.ServiceOptions{
				Dotenv: map[string]string{
					"AZTEC_EXPORTER_OTLP_GRPC_PORT": "4317",
					"AZTEC_EXPORTER_OTLP_HTTP_PORT": "4318",
				},
				Stack: stack,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewAztecExporter(AztecExporterParams{})
			err := svc.Init(tt.options)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, uint16(9464), svc.params.PromPort)
			assert.Equal(t, uint16(4317), svc.params.OtlpGrpcPort)
			assert.Equal(t, uint16(4318), svc.params.OtlpHttpPort)
		})
	}
}

func TestSetupWritesConfig(t *testing.T) {
	afs := afero.NewMemMapFs()

	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)
	locker.EXPECT().New(utils.PathMatcher{Expected: filepath.Join("monitoring", ".lock")}).Return(locker)
	locker.EXPECT().Lock().Return(nil).AnyTimes()
	locker.EXPECT().Locked().Return(true).AnyTimes()
	locker.EXPECT().Unlock().Return(nil).AnyTimes()

	dataDir, err := data.NewDataDir("/", afs, locker)
	require.NoError(t, err)
	stack, err := dataDir.MonitoringStack()
	require.NoError(t, err)

	svc := NewAztecExporter(AztecExporterParams{})
	err = svc.Init(types.ServiceOptions{
		Stack: stack,
		Dotenv: map[string]string{
			"AZTEC_EXPORTER_PORT":           "9464",
			"AZTEC_EXPORTER_OTLP_GRPC_PORT": "4317",
			"AZTEC_EXPORTER_OTLP_HTTP_PORT": "4318",
			"AZTEC_EXPORTER_LOG_LEVEL":      "info",
			"AZTEC_EXPORTER_METRIC_EXPIRY":  "5m",
			"AZTEC_EXPORTER_CONF":           "./aztec-exporter/config.yml",
		},
	})
	require.NoError(t, err)

	err = svc.Setup(svc.DotEnv())
	require.NoError(t, err)

	confPath := filepath.Join(stack.Path(), "aztec-exporter", "config.yml")
	ok, err := afero.Exists(afs, confPath)
	require.NoError(t, err)
	require.True(t, ok)

	raw, err := afero.ReadFile(afs, confPath)
	require.NoError(t, err)
	content := string(raw)
	assert.True(t, strings.Contains(content, "endpoint: 0.0.0.0:9464"))
	assert.True(t, strings.Contains(content, "endpoint: 0.0.0.0:4317"))
	assert.True(t, strings.Contains(content, "endpoint: 0.0.0.0:4318"))
}

func TestSetContainerIP(t *testing.T) {
	svc := NewAztecExporter(AztecExporterParams{})
	ip := net.ParseIP("127.0.0.1")
	svc.SetContainerIP(ip)
	assert.Equal(t, ip, svc.containerIP)
}

func TestContainerName(t *testing.T) {
	svc := NewAztecExporter(AztecExporterParams{})
	assert.Equal(t, monitoring.AztecExporterContainerName, svc.ContainerName())
}

func TestEndpoint(t *testing.T) {
	svc := NewAztecExporter(AztecExporterParams{})
	err := svc.Init(types.ServiceOptions{
		Dotenv: map[string]string{
			"AZTEC_EXPORTER_PORT":           "9464",
			"AZTEC_EXPORTER_OTLP_GRPC_PORT": "4317",
			"AZTEC_EXPORTER_OTLP_HTTP_PORT": "4318",
			"AZTEC_EXPORTER_LOG_LEVEL":      "info",
			"AZTEC_EXPORTER_METRIC_EXPIRY":  "5m",
			"AZTEC_EXPORTER_CONF":           "./aztec-exporter/config.yml",
		},
	})
	require.NoError(t, err)
	svc.SetContainerIP(net.ParseIP("168.77.88.99"))
	assert.Equal(t, "http://168.77.88.99:9464", svc.Endpoint())
}

func TestName(t *testing.T) {
	svc := NewAztecExporter(AztecExporterParams{})
	assert.Equal(t, monitoring.AztecExporterServiceName, svc.Name())
}
