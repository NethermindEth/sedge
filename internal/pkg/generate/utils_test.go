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
package generate

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestParseCompose(t *testing.T) {
	tcs := []struct {
		name      string
		pTestData string
		err       error
	}{
		{
			name:      "Valid compose file",
			pTestData: "valid",
		},
		{
			name:      "Invalid yml file",
			pTestData: "invalid",
			err:       errors.New("error parsing yml file, it seems is not a valid docker-compose script:"),
		},
		{
			name:      "Without compose file",
			pTestData: "no_compose",
			err:       errors.New("no such file or directory"),
		},
		{
			name:      "Empty compose file",
			pTestData: "empty",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			var path string
			if tc.pTestData != "" {
				tmp := t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "utils_tests", tc.pTestData), tmp)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				path = filepath.Join(tmp, "docker-compose.yml")
			}

			cd, err := ParseCompose(path)

			if tc.err != nil {
				assert.ErrorContains(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				// Marshal to yaml and compare with existing file
				out, err := yaml.Marshal(cd)
				if err != nil {
					t.Fatalf("Error marshaling compose: %v", err)
				}
				// Read existing file
				content, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("Error reading compose: %v", err)
				}
				// if empty skip
				if len(content) == 0 {
					t.Skip("Empty compose file")
				}
				// Unmarshal existing file
				var temp interface{}
				err = yaml.Unmarshal(content, &temp)
				if err != nil {
					t.Fatalf("Error unmarshaling compose: %v", err)
				}
				// Marshal to yaml to get the same bytes
				_, err = yaml.Marshal(temp)
				if err != nil {
					t.Fatalf("Error marshaling compose: %v", err)
				}
				assert.YAMLEq(t, string(content), string(out))
			}
		})
	}
}
