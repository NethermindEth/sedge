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
	_, baseTmp := buildTemplate("docker-compose_base", "base")
	_, executionTmp := buildTemplate(executionClient, "execution")
	_, consensusTmp := buildTemplate(consensusClient, "consensus")
	_, validatorTmp := buildTemplate(validatorClient, "validator")

	tmp := template.Must(template.New("docker-compose").Parse(baseTmp))
	template.Must(tmp.Parse(executionTmp))
	template.Must(tmp.Parse(consensusTmp))
	template.Must(tmp.Parse(validatorTmp))

	// Print docker-compose file
	log.Infof(configs.PrintingFile, "docker-compose.yml")
	tmp.Execute(os.Stdout, nil)

	err := writeTemplateToFile(tmp, path+"/docker-compose.yml", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}
	log.Infof(configs.CreatedFile, path+"/docker-compose.yml")
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
	executionEnvTmp, _ := buildTemplate(executionClient, "env")
	consensusEnvTmp, _ := buildTemplate(consensusClient+"_"+"consensus", "env")
	validatorEnvTmp, _ := buildTemplate(validatorClient+"_"+"validator", "env")

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

	// Print .env file
	log.Infof(configs.PrintingFile, ".env")
	executionEnvTmp.Execute(os.Stdout, nil)
	consensusEnvTmp.Execute(os.Stdout, consensusEnv)
	validatorEnvTmp.Execute(os.Stdout, validatorEnv)

	log.Infof(configs.CreatedFile, path+"/.env")
}

/*
buildTemplate :
Gets template given client and template type.

params :-
a. client string
Name of client. E.g geth, lighthouse
b. tmpType string
Kind of the template to be generated. Supported values are "execution", "consensus", "validator" and "env".

returns :-
a. *template.Template
*/
func buildTemplate(client, tmpType string) (tmp *template.Template, rawTmp string) {
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

	return nil
}
