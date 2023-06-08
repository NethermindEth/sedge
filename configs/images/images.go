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
package clientsimages

import (
	"fmt"
)

type Image struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func (i Image) String() string {
	return fmt.Sprintf("%s:%s", i.Name, i.Version)
}

type ExecutionClientsImages interface {
	Geth() Image
	Besu() Image
	Nethermind() Image
	Erigon() Image
}

type ConsensusClientsImages interface {
	Lighthouse() Image
	Lodestar() Image
	Teku() Image
	Prysm() Image
}

type ValidatorClientsImages interface {
	Lighthouse() Image
	Lodestar() Image
	Teku() Image
	Prysm() Image
}

type ClientsImages interface {
	Execution() ExecutionClientsImages
	Consensus() ConsensusClientsImages
	Validator() ValidatorClientsImages
}
