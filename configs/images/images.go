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
