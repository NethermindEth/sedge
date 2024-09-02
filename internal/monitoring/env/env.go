package env

import (
	"strings"

	"github.com/spf13/afero"
)

func LoadEnv(fs afero.Fs, path string) (map[string]string, error) {
	env := make(map[string]string)
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		env[strings.Trim(parts[0], " ")] = strings.Trim(parts[1], " ")
	}
	return env, nil
}
