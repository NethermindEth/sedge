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
package e2e

import (
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	base "github.com/NethermindEth/sedge/e2e"
)

func TestE2E_ValidArgs_NodeOperatorID(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--node-operator-id", "25", "--network", "hoodi", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidArgs_EnvNodeOperatorID(t *testing.T) {
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			return map[string]string{
				"LIDO_EXPORTER_NODE_OPERATOR_ID": "250",
				"LIDO_EXPORTER_PORT":             strconv.Itoa(port),
			}, nil
		},
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

// func TestE2E_ValidArgs_RewardAddress(t *testing.T) {
// 	//t.Parallel()
// 	// Test context
// 	var (
// 		cmd *exec.Cmd
// 	)
// 	// Build test case
// 	e2eTest := newE2ELidoExporterTestCase(
// 		t,
// 		// Arrange
// 		nil,
// 		// Act
// 		func(t *testing.T, binaryPath string) *exec.Cmd {
// 			cmd = base.RunCommandCMD(t, binaryPath, "", "lido-exporter", "--reward-address", "0xe6b5A31d8bb53D2C769864aC137fe25F4989f1fd", "--network", "holesky", "--port", "9981")
// 			time.Sleep(5 * time.Second)

// 			return cmd
// 		},
// 		// Assert
// 		func(t *testing.T) {
// 			// With --reward-address, the test take too long to start the prometheus server due to the time it takes to get the NO ID from the reward address
// 			checkPrometheusServerUp(t, 9981)
// 			checkMetrics(t, 9981)

// 			cmd.Process.Signal(os.Interrupt)

// 			// Wait for the process to exit with a timeout
// 			done := make(chan error, 1)
// 			go func() {
// 				done <- cmd.Wait()
// 			}()

// 			select {
// 			case err := <-done:
// 				assert.NoError(t, err)
// 			case <-time.After(5 * time.Second):
// 				t.Error("Process did not exit within the timeout period")
// 				cmd.Process.Kill() // Force kill if it doesn't exit
// 			}
// 		},
// 	)
// 	// Run test case
// 	e2eTest.run()
// }

func TestE2E_MissingRequiredArgs(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)

	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			return map[string]string{
				"LIDO_EXPORTER_PORT":             strconv.Itoa(port),
				"LIDO_EXPORTER_NODE_OPERATOR_ID": "",
				"LIDO_EXPORTER_REWARD_ADDRESS":   "",
			}, nil
		},
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "", "lido-exporter", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()

			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with missing required arguments")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with missing required arguments")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_NodeOperatorID(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--node-operator-id", "lol_what_a_node_operator_id", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid node operator ID")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid node operator ID")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_RewardAddress(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--reward-address", "lol_what_a_reward_address", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid reward address")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid reward address")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_Network(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--network", "lol_what_a_network", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid network")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid network")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_ScrapeTime(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--scrape-time", "666", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid scrape time")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid scrape time")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_Port(t *testing.T) {
	// Test context
	var (
		cmd *exec.Cmd
	)
	downPort := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--port", "lol_what_a_port", "--network", "hoodi")
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid port")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid port")
			checkPrometheusServerDown(t, downPort)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_RPCEndpoints(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--rpc-endpoints", "lol_what_a_rpc_endpoint", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid RPC endpoints")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid RPC endpoints")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_WSEndpoints(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--ws-endpoints", "lol_what_a_ws_endpoint", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.Error(t, err, "lido-exporter command should fail with invalid WebSocket endpoints")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid WebSocket endpoints")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidFlags_All(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter",
				"--rpc-endpoints", "https://ethereum-hoodi.gateway.tatum.io", "https://0xrpc.io/hoodi",
				"--ws-endpoints", "wss://0xrpc.io/hoodi", // https endpoint should be ignored
				"--port", strconv.Itoa(port),
				"--scrape-time", "1s",
				"--network", "hoodi",
				"--node-operator-id", "25", // should be prioritized over reward address
				"--reward-address", "0x22bA5CaFB5E26E6Fe51f330294209034013A5A4c",
			)
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidEnv_All(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			return map[string]string{
				"LIDO_EXPORTER_RPC_ENDPOINTS":    "'https://ethereum-hoodi.gateway.tatum.io','https://0xrpc.io/hoodi'",
				"LIDO_EXPORTER_WS_ENDPOINTS":     "'wss://0xrpc.io/hoodi'",
				"LIDO_EXPORTER_PORT":             strconv.Itoa(port),
				"LIDO_EXPORTER_SCRAPE_TIME":      "2s",
				"LIDO_EXPORTER_NETWORK":          "hoodi",
				"LIDO_EXPORTER_NODE_OPERATOR_ID": "25",
				"LIDO_EXPORTER_REWARD_ADDRESS":   "0x22bA5CaFB5E26E6Fe51f330294209034013A5A4c",
			}, nil
		},
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_InvalidArgs_NegativeNodeID(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--node-operator-id", "-2", "--network", "hoodi", "--port", strconv.Itoa(port))
			return cmd
		},
		// Assert
		func(t *testing.T) {
			err := cmd.Wait()
			assert.NotContains(t, err.Error(), "killed")
			assert.Error(t, err, "lido-exporter command should fail with invalid node operator ID")
			// cmd should return status code 1
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "lido-exporter command should fail with invalid node operator ID")
			checkPrometheusServerDown(t, port)
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidArgs_NodeOperatorID_Mainnet(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--node-operator-id", "1", "--network", "mainnet", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidEnv_All_Mainnet(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			return map[string]string{
				"LIDO_EXPORTER_RPC_ENDPOINTS":    "'https://eth.llamarpc.com','https://eth-pokt.nodies.app','https://rpc.mevblocker.io'",
				"LIDO_EXPORTER_WS_ENDPOINTS":     "'wss://ethereum-rpc.publicnode.com'",
				"LIDO_EXPORTER_PORT":             strconv.Itoa(port),
				"LIDO_EXPORTER_SCRAPE_TIME":      "2s",
				"LIDO_EXPORTER_NETWORK":          "mainnet",
				"LIDO_EXPORTER_NODE_OPERATOR_ID": "1",
			}, nil
		},
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidFlags_All_Mainnet(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter",
				"--rpc-endpoints", "https://eth.llamarpc.com", "https://eth-pokt.nodies.app", "https://rpc.mevblocker.io",
				"--ws-endpoints", "wss://ethereum-rpc.publicnode.com", // https endpoint should be ignored
				"--port", strconv.Itoa(port),
				"--scrape-time", "1s",
				"--network", "mainnet",
				"--node-operator-id", "1", // should be prioritized over reward address
				"--reward-address", "0x22bA5CaFB5E26E6Fe51f330294209034013A5A4c",
			)
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_ValidArgs_NodeOperatorID_Hoodi(t *testing.T) {
	// t.Parallel()
	// Test context
	var (
		cmd *exec.Cmd
	)
	port := getFreePort(t)
	// Build test case
	e2eTest := newE2ELidoExporterTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = base.RunCommandCMD(t, binaryPath, "lido-exporter", "lido-exporter", "--node-operator-id", "1", "--network", "hoodi", "--port", strconv.Itoa(port))
			time.Sleep(2 * time.Second)
			return cmd
		},
		// Assert
		func(t *testing.T) {
			checkPrometheusServerUp(t, port)
			checkMetrics(t, port)

			_ = cmd.Process.Kill()

			// Wait for the process to exit with a timeout
			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				_ = err // process may exit non-zero when terminated
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill() // Force kill if it doesn't exit
			}
		},
	)
	// Run test case
	e2eTest.run()
}
