package env

import (
	"path/filepath"
	"regexp"

	"github.com/NethermindEth/1click/templates"
)

/*
CheckVariable :
Check whatever a variable exist in a .env template file.

params :-
a. re regexp.Regexp
Regular expression to be used for matchings
b. network string
Target network
c. clientType string
Type of the client, e.g execution, consensus, validator
d. client string
Client's name

returns :-
a. bool
True if variable exists in <client>'s .env template for <network>
b. error
Error if any
*/
func CheckVariable(re *regexp.Regexp, network, clientType, client string) (bool, error) {
	content, err := templates.Envs.ReadFile(filepath.Join("envs", network, clientType, client+".tmpl"))
	if err != nil {
		return false, err
	}

	if m := re.FindStringSubmatch(string(content)); m != nil {
		return true, nil
	}

	return false, nil
}
