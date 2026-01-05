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
package grafana

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
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

//go:embed dashboards
var dashboards embed.FS

// Verify that GrafanaService implements the ServiceAPI interface.
var _ monitoring.ServiceAPI = &GrafanaService{}

// GrafanaService implements the ServiceAPI interface for a Grafana service.
type GrafanaService struct {
	containerIP net.IP
	port        uint16
	stack       *datadir.MonitoringStack
}

// NewGrafana creates a new GrafanaService.
func NewGrafana() *GrafanaService {
	return &GrafanaService{}
}

// Init initializes the Grafana service with the given options.
func (g *GrafanaService) Init(opts types.ServiceOptions) error {
	// Validate dotEnv
	grafanaPort, ok := opts.Dotenv["GRAFANA_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "GRAFANA_PORT")
	} else if grafanaPort == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "GRAFANA_PORT")
	}

	port, err := strconv.ParseUint(opts.Dotenv["GRAFANA_PORT"], 10, 16)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid port", ErrInvalidOptions, "GRAFANA_PORT")
	}
	g.port = uint16(port)
	g.stack = opts.Stack
	return nil
}

func (g *GrafanaService) AddTarget(target types.MonitoringTarget, labels map[string]string, jobName string) error {
	return nil
}

func (g *GrafanaService) RemoveTarget(instanceId string) (string, error) {
	return "", nil
}

// DotEnv returns the dotenv variables and default values for the Grafana service.
func (g *GrafanaService) DotEnv() map[string]string {
	return dotEnv
}

// Setup sets up the Grafana service provisioning and configuration with the given dotenv values.
func (g *GrafanaService) Setup(options map[string]string) error {
	// Validate options
	promPort, ok := options["PROM_PORT"]
	if !ok {
		return fmt.Errorf("%w: %s missing in options", ErrInvalidOptions, "PROM_PORT")
	} else if promPort == "" {
		return fmt.Errorf("%w: %s can't be empty", ErrInvalidOptions, "PROM_PORT")
	}

	// Read config template
	rawTmp, err := config.ReadFile("config/prom.yml")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfigNotFound, err)
	}
	// Load template
	tmp, err := template.New("prom.yml").Parse(string(rawTmp))
	if err != nil {
		return err
	}

	// Create config directory
	grafProvPath := filepath.Join("grafana", "provisioning")
	if err = g.stack.CreateDir(filepath.Join(grafProvPath, "datasources")); err != nil {
		return err
	}
	// Create config file
	configFile, err := g.stack.Create(filepath.Join(grafProvPath, "datasources", "prom.yml"))
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Execute template
	data := struct {
		PromEndpoint string
	}{
		PromEndpoint: fmt.Sprintf("http://%s:%s", monitoring.PrometheusServiceName, options["PROM_PORT"]),
	}
	err = tmp.Execute(configFile, data)
	if err != nil {
		return err
	}

	// Create provisioning dashboards folder
	if err = g.stack.CreateDir(filepath.Join(grafProvPath, "dashboards")); err != nil {
		return err
	}
	// Create dashboards provisioning file
	dashboardsFile, err := g.stack.Create(filepath.Join(grafProvPath, "dashboards", "dashboards.yml"))
	if err != nil {
		return err
	}
	defer dashboardsFile.Close()
	dbs, err := config.ReadFile("config/dashboards.yml")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfigNotFound, err)
	}
	if _, err = dashboardsFile.Write(dbs); err != nil {
		return err
	}

	// Copy dashboards
	if err = g.copyDashboards(filepath.Join("grafana", "data")); err != nil {
		return err
	}

	return nil
}

// copyDashboards copy dashboards to $DATA_DIR/dashboards
func (g *GrafanaService) copyDashboards(dst string) (err error) {
	return fs.WalkDir(dashboards, "dashboards", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			dashboard, err := dashboards.Open(path)
			if err != nil {
				return err
			}
			defer func() {
				cerr := dashboard.Close()
				if err == nil {
					err = cerr
				}
			}()
			data, err := io.ReadAll(dashboard)
			if err != nil {
				return err
			}
			if err = g.stack.WriteFile(filepath.Join(dst, path), data); err != nil {
				return err
			}
		} else {
			if err = g.stack.CreateDir(filepath.Join(dst, path)); err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GrafanaService) SetContainerIP(ip net.IP) {
	g.containerIP = ip
}

func (g *GrafanaService) ContainerName() string {
	return monitoring.GrafanaContainerName
}

func (g *GrafanaService) Endpoint() string {
	if g.containerIP == nil {
		return fmt.Sprintf("http://%s:%d", g.ContainerName(), g.port)
	}
	return fmt.Sprintf("http://%s:%d", g.containerIP, g.port)
}

func (g *GrafanaService) Name() string {
	return monitoring.GrafanaServiceName
}
