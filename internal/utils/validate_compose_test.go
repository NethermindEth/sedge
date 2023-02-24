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
package utils

import (
	"errors"
	"io"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestValidateCompose(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tcs := []struct {
		name      string
		pTestData string
		err       map[string]error
	}{
		{
			name:      "Valid docker-compose",
			pTestData: "valid",
		},
		{
			name:      "Valid docker-compose, no services",
			pTestData: "no_services",
		},
		{
			name:      "Invalid docker-compose, empty services, yaml schema error",
			pTestData: "bad_services",
			err: map[string]error{
				"linux":   errors.New("must be a mapping"),
				"darwin":  errors.New("must be a mapping"),
				"windows": errors.New("must be a mapping"),
			},
		},
		{
			name:      "Valid docker-compose, bad services",
			pTestData: "valid",
		},
		{
			name:      "Without env file",
			pTestData: "no_env",
			err: map[string]error{
				"linux":   errors.New("empty section between colons"),
				"darwin":  errors.New("empty section between colons"),
				"windows": errors.New("empty section between colons"),
			},
		},
		{
			name:      "No compose file in path",
			pTestData: "no_compose",
			err: map[string]error{
				"linux":   errors.New("no such file or directory"),
				"darwin":  errors.New("no such file or directory"),
				"windows": errors.New("The system cannot find the file specified"),
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			var path string
			if tc.pTestData != "" {
				tmp := t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "validate_compose_tests", tc.pTestData), tmp)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				path = filepath.Join(tmp, "docker-compose.yml")
			}

			err := ValidateCompose(path)

			if tc.err != nil {
				assert.ErrorContains(t, err, tc.err[runtime.GOOS].Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseEnv(t *testing.T) {
	tcs := []struct {
		name      string
		pTestData string
		out       map[string]string
		err       map[string]error
	}{
		{
			name:      "Valid env file, two keys",
			pTestData: "env_a",
			out:       map[string]string{"A": "666", "NETWORK": "mainnet"},
		},
		{
			name:      "Valid env file, one key, one comment",
			pTestData: "env_b",
			out:       map[string]string{"B": "007"},
		},
		{
			name:      "Invalid env file",
			pTestData: "bad_env",
			out:       map[string]string{"A ": " B"},
		},
		{
			name:      "Empty env file",
			pTestData: "empty_env",
			out:       map[string]string{},
		},
		{
			name:      "Without env file",
			pTestData: "no_env",
			err: map[string]error{
				"linux":   errors.New("no such file or directory"),
				"darwin":  errors.New("no such file or directory"),
				"windows": errors.New("The system cannot find the file specified"),
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			var path string
			if tc.pTestData != "" {
				tmp := t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "validate_compose_tests", tc.pTestData), tmp)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				path = filepath.Join(tmp, ".env")
			}

			got, err := ParseEnv(path)

			if tc.err != nil {
				assert.ErrorContains(t, err, tc.err[runtime.GOOS].Error())
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, got, tc.out)
			}
		})
	}
}
