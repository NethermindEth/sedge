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
	"os"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
)

/*
GetECBootnodes :
Get the execution bootnodes (list of enodes addresses) from the environment variables in .env.

params :-
a. path to generated env file

returns :-
a. []string
List of bootnodes
b. error
Error if any
*/
func GetECBootnodes(envFilePath string) ([]string, error) {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return nil, err
	}

	if m := ReElBOOTNODES.FindStringSubmatch(string(content)); m != nil {
		m[1] = strings.ReplaceAll(m[1], "\"", "")
		enodes := strings.Split(m[1], ",")
		for i, enode := range enodes {
			enodes[i] = strings.Trim(enode, "\r\n ")
		}
		return enodes, nil
	}

	log.Warnf(configs.NoBootnodesFound, envFilePath)
	return nil, nil
}

/*
GetCCBootnodes :
Get the consensus bootnodes (list of enr addresses) from the environment variables in .env.

params :-
a. path to generated env file

returns :-
a. []string
List of bootnodes
b. error
Error if any
*/
func GetCCBootnodes(envFilePath string) ([]string, error) {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return nil, err
	}

	if m := ReClBOOTNODES.FindStringSubmatch(string(content)); m != nil {
		m[1] = strings.ReplaceAll(m[1], "\"", "")
		enrs := strings.Split(m[1], ",")
		for i, enr := range enrs {
			enrs[i] = strings.Trim(enr, "\r\n ")
		}
		return enrs, nil
	}

	log.Warnf(configs.NoBootnodesFound, envFilePath)
	return nil, nil
}

/*
GetTTD :
Get TTD from the environment variables in .env.

params :-
a. path to generated env file

returns :-
a. []string
List of bootnodes
b. error
Error if any
*/
func GetTTD(envFilePath string) (string, error) {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return "", err
	}

	if m := ReTTD.FindStringSubmatch(string(content)); m != nil {
		m[1] = strings.ReplaceAll(m[1], "\"", "")
		return strings.Trim(m[1], "\r\n "), nil
	}

	log.Warnf(configs.NoBootnodesFound, envFilePath)
	return "", nil
}
