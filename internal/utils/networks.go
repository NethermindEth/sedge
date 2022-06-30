package utils

import (
	"github.com/NethermindEth/1click/templates"
)

/*
SupportedNetworks :
Get supported networks names. A network is supported if it has a folder with the network name in either templates/envs or templates/services forder.

params :-
none

returns :-
a. []string
List of supported network names
b. error
Error if any
*/
func SupportedNetworks() (networkNames []string, err error) {
	files, err := templates.Envs.ReadDir("envs")
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			networkNames = append(networkNames, file.Name())
		}
	}

	return networkNames, nil
}
