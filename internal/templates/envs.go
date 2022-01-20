package templates

const (
	NethermindEnv = `
EXECUTION_IMAGE_VERSION=nethermind/nethermind
NETHERMIND_CONFIG=mainnet_pruned
NETHERMIND_LOG_LEVEL=INFO
NETHERMIND_JSONRPCCONFIG_ENABLEDMODULES=[Web3,Eth,Subscribe,Net]
NETHERMIND_METRICSCONFIG_ENABLED=false
NETHERMIND_METRICSCONFIG_NODENAME=Nethermind
NETHERMIND_METRICSCONFIG_PUSHGATEWAYURL=http://localhost:9090/metrics
NETHERMIND_HEALTHCHECKSCONFIG_ENABLED=false
NETHERMIND_PRUNINGCONFIG_CACHEMB=2048
NETHERMIND_ETHSTATSCONFIG_ENABLED=false
NETHERMIND_ETHSTATSCONFIG_SERVER=http://localhost:3000/api
NETHERMIND_ETHSTATSCONFIG_NAME=Nethermind
NETHERMIND_ETHSTATSCONFIG_SECRET=secret
NETHERMIND_ETHSTATSCONFIG_CONTACT=hello@nethermind.io
`

	LighthouseConsensusEnv = `
LH_PEER_COUNT=50
CC_LOG_LEVEL=info
NETWORK=mainnet
EC_NODE={{.ExecutionNodeURL}}
INSTANCE_NAME=Lighthouse
CC_IMAGE_VERSION=sigp/lighthouse:latest	
`
	LighthouseValidatorEnv = `
CC_NODE={{.ConsensusNodeURL}}
GRAFFITI={{.ExecutionEngineName}}
VL_LOG_LEVEL=info
INSTANCE_NAME=LighthouseValidator
VL_IMAGE_VERSION=
`

	LodestarConsensusEnv = `
LH_PEER_COUNT=50
CC_LOG_LEVEL=info
NETWORK=mainnet
EC_NODE={{.ExecutionNodeURL}}
INSTANCE_NAME=Lodestar
CC_IMAGE_VERSION=
`
	LodestarValidatorEnv = `
CC_NODE={{.ConsensusNodeURL}}
GRAFFITI={{.ExecutionEngineName}}
VL_LOG_LEVEL=info
INSTANCE_NAME=LodestarValidator
VL_IMAGE_VERSION=
`

	PrysmConsensusEnv = `
LH_PEER_COUNT=50
CC_LOG_LEVEL=info
NETWORK=mainnet
EC_NODE={{.ExecutionNodeURL}}
INSTANCE_NAME=Prysm
CC_IMAGE_VERSION=
`
	PrysmValidatorEnv = `
CC_NODE={{.ConsensusNodeURL}}
GRAFFITI={{.ExecutionEngineName}}
VL_LOG_LEVEL=info
INSTANCE_NAME=PrysmValidator
VL_IMAGE_VERSION=
`

	TekuConsensusEnv = `
LH_PEER_COUNT=50
CC_LOG_LEVEL=info
NETWORK=mainnet
EC_NODE={{.ExecutionNodeURL}}
INSTANCE_NAME=Teku
CC_IMAGE_VERSION=
`
	TekuValidatorEnv = `
CC_NODE={{.ConsensusNodeURL}}
NETWORK=mainnet
GRAFFITI={{.ExecutionEngineName}}
VL_LOG_LEVEL=info
INSTANCE_NAME=TekuValidator
VL_IMAGE_VERSION=
JAVA_OPTS=-XX:SoftMaxHeapSize=2g -Xmx4g
`
)
