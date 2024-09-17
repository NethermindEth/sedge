package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLidoStatus_ValidFlags(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "0x22bA5CaFB5E26E6Fe51f330294209034013A5A4c", "--l", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "lido status command should succeed with the given arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestLidoStatus_ValidNodeID(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "--nodeID", "10", "--l", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "lido status command should succeed with the given arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestLidoStatus_InvalidNodeID(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "--nodeID", "-2", "--l", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			// Node ID can't be a negative value
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestLidoStatus_InvalidRewardAddress(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "0xccb", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestLidoStatus_RewardAddressNotFound(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "0xC870Fd7316956C1582A2c8Fd2c42552cCEC70C89", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestLidoStatus_InvalidZeroRewardAddress(t *testing.T) {
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
			runErr = runCommand(t, binaryPath, "lido-status", "0x0000000000000000000000000000000000000000", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}
