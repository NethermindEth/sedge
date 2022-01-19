package utils

import (
	"os"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateDockerComposeScripts :
This function is responsible for generating docker-compose scripts for execution, consensus and
validator clients.

params :-
a. executionClient string
Execution client whose script is to be generated
b. consensusClient string
Execution client whose script is to be generated
c. validatorClient string
Execution client whose script is to be generated
d. path string
Path of generated scripts

returns :-
None
*/
func GenerateDockerComposeScripts(executionClient, consensusClient, validatorClient, path string) {
	executionTmp := getTemplate(executionClient, "execution")
	consensusTmp := getTemplate(consensusClient, "consensus")
	validatorTmp := getTemplate(validatorClient, "validator")

	err := writeTemplateToFile(executionTmp, path+"/docker-compose.execution.yml", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}

	err = writeTemplateToFile(consensusTmp, path+"/docker-compose.consensus.yml", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}

	err = writeTemplateToFile(validatorTmp, path+"/docker-compose.validator.yml", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}
}

/*
GenerateEnvFile :
This function is responsible for generating the environment variable for the
generated docker-compose scripts for execution, consensus and
validator clients.

params :-
a. executionClient string
Execution client whose script was generated
b. consensusClient string
Execution client whose script was generated
c. validatorClient string
Execution client whose script was generated
d. path string
Path of generated scripts

returns :-
None
*/
func GenerateEnvFile(executionClient, consensusClient, validatorClient, path string) {
	executionEnvTmp := getTemplate(executionClient, "env")
	consensusEnvTmp := getTemplate(consensusClient+"_"+"consensus", "env")
	validatorEnvTmp := getTemplate(validatorClient+"_"+"validator", "env")

	consensusEnv := templates.ConsensusEnv{
		ExecutionNodeURL: configs.OnPremiseExecutionURL,
	}

	validatorEnv := templates.ValidatorEnv{
		ConsensusNodeURL:    configs.OnPremiseConsensusURL,
		ExecutionEngineName: executionClient,
	}

	err := writeTemplateToFile(executionEnvTmp, path+"/.env", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}

	err = writeTemplateToFile(consensusEnvTmp, path+"/.env", consensusEnv, true)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}

	err = writeTemplateToFile(validatorEnvTmp, path+"/.env", validatorEnv, true)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}
}

/*
getTemplate :
Gets template given client and template type.

params :-
a. client string
Name of client. E.g geth, lighthouse
b. tmpType string
Kind of the template to be generated. Supported values are "execution", "consensus", "validator" and "env".

returns :-
a. *template.Template
*/
func getTemplate(client, tmpType string) (tmp *template.Template) {
	templates := templates.GetRawTemplates(tmpType)
	if templates == nil {
		log.Fatalf(configs.GetRawTemplatesError, tmpType)
	}

	rawTmp, ok := templates[client]
	if !ok {
		log.Fatalf(configs.ClientNotSupported, client)
	}

	tmp = template.Must(template.New(tmpType).Parse(rawTmp))
	return
}

/*
writeTemplateToFile :
Write template to `file`.

params :-
a. template *template.Template
Template to be written
b. file string
File's complete path
c. data interface{}
Data object to be applied to `template`
d. append bool
True to append the template to `file`. False to create it or overwrite it.

returns :-
a. err error
Error if any
*/
func writeTemplateToFile(template *template.Template, file string, data interface{}, append bool) (err error) {
	var f *os.File

	if append {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Errorf(configs.CreatingFileError, file, err)
			return
		}
	} else {
		f, err = os.Create(file)
		if err != nil {
			log.Errorf(configs.CreatingFileError, file, err)
			return
		}
	}

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			log.Errorf(configs.ClosingFileError, file)
			err = cerr
		}
	}()

	err = template.Execute(f, data)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof(configs.CreatedFile, file)

	return nil
}
