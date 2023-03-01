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
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/configs"
	"github.com/stretchr/testify/assert"
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
	host, strPort := split[1][2:], split[2]
	port64, err := strconv.ParseUint(strPort, 10, 16)
	if err != nil {
		t.Fatalf("cannot convert http server port: %v", err)
	}
	port := uint16(port64)

	tcs := []struct {
		name string
		host string
		port uint16
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
			host, 9999,
			true,
		},
		{
			"Test case 4, good host and available port",
			"localhost", 9999,
			true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := portAvailable(tc.host, tc.port); tc.want != got {
				t.Errorf("portAvailable(%s, %d) failed; expected: %v, got: %v", tc.host, tc.port, tc.want, got)
			}
		})
	}
}

func TestAssignPorts(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
	defer server.Close()
	split := strings.Split(server.URL, ":")
	host, strPort := split[1][2:], split[2]
	port64, err := strconv.ParseUint(strPort, 10, 16)
	if err != nil {
		t.Fatalf("cannot convert http server port: %v", err)
	}
	port := uint16(port64)

	tcs := []struct {
		name     string
		host     string
		defaults map[string]uint16
		want     map[string]uint16
		isErr    bool
	}{
		{
			"Test case 1, good host and defaults",
			host,
			map[string]uint16{"EL": 8545, "CL": port},
			map[string]uint16{"EL": 8545, "CL": port + 1},
			false,
		},
		{
			"Test case 2, good host and bad defaults",
			host,
			map[string]uint16{"EL": 8545, "CL": 0},
			map[string]uint16{},
			true,
		},
		{
			"Test case 3, good host and bad defaults",
			host,
			map[string]uint16{"CL": 0, "EL": 8545},
			map[string]uint16{},
			true,
		},
		{
			"Test case 4, bad host and good defaults",
			"b@dh0$t",
			map[string]uint16{"CL": 9000, "EL": 8545},
			map[string]uint16{"CL": 9000, "EL": 8545},
			false,
		},
		{
			"Test case 5, good host and successive increments",
			host,
			map[string]uint16{"CL": port, "EL": port + 1},
			map[string]uint16{"CL": port + 1, "EL": port + 2},
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := AssignPorts(tc.host, tc.defaults)

			descr := fmt.Sprintf("AssingPorts(%s, %+v)", tc.host, tc.defaults)
			if cerr := CheckErr(descr, tc.isErr, err); cerr != nil {
				t.Error(cerr)
			}

			if err == nil {
				for k := range tc.want {
					if tc.want[k] != got[k] {
						t.Errorf("A mismatch in the result has been found. Expected (key: %s, value: %d); got (key: %s, value %d). Call: %s. Expected object: %+v, Got: %+v", k, tc.want[k], k, got[k], descr, tc.want, got)
					}
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tcs := []struct {
		name   string
		in     []string
		want   []string
		filter func(string) bool
	}{
		{
			"Test case 1, no filter",
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
			func(s string) bool {
				return true
			},
		},
		{
			"Test case 2, filter",
			[]string{"a", "b", "c"},
			[]string{"a", "c"},
			func(s string) bool {
				return s != "b"
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Filter(tc.in, tc.filter)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Filter(%+v) failed; expected: %+v, got: %+v", tc.in, tc.want, got)
			}
		})
	}
}

func TestENodesValidator(t *testing.T) {
	tcs := []struct {
		name      string
		in        []string
		withError bool
		contains  string
	}{
		{
			"Good boot nodes",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",
				"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
			},
			false,
			"",
		},
		{
			"Repeated boot node",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
				"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
			},
			true,
			configs.ErrDuplicatedBootNode,
		},
		{
			"Invalid scheme",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enide://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
			},
			true,
			fmt.Sprintf(configs.InvalidEnode, "enide://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303"),
		},
		{
			"Invalid public key, too short",
			[]string{
				"enode://5fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",
				"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
			},
			true,
			fmt.Sprintf(configs.InvalidEnode, "enode://5fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303"),
		},
		{
			"Invalid public key, too long",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de00000000000@3.209.45.79:30303",
			},
			true,
			fmt.Sprintf(configs.InvalidEnode, "enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de00000000000@3.209.45.79:30303"),
		},
		{
			"Invalid public key, invalid hex",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enode://z499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",
			},
			true,
			fmt.Sprintf(configs.InvalidEnode, "enode://z499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303"),
		},
		{
			"Invalid port",
			[]string{
				"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",
				"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30b303",
			},
			true,
			fmt.Sprintf(configs.InvalidEnode, "enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30b303"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := ENodesValidator(tc.in)
			if tc.withError {
				assert.Contains(t, got.Error(), tc.contains)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

func TestRelayURLsValidator(t *testing.T) {
	tcs := []struct {
		name       string
		in         []string
		invalidUri string
		valid      bool
	}{
		{
			"Valid relay URL",
			[]string{
				"https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com",
				"https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@builder-relay-goerli.flashbots.net",
				"https://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@bloxroute.max-profit.builder.goerli.blxrbdn.com",
				"https://0x8f7b17a74569b7a57e9bdafd2e159380759f5dc3ccbd4bf600414147e8c4e1dc6ebada83c0139ac15850eb6c975e82d0@builder-relay-goerli.blocknative.com",
				"https://0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12@relay-goerli.edennetwork.io",
				"https://0xb1559beef7b5ba3127485bbbb090362d9f497ba64e177ee2c8e7db74746306efad687f2cf8574e38d70067d40ef136dc@relay-stag.ultrasound.money",
			},
			"",
			true,
		},
		{
			"Invalid relay URL, invalid scheme",
			[]string{"htt://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com"},
			"htt://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com",
			false,
		},
		{
			"Invalid relay URL, without domain",
			[]string{"https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@"},
			"https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@",
			false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := UriValidator(tc.in)
			if tc.valid {
				assert.True(t, ok)
				assert.Empty(t, got)
			} else {
				assert.False(t, ok)
				assert.Equal(t, tc.invalidUri, got)
			}
		})
	}
}
