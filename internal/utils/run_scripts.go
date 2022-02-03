package utils

import (
	"fmt"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	log "github.com/sirupsen/logrus"
)

/*
RunDockerCompose :
This function is responsible for running the generated doccker-compose script.

params :-
a. path string
Path of generated script

returns :-
a. error
Error if any
*/
func RunDockerCompose(path string) error {
	log.Info(configs.RunningDockerCompose)
	cmd := fmt.Sprintf(configs.DockerComposeCMD, path)
	tmp, err := template.New("script").Parse(cmd)
	if err != nil {
		return err
	}
	return executeScript(tmp)
}
