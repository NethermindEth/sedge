package generate

// ExecutionEnv : Struct Data object to be applied to execution docker-compose script environment template
type ExecutionEnv struct {
	DataDir       string
	JWTSecretPath string
}

// ConsensusEnv : Struct Data object to be applied to consensus docker-compose script environment template
type ConsensusEnv struct {
	ExecutionNodeURL string
	DataDir          string
	FeeRecipient     string
	JWTSecretPath    string
}

// ValidatorEnv : Struct Data object to be applied to validator docker-compose script environment template
type ValidatorEnv struct {
	ConsensusNodeURL    string
	ExecutionEngineName string
	DataDir             string
	KeystoreDir         string
}

// GenerationData : Struct Data object for script's generation
type GenerationData struct {
	ExecutionClient   string
	ConsensusClient   string
	ValidatorClient   string
	GenerationPath    string
	Network           string
	CheckpointSyncUrl string
	FeeRecipient      string
	JWTSecretPath     string
	FallbackELUrls    []string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	ElTTD             bool
	CcTTD             bool
	CcPrysmCfg        bool
	VlPrysmCfg        bool
	CheckpointSyncUrl string
	FeeRecipient      string
	FallbackELUrls    []string
}
