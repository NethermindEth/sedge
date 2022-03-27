package generate

import (
	"fmt"
	"os"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateScripts :
This function is responsible for generating docker-compose files for execution, consensus and
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
a. error
Error if any
*/
func GenerateScripts(executionClient, consensusClient, validatorClient, path string) (err error) {
	// Create scripts directory if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	log.Info(configs.GeneratingDockerComposeScript)
	err = generateDockerComposeScripts(executionClient, consensusClient, validatorClient, path)
	if err != nil {
		return err
	}

	log.Info(configs.GeneratingEnvFile)
	err = generateEnvFile(executionClient, consensusClient, validatorClient, path)
	if err != nil {
		return err
	}

	return nil
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
a. error
Error if any
*/
func generateDockerComposeScripts(executionClient, consensusClient, validatorClient, path string) (err error) {
	rawBaseTmp, err := templates.Services.ReadFile("services/docker-compose_base.tmpl")
	if err != nil {
		return
	}

	baseTmp, err := template.New("docker-compose").Parse(string(rawBaseTmp))
	if err != nil {
		return
	}

	clients := map[string]string{
		"execution/": executionClient,
		"consensus/": consensusClient,
		"validator/": validatorClient,
	}
	for tmpKind, clientName := range clients {
		tmp, err := templates.Services.ReadFile("services/" + tmpKind + clientName + ".tmpl")
		if err != nil {
			return err
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}

	// Print docker-compose file
	log.Infof(configs.PrintingFile, configs.DefaultDockerComposeScriptName)
	err = baseTmp.Execute(os.Stdout, nil)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, configs.DefaultDockerComposeScriptName, err)
	}
	fmt.Println()

	err = writeTemplateToFile(baseTmp, path+"/"+configs.DefaultDockerComposeScriptName, nil, false)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient, err)
	}
	log.Infof(configs.CreatedFile, path+"/"+configs.DefaultDockerComposeScriptName)

	return nil
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
a. error
Error if any
*/
func generateEnvFile(executionClient, consensusClient, validatorClient, path string) (err error) {
	executionEnvTmp, err := template.ParseFS(templates.Envs, "envs/execution/"+executionClient+".tmpl")
	if err != nil {
		return
	}
	consensusEnvTmp, err := template.ParseFS(templates.Envs, "envs/consensus/"+consensusClient+".tmpl")
	if err != nil {
		return
	}
	validatorEnvTmp, err := template.ParseFS(templates.Envs, "envs/validator/"+validatorClient+".tmpl")
	if err != nil {
		return
	}

	consensusEnv := ConsensusEnv{
		ExecutionNodeURL: configs.OnPremiseExecutionURL,
	}

	validatorEnv := ValidatorEnv{
		ConsensusNodeURL:    configs.OnPremiseConsensusURL,
		ExecutionEngineName: executionClient,
		KeystoreDir:         configs.KeystoreDefaultDataDir,
		DataDir:             configs.ValidatorDefaultDataDir,
	}

	err = writeTemplateToFile(executionEnvTmp, path+"/.env", nil, false)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient, err)
	}

	err = writeTemplateToFile(consensusEnvTmp, path+"/.env", consensusEnv, true)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient, err)
	}

	err = writeTemplateToFile(validatorEnvTmp, path+"/.env", validatorEnv, true)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, executionClient, consensusClient, validatorClient, err)
	}

	// Print .env file
	log.Infof(configs.PrintingFile, ".env")
	err = executionEnvTmp.Execute(os.Stdout, nil)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, ".env", err)
	}

	err = consensusEnvTmp.Execute(os.Stdout, consensusEnv)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, ".env", err)
	}

	err = validatorEnvTmp.Execute(os.Stdout, validatorEnv)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, ".env", err)
	}
	fmt.Println()
	log.Infof(configs.CreatedFile, path+"/.env")

	return nil
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
			return fmt.Errorf(configs.CreatingFileError, file, err)
		}
	} else {
		f, err = os.Create(file)
		if err != nil {
			return fmt.Errorf(configs.CreatingFileError, file, err)
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
		return
	}

	return nil
}
