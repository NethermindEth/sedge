package configs

import (
	_ "embed"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Image struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func (i *Image) String() string {
	return fmt.Sprintf("%s:%s", i.Name, i.Version)
}

type Images struct {
	Get        Image `yaml:"geth"`
	Nethermind Image `yaml:"nethermind"`
	Besu       Image `yaml:"besu"`
	Erigon     Image `yaml:"erigon"`
	Lighthouse Image `yaml:"lighthouse"`
	Prysm      Image `yaml:"prysm"`
	Lodestar   Image `yaml:"lodestar"`
	Teku       Image `yaml:"teku"`
}

var ClientImages struct {
	Execution struct {
		Get        Image `yaml:"geth"`
		Besu       Image `yaml:"besu"`
		Nethermind Image `yaml:"nethermind"`
		Erigon     Image `yaml:"erigon"`
	}
	BeaconChain struct {
		Lighthouse Image `yaml:"lighthouse"`
		Lodestar   Image `yaml:"lodestar"`
		Teku       Image `yaml:"teku"`
		Prysm      Image `yaml:"prysm"`
	}
	Validator struct {
		Lighthouse Image `yaml:"lighthouse"`
		Lodestar   Image `yaml:"lodestar"`
		Teku       Image `yaml:"teku"`
		Prysm      Image `yaml:"prysm"`
	}
}

//go:embed client_images.yaml
var clientImages string

func init() {
	err := yaml.Unmarshal([]byte(clientImages), &ClientImages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", ClientImages)
}

const (
// TODO: Remove this comment
// Execution images
// Geth_Image = "ethereum/client-go:v1.10.26"
// Besu_Image = "hyperledger/besu:22.10.3"
// Nethermind_Image = "nethermind/nethermind:1.14.7"
// Erigon_Image = "thorax/erigon:v2.29.0"
// Consensus images
// Lighthouse_ConsensusImage = "sigp/lighthouse:v3.3.0"
// Lodestar_ConsensusImage   = "chainsafe/lodestar:v1.2.2"
// Prysm_ConsensusImage = "gcr.io/prysmaticlabs/prysm/beacon-chain:v3.2.0"
// Validator images
// Lighthouse_ValidatorImage = "sigp/lighthouse:v3.3.0"
// Lodestar_ValidatorImage   = "chainsafe/lodestar:v1.2.2"
// Teku_ValidatorImage       = "consensys/teku:22.11.0"
// Prysm_ValidatorImage      = "gcr.io/prysmaticlabs/prysm/validator:v3.2.0"
)
