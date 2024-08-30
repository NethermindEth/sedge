package grafana

var dotEnv map[string]string = map[string]string{
	"GRAFANA_IMAGE":          "grafana/grafana-oss:9.4.3",
	"GRAFANA_PORT":           "3000",
	"GRAFANA_ADMIN_USER":     "admin",
	"GRAFANA_ADMIN_PASSWORD": "admin",
	"GRAFANA_PROV":           "./grafana/provisioning",
	"GRAFANA_DATA":           "./grafana/data",
}
