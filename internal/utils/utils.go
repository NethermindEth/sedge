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
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
)

var (
	reAddr     = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	regexEnode = regexp.MustCompile(`^enode:\/\/[0-9a-fA-F]{128}@.*:[1-9][0-9]*$`)
	regexEnr   = regexp.MustCompile(`enr:-.*$`)
)

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
AssignPorts :
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
func AssignPorts(host string, defaults map[string]uint16) (ports map[string]uint16, err error) {
	ports = make(map[string]uint16)
	mask := make(map[uint16]bool)

	keys := make([]string, 0, len(defaults))
	for k := range defaults {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := defaults[k]
		if v == 0 {
			return ports, fmt.Errorf(configs.DefaultPortInvalidError, k)
		}
		for !portAvailable(host, v) || mask[v] {
			v = v + 1
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
func portAvailable(ip string, port uint16) bool {
	log.Debugf("checking occupation of %s:%d", ip, port)
	netIp := net.ParseIP("127.0.0.1")
	if ip != "localhost" && ip != "0.0.0.0" {
		netIp = net.ParseIP(ip)
		if netIp == nil {
			log.Debugf("invalid host ip address")
			return true
		}
	}
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", netIp.String(), port))
	if err != nil {
		log.Debugf("error checking  %s:%d occupation: %v", ip, port, err)
		return false
	}
	sock.Close()
	return true
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
func Filter[K any](list []K, filter func(K) bool) []K {
	n := 0
	for _, v := range list {
		if filter(v) {
			list[n] = v
			n++
		}
	}

	return list[:n]
}

// UriValidator validates a URI and returns true if it is valid.
func UriValidator(input []string) (string, bool) {
	for _, uri := range input {
		u, err := url.Parse(uri)
		if err != nil {
			return uri, false
		}
		wrongScheme := u.Scheme != "http" && u.Scheme != "https"

		if wrongScheme || u.Host == "" {
			return uri, false
		}
	}
	return "", true
} // TODO: Add tests to avoid regression (partially covered by Generate cmd tests with good test cases)

// ENodesValidator validates a list of EL boot nodes and returns an error if any
// of them is invalid.
func ENodesValidator(bootNodes []string) error {
	set := make(map[string]struct{})
	for _, bootNode := range bootNodes {
		if _, ok := set[bootNode]; ok {
			return fmt.Errorf("%s: %s", configs.ErrDuplicatedBootNode, bootNode)
		}
		if !regexEnode.MatchString(bootNode) {
			return fmt.Errorf(configs.InvalidEnode, bootNode)
		}
		set[bootNode] = struct{}{}
	}
	return nil
} // TODO: Add tests to avoid regression (partially covered by Generate cmd tests with good test cases)

// ENRValidator validates a list of CL boot nodes and returns an error if any
// of them is invalid.
func ENRValidator(bootNodes []string) error {
	set := make(map[string]struct{})
	for _, bootNode := range bootNodes {
		if _, ok := set[bootNode]; ok {
			return fmt.Errorf("%s: %s", configs.ErrDuplicatedBootNode, bootNode)
		}
		if !regexEnr.MatchString(bootNode) {
			return fmt.Errorf(configs.InvalidEnr, bootNode)
		}
		set[bootNode] = struct{}{}
	}
	return nil
} // TODO: Add tests to avoid regression (partially covered by Generate cmd tests with good test cases)
