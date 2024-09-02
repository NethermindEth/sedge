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

import (
	"bytes"
	"embed"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring/compose"
	"github.com/NethermindEth/sedge/internal/monitoring/data"
	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	"github.com/NethermindEth/sedge/internal/monitoring/monitoring/services/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	funk "github.com/thoas/go-funk"
)

//go:embed script
var script embed.FS

// MonitoringManager manages the monitoring services. It provides methods for initializing the monitoring stack,
// adding and removing targets, running and stopping the monitoring stack, and checking the status of the monitoring stack.
type MonitoringManager struct {
	services       []ServiceAPI
	composeManager ComposeManager
	dockerManager  DockerManager
	stack          *data.MonitoringStack
}

// NewMonitoringManager creates a new MonitoringManager with the given services, compose manager, docker manager, file system, and locker.
func NewMonitoringManager(
	services []ServiceAPI,
	cmpMgr ComposeManager,
	dockerMgr DockerManager,
	fs afero.Fs,
	locker locker.Locker,
) *MonitoringManager {
	// Create stack
	datadir, err := data.NewDataDirDefault(fs, locker)
	if err != nil {
		log.Fatal(err)
	}
	stack, err := datadir.MonitoringStack()
	if err != nil {
		log.Fatal(err)
	}

	return &MonitoringManager{
		services:       services,
		composeManager: cmpMgr,
		dockerManager:  dockerMgr,
		stack:          stack,
	}
}

// Init initializes the monitoring stack. Assumes that the stack is already installed.
func (m *MonitoringManager) Init() error {
	// Read installed .env
	rawDotEnv, err := m.stack.ReadFile(filepath.Join(".env"))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInitializingMonitoringMngr, err)
	}

	dotEnv := make(map[string]string)
	for _, line := range bytes.Split(rawDotEnv, []byte("\n")) {
		split := bytes.Split(line, []byte("="))
		if len(split) != 2 {
			continue
		}
		dotEnv[string(split[0])] = string(split[1])
	}

	// Initialize stack
	for _, service := range m.services {
		if err := service.Init(types.ServiceOptions{
			Stack:  m.stack,
			Dotenv: dotEnv,
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrInitializingMonitoringMngr, err)
		}
	}

	// Save container IPs of monitoring services
	if err := m.saveServiceIP(); err != nil {
		return fmt.Errorf("%w: %w", ErrInitializingMonitoringMngr, err)
	}

	return nil
}

// InitStack initializes the monitoring stack by merging all environment variables, checking ports, setting up the stack and services, and creating containers.
func (m *MonitoringManager) InstallStack() error {
	// Merge all dotEnv
	dotEnv := make(map[string]string)
	defaultPorts := make(map[string]uint16)
	for _, service := range m.services {
		for k, v := range service.DotEnv() {
			dotEnv[k] = v
			// Grab default ports
			if strings.HasSuffix(k, "_PORT") {
				// Cast string to uint16
				p, err := strconv.ParseUint(v, 10, 16)
				if err != nil {
					return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
				}
				defaultPorts[k] = uint16(p)
			}
		}
	}

	// Check ports
	ports, err := assignPorts("localhost", defaultPorts)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
	}
	for k, v := range ports {
		dotEnv[k] = strconv.Itoa(int(v))
	}

	// Intialize stack
	for _, service := range m.services {
		if err := service.Init(types.ServiceOptions{
			Stack:  m.stack,
			Dotenv: dotEnv,
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
		}
	}

	if err = m.stack.Setup(dotEnv, script); err != nil {
		return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
	}

	// Setup services
	log.Debug("Setting up monitoring stack...")
	for _, service := range m.services {
		if err = service.Setup(dotEnv); err != nil {
			return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
		}
	}

	// Create containers
	if err = m.composeManager.Create(compose.DockerComposeCreateOptions{Path: filepath.Join(m.stack.Path(), "docker-compose.yml")}); err != nil {
		return fmt.Errorf("%w: %w", ErrInstallingMonitoringMngr, err)
	}

	log.Debug("Starting monitoring stack...")
	if err := m.composeManager.Up(compose.DockerComposeUpOptions{Path: filepath.Join(m.stack.Path(), "docker-compose.yml")}); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
	}

	// Save container IPs of monitoring services
	if err := m.saveServiceIP(); err != nil {
		return fmt.Errorf("%w: %w", ErrInitializingMonitoringMngr, err)
	}

	return nil
}

// AddTarget adds a new target to all services in the monitoring stack.
// It also connects the target to the docker network of the monitoring stack if it isn't already connected.
// The labels are added to the service's metrics.
func (m *MonitoringManager) AddTarget(target types.MonitoringTarget, labels map[string]string, dockerNetwork string) error {
	for _, service := range m.services {
		// Check if network was already added to service
		containerName := service.ContainerName()
		if containerName == PrometheusContainerName {
			networks, err := m.dockerManager.ContainerNetworks(containerName)
			if err != nil {
				return err
			}
			if !funk.Contains(networks, dockerNetwork) {
				if err := m.dockerManager.NetworkConnect(containerName, dockerNetwork); err != nil {
					return err
				}
			}
		}
		if err := service.AddTarget(target, labels, labels[InstanceIDLabel]+"--"+containerName+"++"+dockerNetwork); err != nil {
			return err
		}
	}
	return nil
}

// RemoveTarget removes a target from all services in the monitoring stack.
// It also disconnects the target from the docker network of the monitoring stack if it isn't already disconnected.
func (m *MonitoringManager) RemoveTarget(instanceID string) error {
	for _, service := range m.services {
		network, err := service.RemoveTarget(instanceID)
		if err != nil {
			return err
		}
		// Disconnect may fail if the network was already disconnected or if the container was already removed
		// so we ignore the error
		serviceName := service.ContainerName()
		if err := m.dockerManager.NetworkDisconnect(serviceName, network); err != nil {
			log.Debugf("Error disconnecting %s from %s: %s", serviceName, network, err)
		}
	}
	return nil
}

// Run starts the monitoring stack by shutting down any existing stack and starting a new one.
func (m *MonitoringManager) Run() error {
	log.Info("Starting monitoring stack...")
	if err := m.composeManager.Up(compose.DockerComposeUpOptions{Path: filepath.Join(m.stack.Path(), "docker-compose.yml")}); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
	}

	// Save container IPs of monitoring services
	if err := m.saveServiceIP(); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
	}

	return nil
}

// Stop shuts down the monitoring stack.
func (m *MonitoringManager) Stop() error {
	log.Info("Shutting down monitoring stack...")
	if err := m.composeManager.Down(compose.DockerComposeDownOptions{Path: filepath.Join(m.stack.Path(), "docker-compose.yml")}); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
	}

	return nil
}

// Status checks the status of the containers in the monitoring stack and returns the status.
func (m *MonitoringManager) Status() (status common.Status, err error) {
	containers := []string{
		GrafanaContainerName,
		PrometheusContainerName,
		NodeExporterContainerName,
	}

	for _, container := range containers {
		status, err = m.dockerManager.ContainerStatus(container)
		if err != nil {
			return common.Unknown, fmt.Errorf("%w: %w", ErrCheckingMonitoringStack, err)
		}
		// running or restarting means the stack is running
		if status != common.Running && status != common.Restarting {
			return common.Broken,
				fmt.Errorf("%w: %s container is either paused, exited, or dead", ErrCheckingMonitoringStack, container)
		}
	}

	return status, nil
}

// InstallationStatus checks whether the monitoring stack is installed and returns the installation status.
func (m *MonitoringManager) InstallationStatus() (common.Status, error) {
	installed, err := m.stack.Installed()
	if err != nil {
		return common.Unknown, fmt.Errorf("%w: %w", ErrCheckingMonitoringStack, err)
	}
	if installed {
		return common.Installed, nil
	}

	return common.NotInstalled, nil
}

// Cleanup removes the monitoring stack. If force is true, it bypasses locks and removes the stack without running 'docker compose down'.
func (m *MonitoringManager) Cleanup(force bool) error {
	if !force {
		log.Info("Shutting down monitoring stack...")
		if err := m.composeManager.Down(compose.DockerComposeDownOptions{Path: filepath.Join(m.stack.Path(), "docker-compose.yml"), Volumes: true}); err != nil {
			return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
		}
	}

	log.Info("Cleaning up monitoring stack...")
	if err := m.stack.Cleanup(force); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningMonitoringStack, err)
	}

	return nil
}

// ServiceEndpoints returns a map of the service's container names and their endpoints.
func (m *MonitoringManager) ServiceEndpoints() map[string]string {
	endpoints := make(map[string]string)
	for _, service := range m.services {
		endpoints[service.ContainerName()] = service.Endpoint()
	}
	return endpoints
}

func (m *MonitoringManager) idToIP(id string) (string, error) {
	ip, err := m.dockerManager.ContainerIP(id)
	if err != nil {
		return "", err
	}
	return ip, nil
}

func (m *MonitoringManager) saveServiceIP() error {
	for _, service := range m.services {
		name := service.ContainerName()
		ip, err := m.idToIP(name)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInitializingMonitoringMngr, err)
		}
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			return fmt.Errorf("%w: failed to save the IP address of the monitoring service %s: %s is not a valid IP address", ErrInitializingMonitoringMngr, name, ip)
		}
		service.SetContainerIP(parsedIP)
	}
	return nil
}
