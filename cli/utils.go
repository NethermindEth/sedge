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
