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

var dotEnv map[string]string = map[string]string{
	"AZTEC_EXPORTER_IMAGE":          "otel/opentelemetry-collector-contrib:v0.142.0",
	"AZTEC_EXPORTER_PORT":           "9464",
	"AZTEC_EXPORTER_OTLP_GRPC_PORT": "4317",
	"AZTEC_EXPORTER_OTLP_HTTP_PORT": "4318",
	"AZTEC_EXPORTER_CONF":           "./aztec-exporter/config.yml",
	"AZTEC_EXPORTER_LOG_LEVEL":      "info",
	"AZTEC_EXPORTER_METRIC_EXPIRY":  "5m",
}
