package utils

import (
	"fmt"
	"text/template"
)

// Run docker-compose scripts
func RunDockerCompose(path string) error {
	cmd := fmt.Sprintf("sudo docker-compose -f %s up -d", path)
	tmp, err := template.New("script").Parse(cmd)
	if err != nil {
		return err
	}
	return executeScript(tmp)
}
