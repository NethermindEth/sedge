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
func RandomChoice(clients ClientMap) (client Client, err error) {
	if len(clients) == 0 {
		return client, errors.New(configs.EmptyClientMapError)
	}

	list := make([]Client, 0)
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
