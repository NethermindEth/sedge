package cli


import (

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
)

// NodeOptionsFactory is the interface for creating options
type NodeOptionsFactory interface {
	CreateNetworkOptions() []string
	GetWithdrawalAddress(network string) (string, bool)
	GetFeeRecipient(network string) (string, bool)
	GetRelayURLs(network string) ([]string, error)
	EnableMEVBoost(network string) (bool, bool)
}

// GetNodeOptions returns the appropriate NodeOptionsFactory based on node setup
func GetNodeOptions(nodeSetup string) NodeOptionsFactory {
	switch nodeSetup {
	case LidoNode:
		return &LidoNodeOptions{}
	default:
		return &EthereumNodeOptions{}
	}
}

// LidoNodeOptions for Lido Node
type LidoNodeOptions struct{}
func (l *LidoNodeOptions) CreateNetworkOptions() []string {
	options := contracts.GetLidoSupportedNetworks()
	return options
}
func (l *LidoNodeOptions) GetWithdrawalAddress(network string) (string, bool) {
	_, supported := contracts.NetworkSupportedByLidoKeys(network)
	if supported {
		return contracts.WithdrawalAddress[network].WithdrawalAddress, true
	}
	return "", true
}
func (l *LidoNodeOptions) GetFeeRecipient(network string) (string, bool) {
	_, supported := contracts.NetworkSupportedByLido(network)
	if supported {
		return contracts.FeeRecipient[network].FeeRecipientAddress, true
	}
	return "", true
}
func (l *LidoNodeOptions) GetRelayURLs(network string) ([]string, error) {
	relayURLs, err := mevboostrelaylist.GetRelaysURI(network)
		if err != nil {
			return nil, err
		}
	return relayURLs, nil
}
func (l *LidoNodeOptions) EnableMEVBoost(network string) (bool, bool) {
	_, supported := mevboostrelaylist.NetworkSupportedByLidoMevBoost(network)
	if supported {
		return true, true
	}
	return false, true
}

// EthereumNodeOptions for Ethereum Node
type EthereumNodeOptions struct{}
func (e *EthereumNodeOptions) CreateNetworkOptions() []string {
	return []string{NetworkMainnet, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}
}

func (e *EthereumNodeOptions) GetWithdrawalAddress(network string) (string, bool) {
	return "", false 
}
func (e *EthereumNodeOptions) GetFeeRecipient(network string) (string, bool) {
	return "", false // Default implementation returns false for EthereumNode
}
func (e *EthereumNodeOptions) GetRelayURLs(network string) ([]string, error) {
	return nil, nil
}
func (e *EthereumNodeOptions) EnableMEVBoost(network string) (bool, bool) {
	return false, false
}
