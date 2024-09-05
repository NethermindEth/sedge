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
package compose

// ComposeService defines the structure of a service in the output of the 'docker compose ps' command.
type ComposeService struct {
	// Id is the ID of the container.
	Id string `json:"ID"`
	// Service is the name of the service.
	Service string `json:"Service"`
	// Name is the name of the container.
	Name string `json:"Name"`
	// State is the state of the container.
	State string `json:"State"`
}

// DockerComposeStopOptions defines the options for the 'docker compose stop' command.
type DockerComposeStopOptions struct {
	// Path specifies the location of the docker-compose.yaml file.
	Path string
}
