package node_exporter

import (
	"net"
	"strconv"
	"testing"

	"github.com/NethermindEth/sedge/internal/monitoring/data"
	mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/internal/monitoring/monitoring"
	"github.com/NethermindEth/sedge/internal/monitoring/monitoring/services/types"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	// Create an in-memory filesystem
	afs := afero.NewMemMapFs()

	// Create a mock locker
	ctrl := gomock.NewController(t)
	locker := mocks.NewMockLocker(ctrl)

	// Expect the lock to be acquired
	locker.EXPECT().New("/monitoring/.lock").Return(locker)

	// Create a new DataDir with the in-memory filesystem
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
					"NODE_EXPORTER_PORT": "6666",
				},
				Stack: stack,
			},
		},
		{
			name: "missing node exporter port",
			options: types.ServiceOptions{
				Dotenv: map[string]string{},
				Stack:  stack,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeExporter := NewNodeExporter()
			err := nodeExporter.Init(tt.options)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.options.Dotenv["NODE_EXPORTER_PORT"], strconv.Itoa(int(nodeExporter.port)))
			}
		})
	}
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
			// Create a new Node Exporter service
			nodeExporter := NewNodeExporter()
			nodeExporter.SetContainerIP(tt.ip)
			assert.Equal(t, tt.ip, nodeExporter.containerIP)
		})
	}
}

func TestContainerName(t *testing.T) {
	want := monitoring.NodeExporterContainerName

	// Create a new Node Exporter service
	nodeExporter := NewNodeExporter()
	assert.Equal(t, want, nodeExporter.ContainerName())
}

func TestEndpoint(t *testing.T) {
	dotenv := map[string]string{
		"NODE_EXPORTER_PORT": "6666",
	}
	want := "http://168.77.88.99:6666"

	// Create a new Node exporter service
	nodeExporter := NewNodeExporter()
	err := nodeExporter.Init(types.ServiceOptions{
		Dotenv: dotenv,
	})
	require.NoError(t, err)
	nodeExporter.SetContainerIP(net.ParseIP("168.77.88.99"))

	endpoint := nodeExporter.Endpoint()
	assert.Equal(t, want, endpoint)
}
