package configs

// All the strings that are needed for debugging, info and error logging.
const (
	CheckingDependencies          = "Checking dependencies: %s"
	DependenciesPending           = "Pending dependencies: %s"
	DependencyNotInstalled        = "Dependency %s is not installed on host machine"
	DependenciesOK                = "All dependencies are installed on host machine"
	GeneratingDockerComposeScript = "Generating docker-compose script for current selection of clients"
	InstallingDependenciesError   = "Something went wrong while installing dependencies. Error: %s"
)
