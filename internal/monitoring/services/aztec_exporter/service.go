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

import (
	"embed"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/NethermindEth/sedge/internal/monitoring"
	datadir "github.com/NethermindEth/sedge/internal/monitoring/data"
	"github.com/NethermindEth/sedge/internal/monitoring/services/types"
)

//go:embed config
var config embed.FS

var _ monitoring.ServiceAPI = &AztecExporterService{}

type AztecExporterParams struct {
	PromPort     uint16
	OtlpGrpcPort uint16
	OtlpHttpPort uint16
	LogLevel     string
	MetricExpiry string
}

type AztecExporterService struct {
	containerIP net.IP
	stack       *datadir.MonitoringStack
	params      AztecExporterParams
}

func NewAztecExporter(params AztecExporterParams) *AztecExporterService {
	if params.PromPort != 0 {
		dotEnv["AZTEC_EXPORTER_PORT"] = strconv.Itoa(int(params.PromPort))
	}
	if params.OtlpGrpcPort != 0 {
		dotEnv["AZTEC_EXPORTER_OTLP_GRPC_PORT"] = strconv.Itoa(int(params.OtlpGrpcPort))
	}
	if params.OtlpHttpPort != 0 {
		dotEnv["AZTEC_EXPORTER_OTLP_HTTP_PORT"] = strconv.Itoa(int(params.OtlpHttpPort))
	}
	if params.LogLevel != "" {
		dotEnv["AZTEC_EXPORTER_LOG_LEVEL"] = params.LogLevel
	}
	if params.MetricExpiry != "" {
		dotEnv["AZTEC_EXPORTER_METRIC_EXPIRY"] = params.MetricExpiry
	}
	return &AztecExporterService{
		params: params,
	}
}

func (a *AztecExporterService) Init(opts types.ServiceOptions) error {
	// Validate dotEnv
	promPortStr, ok := opts.Dotenv["AZTEC_EXPORTER_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "AZTEC_EXPORTER_PORT")
	} else if promPortStr == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "AZTEC_EXPORTER_PORT")
	}
	promPort64, err := strconv.ParseUint(promPortStr, 10, 16)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid port", ErrInvalidOptions, "AZTEC_EXPORTER_PORT")
	}
	a.params.PromPort = uint16(promPort64)

	grpcPortStr, ok := opts.Dotenv["AZTEC_EXPORTER_OTLP_GRPC_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_GRPC_PORT")
	} else if grpcPortStr == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_GRPC_PORT")
	}
	grpcPort64, err := strconv.ParseUint(grpcPortStr, 10, 16)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid port", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_GRPC_PORT")
	}
	a.params.OtlpGrpcPort = uint16(grpcPort64)

	httpPortStr, ok := opts.Dotenv["AZTEC_EXPORTER_OTLP_HTTP_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_HTTP_PORT")
	} else if httpPortStr == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_HTTP_PORT")
	}
	httpPort64, err := strconv.ParseUint(httpPortStr, 10, 16)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid port", ErrInvalidOptions, "AZTEC_EXPORTER_OTLP_HTTP_PORT")
	}
	a.params.OtlpHttpPort = uint16(httpPort64)

	a.params.LogLevel = opts.Dotenv["AZTEC_EXPORTER_LOG_LEVEL"]
	a.params.MetricExpiry = opts.Dotenv["AZTEC_EXPORTER_METRIC_EXPIRY"]
	a.stack = opts.Stack

	return nil
}

func (a *AztecExporterService) AddTarget(target types.MonitoringTarget, labels map[string]string, jobName string) error {
	return nil
}

func (a *AztecExporterService) RemoveTarget(instanceID string) (string, error) {
	return "", nil
}

func (a *AztecExporterService) DotEnv() map[string]string {
	return dotEnv
}

func (a *AztecExporterService) Setup(options map[string]string) error {
	// Write collector config
	rawTmp, err := config.ReadFile("config/config.yml.tmpl")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfigNotFound, err)
	}
	tmp, err := template.New("config.yml").Parse(string(rawTmp))
	if err != nil {
		return err
	}

	// Create config directory
	dir := filepath.Join("aztec-exporter")
	if err := a.stack.CreateDir(dir); err != nil {
		return err
	}

	confPath := options["AZTEC_EXPORTER_CONF"]
	if confPath == "" {
		confPath = dotEnv["AZTEC_EXPORTER_CONF"]
	}
	confFile, err := a.stack.Create(confPath)
	if err != nil {
		return err
	}
	defer confFile.Close()

	data := struct {
		PromPort     string
		OtlpGrpcPort string
		OtlpHttpPort string
		LogLevel     string
		MetricExpiry string
	}{
		PromPort:     options["AZTEC_EXPORTER_PORT"],
		OtlpGrpcPort: options["AZTEC_EXPORTER_OTLP_GRPC_PORT"],
		OtlpHttpPort: options["AZTEC_EXPORTER_OTLP_HTTP_PORT"],
		LogLevel:     options["AZTEC_EXPORTER_LOG_LEVEL"],
		MetricExpiry: options["AZTEC_EXPORTER_METRIC_EXPIRY"],
	}

	if err := tmp.Execute(confFile, data); err != nil {
		return err
	}
	return nil
}

func (a *AztecExporterService) SetContainerIP(ip net.IP) {
	a.containerIP = ip
}

func (a *AztecExporterService) ContainerName() string {
	return monitoring.AztecExporterContainerName
}

func (a *AztecExporterService) Endpoint() string {
	return fmt.Sprintf("http://%s:%d", a.containerIP, a.params.PromPort)
}

func (a *AztecExporterService) Name() string {
	return monitoring.AztecExporterServiceName
}
