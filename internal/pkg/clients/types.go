package clients

// Client : Struct Represent a client like geth, prysm, etc
type Client struct {
	Name      string
	Type      string
	Supported bool
}

// Client : Struct Represent a combination of execution, consensus and validator clients
type Clients struct {
	Execution Client
	Consensus Client
	Validator Client
}

type ClientMap map[string]Client

type OrderedClients map[string]ClientMap
