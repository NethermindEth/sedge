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
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed keystore_schema_eip2335.json
var keystoreJSONSchema []byte

var (
	ErrDepositDataNotFound              = errors.New("deposit_data.json not found")
	ErrKeystorePasswordNotFound         = errors.New("keystore_password.txt not found")
	ErrValidatorKeysDirNotFound         = errors.New("validator_keys directory not found")
	ErrValidatorKeysDirWithoutKeystores = errors.New("validator_keys directory does not contain any keystores")
	ErrInvalidKeystoreFile              = errors.New("invalid keystore file")
	ErrInvalidKeystoreFileSchema        = errors.New("file does not match keystore schema (EIP-2335)")
)

const (
	DepositDataFileName      = "deposit_data.json"
	KeystorePasswordFileName = "keystore_password.txt"
	ValidatorKeysDirName     = "validator_keys"
)

func ValidateKeystoreDir(dir string) (errors []error) {
	if err := validateDepositDataFile(dir); err != nil {
		errors = append(errors, err)
	}
	if err := validateKeystorePasswordFile(dir); err != nil {
		errors = append(errors, err)
	}
	if err := validateValidatorKeysFolder(dir); err != nil {
		errors = append(errors, err)
	}
	return
}

func validateDepositDataFile(keystoreDirPath string) error {
	depositDataFile, err := os.Stat(filepath.Join(keystoreDirPath, DepositDataFileName))
	if err != nil || depositDataFile.IsDir() {
		return ErrDepositDataNotFound
	}
	// TODO check deposit_data.json is valid json
	return nil
}

func validateKeystorePasswordFile(keystoreDirPath string) error {
	keystorePasswordFileName, err := os.Stat(filepath.Join(keystoreDirPath, KeystorePasswordFileName))
	if err != nil || keystorePasswordFileName.IsDir() {
		return ErrKeystorePasswordNotFound
	}
	return nil
}

func validateValidatorKeysFolder(keystoreDirPath string) error {
	validatorKeysDirPath := filepath.Join(keystoreDirPath, ValidatorKeysDirName)
	validatorKeysFolder, err := os.Stat(validatorKeysDirPath)
	if err != nil || !validatorKeysFolder.IsDir() {
		return ErrValidatorKeysDirNotFound
	}
	var jsonFiles []string
	filepath.WalkDir(validatorKeysDirPath, func(path string, d os.DirEntry, err error) error {
		if filepath.Ext(path) == ".json" {
			jsonFiles = append(jsonFiles, path)
		}
		return nil
	})
	if len(jsonFiles) == 0 {
		return ErrValidatorKeysDirWithoutKeystores
	}
	for _, jsonFile := range jsonFiles {
		if err := validateKeystoreFile(jsonFile); err != nil {
			return fmt.Errorf("%w: %s", err, jsonFile)
		}
	}
	return nil
}

func validateKeystoreFile(path string) error {
	// TODO validate file name
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("keystore_schema", bytes.NewReader(keystoreJSONSchema)); err != nil {
		return err
	}
	schema := compiler.MustCompile("keystore_schema")

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if err := schema.Validate(value); err != nil {
		return ErrInvalidKeystoreFileSchema
	}
	return nil
}
