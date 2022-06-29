package generate

// EnvData : Struct Data object to be applied to the docker-compose script environment (.env) template
type EnvData struct {
	ElImage             string
	ElDataDir           string
	CcImage             string
	CcDataDir           string
	VlImage             string
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
	ExecutionImage    string
	ConsensusClient   string
	ConsensusImage    string
	ValidatorClient   string
	ValidatorImage    string
	GenerationPath    string
	Network           string
	CheckpointSyncUrl string
	FeeRecipient      string
	JWTSecretPath     string
	FallbackELUrls    []string
	ElExtraFlags      []string
	ClExtraFlags      []string
	VlExtraFlags      []string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	TTD               bool
	CcPrysmCfg        bool
	VlPrysmCfg        bool
	CheckpointSyncUrl string
	FeeRecipient      string
	ElDiscoveryPort   string
	ElMetricsPort     string
	ClDiscoveryPort   string
	ClMetricsPort     string
	VlMetricsPort     string
	FallbackELUrls    []string
	ElExtraFlags      []string
	ClExtraFlags      []string
	VlExtraFlags      []string
}
