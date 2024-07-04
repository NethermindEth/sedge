package contracts

import (
	"github.com/NethermindEth/sedge/configs"
)

type FeeRecipientConfig struct {
	Network             string
	FeeRecipientAddress string
}

var FeeRecipient = map[string]FeeRecipientConfig{
	configs.NetworkMainnet: {
		Network:             configs.NetworkMainnet,
		FeeRecipientAddress: "0x388C818CA8B9251b393131C08a736A67ccB19297",
	},
	configs.NetworkHolesky: {
		Network:             configs.NetworkHolesky,
		FeeRecipientAddress: "0xE73a3602b99f1f913e72F8bdcBC235e206794Ac8",
	},
	configs.NetworkSepolia: {
		Network:             configs.NetworkSepolia,
		FeeRecipientAddress: "0x94B1B8e2680882f8652882e7F196169dE3d9a3B2",
	},
}
