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
package types

import (
	"strconv"

	"github.com/NethermindEth/sedge/internal/monitoring/data"
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

// ServiceTemplateData: Struct Data object to be applied to docker-compose script
type ServiceTemplateData struct {
	LidoExporter bool
}