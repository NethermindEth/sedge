package templates

import "embed"

//go:embed services
var Services embed.FS

//go:embed envs
var Envs embed.FS

//go:embed setup
var Setup embed.FS

//go:embed config
var Config embed.FS

//go:embed deposit-cli
var DepositCLI embed.FS
