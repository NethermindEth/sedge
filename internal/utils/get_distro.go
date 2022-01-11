package utils

import (
	"bufio"
	"os"
	"regexp"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

var reID = regexp.MustCompile(`^ID=(.*)$`)

/*
GetOSInfo :
This function is responsible for gathering information like architecture
and name of the linux distribution of the host machine.

params :-
None

returns :-
a. distro DistroInfo
DistroInfo struct containing name and architecture of the host machine
a. err error
Error if any
*/
func GetOSInfo() (distro DistroInfo, err error) {
	// Get the architecture
	distro.Architecture = runtime.GOARCH

	f, err := os.Open("/etc/os-release")
	if err != nil {
		return
	}

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			log.Errorf("Failed to close file /etc/os-release")
			err = cerr
		}
	}()

	// Get the distro name
	s := bufio.NewScanner(f)
	for s.Scan() {
		if m := reID.FindStringSubmatch(s.Text()); m != nil {
			distro.Name = strings.Trim(m[1], `"`)
		}
	}

	return distro, nil
}
