package configs

// All the strings that are needed for debugging, info and error logging.
const (
	DependencyNotInstalled        = "Dependency %s is not installed on host machine"
	DependenciesOK                = "All dependencies are installed on host machine"
	GeneratingDockerComposeScript = "Generating docker-compose script for current selection of clients"
	InstallingDependenciesError   = "Something went wrong while installing dependencies. Error: %s"
)

//TODO: Move all this to a config file. No supported clients for now

var Eth1ClientsSupported = []string{}

var ConsensusClientsSupported = []string{}

var Dependencies = []string{
	"docker",
	"docker-compose",
}
