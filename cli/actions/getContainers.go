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
	"context"
	"fmt"

	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
)

type GetContainersDataOptions struct {
	DockerComposePath string
}

func (actions *sedgeActions) GetContainersData(options GetContainersDataOptions) (ContainersData, error) {
	log.Info("Showing existing containers information")

	composeData, err := generate.ParseCompose(options.DockerComposePath)
	if err != nil {
		log.Errorf("Failed to parse %s docker compose file: %v", options.DockerComposePath, err)
		return ContainersData{}, err
	}

	if composeData.Services != nil {
		// To inspect names
		containersNames := make([]string, 0, 5)
		if composeData.Services.Execution != nil && composeData.Services.Execution.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Execution.ContainerName)
		}
		if composeData.Services.Mevboost != nil && composeData.Services.Mevboost.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Mevboost.ContainerName)
		}
		if composeData.Services.Consensus != nil && composeData.Services.Consensus.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Consensus.ContainerName)
		}
		if composeData.Services.Validator != nil && composeData.Services.Validator.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Validator.ContainerName)
		}

		outputData := ContainersData{}
		outputData.Containers = make([]ContainerData, 0, len(containersNames))
		for _, containerName := range containersNames {
			containerData, err := actions.dockerClient.ContainerInspect(context.Background(), containerName)
			if err != nil {
				log.Warnf("Failed to inspect container %s: %v", containerName, err)
				continue
			}
			data, err := getContainerData(containerData)
			if err != nil {
				log.Warnf("Failed to get container %s data: %v", containerName, err)
				continue
			}

			outputData.Containers = append(outputData.Containers, data)
		}

		return outputData, nil

	} else {
		log.Warnf("No containers found in %s", options.DockerComposePath)
	}
	return ContainersData{
		Containers: []ContainerData{},
	}, err
}

type ContainersData struct {
	Containers []ContainerData `yaml:"containers"`
}

type ContainerData struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Ip    string `yaml:"internal_ip"`
}

func getContainerData(containerData types.ContainerJSON) (ContainerData, error) {
	data := ContainerData{}

	sedgeNetwork, ok := containerData.NetworkSettings.Networks["sedge-network"] // FIXME: fix in case of network data renaming
	if !ok {
		return ContainerData{}, fmt.Errorf("failed to get sedge-network for container %s", containerData.Name)
	}

	data.Name = containerData.Name
	data.Image = containerData.Config.Image
	data.Ip = sedgeNetwork.IPAddress

	return data, nil
}
