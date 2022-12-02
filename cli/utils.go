package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readFileContent(path string) (string, error) {
	raw, err := os.ReadFile(path)
	content := strings.TrimSpace(strings.TrimSuffix(string(raw), "\n"))
	return content, err
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
