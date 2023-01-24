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
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/NethermindEth/sedge/configs"
)

/*
RandomChoice :
Select a random supported client from a ClientMap

params :-
a. clients ClientMap
Target clients

returns :-
a. Client
Random element from list
b. error
Error if any
*/
func RandomChoice(clients ClientMap) (client *Client, err error) {
	if len(clients) == 0 {
		return client, errors.New(configs.EmptyClientMapError)
	}

	list := make([]*Client, 0)
	for _, client := range clients {
		if client.Supported {
			list = append(list, client)
		}
	}

	if len(list) == 0 {
		return client, errors.New(configs.NoSupportedClientsError)
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return
	}
	return list[n.Int64()], nil
}
