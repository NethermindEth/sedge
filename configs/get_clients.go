package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

/*
ConfigClients :
Get the client names of type <clientType> from configuration file

params :-
a. clientType string
Type of client to be returned

returns :-
a. []string
List of clients names of type <clientType>
*/
func ConfigClients(clientType string) ([]string, error) {
	clients := viper.GetStringSlice(clientType + "Clients")
	if len(clients) == 0 {
		return nil, fmt.Errorf(NoClientsFoundError, clientType)
	}
	return strings.Split(strings.ToLower(strings.Join(clients, ",")), ","), nil
}
