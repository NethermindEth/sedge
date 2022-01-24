package utils

import (
	"fmt"
	"os"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/templates"
	log "github.com/sirupsen/logrus"
)

func GenerateScripts(executionClient, consensusClient, validatorClient, path string) {
	// Create scripts directory if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Info(configs.GeneratingDockerComposeScript)
	generateDockerComposeScripts(executionClient, consensusClient, validatorClient, path)

	log.Info(configs.GeneratingEnvFile)
	generateEnvFile(executionClient, consensusClient, validatorClient, path)
}

/*
generateDockerComposeScripts :
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
func generateDockerComposeScripts(executionClient, consensusClient, validatorClient, path string) error {
	baseTmp, err := templates.Services.ReadFile("services/docker-compose_base.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	executionTmp, err := templates.Services.ReadFile("services/execution/" + executionClient + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}
	consensusTmp, err := templates.Services.ReadFile("services/consensus/" + consensusClient + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}
	validatorTmp, err := templates.Services.ReadFile("services/validator/" + validatorClient + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmp := template.Must(template.New("docker-compose").Parse(string(baseTmp)))
	template.Must(tmp.Parse(string(executionTmp)))
	template.Must(tmp.Parse(string(consensusTmp)))
	template.Must(tmp.Parse(string(validatorTmp)))

	// Print docker-compose file
	log.Infof(configs.PrintingFile, "docker-compose.yml")
	tmp.Execute(os.Stdout, nil)
	fmt.Println()

	err = writeTemplateToFile(tmp, path+"/docker-compose.yml", nil, false)
	if err != nil {
		log.Fatalf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient)
	}
	log.Infof(configs.CreatedFile, path+"/docker-compose.yml")
}

/*
generateEnvFile :
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
func generateEnvFile(executionClient, consensusClient, validatorClient, path string) error {
	executionEnvTmp, err := template.ParseFS(templates.Envs, "envs/execution/"+executionClient+".tmpl")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	consensusEnvTmp, err := template.ParseFS(templates.Envs, "envs/consensus/"+consensusClient+".tmpl")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	validatorEnvTmp, err := template.ParseFS(templates.Envs, "envs/validator/"+validatorClient+".tmpl")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	consensusEnv := ConsensusEnv{
		ExecutionNodeURL: configs.OnPremiseExecutionURL,
	}

	validatorEnv := ValidatorEnv{
		ConsensusNodeURL:    configs.OnPremiseConsensusURL,
		ExecutionEngineName: executionClient,
	}

	err = writeTemplateToFile(executionEnvTmp, path+"/.env", nil, false)
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

	fmt.Println()
	log.Infof(configs.CreatedFile, path+"/.env")
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
