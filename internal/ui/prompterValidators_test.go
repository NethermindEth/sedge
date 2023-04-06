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
package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/stretchr/testify/assert"
)

func TestEthAddressValidator(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		allowEmpty bool
		want       error
	}{
		{
			name:  "valid address",
			input: "0x4675c7e5baafbffbca748158becba61ef3b0a263",
			want:  nil,
		},
		{
			name:  "invalid address, too long",
			input: "0x4675c7e5baafbffbca748158becba61ef3b0a2630x4675c7e5baafbffbca748158becba61ef3b0a263",
			want:  ErrInvalidEthereumAddress,
		},
		{
			name:  "invalid address, too short",
			input: "0x4675c7e5baafbffbca748158becba61",
			want:  ErrInvalidEthereumAddress,
		},
		{
			name:  "non-address string",
			input: "bananas",
			want:  ErrInvalidEthereumAddress,
		},
		{
			name:       "empty string, allowed",
			input:      "",
			want:       nil,
			allowEmpty: true,
		},
		{
			name:  "empty string, disallowed",
			input: "",
			want:  ErrInvalidEthereumAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EthAddressValidator(tt.input, tt.allowEmpty)
			if tt.want != nil {
				assert.ErrorIs(t, tt.want, got)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestFilePathValidator(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFilePath := filepath.Join(tmpDir, "test.txt")
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpFile.Close()

	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "valid file path",
			input:   tmpFilePath,
			wantErr: false,
		},
		{
			name:    "invalid file path, not a file",
			input:   filepath.Join(tmpDir, "test_2.txt"),
			wantErr: true,
		},
		{
			name:    "invalid file path, is a directory",
			input:   tmpDir,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilePathValidator(tt.input)
			if tt.wantErr {
				assert.Error(t, got)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestInt64Validator(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  error
	}{
		{
			name:  "valid int64",
			input: "123456789",
			want:  nil,
		},
		{
			name:  "invalid int64, invalid string",
			input: "14084s",
			want:  ErrInvalidInt64String,
		},
		{
			name:  "invalid int64, too long",
			input: "123456789123456789",
			want:  ErrInvalidInt64String,
		},
		{
			name:  "invalid int64, float",
			input: "123456789.123456789",
			want:  ErrInvalidInt64String,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64Validator(tt.input)
			if tt.want != nil {
				assert.ErrorIs(t, tt.want, got)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestDirPathValidator(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid directory path",
			input:   tmpDir,
			wantErr: false,
		},
		{
			name:    "invalid directory path, not a directory",
			input:   filepath.Join(tmpDir, "test.txt"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DirPathValidator(tt.input)
			if tt.wantErr {
				assert.Error(t, got)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestURLValidator(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "valid URL",
			input: "https://www.google.com",
			want:  nil,
		},
		{
			name:  "valid URL, http",
			input: "http://www.google.com",
			want:  nil,
		},
		{
			name:  "invalid URL protocol",
			input: "https//www.google",
			want:  ErrInvalidURL,
		},
		{
			name:  "invalid URL, no protocol",
			input: "www.google.com",
			want:  ErrInvalidURL,
		},
		{
			name:  "invalid URL, no domain",
			input: "https://",
			want:  ErrInvalidURL,
		},
		{
			name:  "invalid URL, no domain, no protocol",
			input: "www.",
			want:  ErrInvalidURL,
		},
		{
			name:  "empty url",
			input: "",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := URLValidator(tt.input)
			if tt.want != nil {
				assert.ErrorContains(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestFileExtensionValidator(t *testing.T) {
	tests := []struct {
		name         string
		input        []string
		handlerInput string
		want         error
	}{
		{
			name:         "valid file extension",
			input:        []string{".txt"},
			handlerInput: "test.txt",
			want:         nil,
		},
		{
			name:         "valid file extensions, test.txt",
			input:        []string{".txt", ".csv", ""},
			handlerInput: "test.txt",
			want:         nil,
		},
		{
			name:         "valid file extensions, test",
			input:        []string{".txt", ".csv", ""},
			handlerInput: "test",
			want:         nil,
		},
		{
			name:         "invalid file extension, no extension",
			input:        []string{".txt", ".csv"},
			handlerInput: "test.json",
			want:         ErrInvalidFileExtension,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fileExtensionValidator(tt.input)(tt.handlerInput)
			if tt.want != nil {
				assert.ErrorContains(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
func TestDigitsStringValidator(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "valid digits string",
			input: "123456789",
			want:  nil,
		},
		{
			name:  "invalid digits string, invalid string",
			input: "14084s",
			want:  ErrInvalidDigitString,
		},
		{
			name:  "invalid digits string, too long",
			input: "123456789123456789123456789123456789",
			want:  nil,
		},
		{
			name:  "invalid digits string, float",
			input: "123456789.123456789",
			want:  ErrInvalidDigitString,
		},
		{
			name:  "invalid digits string, negative",
			input: "-123456789",
			want:  ErrInvalidDigitString,
		},
		{
			name:  "invalid digits string, empty",
			input: "",
			want:  ErrInvalidDigitString,
		},
		{
			name:  "valid digits string, zero",
			input: "0",
			want:  nil,
		},
		{
			name:  "invalid digits string, starting with zero",
			input: "0123456789",
			want:  ErrInvalidDigitString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DigitsStringValidator(tt.input)
			if tt.want != nil {
				assert.ErrorContains(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestGraffitiValidator(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{
			name:  "valid graffiti",
			input: "test",
		},
		{
			name:  "invalid graffiti, too long",
			input: "GraffitiTooLongSupercalifragilisticexpialidocious",
			want:  fmt.Errorf(configs.ErrGraffitiLength, "GraffitiTooLongSupercalifragilisticexpialidocious", 49),
		},
		{
			name:  "valid graffiti, empty",
			input: "",
		},
		{
			name:  "valid graffiti, exact 16 characters",
			input: "GrafitiExactly16",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GraffitiValidator(tt.input)
			if tt.want != nil {
				assert.ErrorContains(t, got, tt.want.Error())
			} else {
				assert.NoError(t, got)
			}
		})
	}
}
