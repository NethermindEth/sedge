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
package testdata

import (
	"embed"
	"io"
	"io/fs"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

//go:embed all:*
var TestData embed.FS

//go:embed empty
var Empty embed.FS

func SetupProfileFS(t *testing.T, instanceName string, afs afero.Fs) string {
	t.Helper()
	instanceFs, err := fs.Sub(TestData, instanceName)
	if err != nil {
		t.Fatalf("failed to setup instance filesystem: %v", err)
	}

	tempPath := t.TempDir()

	err = fs.WalkDir(instanceFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		targetPath := tempPath + "/" + path
		if d.IsDir() {
			if err := afs.MkdirAll(targetPath, 0o755); err != nil {
				return err
			}
		} else {
			file, err := instanceFs.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			targetFile, err := afs.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()
			if _, err := io.Copy(targetFile, file); err != nil {
				return err
			}
		}
		return nil
	})
	require.NoError(t, err)

	return tempPath
}

func GetEnv(t *testing.T, envName string) io.ReadCloser {
	t.Helper()
	file, err := TestData.Open("env/" + envName)
	require.NoError(t, err, "failed to open env file %s", envName)
	return file
}
