package prompts

type Prompt interface {
	Passphrase() string
	ExistingVal() int64
	NumberVal() int64
	Eth1Withdrawal() (string, error)
	FeeRecipient() (string, error)
}
