package factory

import (
	"fmt"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/utils"
	log "github.com/sirupsen/logrus"
)

const (
	execution            = "execution"
	consensus            = "consensus"
	validator            = "validator"
	distributedValidator = "distributedValidator"
	optimism             = "optimism"
	opExecution          = "opexecution"
	taiko                = "taiko"
	tExecution           = "texecution"
	NetworkGnosis        = "gnosis"
	NetworkChiado        = "chiado"
)

// NodeInitializer defines how a specific type of node should be initialized
type NodeInitializer interface {
	Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error)
	ShouldInitialize(services []string, flags ClientFlags) bool
	UpdateResult(result *clients.Clients, client *clients.Client)
}

// BaseNodeInitializer provides common functionality for node initializers
type BaseNodeInitializer struct {
	serviceType string
	config      clientConfig
}

func (b *BaseNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	return initializeClient(allClients, b.config)
}

func (b *BaseNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, b.serviceType)
}

// clientConfig holds the configuration for initializing a client
type clientConfig struct {
	clientType string
	flagName   string
	forceName  string
	warningMsg string
}

// initializeClient handles the common logic for client initialization
func initializeClient(allClients clients.OrderedClients, config clientConfig) (*clients.Client, error) {
	client, err := clients.RandomChoice(allClients[config.clientType])
	if err != nil {
		return nil, fmt.Errorf("failed to choose random client for %s: %w", config.clientType, err)
	}

	if config.flagName == "" {
		client.SetImageOrDefault("")
		return client, clients.ValidateClient(client, config.clientType)
	}

	parts := strings.Split(config.flagName, ":")
	if config.forceName != "" {
		client.Name = config.forceName
	} else {
		client.Name = parts[0]
	}

	if len(parts) > 1 {
		if config.warningMsg != "" {
			log.Warn(config.warningMsg)
		}
		client.Image = strings.Join(parts[1:], ":")
		client.Modified = true
	}

	client.SetImageOrDefault(strings.Join(parts[1:], ":"))
	if err := clients.ValidateClient(client, config.clientType); err != nil {
		return nil, fmt.Errorf("client validation failed for %s: %w", config.clientType, err)
	}
	return client, nil
}

// ExecutionNodeInitializer handles execution client initialization
type ExecutionNodeInitializer struct {
	BaseNodeInitializer
}

func NewExecutionNodeInitializer() *ExecutionNodeInitializer {
	return &ExecutionNodeInitializer{
		BaseNodeInitializer{
			serviceType: execution,
			config: clientConfig{
				clientType: execution,
				warningMsg: configs.CustomExecutionImagesWarning,
			},
		},
	}
}

func (e *ExecutionNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	e.config.flagName = flags.GetExecutionName()
	return e.BaseNodeInitializer.Initialize(allClients, flags)
}

func (e *ExecutionNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, e.serviceType) && flags.GetExecutionApiUrl() == ""
}

func (e *ExecutionNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Execution = client
}

// ConsensusNodeInitializer handles consensus client initialization
type ConsensusNodeInitializer struct {
	BaseNodeInitializer
}

func NewConsensusNodeInitializer() *ConsensusNodeInitializer {
	return &ConsensusNodeInitializer{
		BaseNodeInitializer{
			serviceType: consensus,
			config: clientConfig{
				clientType: consensus,
				warningMsg: configs.CustomConsensusImagesWarning,
			},
		},
	}
}

func (c *ConsensusNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	c.config.flagName = flags.GetConsensusName()
	// Special handling for Gnosis and Chiado networks
	if flags.GetNetwork() == NetworkGnosis || flags.GetNetwork() == NetworkChiado {
		if flags.GetConsensusName() == "nimbus" {
			c.config.flagName = "nimbus:ghcr.io/gnosischain/gnosis-nimbus-eth2:v24.9"
		}
	}

	client, err := c.BaseNodeInitializer.Initialize(allClients, flags)
	if err != nil {
		return nil, err
	}

	// Handle custom image warning
	if client != nil && strings.Contains(flags.GetConsensusName(), ":") {
		log.Warn(configs.CustomConsensusImagesWarning)
	}

	return client, nil
}

func (c *ConsensusNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, c.serviceType) && flags.GetExecutionApiUrl() == ""
}

func (c *ConsensusNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Consensus = client
}

// ValidatorNodeInitializer handles validator client initialization
type ValidatorNodeInitializer struct {
	BaseNodeInitializer
}

func NewValidatorNodeInitializer() *ValidatorNodeInitializer {
	return &ValidatorNodeInitializer{
		BaseNodeInitializer{
			serviceType: validator,
			config: clientConfig{
				clientType: validator,
				warningMsg: configs.CustomValidatorImagesWarning,
			},
		},
	}
}

func (v *ValidatorNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	v.config.flagName = flags.GetValidatorName()
	return v.BaseNodeInitializer.Initialize(allClients, flags)
}

func (v *ValidatorNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, v.serviceType) && !flags.IsNoValidator()
}

func (v *ValidatorNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Validator = client
}

// L2BaseInitializer provides common functionality for L2 nodes (Optimism and Taiko)
type L2BaseInitializer struct {
	BaseNodeInitializer
	l2ExecutionType string
}

func (l *L2BaseInitializer) initializeL2(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, *clients.Client, error) {
	l2Client, err := l.BaseNodeInitializer.Initialize(allClients, flags)
	if err != nil {
		return nil, nil, fmt.Errorf("%s client initialization failed: %w", l.serviceType, err)
	}

	l2ExecConfig := clientConfig{
		clientType: l.l2ExecutionType,
		flagName:   flags.GetL2ExecutionName(),
	}
	execClient, err := initializeClient(allClients, l2ExecConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("%s L2 execution client initialization failed: %w", l.serviceType, err)
	}

	return l2Client, execClient, nil
}

// OptimismNodeInitializer handles optimism client initialization
type OptimismNodeInitializer struct {
	L2BaseInitializer
	flags      ClientFlags
	execClient *clients.Client
}

func NewOptimismNodeInitializer() *OptimismNodeInitializer {
	return &OptimismNodeInitializer{
		L2BaseInitializer: L2BaseInitializer{
			BaseNodeInitializer: BaseNodeInitializer{
				serviceType: optimism,
				config: clientConfig{
					clientType: optimism,
					forceName:  "opnode",
				},
			},
			l2ExecutionType: opExecution,
		},
	}
}

func (o *OptimismNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	o.flags = flags
	client, execClient, err := o.initializeL2(allClients, flags)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize optimism node: %w", err)
	}
	o.execClient = execClient

	// Force opnode name for optimism client
	if client != nil {
		client.Name = "opnode"
		if strings.Contains(flags.GetOptimismName(), ":") {
			parts := strings.Split(flags.GetOptimismName(), ":")
			client.Image = strings.Join(parts[1:], ":")
			client.Modified = true
		}
		client.SetImageOrDefault(strings.Join(strings.Split(flags.GetOptimismName(), ":")[1:], ":"))
		if err = clients.ValidateClient(client, optimism); err != nil {
			return nil, err
		}
	}

	// Handle L2 execution client name and image
	if execClient != nil {
		parts := strings.Split(flags.GetL2ExecutionName(), ":")
		if len(parts) > 1 {
			execClient.Name = strings.ReplaceAll(parts[0], "-", "")
			if len(parts) > 1 {
				execClient.Image = strings.Join(parts[1:], ":")
				execClient.Modified = true
			}
			execClient.SetImageOrDefault(strings.Join(parts[1:], ":"))
			if err = clients.ValidateClient(execClient, opExecution); err != nil {
				return nil, err
			}
		}
	}

	return client, nil
}

func (o *OptimismNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, o.serviceType)
}

func (o *OptimismNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Optimism = client
	if o.execClient != nil {
		result.L2Execution = o.execClient
	}
	if client != nil && o.flags != nil && o.flags.GetExecutionApiUrl() != "" {
		result.Execution = nil
		result.Consensus = nil
	}
}

// TaikoNodeInitializer handles taiko client initialization
type TaikoNodeInitializer struct {
	L2BaseInitializer
	flags      ClientFlags
	execClient *clients.Client
}

func NewTaikoNodeInitializer() *TaikoNodeInitializer {
	return &TaikoNodeInitializer{
		L2BaseInitializer: L2BaseInitializer{
			BaseNodeInitializer: BaseNodeInitializer{
				serviceType: taiko,
				config: clientConfig{
					clientType: taiko,
					flagName:   "taiko-name",
					forceName:  "taikoclient",
				},
			},
			l2ExecutionType: tExecution,
		},
	}
}

func (t *TaikoNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	t.flags = flags
	client, execClient, err := t.initializeL2(allClients, flags)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize taiko node: %w", err)
	}
	t.execClient = execClient

	// Force taikoclient name for taiko client
	if client != nil {
		client.Name = "taikoclient"
		if strings.Contains(flags.GetTaikoName(), ":") {
			parts := strings.Split(flags.GetTaikoName(), ":")
			client.Image = strings.Join(parts[1:], ":")
			client.Modified = true
		}
		client.SetImageOrDefault(strings.Join(strings.Split(flags.GetTaikoName(), ":")[1:], ":"))
		if err = clients.ValidateClient(client, taiko); err != nil {
			return nil, err
		}
	}

	// Handle L2 execution client name and image
	if execClient != nil {
		parts := strings.Split(flags.GetL2ExecutionName(), ":")
		execClient.Name = strings.ReplaceAll(parts[0], "-", "")
		if len(parts) > 1 {
			execClient.Image = strings.Join(parts[1:], ":")
			execClient.Modified = true
		}
		execClient.SetImageOrDefault(strings.Join(parts[1:], ":"))
		if err = clients.ValidateClient(execClient, tExecution); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (t *TaikoNodeInitializer) ShouldInitialize(services []string, flags ClientFlags) bool {
	return utils.Contains(services, t.serviceType)
}

func (t *TaikoNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.Taiko = client
	result.L2Execution = t.execClient
}

// DistributedValidatorNodeInitializer handles distributed validator client initialization
type DistributedValidatorNodeInitializer struct {
	BaseNodeInitializer
}

func NewDistributedValidatorNodeInitializer() *DistributedValidatorNodeInitializer {
	return &DistributedValidatorNodeInitializer{
		BaseNodeInitializer{
			serviceType: distributedValidator,
			config: clientConfig{
				clientType: distributedValidator,
				forceName:  "charon",
			},
		},
	}
}

func (d *DistributedValidatorNodeInitializer) Initialize(allClients clients.OrderedClients, flags ClientFlags) (*clients.Client, error) {
	d.config.flagName = flags.GetDistributedValidatorName()
	return d.BaseNodeInitializer.Initialize(allClients, flags)
}

func (d *DistributedValidatorNodeInitializer) UpdateResult(result *clients.Clients, client *clients.Client) {
	result.DistributedValidator = client
}

// NodeFactory manages the creation of different node types
type NodeFactory struct {
	initializers []NodeInitializer
}

func NewNodeFactory() *NodeFactory {
	return &NodeFactory{
		initializers: []NodeInitializer{
			NewExecutionNodeInitializer(),
			NewConsensusNodeInitializer(),
			NewValidatorNodeInitializer(),
			NewOptimismNodeInitializer(),
			NewTaikoNodeInitializer(),
			NewDistributedValidatorNodeInitializer(),
		},
	}
}

// InitializeClients initializes all required clients based on the services list
func (f *NodeFactory) InitializeClients(allClients clients.OrderedClients, flags ClientFlags, services []string) (*clients.Clients, error) {
	result := &clients.Clients{}

	for _, initializer := range f.initializers {
		if initializer.ShouldInitialize(services, flags) {
			client, err := initializer.Initialize(allClients, flags)
			if err != nil {
				return nil, err
			}
			initializer.UpdateResult(result, client)
		}
	}

	return result, nil
}
