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
		func(t *testing.T, binaryPath string) {
			runErr = runSedge(t, binaryPath, "generate", "full-node", "--network", "goerli")
		},
		// Assert
		func(t *testing.T) {
			assert.Error(t, runErr, "generate command should fail without arguments")
		},
	)
	// Run test case
	e2eTest.run()
}
