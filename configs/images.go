package configs

const (
	// Execution images
	Geth_Image       = "ethereum/client-go:v1.10.26"
	Besu_Image       = "hyperledger/besu:22.10.3"
	Nethermind_Image = "nethermind/nethermind:1.14.7"
	Erigon_Image     = "thorax/erigon:v2.29.0"
	// Consensus images
	Lighthouse_ConsensusImage = "sigp/lighthouse:v3.3.0"
	Lodestar_ConsensusImage   = "chainsafe/lodestar:v1.2.2"
	// TODO add teku
	Prysm_ConsensusImage = "gcr.io/prysmaticlabs/prysm/beacon-chain:v3.2.0"
	// Validator images
	Lighthouse_ValidatorImage = "sigp/lighthouse:v3.3.0"
	Lodestar_ValidatorImage   = "chainsafe/lodestar:v1.2.2"
	Teku_ValidatorImage       = "consensys/teku:22.11.0"
	Prysm_ValidatorImage      = "gcr.io/prysmaticlabs/prysm/validator:v3.2.0"
)
