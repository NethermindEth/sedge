package clientsimages

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed client_images.yaml
var clientImages string

type DefaultExecutionImages struct {
	geth       Image `yaml:"geth"`
	besu       Image `yaml:"besu"`
	nethermind Image `yaml:"nethermind"`
	erigon     Image `yaml:"erigon"`
}

func (dei DefaultExecutionImages) Geth() Image {
	return dei.geth
}

func (dei DefaultExecutionImages) Besu() Image {
	return dei.besu
}

func (dei DefaultExecutionImages) Nethermind() Image {
	return dei.nethermind
}

func (dei DefaultExecutionImages) Erigon() Image {
	return dei.erigon
}

type DefaultConsensusImages struct {
	lighthouse Image `yaml:"lighthouse"`
	lodestar   Image `yaml:"lodestar"`
	teku       Image `yaml:"teku"`
	prysm      Image `yaml:"prysm"`
}

func (dci DefaultConsensusImages) Lighthouse() Image {
	return dci.lighthouse
}

func (dci DefaultConsensusImages) Lodestar() Image {
	return dci.lodestar
}

func (dci DefaultConsensusImages) Teku() Image {
	return dci.teku
}

func (dci DefaultConsensusImages) Prysm() Image {
	return dci.prysm
}

type DefaultValidatorImages struct {
	lighthouse Image `yaml:"lighthouse"`
	lodestar   Image `yaml:"lodestar"`
	teku       Image `yaml:"teku"`
	prysm      Image `yaml:"prysm"`
}

func (dvi DefaultValidatorImages) Lighthouse() Image {
	return dvi.lighthouse
}

func (dvi DefaultValidatorImages) Lodestar() Image {
	return dvi.lodestar
}

func (dvi DefaultValidatorImages) Teku() Image {
	return dvi.teku
}

func (dvi DefaultValidatorImages) Prysm() Image {
	return dvi.prysm
}

type DefaultClientsImages struct {
	execution DefaultExecutionImages `yaml:"execution"`
	consensus DefaultConsensusImages `yaml:"consensus"`
	validator DefaultValidatorImages `yaml:"validator"`
}

func (dci DefaultClientsImages) Execution() ExecutionClientsImages {
	return dci.execution
}

func (dci DefaultClientsImages) Consensus() ConsensusClientsImages {
	return dci.consensus
}

func (dci DefaultClientsImages) Validator() ValidatorClientsImages {
	return dci.validator
}

func NewDefaultClientsImages() (ClientsImages, error) {
	dci := DefaultClientsImages{}

	err := yaml.Unmarshal([]byte(clientImages), &dci)
	if err != nil {
		return nil, err
	}

	return &dci, nil
}
