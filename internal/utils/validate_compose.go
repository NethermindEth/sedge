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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	cmpcli "github.com/compose-spec/compose-go/cli"
	log "github.com/sirupsen/logrus"
)

// ParseEnv parse a .env file and return a map of the environment variables
func ParseEnv(path string) (map[string]string, error) {
	envFile, err := os.Open(path)
	if err != nil {
		// return nil, fmt.Errorf(configs.OpenEnvErr, err)
		return nil, err
	}
	defer envFile.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		envMap[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		// return nil, fmt.Errorf(configs.ReadEnvErr, err)
		return nil, err
	}

	return envMap, nil
}

// ValidateCompose validate a docker-compose script according to the docker-compose specification
func ValidateCompose(path string) error {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(configs.GetPWDError, err)
	}
	envFilePath := filepath.Join(filepath.Dir(path), ".env")
	var environment map[string]string
	if env, err := ParseEnv(envFilePath); err != nil {
		log.Warnf("Could not parse .env file: %s", err.Error())
		environment = make(map[string]string)
	} else {
		environment = env
	}
	// Check if provided docker-compose script is a valid script according to the docker-compose specification
	_, err = cmpcli.ProjectFromOptions(&cmpcli.ProjectOptions{
		WorkingDir:  wd,
		ConfigPaths: []string{path},
		Environment: environment,
	})
	return err
}
