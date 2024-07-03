package e2e

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGenerate_FullNode_GoerliNotSupported(t *testing.T) {
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
		func(t *testing.T, binaryPath, dataDirPath string) {
			runErr = runSedge(t, binaryPath, "generate", "full-node", "--network", "goerli")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail without arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestGenerate_FullNode_Lido_GnosisNotSupported(t *testing.T) {
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
			runErr = runSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "gnosis")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestGenerate_FullNode_Lido_SepoliaNotSupported(t *testing.T) {
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
			runErr = runSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "sepolia")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "generate command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestGenerate_FullNode_Lido_Sepolia_NoMEV(t *testing.T) {
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
			runErr = runSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "sepolia", "--no-mev-boost")
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
			expectedFeeRecipient := contracts.FeeRecipient["sepolia"].FeeRecipientAddress
			assert.Equal(t, expectedFeeRecipient, feeRecipient, "FEE_RECIPIENT value should match expected value")

			// Read RELAY_URLS value
			_, exists = envMap["RELAY_URLS"]
			assert.False(t, exists, "RELAY_URLS shouldn't exist in .env file")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestGenerate_FullNode_Lido_Mainnet(t *testing.T) {
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
			runErr = runSedge(t, binaryPath, "generate", "--lido", "full-node", "--network", "mainnet")
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
			expectedFeeRecipient := contracts.FeeRecipient["mainnet"].FeeRecipientAddress
			assert.Equal(t, expectedFeeRecipient, feeRecipient, "FEE_RECIPIENT value should match expected value")

			// Read RELAY_URLS value
			relayURLs, exists := envMap["RELAY_URLS"]
			assert.True(t, exists, "RELAY_URLS should exist in .env file")
			relayURLsList := strings.Split(relayURLs, ",")
			expectedRelayURLs, _ := mevboostrelaylist.GetRelaysURI("mainnet")
			assert.Equal(t, expectedRelayURLs, relayURLsList, "RELAY_URLS value should match expected value")
		},
	)
	// Run test case
	e2eTest.run()
}
