package node_exporter

var dotEnv map[string]string = map[string]string{
	"NODE_EXPORTER_IMAGE": "prom/node-exporter:v1.1.2",
	"NODE_EXPORTER_PORT":  "9100",
}
