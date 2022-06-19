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
	OnPremiseConsensusURL           = "http://consensus:4000"
	ClientNotSupported              = "client %s is not supported. Please use 'clients' command to see the list of supported clients"
	PrintingFile                    = "File %s:"
	SupportedClients                = "Supported clients of type %s: %v"
	ConfigClientsMsg                = "Provided clients of type %s in configuration file: %v"
	RunningDockerCompose            = "Running docker-compose script"
	Component                       = "component"
	RunningCommand                  = "Running command: %s"
	ConfigFileName                  = ".1click"
	UnableToProceed                 = "Unable to proceed. Please check the logs for more details"
	DefaultDockerComposeScriptName  = "docker-compose.yml"
	CheckingDockerEngine            = "Checking if docker engine is on"
	DepositCLIDockerImageUrl        = "github.com/ethereum/staking-deposit-cli"
	DepositCLIDockerImageName       = "deposit-cli:local"
	GeneratingKeystore              = "Generating keystore folder"
	KeysFoundAt                     = "If everything went well, your keys can be found at: %s"
	ImageNotFound                   = "Image %s not found, building it"
	ExecutionDefaultDataDir         = "./execution-data"
	ConsensusDefaultDataDir         = "./consensus-data"
	ValidatorDefaultDataDir         = "./validator-data"
	KeystoreDefaultDataDir          = "./keystore"
	ReviewKeystorePath              = "In case you used custom paths for the 'cli' or the 'keys' commands, please review if the keystore path in the generated .env file points to the generated keystore folder (the .env key should be KEYSTORE_DIR). If not, change the path in the .env file to the correct one."
	NodesSynced                     = "Execution and Consensus clients are synced, proceding to start validator node"
	RemoteNodeNeeded                = `If you want to run a validator, make sure you have access to a remote or external consensus node.
	
	You can use one of your own, a friend's node or providers such as Infura. Edit the .env file accordinly please, check that variable CC_NODE <-> (consensus endpoint) have the correct value. The validator node requires a high available consensus node, and consensus in turn needs a high available execution node.`
	ValidatorTips = `A validator node needs to connect to a synced consensus node, and the consensus node in turn needs to connect to a synced execution node. 
	
	While these required nodes (execution/consensus) are syncing, you can setup the keystore folder(s) using the staking-deposit-cli tool (https://github.com/ethereum/staking-deposit-cli) or the command '1click keys'. If you don't want to use '1click keys', make sure to set .env variables KEYSTORE_DIR and VL_DATA_DIR to correct values. You can also check https://launchpad.ethereum.org/ for tips and more instructions.
	
	1click will track the syncing progress of the required nodes and let you run the validator after those nodes are synced (as we recommend). This takes a while, so you have time to prepare the keystore folder.
	
	Don't make the deposit to register the validator until its ready to run!!!
	
	After you complete the above steps follow https://launchpad.ethereum.org/ instructions to register the validator addresses of the validators you want to setup and manage using the validator node.`
	HappyStaking = `Validator is up. Remember to setup the keystore folder and to make the deposit. If something went wrong and your validator node is down, don't panic! Short downtimes are not very bad. Check the logs and try to fix the errors. Make sure errors are not related to connection issues with the consensus nodes nor related to the keystore folder path or validator data directory. You can start the validator again using the instructions displayed by the tool.
	
	Happy Staking!
	`
	HappyStaking2 = `You set the flag --run-clients=none which means the docker-compose scripts are generated but they will not be executed by 1click and the setup is stopped here. Normally 1click will run the execution and consensus services, wait for the execution and consensus client to sync and after that start the validator node, giving you instructions/recommendations about what to do in every step.
	
	In case you don't know what to do next, please consider running 1click without the --run-clients flag (default behavior) and without the -r flag as well.
	
	Follow https://launchpad.ethereum.org/ and happy staking!`
	ExecutionClientNotSpecifiedWarn = "Execution client not provided. A random client will be selected. Random client: %s"
	ConsensusClientNotSpecifiedWarn = "Consensus client not provided. Selecting same pair of clients for consensus and validator clients"
	ValidatorClientNotSpecifiedWarn = "Validator client not provided. Selecting same pair of clients for consensus and validator clients"
	CLNotSpecifiedWarn              = "Consensus and validator clients not provided. Selecting same pair of clients for consensus and validator clients using a random client. Random client: %s"
	GeneratingJWTSecret             = "Generating JWT secret for client authentication"
	JWTSecretGenerated              = "JWT secret generated"
)
