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
package actions

import (
	"errors"
	"fmt"
	"os"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
)

var ErrComposeFileNotFound = errors.New("docker-compose file not found")

// ValidateCompose check if the docker-compose file is valid
func (a *sedgeActions) ValidateDockerComposeFile(path string, services ...string) error {
	fileStat, err := os.Stat(path)
	if err != nil {
		return ErrComposeFileNotFound
	}
	if fileStat.IsDir() {
		return ErrComposeFileNotFound
	}
	if err := utils.ValidateCompose(path); err != nil {
		return fmt.Errorf(configs.InvalidComposeErr, err)
	}
	if len(services) == 0 {
		return nil
	}
	composeServices, err := utils.LoadDockerComposeServices(path)
	if err != nil {
		return err
	}
	for _, service := range services {
		if !utils.Contains(composeServices, service) {
			return fmt.Errorf(configs.InvalidService, service)
		}
	}
	return nil
}
