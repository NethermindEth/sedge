package utils

import (
	"bufio"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/NethermindEth/1click/configs"
	log "github.com/sirupsen/logrus"
)

var (
	reID        = regexp.MustCompile(`^ID=(.*)$`)
	reVersionID = regexp.MustCompile(`^VERSION_ID=(.*)$`)
)

/*
GetOSInfo :
Gather information like architecture
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
	file := "/etc/os-release"

	f, err := os.Open(file)
	if err != nil {
		return
	}

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			log.Errorf(configs.ClosingFileError, file)
			err = cerr
		}
	}()

	// Get the distro name
	s := bufio.NewScanner(f)
	for s.Scan() {
		if m := reID.FindStringSubmatch(s.Text()); m != nil {
			distro.Name = strings.Trim(m[1], `"`)
		} else if m := reVersionID.FindStringSubmatch(s.Text()); m != nil {
			distro.Version = strings.Trim(m[1], `"`)
		}
	}

	return distro, nil
}
