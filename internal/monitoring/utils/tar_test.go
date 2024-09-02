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
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompressToTarGz(t *testing.T) {
	testDir := t.TempDir()
	pkgDir := filepath.Join(testDir, "mock-avs")
	outTarPath := filepath.Join(testDir, "out.tar.gz")
	outTarContentDir := filepath.Join(testDir, "out")

	err := os.MkdirAll(pkgDir, 0o755)
	require.NoError(t, err, "failed to create mock-avs dir")
	err = exec.Command("git", "clone", "--single-branch", "-b", common.MockAvsPkg.Version(), common.MockAvsPkg.Repo(), pkgDir).Run()
	require.NoError(t, err, "failed to clone mock-avs repo")

	outTar, err := os.OpenFile(outTarPath, os.O_CREATE|os.O_RDWR, 0o755)
	require.NoError(t, err)

	err = CompressToTarGz(pkgDir, outTar)
	require.NoError(t, err)

	err = os.MkdirAll(outTarContentDir, 0o755)
	require.NoError(t, err, "failed to create out dir")
	err = exec.Command("tar", "-xf", outTarPath, "-C", outTarContentDir).Run()
	require.NoError(t, err, "failed to create mock-avs.tar.gz")

	assertEqualDirs(t, pkgDir, outTarContentDir)
}

func TestDecompressTarGz(t *testing.T) {
	testDir := t.TempDir()
	pkgDir := filepath.Join(testDir, "mock-avs")
	tarPath := filepath.Join(testDir, "mock-avs.tar.gz")
	outDir := filepath.Join(testDir, "out")

	err := os.MkdirAll(pkgDir, 0o755)
	require.NoError(t, err, "failed to create mock-avs dir")
	err = exec.Command("git", "clone", "--single-branch", "-b", common.MockAvsPkg.Version(), common.MockAvsPkg.Repo(), pkgDir).Run()
	require.NoError(t, err, "failed to clone mock-avs repo")

	err = exec.Command("tar", "-czf", tarPath, "-C", pkgDir, ".").Run()
	require.NoError(t, err, "failed to create mock-avs.tar.gz")

	tarFile, err := os.Open(tarPath)
	require.NoError(t, err, "failed to open mock-avs.tar.gz")

	err = DecompressTarGz(tarFile, outDir)
	require.NoError(t, err, "failed to decompress mock-avs.tar.gz")

	assertEqualDirs(t, pkgDir, outDir)
}

func assertEqualDirs(t *testing.T, dir1, dir2 string) {
	err := filepath.Walk(dir1, func(path1 string, info1 os.FileInfo, err1 error) error {
		if err1 != nil {
			return err1
		}

		path2 := filepath.Join(dir2, path1[len(dir1):])

		if info1.IsDir() {
			assert.DirExists(t, path2)
		} else {
			assert.FileExists(t, path2)
			assertEqualFiles(t, path1, path2)
		}
		return nil
	})
	require.NoError(t, err, "failed to walk dir %s", dir1)
}

func assertEqualFiles(t *testing.T, f1, f2 string) {
	file1, err := os.ReadFile(f1)
	require.NoError(t, err, "failed to read file %s", f1)
	file2, err := os.ReadFile(f2)
	require.NoError(t, err, "failed to read file %s", f2)
	assert.Equal(t, file1, file2)
}
