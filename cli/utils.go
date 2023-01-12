package cli

import (
	"os"
	"strings"
)

func readFileContent(path string) (string, error) {
	raw, err := os.ReadFile(path)
	content := strings.TrimSpace(strings.TrimSuffix(string(raw), "\n"))
	return content, err
}

func fileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), nil
}
