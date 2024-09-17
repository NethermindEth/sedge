package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMonitoringStack_Init tests that the monitoring stack is not initialized if the user does not run the init-monitoring command
func TestE2E_MonitoringStack_NotInitialized(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = runCommand(t, binaryPath, "--help")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)

			checkMonitoringStackNotInstalled(t)
			checkMonitoringStackContainersNotRunning(t)
		},
	)
	// Run test case
	e2eTest.run()
}

// TestMonitoringStack_Init tests the monitoring stack initialization
func TestE2E_MonitoringStack_Init(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = runCommand(t, binaryPath, "monitoring", "init")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)
			checkMonitoringStackDir(t)
			checkMonitoringStackContainers(t)
			checkPrometheusTargetsUp(t, "sedge_node_exporter:9100")
			checkGrafanaHealth(t)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_MonitoringStack_NotReinstalled(t *testing.T) {
	// Test context
	var (
		grafanaContainerID      string
		prometheusContainerID   string
		nodeExporterContainerID string
		runErr                  error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		func(t *testing.T, sedgePath string) error {
			err := runCommand(t, sedgePath, "monitoring", "init")
			if err != nil {
				return err
			}
			grafanaContainerID, err = getContainerIDByName("sedge_grafana")
			if err != nil {
				return err
			}
			prometheusContainerID, err = getContainerIDByName("sedge_prometheus")
			if err != nil {
				return err
			}
			nodeExporterContainerID, err = getContainerIDByName("sedge_node_exporter")
			return err
		},
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = runCommand(t, binaryPath, "monitoring", "init")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)

			checkMonitoringStackDir(t)
			checkMonitoringStackContainers(t)
			checkGrafanaHealth(t)
			newGrafanaContainerID, err := getContainerIDByName("sedge_grafana")
			assert.NoError(t, err)
			assert.Equal(t, grafanaContainerID, newGrafanaContainerID, "grafana container ID has changed")
			newPrometheusContainerID, err := getContainerIDByName("sedge_prometheus")
			assert.NoError(t, err)
			assert.Equal(t, prometheusContainerID, newPrometheusContainerID, "prometheus container ID has changed")
			newNodeExporterContainerID, err := getContainerIDByName("sedge_node_exporter")
			assert.NoError(t, err)
			assert.Equal(t, nodeExporterContainerID, newNodeExporterContainerID, "node-exporter container ID has changed")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_MonitoringStack_Clean(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		func(t *testing.T, sedgePath string) error {
			return runCommand(t, sedgePath, "monitoring", "init")
		},
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = runCommand(t, binaryPath, "monitoring", "clean")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)

			// Check that monitoring stack directory is removed
			assert.NoDirExists(t, dataDirPath)

			// Check that monitoring stack containers are removed
			checkMonitoringStackContainersNotRunning(t)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_MonitoringStack_CleanNonExistent(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = runCommand(t, binaryPath, "monitoring", "clean")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)

			// Check that monitoring stack directory doesn't exist
			assert.NoDirExists(t, dataDirPath)

			// Check that monitoring stack containers don't exist
			checkMonitoringStackContainersNotRunning(t)
		},
	)
	// Run test case
	e2eTest.run()
}
