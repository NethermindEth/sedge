package types

import (
	"strconv"

	"github.com/NethermindEth/sedge/internal/data"
)

// ServiceOptions defines the options for initializing a monitoring service. It includes a reference to the monitoring stack
// and a map of environment variables.
type ServiceOptions struct {
	// Stack is a reference to the monitoring stack that the service is a part of.
	Stack *data.MonitoringStack

	// Dotenv is a map of environment variables for the service. The keys are the variable names and the values are the variable values.
	Dotenv map[string]string
}

type MonitoringTarget struct {
	// Host is the host of the monitoring target endpoint, e.g. localhost
	Host string
	// Port is the port of the monitoring target endpoint, e.g. 8080
	Port uint16
	// Path is the path of the monitoring target endpoint, e.g. /metrics
	Path string
}

func (t MonitoringTarget) String() string {
	return t.Host + ":" + strconv.Itoa(int(t.Port)) + t.Path
}

func (t MonitoringTarget) Endpoint() string {
	return t.Host + ":" + strconv.Itoa(int(t.Port))
}
