package grafana

import "errors"

var (
	ErrConfigNotFound = errors.New("configuration file not found")
	ErrInvalidOptions = errors.New("invalid options for grafana setup")
)
