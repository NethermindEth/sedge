package services

import "errors"

var (
	ErrStoppingContainer  = errors.New("error stopping container")
	ErrStartingContainer  = errors.New("error starting container")
	ErrContainerNotFound  = errors.New("container not found")
	ErrMultipleContainers = errors.New("multiple containers")
)
