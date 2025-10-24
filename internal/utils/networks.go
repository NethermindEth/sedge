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
package utils

import (
	"github.com/NethermindEth/sedge/templates"
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
		return networkNames, err
	}

	for _, file := range files {
		if file.IsDir() {
			networkNames = append(networkNames, file.Name())
		}
	}

	return networkNames, nil
}
