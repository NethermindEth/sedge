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
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/NethermindEth/sedge/configs"
)

/*
HandleUrlOrPath :
Get a source string and call related handler if the source

params :-
a. string
Source that will be tried to convert to url or path
b. func(url string) error
Handler function to call if source is an url
c. func(path string) error
Handler function to call if source is a filepath

returns :-
a. error occurred during any of the handlers if any or if provided source is neither file or url.
*/
func HandleUrlOrPath(
	src string,
	handleUrl func(url string) error,
	handlePath func(path string) error,
) error {
	uri, err := url.ParseRequestURI(src)
	if err != nil {
		return fmt.Errorf(configs.InvalidFilePathOrUrl, src)
	}

	if uri.Scheme == "http" || uri.Scheme == "https" {
		return handleUrl(src)
	} else if _, err := os.Stat(src); err == nil {
		return handlePath(src)
	}

	return fmt.Errorf(configs.InvalidFilePathOrUrl, src)
}

/*
CheckUrlOrPath :
Verifies if provided source string is an url or filepath returning relevant error if neither

params :-
a. string
Source that will be tried to convert to url or path

returns :-
a. error if provided source is neither file or url.
*/
func CheckUrlOrPath(src string) error {
	return HandleUrlOrPath(
		src,
		func(url string) error { return nil },
		func(path string) error { return nil },
	)
}

/*
GetUrlOrPathContent :
Get content of the source url or filepath provided

params :-
a. string
Source that will be used as url or filepath

returns :-
a. error if any
*/
func GetUrlOrPathContent(src string) (string, error) {
	urlContent := ""
	fileContent := ""

	err := HandleUrlOrPath(
		src,
		func(url string) error {
			resp, err := GetRequest(url, time.Minute)
			if err != nil || resp.StatusCode != 200 {
				return fmt.Errorf(configs.CannotGetUrlContent, src, err)
			}
			defer resp.Body.Close()

			rawContent, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf(configs.CannotGetUrlContent, src, err)
			}

			urlContent = string(rawContent)
			return nil
		},
		func(path string) error {
			rawContent, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf(configs.CannotReadFileContent, src, err)
			}

			fileContent = string(rawContent)
			return nil
		},
	)

	return strings.ReplaceAll(strings.TrimSpace(urlContent+fileContent), "\r", ""), err
}

/*
DownloadOrCopy :
Get content of the source url or filepath provided

params :-
a. string
Source that will be used as url or filepath
b. string
Destination filepath to downloaded or copied file
c. bool
Delete existing file in the destination filepath

returns :-
a. error if any
*/
func DownloadOrCopy(src, dest string, autoremove bool) error {
	_, err := os.Stat(dest)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf(configs.ErrorCheckingFile, dest, err)
	} else if err == nil && !autoremove {
		return fmt.Errorf(configs.DestFileAlreadyExist, src)
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf(configs.ErrorCreatingFile, dest, err)
	}
	defer destFile.Close()

	return HandleUrlOrPath(
		src,
		func(url string) error {
			resp, err := GetRequest(url, time.Minute)
			if err != nil || resp.StatusCode != 200 {
				return fmt.Errorf(configs.ErrorDownloadingFile, url, err)
			}
			defer resp.Body.Close()

			_, err = io.Copy(destFile, resp.Body)
			if err != nil {
				return fmt.Errorf(configs.ErrorDownloadingFile, url, err)
			}

			return nil
		},
		func(path string) error {
			srcFile, err := os.Open(path)
			if err != nil {
				return fmt.Errorf(configs.ErrorCopyingFile, path, err)
			}
			defer srcFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return fmt.Errorf(configs.ErrorCopyingFile, path, err)
			}

			return nil
		},
	)
}
