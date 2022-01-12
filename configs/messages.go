package configs

// All the strings that are needed for debugging, info and error logging.
const (
	CheckingDependencies          = "Checking dependencies: %s"
	DependenciesPending           = "Pending dependencies: %s"
	DependencyNotInstalled        = "Dependency %s is not installed on host machine"
	DependenciesOK                = "All dependencies are installed on host machine"
	GeneratingDockerComposeScript = "Generating docker-compose script for current selection of clients"
	InstallingDependenciesError   = "Something went wrong while installing dependencies. %s"
	InstallScriptsPath            = "scripts/setup"
	InstallNotSupported           = "install support for %s is not available for %s. Please install it and try again"
	Exiting                       = "Exiting..."
	ShowingInstructionsError      = "Something went wrong while showing the instructions for installing %s"
	ScriptPathError               = "Failed to get path for instructions file. Error: %s"
	ReadingInstructionError       = "Failed to read instructions from file %s"
	InstructionsFor               = "Instructions for %s"
	OSNotSupported                = "installation not supported for %s"
	ProvideClients                = "Please provide both execution client and consensus client"
	IncorrectClient               = "Incorrect client name %s. Please provide correct client name. Use 'listClient' command to see the list of supported clients"
	NoClientsFound                = "No client found. Please check your configuration file"
)
