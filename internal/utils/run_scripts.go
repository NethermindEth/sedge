package utils

import (
	"os"
	"os/exec"
)

// Run docker-compose scripts
func RunDockerCompose(path string) error {
	cmd := exec.Command("docker-compose", "-f", path, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
