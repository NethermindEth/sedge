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
	"os"

	"github.com/NethermindEth/sedge/configs"
	"gopkg.in/yaml.v2"
)

// LoadDockerComposeServices get the services from a Sedge valid docker-compose script
func LoadDockerComposeServices(path string) ([]string, error) {
	type servicesMap struct {
		Services map[string]interface{} `yaml:"services"`
	}
	composeBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(configs.ReadingYmlErr, err)
	}
	var composeData servicesMap
	// It was already parsed by parseCompose
	if err := yaml.Unmarshal(composeBytes, &composeData); err != nil {
		return nil, fmt.Errorf(configs.ParsingYmlErr, err)
	}

	if composeData.Services != nil {
		var servicesNames []string
		for service := range composeData.Services {
			servicesNames = append(servicesNames, service)
		}
		return servicesNames, nil
	}

	return nil, fmt.Errorf(configs.ServicesNotFoundErr)
}
