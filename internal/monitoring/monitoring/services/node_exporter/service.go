package node_exporter

import (
	"fmt"
	"net"
	"strconv"

	"github.com/NethermindEth/sedge/internal/pkg/monitoring"
	"github.com/NethermindEth/sedge/internal/pkg/monitoring/services/types"
)

var _ monitoring.ServiceAPI = &NodeExporterService{}

type NodeExporterService struct {
	containerIP net.IP
	port        uint16
}

func NewNodeExporter() *NodeExporterService {
	return &NodeExporterService{}
}

func (n *NodeExporterService) Init(opts types.ServiceOptions) error {
	// Validate dotEnv
	nodeExporterPort, ok := opts.Dotenv["NODE_EXPORTER_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "NODE_EXPORTER_PORT")
	} else if nodeExporterPort == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "NODE_EXPORTER_PORT")
	}

	port, err := strconv.ParseUint(opts.Dotenv["NODE_EXPORTER_PORT"], 10, 16)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid port", ErrInvalidOptions, "NODE_EXPORTER_PORT")
	}
	n.port = uint16(port)
	return nil
}

func (n *NodeExporterService) AddTarget(target types.MonitoringTarget, labels map[string]string, jobName string) error {
	return nil
}

func (n *NodeExporterService) RemoveTarget(instanceID string) (string, error) {
	return "", nil
}

func (n *NodeExporterService) DotEnv() map[string]string {
	return dotEnv
}

func (n *NodeExporterService) Setup(options map[string]string) error {
	return nil
}

func (n *NodeExporterService) SetContainerIP(ip net.IP) {
	n.containerIP = ip
}

func (n *NodeExporterService) ContainerName() string {
	return monitoring.NodeExporterContainerName
}

func (n *NodeExporterService) Endpoint() string {
	return fmt.Sprintf("http://%s:%d", n.containerIP, n.port)
}
