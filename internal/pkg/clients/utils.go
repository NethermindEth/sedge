package clients

import (
	"crypto/rand"
	"math/big"
)

/*
RandomChoice :
Select a random element from a Client list

params :-
a. list [].Client
Target list

returns :-
a. Client
Random element from list
b. error
Error if any
*/
func RandomChoice(list []Client) (client Client, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return
	}
	return list[n.Int64()], nil
}

/*
Select :
Select an element from a Client list following the predicate (n => n.Name == name)

params :-
a. list [].Client
Target list
b. clientName string
Name of the client to be selected

returns :-
a. Client
Selected element from list
*/
func Select(list []Client, clientName string) Client {
	for _, client := range list {
		if client.Name == clientName {
			return client
		}
	}
	return Client{Name: clientName, Type: ""}
}
