package compose

// ComposeService defines the structure of a service in the output of the 'docker compose ps' command.
type ComposeService struct {
	// Id is the ID of the container.
	Id string `json:"ID"`
	// Service is the name of the service.
	Service string `json:"Service"`
	// Name is the name of the container.
	Name string `json:"Name"`
	// State is the state of the container.
	State string `json:"State"`
}

// DockerComposeStopOptions defines the options for the 'docker compose stop' command.
type DockerComposeStopOptions struct {
	// Path specifies the location of the docker-compose.yaml file.
	Path string
}

