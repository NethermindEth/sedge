package prometheus

var dotEnv map[string]string = map[string]string{
	"PROM_IMAGE": "prom/prometheus:v2.37.0",
	"PROM_PORT":  "9090",
	"PROM_CONF":  "./prometheus/prometheus.yml",
}
