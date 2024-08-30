package prometheus

import "errors"

var (
	ErrReloadFailed   = errors.New("failed to reload Prometheus config")
	ErrInvalidOptions = errors.New("invalid options for grafana setup")
)
