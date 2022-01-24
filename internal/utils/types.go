package utils

// DistroInfo : Struct Contains name, architecture and version of the linux distribution of the host machine
type DistroInfo struct {
	Name         string
	Architecture string
	Version      string
}

// ConsensusEnv : Struct Data object to be applied to consensus docker-compose script environment template
type ConsensusEnv struct {
	ExecutionNodeURL string
}

// ValidatorEnv : Struct Data object to be applied to validator docker-compose script environment template
type ValidatorEnv struct {
	ConsensusNodeURL    string
	ExecutionEngineName string
}
