package templates

import "embed"

//go:embed services
var Services embed.FS

//go:embed envs
var Envs embed.FS
