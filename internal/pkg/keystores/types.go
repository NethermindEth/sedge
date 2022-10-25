package keystores

type ValidatorKeysGenData struct {
	Mnemonic            string
	UseUniquePassphrase bool
	Passphrase          string
	OutputPath          string
	MinIndex            uint64
	MaxIndex            uint64
	Insecure            bool
	NetworkName         string
	ForkVersion         string
	AmountGwei          uint64
	AsJsonList          bool
}
