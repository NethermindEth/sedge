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
package env

import (
	"regexp"
	"strings"

	"github.com/NethermindEth/sedge/templates"
)

/*
CheckVariable :
Check whatever a variable exist in the base .env template file.

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
	//TODO: change usage to check variables from generated config instead of template
	content, err := templates.Envs.ReadFile(strings.Join([]string{"envs", network, clientType, client + ".tmpl"}, "/"))
	if err != nil {
		return false, err
	}

	if m := re.FindStringSubmatch(string(content)); m != nil {
		return true, nil
	}

	return false, nil
}

/*
CheckVariableBase :
Check whatever a variable exist in the base .env template file.

params :-
a. re regexp.Regexp
Regular expression to be used for matchings
b. network string
Target network

returns :-
a. bool
True if variable exists in base .env template for <network>
b. error
Error if any
*/
func CheckVariableBase(re *regexp.Regexp, network string) (bool, error) {
	//TODO: change usage to check variables from generated config instead of template
	content, err := templates.Envs.ReadFile(strings.Join([]string{"envs", network, "env_base.tmpl"}, "/"))
	if err != nil {
		return false, err
	}

	if m := re.FindStringSubmatch(string(content)); m != nil {
		return true, nil
	}

	return false, nil
}
