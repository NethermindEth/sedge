package e2e

import (
	"testing"

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
		func(t *testing.T, binaryPath string) {
			runErr = runSedge(t, binaryPath, "generate", "full-node", "--lido", "--network", "sepolia", "--no-mev-boost")
		},
		// Assert
		func(t *testing.T) {
			assert.NoError(t, runErr, "generate command should succeed with the given arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestGenerate_FullNode_Lido_Sepolia(t *testing.T) {
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
		func(t *testing.T, binaryPath string) {
			runErr = runSedge(t, binaryPath, "generate", "full-node", "--lido", "--network", "sepolia")
		},
		// Assert
		func(t *testing.T) {
			assert.Error(t, runErr, "generate command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}
