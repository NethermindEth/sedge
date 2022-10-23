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
package configs

// All the strings that are needed for debugging and info logging, and constant strings.
const (
	CheckingDependencies            = "Checking dependencies: %s"
	DependenciesPending             = "pending dependencies: %s"
	DependenciesOK                  = "All dependencies are installed on host machine"
	GeneratingDockerComposeScript   = "Generating docker-compose script for current selection of clients"
	GeneratingEnvFile               = "Generating environment file for current selection of clients"
	Exiting                         = "Exiting..."
	InstructionsFor                 = "Instructions for %s"
	OSNotSupported                  = "installation not supported for %s"
	ProvideClients                  = "Please provide both execution client and consensus client"
	CreatedFile                     = "Created file %s"
	DefaultDockerComposeScriptsPath = "./docker-compose-scripts"
	OnPremiseExecutionURL           = "http://execution"
	OnPremiseConsensusURL           = "http://consensus"
	ClientNotSupported              = "client %s is not supported. Please use 'clients' command to see the list of supported clients"
	PrintingFile                    = "File %s:"
	SupportedClients                = "Supported clients of type %s: %v"
	ConfigClientsMsg                = "Provided clients of type %s in configuration file: %v"
	RunningDockerCompose            = "Running docker-compose script"
	Component                       = "component"
	RunningCommand                  = "Running command: %s"
	ConfigFileName                  = ".sedge"
	UnableToProceed                 = "Unable to proceed. Please check the logs for more details"
	DefaultDockerComposeScriptName  = "docker-compose.yml"
	CheckingDockerEngine            = "Checking if docker engine is on"
	DepositCLIDockerImageUrl        = "nethermindeth/staking-deposit-cli" //"github.com/ethereum/staking-deposit-cli"
	DepositCLIDockerImageName       = "nethermindeth/staking-deposit-cli" //"deposit-cli:local"
	GeneratingKeystores             = "Generating keystores..."
	GeneratingKeystoresLegacy       = "Generating keystore folder"
	KeystoresGenerated              = "Keystores generated."
	GeneratingDepositData           = "Generating deposit data..."
	DepositDataGenerated            = "Deposit data generated."
	KeysFoundAt                     = "If everything went well, your keys can be found at: %s"
	ImageNotFoundBuilding           = "Image %s not found, building it"
	ImageNotFoundPulling            = "Image %s not found, pulling it"
	ExecutionDefaultDataDir         = "./execution-data"
	ConsensusDefaultDataDir         = "./consensus-data"
	ValidatorDefaultDataDir         = "./validator-data"
	KeystoreDefaultDataDir          = "./keystore"
	ReviewKeystorePath              = "In case you used custom paths for the 'cli' or the 'keys' commands, please review if the keystore path in the generated .env file points to the generated keystore folder (the .env key should be KEYSTORE_DIR). If not, change the path in the .env file to the correct one."
	NodesSynced                     = "Execution and Consensus clients are synced, proceding to start validator node"
	RemoteNodeNeeded                = `
If you want to run a validator, make sure you have access to a remote or external consensus node.
	
You can use one of your own, a friend's node or providers such as Infura. Edit the .env file accordinly please, check that variable CC_NODE <-> (consensus endpoint) have the correct value. The validator node requires a high available consensus node, and consensus in turn needs a high available execution node.`
	ValidatorTips = `
A validator node needs to connect to a synced consensus node, and the consensus node in turn needs to connect to a synced execution node. 
	
While these required nodes (execution/consensus) are syncing, you can setup the keystore folder(s) using the staking-deposit-cli tool (https://github.com/ethereum/staking-deposit-cli) or the command 'sedge keys'. If you don't want to use 'sedge keys', make sure to set .env variables KEYSTORE_DIR and VL_DATA_DIR to correct values. You can also check https://launchpad.ethereum.org/ for tips and more instructions.
	
sedge will track the syncing progress of the required nodes and let you run the validator after those nodes are synced (as we recommend). This takes a while, so you have time to prepare the keystore folder.
	
Don't make the deposit to register the validator until its ready to run!!!
	
After you complete the above steps follow https://launchpad.ethereum.org/ instructions to register the validator addresses of the validators you want to setup and manage using the validator node.`
	HappyStaking = `
Validator is up. Remember to setup the keystore folder and to make the deposit. If something went wrong and your validator node is down, don't panic! Short downtimes are not very bad. Check the logs and try to fix the errors. Make sure errors are not related to connection issues with the consensus nodes nor related to the keystore folder path or validator data directory. You can start the validator again using the instructions displayed by the tool.
	
Happy Staking!
	`
	HappyStaking2 = `
You set the flag --run-clients=none which means the docker-compose scripts are generated but they will not be executed by sedge and the setup is stopped here. Normally sedge will run the execution and consensus services, wait for the execution and consensus client to sync and after that start the validator node, giving you instructions/recommendations about what to do in every step.
	
In case you don't know what to do next, please consider running sedge without the --run-clients flag (default behavior) and without the -r flag as well.
	
Follow https://launchpad.ethereum.org/ and happy staking!`
	ExecutionClientNotSpecifiedWarn = "Execution client not provided. A random client will be selected. Random client: %s"
	ConsensusClientNotSpecifiedWarn = "Consensus client not provided. Selecting same pair of clients for consensus and validator clients"
	ValidatorClientNotSpecifiedWarn = "Validator client not provided. Selecting same pair of clients for consensus and validator clients"
	CLNotSpecifiedWarn              = "Consensus and validator clients not provided. Selecting same pair of clients for consensus and validator clients using a random client. Random client: %s"
	GeneratingJWTSecret             = "Generating JWT secret for client authentication"
	JWTSecretGenerated              = "JWT secret generated"
	CreatingKeystorePassword        = "Creating keystore_password.txt on keystore folder"
	KeystorePasswordCreated         = "keystore_password.txt on keystore folder created with provided password"
	MnemonicTips                    = "The following mnemonic is going to be used to create the validator keystore. Please save it carefully. It can be used to generate the keystore folder again. If you lose the password and mnemonic, access to your keystore will be lost forever!"
	GeneratingMnemonic              = "Existing mnemonic not provided. Generating mnemonic for validator keystore:"
	StoreMnemonic                   = "Make sure to store your mnemonic somewhere safe. Losing it could end in the lost of your validators. Press enter to continue" // TODO: improve warning message
	PreparingTekuDatadir            = "Preparing teku datadirs (must have full read/write/execute permissions to work)"
	GettingContainersIP             = "Proceding to get execution and consensus containers IP address for the monitoring tool"
	WaitingForNodesToStart          = "Waiting a minute for nodes to start"
	CustomImagesWarning             = "You are using custom images for the execution, consensus or validator clients!!! Make sure this is intended. Also check these images are correct and available from this device otherwise the setup will fail or have an unexpected behavior."
	DefaultDiscoveryPortEL          = "30303"
	DefaultMetricsPortEL            = "8008"
	DefaultApiPortEL                = "8545"
	DefaultAuthPortEL               = "8551"
	DefaultWSPortEL                 = "8546"
	DefaultDiscoveryPortCL          = "9000"
	DefaultMetricsPortCL            = "5054"
	DefaultApiPortCL                = "4000"
	DefaultAdditionalApiPortCL      = "4001"
	DefaultMetricsPortVL            = "5056"
	DefaultMevPort                  = "18550"
	MapAllPortsWarning              = "You are mapping all ports for the clients!!! Make sure this is intended. This could make the clients vulnerable to attacks. Be sure to setup a firewall."
	CheckpointUrlUsedWarning        = "A Checkpoint Sync Url will be used for the consensus node. Using %s ."
	NoBootnodesFound                = "No bootnodes found for %s/%s/%s"
	UnableToCheckVersion            = "Unable to check for new Version. Please check manually at " +
		"https://github.com/NethermindEth/sedge/releases, with error:"
	NeedVersionUpdate = "A new Version of sedge is available. Please update to the latest Version. See " +
		"https://github.com/NethermindEth/sedge/releases for more information. Latest detected tag:"
	VersionUpdated = "You are running the latest version of sedge. Version: "
)
