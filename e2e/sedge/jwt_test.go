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
	"os"
	"path/filepath"
	"testing"

	base "github.com/NethermindEth/sedge/e2e"
	"github.com/stretchr/testify/assert"
)

const (
	// validJWTSecret is a 32-byte hex string used as a valid JWT secret
	validJWTSecret = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	// jwtSecretPath is the default path for JWT secrets
	jwtSecretPath = "jwtsecret"
)

func TestE2E_Generate_JWTSecret_RelativePath(t *testing.T) {
	// Test context
	var (
		runErr  error
		jwtPath = "custom/path/jwtsecret"
	)

	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		func(t *testing.T, sedgePath string) error {
			// Create the JWT secret file in the data directory
			dataDir := filepath.Join(filepath.Dir(sedgePath), "sedge-data")
			fullJWTPath := filepath.Join(dataDir, jwtPath)
			if err := os.MkdirAll(filepath.Dir(fullJWTPath), 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(fullJWTPath, []byte(validJWTSecret), 0o644); err != nil {
				return err
			}
			return nil
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			// Use the full path when running the generate command
			fullJWTPath := filepath.Join(dataDirPath, jwtPath)
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "mainnet", "--jwt-secret-path", fullJWTPath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)
			// Check if JWT file exists in the relative path
			fullJWTPath := filepath.Join(dataDirPath, jwtPath)
			_, err := os.Stat(fullJWTPath)
			assert.NoError(t, err, "JWT secret file should exist at relative path")

			// Verify the JWT secret content
			content, err := os.ReadFile(fullJWTPath)
			assert.NoError(t, err)
			assert.Equal(t, validJWTSecret, string(content), "JWT secret content should match")
		},
	)
	e2eTest.run()
}

func TestE2E_Generate_JWTSecret_AbsolutePath(t *testing.T) {
	// Test context
	var (
		runErr     error
		tempJWTDir string
	)

	// Build test case
	e2eTest := newE2ESedgeTestCase(
		t,
		// Arrange
		func(t *testing.T, sedgePath string) error {
			// Create a temporary directory for JWT secret
			tempJWTDir = t.TempDir()
			// Create the JWT secret file
			jwtPath := filepath.Join(tempJWTDir, "jwtsecret")
			return os.WriteFile(jwtPath, []byte(validJWTSecret), 0o644)
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			jwtPath := filepath.Join(tempJWTDir, "jwtsecret")
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "mainnet", "--jwt-secret-path", jwtPath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)
			// Check if JWT file exists in the absolute path
			jwtPath := filepath.Join(tempJWTDir, "jwtsecret")
			_, err := os.Stat(jwtPath)
			assert.NoError(t, err, "JWT secret file should exist at absolute path")

			// Verify the JWT secret content
			content, err := os.ReadFile(jwtPath)
			assert.NoError(t, err)
			assert.Equal(t, validJWTSecret, string(content), "JWT secret content should match")
		},
	)
	e2eTest.run()
}

func TestE2E_Generate_JWTSecret_DefaultPath(t *testing.T) {
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
			runErr = base.RunSedge(t, binaryPath, "generate", "full-node", "--network", "mainnet")
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr)
			// Check if JWT file exists in the default path
			defaultJWTPath := filepath.Join(dataDirPath, "jwtsecret")
			_, err := os.Stat(defaultJWTPath)
			assert.NoError(t, err, "JWT secret file should exist at default path")

			// Read the file to ensure it's not empty
			content, err := os.ReadFile(defaultJWTPath)
			assert.NoError(t, err)
			assert.NotEmpty(t, content, "JWT secret file should not be empty")
		},
	)
	e2eTest.run()
}
