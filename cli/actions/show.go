package actions

import (
	"context"

	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ShowContainersOptions struct {
	DockerComposePath string
}

func (actions *sedgeActions) ShowContainers(options ShowContainersOptions) error {

	log.Info("Showing existing containers information")

	composeData, err := generate.ParseCompose(options.DockerComposePath)
	if err != nil {
		log.Errorf("Failed to parse %s docker compose file: %v", options.DockerComposePath, err)
		return err
	}

	if composeData.Services != nil {
		// To inspect names
		containersNames := make([]string, 0, 5)
		if composeData.Services.Execution != nil && composeData.Services.Execution.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Execution.ContainerName)
		} else if composeData.Services.Consensus != nil && composeData.Services.Consensus.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Consensus.ContainerName)
		} else if composeData.Services.Validator != nil && composeData.Services.Validator.ContainerName != "" {
			containersNames = append(containersNames, composeData.Services.Validator.ContainerName)
		}

		for _, containerName := range containersNames {
			containerData, err := actions.dockerClient.ContainerInspect(context.Background(), containerName)
			if err != nil {
				log.Warnf("Failed to inspect container %s: %v", containerName, err)
			}
			printContainerData(containerData)
		}

	} else {
		log.Warnf("No containers found in %s", options.DockerComposePath)
	}

	return nil
}

type ContainerData struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Ip    string `yaml:"internal_ip"`
}

func printContainerData(containerData types.ContainerJSON) error {
	data := ContainerData{
		Name:  containerData.Name,
		Image: containerData.Config.Image,
		Ip:    containerData.NetworkSettings.IPAddress,
	}

	rawData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	log.Info(string(rawData))

	return nil
}
