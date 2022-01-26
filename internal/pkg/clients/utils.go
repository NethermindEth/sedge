package clients

import "math/rand"

/*
RandomChoice :
Select a random element from a Client list

params :-
a. list [].Client
Target list

returns :-
a. Client
Random element from list
*/
func RandomChoice(list []Client) Client {
	item := list[rand.Intn(len(list))]
	return item
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
