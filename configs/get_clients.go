package configs

import "github.com/spf13/viper"

/*
GetClients :
This function is responsible for giving the clients of type <clientType> from configuration file

params :-
a. clientType :- string
Type of client to be returned

returns :-
a. []string
List of clients of type <clientType>
*/
func GetClients(clientType string) []string {
	return viper.GetStringSlice(clientType)
}
