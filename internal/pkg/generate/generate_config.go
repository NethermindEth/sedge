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
package generate

import (
	"path/filepath"
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/templates"
)

/*
GenerateConfig :
This function is responsible for generating the default configuration file containing all
supported clients.

params :-
a. path string
Path of generated config file

returns :-
a. error
Error if any
*/
func GenerateConfig(path string) (err error) {
	rawTmp, err := templates.Config.ReadFile(filepath.Join("config", "config.tmpl"))
	if err != nil {
		return
	}

	tmp, err := template.New("config").Parse(string(rawTmp))
	if err != nil {
		return
	}

	// Get supported clients
	clientsMap := make(map[string][]string)
	for _, clientType := range []string{"execution", "consensus", "validator"} {
		c := clients.ClientInfo{Network: "mainnet"}
		supportedClients, err := c.SupportedClients(clientType)
		if err != nil {
			return err
		}
		clientsMap[clientType] = supportedClients
	}

	if err = writeTemplateToFile(tmp, path+"/"+configs.ConfigFileName+".yaml", clientsMap, false); err != nil {
		return
	}

	return nil
}
