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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/utils"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	base "github.com/NethermindEth/sedge/e2e"
)

func TestE2E_Generate_FullNode_GoerliNotSupported(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "goerli")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail without arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lido_GnosisNotSupported(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "gnosis")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lido_SepoliaNotSupported(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "sepolia")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			// Sepolia not supported Lido with MEV-Boost
			assert.Error(t, runErr, "generate command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lido_Sepolia_NoMEV(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "sepolia", "--no-mev-boost")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed with the given arguments")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read FEE_RECIPIENT value
			feeRecipient, exists := envMap["FEE_RECIPIENT"]
			assert.True(t, exists, "FEE_RECIPIENT should exist in .env file")
			expectedFeeRecipient, ok := contracts.FeeRecipient("sepolia")
			assert.True(t, ok, "FeeRecipient should be found")
			assert.Equal(t, expectedFeeRecipient, feeRecipient, "FEE_RECIPIENT value should match expected value")

			// Read RELAY_URLS value
			_, exists = envMap["RELAY_URLS"]
			assert.False(t, exists, "RELAY_URLS shouldn't exist in .env file")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lido_Mainnet(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "mainnet")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read FEE_RECIPIENT value
			feeRecipient, exists := envMap["FEE_RECIPIENT"]
			assert.True(t, exists, "FEE_RECIPIENT should exist in .env file")
			expectedFeeRecipient, ok := contracts.FeeRecipient("mainnet")
			assert.True(t, ok, "FeeRecipient should be found")
			assert.Equal(t, expectedFeeRecipient, feeRecipient, "FEE_RECIPIENT value should match expected value")

			// Read RELAY_URLS value
			relayURLs, exists := envMap["RELAY_URLS"]
			assert.True(t, exists, "RELAY_URLS should exist in .env file")
			relayURLsList := strings.Split(relayURLs, ",")
			expectedRelayURLs, _ := mevboostrelaylist.RelaysURI("mainnet")
			assert.Equal(t, expectedRelayURLs, relayURLsList, "RELAY_URLS value should match expected value")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Nimbus_Mainnet(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "mainnet", "-v", "nimbus")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read CC_IMAGE_VERSION value
			ccImageVersion, exists := envMap["CC_IMAGE_VERSION"]
			assert.True(t, exists, "FEE_RECIPIENT should exist in .env file")
			assert.Contains(t, ccImageVersion, "nimbus", "CC_IMAGE_VERSION value should contain 'nimbus'")

			// Read VL_IMAGE_VERSION value
			vlImageVersion, exists := envMap["VL_IMAGE_VERSION"]
			assert.True(t, exists, "VL_IMAGE_VERSION should exist in .env file")
			assert.Contains(t, vlImageVersion, "nimbus", "VL_IMAGE_VERSION value should contain 'nimbus'")

			// Check that the docker-compose contains the consensus and validator services
			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			err = utils.ValidateCompose(dockerComposeFilePath)
			assert.NoError(t, err, "docker-compose file should be valid")
			composeServices, err := utils.LoadDockerComposeServices(dockerComposeFilePath)
			for _, service := range []string{"execution", "consensus", "validator"} {
				assert.Contains(t, composeServices, service, fmt.Sprintf("docker-compose file should contain service %s", service))
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lighthouse_Hoodi(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "hoodi", "-v", "lighthouse", "--no-mev-boost")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read Network value
			network, exists := envMap["NETWORK"]
			assert.True(t, exists, "NETWORK should exist in .env file")
			assert.Equal(t, "hoodi", network, "NETWORK value should be 'hoodi'")

			// Read CC_IMAGE_VERSION value
			ccImageVersion, exists := envMap["CC_IMAGE_VERSION"]
			assert.True(t, exists, "FEE_RECIPIENT should exist in .env file")
			assert.Contains(t, ccImageVersion, "lighthouse", "CC_IMAGE_VERSION value should contain 'lighthouse'")

			// Read VL_IMAGE_VERSION value
			vlImageVersion, exists := envMap["VL_IMAGE_VERSION"]
			assert.True(t, exists, "VL_IMAGE_VERSION should exist in .env file")
			assert.Contains(t, vlImageVersion, "lighthouse", "VL_IMAGE_VERSION value should contain 'lighthouse'")

			// Check that the docker-compose contains the consensus and validator services
			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			err = utils.ValidateCompose(dockerComposeFilePath)
			assert.NoError(t, err, "docker-compose file should be valid")
			composeServices, err := utils.LoadDockerComposeServices(dockerComposeFilePath)
			for _, service := range []string{"execution", "consensus", "validator"} {
				assert.Contains(t, composeServices, service, fmt.Sprintf("docker-compose file should contain service %s", service))
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_FullNode_Lido_Hoodi(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "hoodi")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read FEE_RECIPIENT value
			feeRecipient, exists := envMap["FEE_RECIPIENT"]
			assert.True(t, exists, "FEE_RECIPIENT should exist in .env file")
			expectedFeeRecipient, ok := contracts.FeeRecipient("hoodi")
			assert.True(t, ok, "FeeRecipient should be found")
			assert.Equal(t, expectedFeeRecipient, feeRecipient, "FEE_RECIPIENT value should match expected value")

			// Read RELAY_URLS value
			relayURLs, exists := envMap["RELAY_URLS"]
			assert.True(t, exists, "RELAY_URLS should exist in .env file")
			relayURLsList := strings.Split(relayURLs, ",")
			expectedRelayURLs, _ := mevboostrelaylist.RelaysURI("hoodi")
			assert.Equal(t, expectedRelayURLs, relayURLsList, "RELAY_URLS value should match expected value")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecSequencer_Sepolia(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			// Create data directory if it doesn't exist
			err := os.MkdirAll(dataDirPath, 0o755)
			if err != nil {
				t.Fatalf("Failed to create data directory: %v", err)
			}

			// Prepare valid keystore file
			testDataPath := filepath.Join("cli", "testdata", "cli_tests", "aztec_keystore", "valid_single")
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get working directory: %v", err)
			}
			repoPath := filepath.Dir(filepath.Dir(wd))
			sourcePath := filepath.Join(repoPath, testDataPath)
			destPath := filepath.Join(dataDirPath, "keystore.json")

			// Copy keystore file
			sourceFile, err := os.Open(filepath.Join(sourcePath, "keystore.json"))
			if err != nil {
				t.Fatalf("Failed to open source keystore: %v", err)
			}
			defer sourceFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				t.Fatalf("Failed to create dest keystore: %v", err)
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				t.Fatalf("Failed to copy keystore: %v", err)
			}

			keystorePath := filepath.Join(dataDirPath, "keystore.json")
			runErr = base.RunSedge(t, binaryPath, "generate", "aztec", "--network", "sepolia", "--type", "sequencer", "--aztec-keystore-path", keystorePath, "--aztec-p2p-ip", "192.168.1.100")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read Network value
			network, exists := envMap["NETWORK"]
			assert.True(t, exists, "NETWORK should exist in .env file")
			assert.Equal(t, "sepolia", network, "NETWORK value should be 'sepolia'")

			// Read AZTEC_KEYSTORE_PATH value
			aztecKeystorePath, exists := envMap["AZTEC_KEYSTORE_PATH"]
			assert.True(t, exists, "AZTEC_KEYSTORE_PATH should exist in .env file")
			assert.NotEmpty(t, aztecKeystorePath, "AZTEC_KEYSTORE_PATH should not be empty")

			// Read AZTEC_IMAGE_VERSION value
			aztecImage, exists := envMap["AZTEC_IMAGE_VERSION"]
			assert.True(t, exists, "AZTEC_IMAGE_VERSION should exist in .env file")
			assert.Contains(t, aztecImage, "aztec", "AZTEC_IMAGE_VERSION should contain 'aztec'")

			// Check that the docker-compose contains the required services
			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			err = utils.ValidateCompose(dockerComposeFilePath)
			assert.NoError(t, err, "docker-compose file should be valid")
			composeServices, err := utils.LoadDockerComposeServices(dockerComposeFilePath)
			assert.NoError(t, err, "should be able to load docker-compose services")
			for _, service := range []string{"execution", "consensus", "aztec"} {
				assert.Contains(t, composeServices, service, fmt.Sprintf("docker-compose file should contain service %s", service))
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecSequencer_Mainnet(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			// Create data directory if it doesn't exist
			err := os.MkdirAll(dataDirPath, 0o755)
			if err != nil {
				t.Fatalf("Failed to create data directory: %v", err)
			}

			// Prepare valid keystore file
			testDataPath := filepath.Join("cli", "testdata", "cli_tests", "aztec_keystore", "valid_single")
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get working directory: %v", err)
			}
			repoPath := filepath.Dir(filepath.Dir(wd))
			sourcePath := filepath.Join(repoPath, testDataPath)
			destPath := filepath.Join(dataDirPath, "keystore.json")

			// Copy keystore file
			sourceFile, err := os.Open(filepath.Join(sourcePath, "keystore.json"))
			if err != nil {
				t.Fatalf("Failed to open source keystore: %v", err)
			}
			defer sourceFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				t.Fatalf("Failed to create dest keystore: %v", err)
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				t.Fatalf("Failed to copy keystore: %v", err)
			}

			keystorePath := filepath.Join(dataDirPath, "keystore.json")
			runErr = base.RunSedge(t, binaryPath, "generate", "aztec", "--network", "mainnet", "--type", "sequencer", "--aztec-keystore-path", keystorePath, "--aztec-p2p-ip", "192.168.1.100")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read Network value
			network, exists := envMap["NETWORK"]
			assert.True(t, exists, "NETWORK should exist in .env file")
			assert.Equal(t, "mainnet", network, "NETWORK value should be 'mainnet'")

			// Read AZTEC_KEYSTORE_PATH value
			aztecKeystorePath, exists := envMap["AZTEC_KEYSTORE_PATH"]
			assert.True(t, exists, "AZTEC_KEYSTORE_PATH should exist in .env file")
			assert.NotEmpty(t, aztecKeystorePath, "AZTEC_KEYSTORE_PATH should not be empty")

			// Check that the docker-compose contains the required services
			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			err = utils.ValidateCompose(dockerComposeFilePath)
			assert.NoError(t, err, "docker-compose file should be valid")
			composeServices, err := utils.LoadDockerComposeServices(dockerComposeFilePath)
			assert.NoError(t, err, "should be able to load docker-compose services")
			for _, service := range []string{"execution", "consensus", "aztec"} {
				assert.Contains(t, composeServices, service, fmt.Sprintf("docker-compose file should contain service %s", service))
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecSequencer_ExtraFlag(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			// Create data directory if it doesn't exist
			err := os.MkdirAll(dataDirPath, 0o755)
			if err != nil {
				t.Fatalf("Failed to create data directory: %v", err)
			}

			// Prepare valid keystore file
			testDataPath := filepath.Join("cli", "testdata", "cli_tests", "aztec_keystore", "valid_single")
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get working directory: %v", err)
			}
			repoPath := filepath.Dir(filepath.Dir(wd))
			sourcePath := filepath.Join(repoPath, testDataPath)
			destPath := filepath.Join(dataDirPath, "keystore.json")

			// Copy keystore file
			sourceFile, err := os.Open(filepath.Join(sourcePath, "keystore.json"))
			if err != nil {
				t.Fatalf("Failed to open source keystore: %v", err)
			}
			defer sourceFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				t.Fatalf("Failed to create dest keystore: %v", err)
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				t.Fatalf("Failed to copy keystore: %v", err)
			}

			keystorePath := filepath.Join(dataDirPath, "keystore.json")
			runErr = base.RunSedge(t, binaryPath, "generate", "aztec",
				"--network", "sepolia",
				"--type", "sequencer",
				"--aztec-keystore-path", keystorePath,
				"--aztec-p2p-ip", "192.168.1.100",
				"--aztec-extra-flag", "p2p.maxPeers=200",
			)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")

			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			composeBytes, err := os.ReadFile(dockerComposeFilePath)
			assert.NoError(t, err, "should be able to read docker-compose.yml")
			assert.Contains(t, string(composeBytes), "--p2p.maxPeers=200", "docker-compose should include aztec extra flag")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecSequencer_MissingKeystorePath(t *testing.T) {
	// Test context
	var (
		runErr error
		output []byte
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			output, runErr = base.RunSedgeWithOutput(t, binaryPath, "generate", "aztec", "--network", "sepolia", "--type", "sequencer", "--aztec-p2p-ip", "192.168.1.100")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail without keystore path")
			assert.Contains(t, string(output), "aztec-keystore-path is required", "error should mention missing keystore path")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecSequencer_MissingP2pIp(t *testing.T) {
	// Test context
	var (
		runErr error
		output []byte
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			// Create data directory if it doesn't exist
			err := os.MkdirAll(dataDirPath, 0o755)
			if err != nil {
				t.Fatalf("Failed to create data directory: %v", err)
			}

			// Prepare valid keystore file
			testDataPath := filepath.Join("cli", "testdata", "cli_tests", "aztec_keystore", "valid_single")
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get working directory: %v", err)
			}
			repoPath := filepath.Dir(filepath.Dir(wd))
			sourcePath := filepath.Join(repoPath, testDataPath)
			destPath := filepath.Join(dataDirPath, "keystore.json")

			// Copy keystore file
			sourceFile, err := os.Open(filepath.Join(sourcePath, "keystore.json"))
			if err != nil {
				t.Fatalf("Failed to open source keystore: %v", err)
			}
			defer sourceFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				t.Fatalf("Failed to create dest keystore: %v", err)
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				t.Fatalf("Failed to copy keystore: %v", err)
			}

			keystorePath := filepath.Join(dataDirPath, "keystore.json")
			output, runErr = base.RunSedgeWithOutput(t, binaryPath, "generate", "aztec", "--network", "sepolia", "--type", "sequencer", "--aztec-keystore-path", keystorePath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail without P2P IP")
			assert.Contains(t, string(output), "aztec-p2p-ip is required when generating aztec configuration", "error should mention missing P2P IP")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_Generate_AztecFullNode_Sepolia(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath string, dataDirPath string) {
			runErr = base.RunSedge(t, binaryPath, "generate", "aztec", "--network", "sepolia", "--type", "node")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "generate command should succeed")
			generateDataFilePath := filepath.Join(dataDirPath, ".env")
			assert.FileExists(t, generateDataFilePath, ".env file should be created")

			// Read .env file
			envMap, err := godotenv.Read(generateDataFilePath)
			assert.NoError(t, err, "should be able to read .env file")

			// Read Network value
			network, exists := envMap["NETWORK"]
			assert.True(t, exists, "NETWORK should exist in .env file")
			assert.Equal(t, "sepolia", network, "NETWORK value should be 'sepolia'")

			// AZTEC_IMAGE_VERSION should exist
			aztecImage, exists := envMap["AZTEC_IMAGE_VERSION"]
			assert.True(t, exists, "AZTEC_IMAGE_VERSION should exist in .env file")
			assert.Contains(t, aztecImage, "aztec", "AZTEC_IMAGE_VERSION should contain 'aztec'")

			// Full node should not require keystore; env may include AZTEC_KEYSTORE_PATH but it should be empty.
			aztecKeystorePath, exists := envMap["AZTEC_KEYSTORE_PATH"]
			assert.True(t, exists, "AZTEC_KEYSTORE_PATH should exist in .env file")
			assert.Empty(t, aztecKeystorePath, "AZTEC_KEYSTORE_PATH should be empty for full node")

			// Check compose services and ensure it is NOT in sequencer mode
			dockerComposeFilePath := filepath.Join(dataDirPath, "docker-compose.yml")
			assert.FileExists(t, dockerComposeFilePath, "docker-compose.yml file should be created")
			err = utils.ValidateCompose(dockerComposeFilePath)
			assert.NoError(t, err, "docker-compose file should be valid")

			composeServices, err := utils.LoadDockerComposeServices(dockerComposeFilePath)
			assert.NoError(t, err, "should be able to load docker-compose services")
			for _, service := range []string{"execution", "consensus", "aztec"} {
				assert.Contains(t, composeServices, service, fmt.Sprintf("docker-compose file should contain service %s", service))
			}

			composeBytes, err := os.ReadFile(dockerComposeFilePath)
			assert.NoError(t, err, "should be able to read docker-compose.yml")
			assert.NotContains(t, string(composeBytes), "--sequencer", "full node docker-compose should not include --sequencer")
			assert.NotContains(t, string(composeBytes), "/var/lib/keystore", "full node docker-compose should not mount keystore")
		},
	)
	// Run test case
	e2eTest.run()
}
