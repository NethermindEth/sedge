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
