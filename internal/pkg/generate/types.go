package generate

// EnvData : Struct Data object to be applied to the docker-compose script environment (.env) template
type EnvData struct {
	ElDataDir           string
	CcDataDir           string
	VlDataDir           string
	ExecutionNodeURL    string
	ConsensusNodeURL    string
	FeeRecipient        string
	JWTSecretPath       string
	ExecutionEngineName string
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
