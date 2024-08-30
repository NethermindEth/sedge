package monitoring

const (
	PrometheusServiceName     = "prometheus"
	PrometheusContainerName   = "sedge_prometheus"
	GrafanaServiceName        = "grafana"
	GrafanaContainerName      = "sedge_grafana"
	NodeExporterServiceName   = "node_exporter"
	NodeExporterContainerName = "sedge_node_exporter"
	monitoringPath            = "monitoring"
	InstanceIDLabel           = "instance_id"
	CommitHashLabel           = "instance_commit_hash"
	AVSNameLabel              = "avs_name"
	AVSVersionLabel           = "avs_version"
	SpecVersionLabel          = "spec_version"
)
