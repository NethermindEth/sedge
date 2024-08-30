package services

import "errors"

var (
	ErrStartingContainer = errors.New("error starting container")
	ErrContainerNotFound = errors.New("container not found")
	ErrStoppingContainer = errors.New("error stopping container")
	ErrNetworksNotFound  = errors.New("networks not found")
)
