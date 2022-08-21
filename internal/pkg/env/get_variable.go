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
	"path/filepath"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/templates"
	log "github.com/sirupsen/logrus"
)

/*
GetBootnodes :
Get the bootnodes (list of enr addresses) from the environment variable.

params :-
a. network string
Target network
b. client string
Client's name

returns :-
a. []string
List of bootnodes
b. error
Error if any
*/
func GetBootnodes(network, client string) ([]string, error) {
	content, err := templates.Envs.ReadFile(filepath.Join("envs", network, "consensus", client+".tmpl"))
	if err != nil {
		return nil, err
	}

	if m := ReBOOTNODES.FindStringSubmatch(string(content)); m != nil {
		m[1] = strings.ReplaceAll(m[1], "\"", "")
		enrs := strings.Split(m[1], ",")
		return enrs, nil
	}

	log.Warnf(configs.NoBootnodesFound, network, "consensus", client)
	return nil, nil
}
