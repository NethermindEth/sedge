package utils

// DistroInfo : Struct Contains name and architecture of the linux distribution of the host machine
type DistroInfo struct {
	Name         string `json:"name,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}
