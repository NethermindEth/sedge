package common

// Status is an enum used to represent the status of a service.
type Status int

const (
	Created Status = iota + 1
	Running
	Paused
	Restarting
	Removing
	Exited
	Dead
	Unknown
	Installed
	NotInstalled
	Broken
)
