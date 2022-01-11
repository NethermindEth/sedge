package utils

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

//WIP
func InstallDependencies(dependencies []string) error {
	switch os := runtime.GOOS; os {
	case "linux":
		return fmt.Errorf("dependencies are not installed on your machine. Please install them and try again")
	case "windows":
		return fmt.Errorf("dependencies are not installed on your machine. Please install them and try again")
	default:
		log.Fatalf("Dependencies %s are not installed on your machine. Please install them and try again.", dependencies)
	}
	return nil
}
