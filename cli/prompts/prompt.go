package prompts

type Prompt interface {
	Passphrase() string
	ExistingVal() int64
	NumberVal() int64
	Eth1Withdrawal() (string, error)
	FeeRecipient() (string, error)

	Select(label string, options ...string) (string, error)
	Confirm(label string) (bool, error)
	Input(label string, required bool) (string, error)
	InputHide(label string) (string, error)
	InputNumber(label string) (int64, error)
	InputFilePath(label string, required bool) (string, error)
}
