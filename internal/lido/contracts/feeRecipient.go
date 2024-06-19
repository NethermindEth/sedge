package contracts

type FeeRecipientConfig struct {
	Network             string
	FeeRecipientAddress string
}

var FeeRecipient = []FeeRecipientConfig{
	{
		Network:             "mainnet",
		FeeRecipientAddress: "0x388C818CA8B9251b393131C08a736A67ccB19297",
	},
	{
		Network:             "holesky",
		FeeRecipientAddress: "0xE73a3602b99f1f913e72F8bdcBC235e206794Ac8",
	},
	{
		Network:             "sepolia",
		FeeRecipientAddress: "0x94B1B8e2680882f8652882e7F196169dE3d9a3B2",
	},
}
