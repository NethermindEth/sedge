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
package package_handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/monitoring/package_handler/testdata"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func FuzzHashFile(f *testing.F) {
	for i := 0; i < 10; i++ {
		f.Add([]byte(fmt.Sprintf("file content %d\n", i)))
	}

	filePath := filepath.Join(f.TempDir(), "file.txt")
	file, err := os.Create(filePath)
	if err != nil {
		f.Fatalf("failed to create temp file: %v", err)
	}
	defer file.Close()

	f.Fuzz(func(t *testing.T, fileContent []byte) {
		if _, err := file.Write(fileContent); err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}
		sha256sum := exec.Command("sha256sum", filePath)
		output, err := sha256sum.Output()
		if err != nil {
			t.Fatalf("failed to run sha256sum: %v", err)
		}
		fileHash, err := hashFile(filePath, afero.NewOsFs())
		assert.NoError(t, err)
		assert.Equal(t, strings.Split(string(output), " ")[0], fileHash)
	})
}

func TestCheckPackageFileExist(t *testing.T) {
	afs := afero.NewMemMapFs()
	testDir, err := afero.TempDir(afs, "", "test")
	require.NoError(t, err)
	testdata.SetupDir(t, "mock-avs", testDir, afs)

	ts := []struct {
		name     string
		filePath string
		err      error
	}{
		{
			name:     "file exists",
			filePath: "pkg/manifest.yml",
			err:      nil,
		},
		{
			name:     "file does not exist",
			filePath: "pkg/manifest2.yml",
			err: PackageFileNotFoundError{
				fileRelativePath: "pkg/manifest2.yml",
				packagePath:      filepath.Join(testDir, "mock-avs"),
			},
		},
		{
			name:     "is not a file",
			filePath: "pkg",
			err:      ErrInvalidFilePath,
		},
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			err := checkPackageFileExist(filepath.Join(testDir, "mock-avs"), tc.filePath, afs)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCheckPackageDirExist(t *testing.T) {
	afs := afero.NewMemMapFs()
	testDir, err := afero.TempDir(afs, "", "test")
	require.NoError(t, err)
	testdata.SetupDir(t, "mock-avs", testDir, afs)

	ts := []struct {
		name    string
		dirPath string
		err     error
	}{
		{
			name:    "dir exists",
			dirPath: "pkg",
			err:     nil,
		},
		{
			name:    "does not exist",
			dirPath: "pkg2",
			err: PackageDirNotFoundError{
				dirRelativePath: "pkg2",
				packagePath:     filepath.Join(testDir, "mock-avs"),
			},
		},
		{
			name:    "is not a directory",
			dirPath: "pkg/manifest.yml",
			err:     ErrInvalidDirPath,
		},
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			err := checkPackageDirExist(filepath.Join(testDir, "mock-avs"), tc.dirPath, afs)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestParseChecksumFile(t *testing.T) {
	testFile := t.TempDir()
	ts := []struct {
		name    string
		content string
		out     map[string]string
		err     error
	}{
		{
			name:    "valid checksum file",
			content: "9e9d08613004818012fb1b72b427581d8e00c4e09f13e8899c00e8b6228608ed  pkg/manifest.yml\n6f9cf01b1996cdb179ac7a0776ddf907871197afe10d19b9d10cbb5faa141c56  pkg/sepolia/.env\n",
			out: map[string]string{
				"pkg/manifest.yml": "9e9d08613004818012fb1b72b427581d8e00c4e09f13e8899c00e8b6228608ed",
				"pkg/sepolia/.env": "6f9cf01b1996cdb179ac7a0776ddf907871197afe10d19b9d10cbb5faa141c56",
			},
			err: nil,
		},
		{
			name:    "invalid checksum file, invalid separator in line",
			content: "9e9d08613004818012fb1b72b427581d8e00c4e09f13e8899c00e8b6228608ed pkg/manifest.yml\n6f9cf01b1996cdb179ac7a0776ddf907871197afe10d19b9d10cbb 5faa141c56 pkg/sepolia/.env\n",
			out:     nil,
			err:     fmt.Errorf("invalid checksum file format"),
		},
		{
			name:    "invalid checksum file, invalid separator in line",
			content: "9e9d08613004818012fb1b72b427581d8e00c4e09f13e8899c00e8b6228608ed pkg/manifest.yml\n6f9cf01b1996cdb179ac7a0776ddf907871197afe10d19b9d10cbb\n",
			out:     nil,
			err:     fmt.Errorf("invalid checksum file format"),
		},
	}

	for i, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			filePath := filepath.Join(testFile, fmt.Sprintf("checksum_%d.txt", i))
			file, err := os.Create(filePath)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer file.Close()
			if _, err := file.Write([]byte(tc.content)); err != nil {
				t.Fatal("failed to write to temp file: " + err.Error())
			}
			out, err := parseChecksumFile(filePath, afero.NewOsFs())
			if tc.err == nil {
				assert.NoError(t, err)
				assert.Len(t, out, len(tc.out))
				for k, v := range tc.out {
					assert.Equal(t, v, out[k])
				}
			} else {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}
