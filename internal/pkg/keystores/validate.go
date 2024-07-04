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
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed keystore_schema_eip2335.json
var keystoreJSONSchema []byte

//go:embed keystore_schema_deposit_data.json
var depositDataJSONSchema []byte

var (
	ErrDepositDataNotFound              = errors.New("deposit_data.json not found")
	ErrKeystorePasswordNotFound         = errors.New("keystore_password.txt not found")
	ErrValidatorKeysDirNotFound         = errors.New("validator_keys directory not found")
	ErrValidatorKeysDirWithoutKeystores = errors.New("validator_keys directory does not contain any keystores")
	ErrInvalidKeystoreFile              = errors.New("invalid keystore file")
	ErrInvalidKeystoreFileName          = errors.New("invalid keystore file name")
	ErrInvalidKeystoreFileSchema        = errors.New("file does not match keystore schema (EIP-2335)")
	ErrInvalidDepositDataFileSchema     = errors.New("file does not match deposit data schema")
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
	pattern := `^keystore-m_12381_3600_\d+_0_0\.json$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(filepath.Base(path)) {
		return ErrInvalidKeystoreFileName
	}

	return validateFileWithSchema(path, "keystore_schema", keystoreJSONSchema, ErrInvalidKeystoreFileSchema)
}

func validateDepositDataFile(path string) error {
	depositDataFile, err := os.Stat(filepath.Join(path, DepositDataFileName))
	if err != nil || depositDataFile.IsDir() {
		return ErrDepositDataNotFound
	}

	return nil
	return validateFileWithSchema(path, "deposit_data_schema", depositDataJSONSchema, ErrInvalidDepositDataFileSchema)
}

func validateFileWithSchema(path, schemaName string, schema []byte, schemaErr error) error {
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(schemaName, bytes.NewReader(schema)); err != nil {
		return err
	}
	schemaCmp := compiler.MustCompile(schemaName)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if err := schemaCmp.Validate(value); err != nil {
		return schemaErr
	}
	return nil
}
