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
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/test"
)

func TestHandleUrlOrPath(t *testing.T) {
	urlErr := errors.New("url error")
	pathErr := errors.New("path error")

	tmpDir := t.TempDir()

	tcs := []struct {
		name string
		src  string
		want error
	}{
		{
			name: "valid url",
			src:  "https://www.google.com",
			want: urlErr,
		},
		{
			name: "valid path",
			src:  tmpDir,
			want: pathErr,
		},
		{
			name: "invalid url",
			src:  "::/www.google.com/invalid",
			want: fmt.Errorf(configs.InvalidFilePathOrUrl, "::/www.google.com/invalid"),
		},
		{
			name: "invalid url, not http",
			src:  "ftp://www.google.com/invalid",
			want: fmt.Errorf(configs.InvalidFilePathOrUrl, "ftp://www.google.com/invalid"),
		},
		{
			name: "invalid path, non existing",
			src:  filepath.Join(tmpDir, "invalid"),
			want: fmt.Errorf(configs.InvalidFilePathOrUrl, filepath.Join(tmpDir, "invalid")),
		},
		{
			name: "invalid path, invalid char",
			src:  "./tmp\\invalid",
			want: fmt.Errorf(configs.InvalidFilePathOrUrl, "./tmp\\invalid"),
		},
	}
	handleUrl := func(url string) error {
		return urlErr
	}
	handlePath := func(path string) error {
		return pathErr
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleUrlOrPath(tc.src, handleUrl, handlePath)
			if err.Error() != tc.want.Error() {
				t.Errorf("expected error %v, got %v", tc.want, err)
			}
		})
	}
}

func TestCheckUrlOrPath(t *testing.T) {
	tmpDir := t.TempDir()

	tcs := []struct {
		name  string
		src   string
		isErr bool
	}{
		{
			name:  "valid url",
			src:   "https://www.google.com",
			isErr: false,
		},
		{
			name:  "valid path",
			src:   tmpDir,
			isErr: false,
		},
		{
			name:  "invalid url",
			src:   "::/www.google.com/invalid",
			isErr: true,
		},
		{
			name:  "invalid url, not http",
			src:   "ftp://www.google.com/invalid",
			isErr: true,
		},
		{
			name:  "invalid path, non existing",
			src:   filepath.Join(tmpDir, "invalid"),
			isErr: true,
		},
		{
			name:  "invalid path, invalid char",
			src:   "./tmp\\invalid",
			isErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckUrlOrPath(tc.src)
			if err = CheckErr("expected error %v, got %v", tc.isErr, err); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetUrlOrPathContent(t *testing.T) {
	readConfig := func(path string) string {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}
		return string(raw)
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	emptyPath := filepath.Join(pwd, "testdata", "remote_files", "empty_file", "config.yml")
	goodPath := filepath.Join(pwd, "testdata", "remote_files", "good_file", "config.yml")
	empty := readConfig(emptyPath)
	good := readConfig(goodPath)

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.Method != "GET" {
				t.Errorf("Unexpected HTTP method, expected GET, got %s", req.Method)
			}

			query := req.URL.Query().Get("test")
			switch query {
			case "ERROR":
				rw.WriteHeader(http.StatusBadRequest)
			case "OK":
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(good))
			case "EMPTY":
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(empty))
			default:
				rw.WriteHeader(http.StatusNotFound)
			}
		}))
	defer server.Close()

	tcs := []struct {
		name  string
		src   string
		want  string
		isErr bool
	}{
		{
			name:  "valid url, good config",
			src:   server.URL + "?test=OK",
			want:  good,
			isErr: false,
		},
		// {
		// 	name:  "valid url, empty config",
		// 	src:   server.URL + "?test=EMPTY",
		// 	want:  empty,
		// 	isErr: false,
		// },
		// {
		// 	name:  "valid url, error",
		// 	src:   server.URL + "?test=ERROR",
		// 	want:  "",
		// 	isErr: true,
		// },
		// {
		// 	name:  "valid url, error",
		// 	src:   server.URL + "?test=ERROR",
		// 	want:  "",
		// 	isErr: true,
		// },
		// {
		// 	name:  "valid path, good config",
		// 	src:   goodPath,
		// 	want:  good,
		// 	isErr: false,
		// },
		// {
		// 	name:  "valid path, empty config",
		// 	src:   emptyPath,
		// 	want:  empty,
		// 	isErr: false,
		// },
		// {
		// 	name:  "invalid url",
		// 	src:   "::/www.google.com/invalid",
		// 	want:  "",
		// 	isErr: true,
		// },
		// {
		// 	name:  "invalid path, non existing",
		// 	src:   filepath.Join(t.TempDir(), "invalid"),
		// 	want:  "",
		// 	isErr: true,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUrlOrPathContent(tc.src)
			if err = CheckErr("expected error %v, got %v", tc.isErr, err); err != nil {
				t.Error(err)
			}

			if got != tc.want {
				t.Errorf("expected content %v, got %v", tc.want, got)
			}
		})
	}
}

func TestDownloadOrCopy(t *testing.T) {
	readConfig := func(path string) string {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}
		return string(raw)
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	emptyPath := filepath.Join(pwd, "testdata", "remote_files", "empty_file")
	goodPath := filepath.Join(pwd, "testdata", "remote_files", "good_file")
	empty := readConfig(filepath.Join(emptyPath, "config.yml"))
	good := readConfig(filepath.Join(goodPath, "config.yml"))

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.Method != "GET" {
				t.Errorf("Unexpected HTTP method, expected GET, got %s", req.Method)
			}

			query := req.URL.Query().Get("test")
			switch query {
			case "ERROR":
				rw.WriteHeader(http.StatusBadRequest)
			case "OK":
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(good))
			case "EMPTY":
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(empty))
			default:
				rw.WriteHeader(http.StatusNotFound)
			}
		}))
	defer server.Close()

	tcs := []struct {
		name    string
		src     string
		testdir string
		want    string
		isErr   bool
	}{
		{
			name:  "valid url, good config",
			src:   server.URL + "?test=OK",
			want:  good,
			isErr: false,
		},
		{
			name:  "valid url, empty config",
			src:   server.URL + "?test=EMPTY",
			want:  empty,
			isErr: false,
		},
		{
			name:  "valid url, error",
			src:   server.URL + "?test=ERROR",
			want:  "",
			isErr: true,
		},
		{
			name:  "valid url, error",
			src:   server.URL + "?test=ERROR",
			want:  "",
			isErr: true,
		},
		{
			name:    "valid path, good config",
			src:     filepath.Join(goodPath, "config.yml"),
			testdir: goodPath,
			want:    good,
			isErr:   false,
		},
		{
			name:    "valid path, empty config",
			src:     filepath.Join(emptyPath, "config.yml"),
			testdir: emptyPath,
			want:    empty,
			isErr:   false,
		},
		{
			name:  "invalid url",
			src:   "::/www.google.com/invalid",
			want:  "",
			isErr: true,
		},
		{
			name:  "invalid path, non existing",
			src:   filepath.Join(t.TempDir(), "invalid"),
			want:  "",
			isErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// prepare temp dir
			dcPath := t.TempDir()
			err := HandleUrlOrPath(tc.src,
				func(url string) error {
					if strings.Contains(url, "EMPTY") {
						return test.PrepareTestCaseDir(emptyPath, dcPath)
					} else if strings.Contains(url, "OK") {
						return test.PrepareTestCaseDir(goodPath, dcPath)
					}
					return nil
				},
				func(path string) error {
					return test.PrepareTestCaseDir(tc.testdir, dcPath)
				},
			)
			if err != nil {
				t.Log(err)
			}

			// autoremove false, must fail
			err = DownloadOrCopy(tc.src, filepath.Join(dcPath, "config.yml"), false)
			if err = CheckErr("expected error %v, got %v", true, err); err != nil {
				t.Error(err)
			}

			err = DownloadOrCopy(tc.src, filepath.Join(dcPath, "config.yml"), true)
			if err = CheckErr("expected error %v, got %v", tc.isErr, err); err != nil {
				t.Error(err)
			}

			got := readConfig(filepath.Join(dcPath, "config.yml"))
			if got != tc.want {
				t.Errorf("expected content %v, got %v", tc.want, got)
			}
		})
	}
}
