package generate

import (
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/pkg/clients"
	"github.com/NethermindEth/1Click/templates"
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
	rawTmp, err := templates.Config.ReadFile("config/config.tmpl")
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
		supportedClients, err := clients.GetSupportedClients(clientType)
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
