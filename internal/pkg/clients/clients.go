package clients

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/utils"
	"github.com/NethermindEth/1click/templates"
	log "github.com/sirupsen/logrus"
)

// ClientInfo : Struct Interface for listing available clients
type ClientInfo struct {
	Network string
}

/*
SupportedClients :
Get supported client names of type <clientType> for network <network>. A client is supported if it has a docker-compose service template

params :-
a. clientType string
Type of client to be returned
b .network string
Target network

returns :-
a. []string
List of supported clients names of type <clientType>
b. error
Error if any
*/
func (c ClientInfo) SupportedClients(clientType string) (clientsNames []string, err error) {
	files, err := templates.Envs.ReadDir(filepath.Join("envs", c.Network, clientType))
	if err != nil {
		return
	}

	for _, file := range files {
		clientsNames = append(clientsNames, strings.TrimSuffix(file.Name(), ".tmpl"))
	}

	return clientsNames, nil
}

/*
Clients :
Get all the supported and configured clients

params :-
a. clientTypes []string
Types of client supported. E.g execution, consensus, validator
b. network
Target network

returns :-
a. OrderedClients
Map of <clientType>: map of <clientName>: Client
b. []error
List of errors
*/
func (c ClientInfo) Clients(clientTypes []string) (clients OrderedClients, errs []error) {
	clients = make(OrderedClients)

	for _, clientType := range clientTypes {
		clients[clientType] = make(ClientMap)
		// Get the clients with a docker-compose service template
		supportedClients, err := c.SupportedClients(clientType)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		log.Debugf(configs.SupportedClients, clientType, strings.ToLower(strings.Join(supportedClients, ", ")))

		// Get the clients from the configuration file
		configClients, err := configs.ConfigClients(clientType)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		log.Debugf(configs.ConfigClientsMsg, clientType, strings.ToLower(strings.Join(configClients, ", ")))

		for _, client := range configClients {
			// Check if the client is supported
			supported := utils.Contains(supportedClients, client)
			clients[clientType][client] = Client{Name: client, Type: clientType, Supported: supported}
		}
	}

	return
}

/*
ValidateClient :
Validate if the client is supported and configured

params :-
a. client Client
Client to be validated

returns :-
a. error
Error if client is not supported or configured
*/
func ValidateClient(client Client, currentType string) error {
	if client.Type == "" {
		return fmt.Errorf(configs.IncorrectClientError, currentType, client.Name)
	}

	if !client.Supported {
		return fmt.Errorf(configs.ClientNotSupported, client.Name)
	}
	return nil
}
