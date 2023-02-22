package ui

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEthAddressValidator(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  error
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EthAddressValidator(tt.input)
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
		withErr bool
	}{
		{
			name:    "valid file path",
			input:   tmpFilePath,
			withErr: false,
		},
		{
			name:    "invalid file path, not a file",
			input:   filepath.Join(tmpDir, "test_2.txt"),
			withErr: true,
		},
		{
			name:    "invalid file path, is a directory",
			input:   tmpDir,
			withErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilePathValidator(tt.input)
			if tt.withErr {
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
