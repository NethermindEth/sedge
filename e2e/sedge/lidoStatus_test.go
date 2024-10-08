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
	"testing"

	"github.com/stretchr/testify/assert"

	base "github.com/NethermindEth/sedge/e2e"
)

func TestE2E_LidoStatus_ValidFlags(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "0x22bA5CaFB5E26E6Fe51f330294209034013A5A4c", "--l", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "lido status command should succeed with the given arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_LidoStatus_ValidNodeID(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "--nodeID", "10", "--l", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "lido status command should succeed with the given arguments")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_LidoStatus_InvalidNodeID(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "--nodeID", "-2", "--l", "--network", "holesky")
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

func TestE2E_LidoStatus_InvalidRewardAddress(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "0xccb", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_LidoStatus_RewardAddressNotFound(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "0xC870Fd7316956C1582A2c8Fd2c42552cCEC70C89", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestE2E_LidoStatus_InvalidZeroRewardAddress(t *testing.T) {
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
			runErr = base.RunCommand(t, binaryPath, "sedge", "lido-status", "0x0000000000000000000000000000000000000000", "--network", "holesky")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "lido status command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}
