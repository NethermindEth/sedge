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
	// Remove \r from keys and values (for Windows line endings)
	for key, value := range env {
		newKey := strings.TrimRight(key, "\r")
		newValue := strings.TrimRight(value, "\r")
		if newKey != key {
			delete(env, key)
			env[newKey] = newValue
		} else {
			env[key] = newValue
		}
	}
	return env, nil
}
