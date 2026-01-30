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
package monitoring

const (
	PrometheusServiceName      = "prometheus"
	PrometheusContainerName    = "sedge_prometheus"
	GrafanaServiceName         = "grafana"
	GrafanaContainerName       = "sedge_grafana"
	NodeExporterServiceName    = "node_exporter"
	NodeExporterContainerName  = "sedge_node_exporter"
	LidoExporterServiceName    = "lido_exporter"
	LidoExporterContainerName  = "sedge_lido_exporter"
	AztecExporterServiceName   = "aztec_exporter"
	AztecExporterContainerName = "sedge_aztec_exporter"
	monitoringPath             = "monitoring"
	InstanceIDLabel            = "instance_id"
	SedgeNetworkName           = "sedge-network"
)
