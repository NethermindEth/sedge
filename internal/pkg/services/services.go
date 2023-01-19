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
package services

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	ServiceCtSlashingData       = "validator-slashing-data"
	ServiceCtValidatorImport    = "validator-import-client"
	DefaultSedgeValidatorClient = "sedge-validator-client"
	DefaultSedgeExecutionClient = "sedge-execution-client"
	DefaultSedgeConsensusClient = "sedge-consensus-client"
)

type ServiceManager interface {
	Image(service string) (string, error)
	Stop(service string) error
	Start(service string) error
	IsRunning(service string) (bool, error)
	Wait(service string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
	ContainerId(service string) (string, error)
}

func NewServiceManager(dockerClient client.APIClient) ServiceManager {
	return &serviceManager{dockerClient: dockerClient}
}

func ContainerNameWithTag(containerName, tag string) string {
	if tag == "" {
		return containerName
	}
	return containerName + "_" + tag
}

type serviceManager struct {
	dockerClient client.APIClient
}
