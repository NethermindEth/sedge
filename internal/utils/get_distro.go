/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
)

var (
	reID        = regexp.MustCompile(`^ID=(.*)$`)
	reVersionID = regexp.MustCompile(`^VERSION_ID=(.*)$`)
)

/*
getOSInfo :
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
func getOSInfo() (distro DistroInfo, err error) {
	// Get the architecture
	distro.Architecture = runtime.GOARCH
	file := "/etc/os-release"

	f, err := os.Open(file)
	if err != nil {
		return distro, err
	}

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			// notest
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

/*
GetDistroName :
Get the name and version of the linux distribution of the host machine.

params :-
None

returns :-
a. string
"<Distribution_Name> <Distribution_Version>"
a. error
Error if any
*/
func GetDistroName() (_ string, err error) {
	var distro DistroInfo
	if distro, err = getOSInfo(); err != nil {
		return "", fmt.Errorf(configs.DistroInfoError, err)
	}
	return fmt.Sprintf("%s %s", distro.Name, distro.Version), nil
}
