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
package monitoring

import "github.com/NethermindEth/sedge/internal/common"

// DockerServiceManager is an interface for managing Docker containers.
type DockerServiceManager interface {
	// ContainerStatus returns the status of a container.
	ContainerStatus(container string) (common.Status, error)

	// ContainerIP returns the IP address of the container.
	ContainerIP(container string) (string, error)

	// ContainerNetworks returns the networks of a container.
	ContainerNetworks(container string) ([]string, error)

	// NetworkConnect connects a container to a network.
	NetworkConnect(container, network string) error

	// NetworkDisconnect disconnects a container from a network.
	NetworkDisconnect(container, network string) error
}
