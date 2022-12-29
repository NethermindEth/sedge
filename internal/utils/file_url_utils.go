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

func CheckUrlOrPath(src string) error {
	return HandleUrlOrPath(
		src,
		func(url string) error { return nil },
		func(path string) error { return nil },
	)
}

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
