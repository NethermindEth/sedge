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
package monitoring

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

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
			got, err := assignPorts(tc.host, tc.defaults)

			descr := fmt.Sprintf("AssingPorts(%s, %+v)", tc.host, tc.defaults)
			if cerr := checkErr(descr, tc.isErr, err); cerr != nil {
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

func checkErr(descr string, isErr bool, err error) error {
	l := err == nil && isErr
	r := err != nil && !isErr

	if l || r {
		return fmt.Errorf("%s failed, unexpected error value: %v", descr, err)
	}
	return nil
}
