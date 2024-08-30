package services

type ContainerInfo struct {
	ID      string `json:"Id"`
	Names   []string
	Image   string
	Command string
	Created int64
	Ports   []Port
	Status  string
}

type Port struct {
	IP          string `json:"ip"`
	PrivatePort uint16 `json:"private_port"`
	PublicPort  uint16 `json:"public_port"`
}
