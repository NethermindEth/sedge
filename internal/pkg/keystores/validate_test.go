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
package keystores

import (
	"embed"
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata
var testdata embed.FS

func TestValidateKeystoreDir(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T) (string, error)
		expected []error
	}{
		{
			name:  "empty dir",
			setup: setupEmptyDir,
			expected: []error{
				ErrDepositDataNotFound,
				ErrKeystorePasswordNotFound,
				ErrValidatorKeysDirNotFound,
			},
		},
		{
			name:     "valid keystore dir",
			setup:    setupTestDataDir,
			expected: []error{},
		},
		{
			name:  "invalid keystore schema",
			setup: setupTestDataDirWithInvalidKeyFile,
			expected: []error{
				ErrInvalidKeystoreFileSchema,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := tt.setup(t)
			if err != nil {
				t.Fatal(err)
			}
			got := ValidateKeystoreDir(dir)
			assert.Equal(t, len(tt.expected), len(got))
			for _, err := range got {
				var found bool
				for _, e := range tt.expected {
					if errors.Is(err, e) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("error %v not found in set %v", err, tt.expected)
				}
			}
		})
	}
}

func setupEmptyDir(t *testing.T) (string, error) {
	t.Helper()
	return t.TempDir(), nil
}

func setupTestDataDir(t *testing.T) (string, error) {
	t.Helper()
	dir := t.TempDir()
	depositDataContent, err := testdata.ReadFile(path.Join("testdata/valid", depositDataFileName))
	if err != nil {
		return "", err
	}
	if f, err := os.Create(filepath.Join(dir, depositDataFileName)); err != nil {
		return "", err
	} else {
		if _, err := f.Write(depositDataContent); err != nil {
			return "", err
		}
		f.Close()
	}
	keystorePasswordContent, err := testdata.ReadFile(path.Join("testdata/valid", keystorePasswordFileName))
	if err != nil {
		return "", err
	}
	if f, err := os.Create(filepath.Join(dir, keystorePasswordFileName)); err != nil {
		return "", err
	} else {
		if _, err := f.Write(keystorePasswordContent); err != nil {
			return "", err
		}
		f.Close()
	}
	validatorKeysDir, err := testdata.ReadDir(path.Join("testdata/valid", validatorKeysDirName))
	if err != nil {
		return "", err
	}
	os.MkdirAll(filepath.Join(dir, validatorKeysDirName), 0o755)
	for _, source := range validatorKeysDir {
		if dest, err := os.Create(filepath.Join(dir, validatorKeysDirName, source.Name())); err != nil {
			return "", err
		} else {
			sourceContent, err := testdata.ReadFile(path.Join("testdata/valid", validatorKeysDirName, source.Name()))
			if err != nil {
				return "", err
			}
			if _, err := dest.Write(sourceContent); err != nil {
				return "", err
			}
			dest.Close()
		}
	}
	return dir, nil
}

func setupTestDataDirWithInvalidKeyFile(t *testing.T) (string, error) {
	t.Helper()
	dir, err := setupTestDataDir(t)
	if err != nil {
		return "", err
	}
	invalidKeyFileContent, err := testdata.ReadFile(path.Join("testdata/keystore_invalid.json"))
	if err != nil {
		return "", err
	}
	if dest, err := os.Create(filepath.Join(dir, validatorKeysDirName, "keystore-m_12381_3600_5_0_0.json")); err != nil {
		return "", err
	} else {
		if _, err := dest.Write(invalidKeyFileContent); err != nil {
			return "", err
		}
		defer dest.Close()
	}
	return dir, nil
}
