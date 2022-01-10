package onpremise

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

//WIP
func InstallDependencies(dependencies []string) error {
	switch os := runtime.GOOS; os {
	case "linux":
		log.Error("Dependencies are not installed on your machine. Please install them and try again.")
	case "windows":
		log.Error("Dependencies are not installed on your machine. Please install them and try again.")
	default:
		log.Fatalf("Dependencies %s are not installed on your machine. Please install them and try again.", dependencies)
	}
	return nil
}
