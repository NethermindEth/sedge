package factory

// ClientFlags defines the interface for accessing client configuration flags
type ClientFlags interface {
	GetExecutionName() string
	GetConsensusName() string
	GetValidatorName() string
	GetOptimismName() string
	GetTaikoName() string
	GetL2ExecutionName() string
	GetDistributedValidatorName() string
	GetExecutionApiUrl() string
	GetNetwork() string
	IsNoValidator() bool
}
