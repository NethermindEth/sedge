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
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
)

var reAddr = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

/*
SkipLines :
Skips lines starting with `symbol`

params :-
a. content string
String to be processed
b. symbol string
Symbol to be skipped

returns :-
a. string
Processed string
*/
func SkipLines(content string, symbol string) string {
	lines := strings.Split(content, "\n")
	var trimmedLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, symbol) {
			continue
		}
		trimmedLines = append(trimmedLines, line)
	}
	return strings.Join(trimmedLines, "\n")
}

/*
Contains :
Checks if a string slice contains a string

params :-
a. list []string
String slice to be checked
b. str string
String to be checked

returns :-
a. bool
True if str is in list, false otherwise
*/
func Contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

/*
ContainsOnly :
Checks if a string slice contains only strings of a given string slice

params :-
a. list []string
String slice to be checked
b. target []string
String slice to be checked

returns :-
a. bool
True if every string in list is in target, false otherwise
*/
func ContainsOnly(list []string, target []string) bool {
	for _, s := range list {
		if !Contains(target, s) {
			return false
		}
	}
	return true
}

/*
IsAddress :
Checks if a string is an Ethereum address

params :-
a. a string
String to be checked

returns :-
a. bool
True if <a> is a valid Ethereum address. False otherwise
*/
func IsAddress(a string) bool {
	return reAddr.MatchString(a)
}

/*
AssingPorts :
Checks if port is occupied in a given host

params :-
a. host string
Host which port is to be checked
b. port string
Port to be checked

returns :-
a. bool
True if <port> is available. False otherwise
*/
func AssingPorts(host string, defaults map[string]string) (ports map[string]string, err error) {
	ports = make(map[string]string)
	mask := make(map[string]bool)

	for k, v := range defaults {
		if v == "" {
			return ports, fmt.Errorf(configs.DefaultPortEmptyError, k)
		}

		for !portAvailable(host, v, time.Second*5) || mask[v] {
			i, err := strconv.Atoi(v)
			if err != nil {
				return ports, err
			}
			v = strconv.Itoa(i + 1)
		}
		ports[k] = v
		mask[v] = true
	}

	return
}

/*
portAvailable :
Checks if port is occupied in a given host

params :-
a. host string
Host which port is to be checked
b. port string
Port to be checked

returns :-
a. bool
True if <port> is available. False otherwise
*/
func portAvailable(host, port string, timeout time.Duration) bool {
	log.Debugf("Checking port ocuppation of %s\n", net.JoinHostPort(host, port))
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		log.Debugf("Port seems available, got connecting error: %v", err)
		return true
	}
	if conn != nil {
		defer conn.Close()
		log.Debugf("Port open at %s", net.JoinHostPort(host, port))
	}
	return conn == nil
}

/*
Filter :
Filter a slice given a predicate

params :-
a. list []K
List to be filtered
b. filter func(K) bool
Predicate to be applied to each element of the list

returns :-
a. []K
Filtered list
*/
func Filter[K comparable](list []K, filter func(K) bool) []K {
	n := 0
	for _, v := range list {
		if filter(v) {
			list[n] = v
			n++
		}
	}

	return list[:n]
}
