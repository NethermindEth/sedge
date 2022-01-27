package clients

import (
	"crypto/rand"
	"math/big"
)

/*
RandomChoice :
Select a random element from a ClientMap

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
	list := make([]Client, 0)
	for _, client := range clients {
		list = append(list, client)
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return
	}
	return list[n.Int64()], nil
}
