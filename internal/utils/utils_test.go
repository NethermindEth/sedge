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
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSkipLines(t *testing.T) {
	inputs := []struct {
		content string
		symbol  string
		result  string
	}{
		{},
		{"bbbb", "a", "bbbb"},
		{"aaaa", "a", ""},
		{"bbbb\naaaa", "a", "bbbb"},
		{"aaaa\nbbbb", "a", "bbbb"},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("SkipLines(%s, %s)", input.content, input.symbol)
		got := SkipLines(input.content, input.symbol)

		if got != input.result {
			t.Errorf("%s expected %s but got %s", descr, input.result, got)
		}
	}
}

func TestContains(t *testing.T) {
	inputs := [...]struct {
		list []string
		str  string
		want bool
	}{
		{},
		{[]string{"a", "A", "b", "B"}, "b", true},
		{[]string{"a", "A", "b", "B"}, "c", false},
		{[]string{}, "a", false},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("Contains(%s)", input.list)
		if got := Contains(input.list, input.str); got != input.want {
			t.Errorf("%s expected %t but got %t", descr, input.want, got)
		}
	}
}

func TestContainsOnly(t *testing.T) {
	inputs := [...]struct {
		list   []string
		target []string
		want   bool
	}{
		{[]string{}, []string{}, true},
		{[]string{}, []string{"a"}, true},
		{[]string{"b"}, []string{"a", "A", "b", "B"}, true},
		{[]string{"c"}, []string{"a", "A", "b", "B"}, false},
		{[]string{"a"}, []string{}, false},
		{[]string{"execution", "validator"}, []string{"execution", "consensus"}, false},
		{[]string{"execution", "consensus", "validator"}, []string{"execution", "consensus"}, false},
		{[]string{"execution", "validator"}, []string{"execution", "consensus", "validator"}, true},
		{[]string{"execution", "consensus", "validator"}, []string{"execution", "consensus", "validator"}, true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("Contains(%s)", input.list)
		if got := ContainsOnly(input.list, input.target); got != input.want {
			t.Errorf("%s expected %t but got %t", descr, input.want, got)
		}
	}
}

func TestIsAddress(t *testing.T) {
	tcs := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"2131", false},
		{"dasd31gsd1231", false},
		{"0x2312313aaef2312312", false},
		{"0x5c00ABEf07604C59Ac72E859E5F93D5abZXCVF83", false},
		{"5c00ABEf07604C59Ac72E859E5F93D5ab8546F83", false},
		{"0x5c00ABEf07604C59Ac72E859E5F93D5ab8546F83", true},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("IsAddress(%s)", tc.input), func(t *testing.T) {
			if got := IsAddress(tc.input); got != tc.want {
				t.Errorf("got != want. Expected %v, got %v", tc.want, tc.input)
			}
		})
	}
}

func TestPortAvailable(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
	defer server.Close()
	split := strings.Split(server.URL, ":")
	host, port := split[1][2:], split[2]

	tcs := []struct {
		name string
		host string
		port string
		want bool
	}{
		{
			"Test case 1, good host and unavailable port",
			host, port,
			false,
		},
		{
			"Test case 2, bad host and port",
			"b@dh0$t", port,
			true,
		},
		{
			"Test case 3, good host and bad port",
			host, "666666666",
			true,
		},
		{
			"Test case 4, good host and available port",
			"localhost", "9999",
			true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := portAvailable(tc.host, tc.port, time.Millisecond*200); tc.want != got {
				t.Errorf("portAvailable(%s, %s, %s) failed; expected: %v, got: %v", tc.host, tc.port, "200ms", tc.want, got)
			}
		})
	}
}

func TestVerifyPortValid(t *testing.T) {
	tcs := []struct {
		name string
		port string
		want bool
	}{
		{
			"Test case 1, good port",
			"9999",
			true,
		},
		{
			"Test case 2, bad port",
			"b@dport",
			false,
		},
		{
			"Test case 3, empty port",
			"",
			false,
		},
		{
			"Test case 4, port too high",
			"68000",
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := VerifyPortValid(tc.port); tc.want != got {
				t.Errorf("VerifyPortValid(%s) failed; expected: %v, got: %v", tc.port, tc.want, got)
			}
		})
	}
}

func TestAssingPorts(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
	defer server.Close()
	split := strings.Split(server.URL, ":")
	host, port := split[1][2:], split[2]
	portN, _ := strconv.Atoi(port)

	tcs := []struct {
		name     string
		host     string
		defaults map[string]string
		want     map[string]string
		isErr    bool
	}{
		{
			"Test case 1, good host and defaults",
			host,
			map[string]string{"EL": "8545", "CL": port},
			map[string]string{"EL": "8545", "CL": strconv.Itoa(portN + 1)},
			false,
		},
		{
			"Test case 2, good host and bad defaults",
			host,
			map[string]string{"EL": "8545", "CL": ""},
			map[string]string{},
			true,
		},
		{
			"Test case 3, good host and bad defaults",
			host,
			map[string]string{"CL": "", "EL": "8545"},
			map[string]string{},
			true,
		},
		{
			"Test case 4, bad host and good defaults",
			"b@dh0$t",
			map[string]string{"CL": "9000", "EL": "8545"},
			map[string]string{"CL": "9000", "EL": "8545"},
			false,
		},
		{
			"Test case 5, good host and succesive increments",
			host,
			map[string]string{"CL": port, "EL": strconv.Itoa(portN + 1)},
			map[string]string{"CL": strconv.Itoa(portN + 1), "EL": strconv.Itoa(portN + 2)},
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := AssingPorts(tc.host, tc.defaults)

			descr := fmt.Sprintf("AssingPorts(%s, %+v)", tc.host, tc.defaults)
			if cerr := CheckErr(descr, tc.isErr, err); cerr != nil {
				t.Error(cerr)
			}

			if err == nil {
				for k := range tc.want {
					if tc.want[k] != got[k] {
						t.Errorf("A mismatch in the result has been found. Expected (key: %s, value: %s); got (key: %s, value %s). Call: %s. Expected object: %+v, Got: %+v", k, tc.want[k], k, got[k], descr, tc.want, got)
					}
				}
			}
		})
	}
}
