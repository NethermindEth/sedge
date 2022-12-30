/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package clients

import (
	"fmt"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/templates"
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
	files, err := templates.Envs.ReadDir(strings.Join([]string{"envs", c.Network, clientType}, "/"))
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
		for _, client := range supportedClients {
			// Check if the client is supported
			//supported := utils.Contains(supportedClients, client)
			clients[clientType][client] = Client{Name: client, Type: clientType, Supported: true}
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
