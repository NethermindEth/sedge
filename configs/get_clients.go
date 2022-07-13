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
