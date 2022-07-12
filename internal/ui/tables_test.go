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
package ui

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexeyco/simpletable"
)

func ensureResult(t *testing.T, fdOut *bytes.Buffer, outputDir string) {
	table, err := io.ReadAll(fdOut)
	if err != nil {
		t.Errorf("Can't read output: %v", err)
	}
	file, err := os.Open(outputDir)
	if err != nil {
		t.Fatalf("Can't read test case output: %v", err)
	}
	output, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Can't read test case output: %v", err)
	}
	if stable, soutput := string(table), string(output); stable != soutput {
		t.Errorf("Wrong output.\nExpected:\n%s\nBut got:\n%s", soutput, stable)
	}
}

type listClientsTableTestCase struct {
	name      string
	fdOut     *bytes.Buffer
	data      ListClientsTable
	outputDir string
}

func buildListClientsTableTestCase(
	t *testing.T,
	name string,
	data ListClientsTable,
	outputDir string,
) listClientsTableTestCase {
	tc := listClientsTableTestCase{}
	tc.name = name
	tc.fdOut = new(bytes.Buffer)
	tc.data = data
	tc.outputDir = filepath.Join("testdata", "table_tests", "clients_tables", outputDir, "output")
	return tc
}

func TestWriteListClientsTable(t *testing.T) {
	tcs := []listClientsTableTestCase{
		buildListClientsTableTestCase(
			t,
			"Empty v1",
			ListClientsTable{},
			"case_1",
		),
		buildListClientsTableTestCase(
			t,
			"Empty v2",
			ListClientsTable{Clients: [][]string{{}, {}, {}}},
			"case_2",
		),
		buildListClientsTableTestCase(
			t,
			"No types",
			ListClientsTable{Clients: [][]string{{"a"}, {"b"}, {}}},
			"case_3",
		),
		buildListClientsTableTestCase(
			t,
			"Ok",
			ListClientsTable{ClientTypes: []string{"A", "C"}, Clients: [][]string{{"a", "b"}, {"c"}}},
			"case_4",
		),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			WriteListClientsTable(tc.fdOut, &tc.data)
			ensureResult(t, tc.fdOut, tc.outputDir)
		})
	}
}

type randomizedClientsTableTestCase struct {
	name      string
	fdOut     *bytes.Buffer
	data      RandomizedClientsTable
	outputDir string
}

func buildRandomizedClientsTestCase(
	t *testing.T,
	name string,
	data RandomizedClientsTable,
	outputDir string,
) randomizedClientsTableTestCase {
	tc := randomizedClientsTableTestCase{}
	tc.name = name
	tc.data = data
	tc.outputDir = filepath.Join("testdata", "table_tests", "randomized_tables", outputDir, "output")
	tc.fdOut = new(bytes.Buffer)
	return tc
}

func TestWriteRandomizedClientsTable(t *testing.T) {
	tcs := []randomizedClientsTableTestCase{
		buildRandomizedClientsTestCase(
			t,
			"Nil",
			RandomizedClientsTable{},
			"case_1",
		),
		buildRandomizedClientsTestCase(
			t,
			"Nil clients",
			RandomizedClientsTable{ClientTypes: []string{}},
			"case_2",
		),
		buildRandomizedClientsTestCase(
			t,
			"Nil types",
			RandomizedClientsTable{Clients: []string{}},
			"case_3",
		),
		buildRandomizedClientsTestCase(
			t,
			"Empty",
			RandomizedClientsTable{ClientTypes: []string{}, Clients: []string{}},
			"case_4",
		),
		buildRandomizedClientsTestCase(
			t,
			"More clients than types",
			RandomizedClientsTable{ClientTypes: []string{"A", "B"}, Clients: []string{"a", "b", "c"}},
			"case_5",
		),
		buildRandomizedClientsTestCase(
			t,
			"More types than clients",
			RandomizedClientsTable{ClientTypes: []string{"A", "B", "C"}, Clients: []string{"a", "b"}},
			"case_6",
		),
	}

	for _, tc := range tcs {
		WriteRandomizedClientsTable(tc.fdOut, tc.data)
		ensureResult(t, tc.fdOut, tc.outputDir)
	}
}

type simpleTableTestCase struct {
	name      string
	fdOut     *bytes.Buffer
	data      SimpleTableData
	outputDir string
}

func buildSimpleTableTestCase(
	t *testing.T,
	name string,
	data SimpleTableData,
	outputDir string,
) simpleTableTestCase {
	tc := simpleTableTestCase{}
	tc.name = name
	tc.data = data
	tc.outputDir = filepath.Join("testdata", "table_tests", "simple_tables", outputDir, "output")
	tc.fdOut = new(bytes.Buffer)
	return tc
}

func TestWriteSimpleTable(t *testing.T) {
	tcs := []simpleTableTestCase{
		buildSimpleTableTestCase(
			t,
			"Empty headers",
			SimpleTableData{
				Headers: []*simpletable.Cell{},
				Columns: [][]*simpletable.Cell{
					{
						{Text: "A"},
					},
				},
			},
			"case_1",
		),
		buildSimpleTableTestCase(
			t,
			"Only enumerate",
			SimpleTableData{
				Headers: []*simpletable.Cell{},
				Columns: [][]*simpletable.Cell{
					{
						{Text: "A"},
					},
				},
				Enumerate: true,
			},
			"case_2",
		),
		buildSimpleTableTestCase(
			t,
			"Empty columns",
			SimpleTableData{
				Headers: []*simpletable.Cell{
					{Text: "A"},
					{Text: "B"},
				},
				Columns: [][]*simpletable.Cell{},
			},
			"case_3",
		),
		buildSimpleTableTestCase(
			t,
			"Ok",
			SimpleTableData{
				Headers: []*simpletable.Cell{
					{Text: "A"},
					{Text: "B"},
				},
				Columns: [][]*simpletable.Cell{
					{
						{Text: "a"},
						{Text: "b"},
					},
				},
			},
			"case_4",
		),
		buildSimpleTableTestCase(
			t,
			"Ok with enumerate",
			SimpleTableData{
				Headers: []*simpletable.Cell{
					{Text: "A"},
					{Text: "B"},
				},
				Columns: [][]*simpletable.Cell{
					{
						{Text: "a"},
					},
					{
						{Text: "b"},
					},
				},
			},
			"case_5",
		),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			WriteSimpleTable(tc.fdOut, &tc.data)
			ensureResult(t, tc.fdOut, tc.outputDir)
		})
	}
}

type listNetworksTableTestCase struct {
	name      string
	fdOut     *bytes.Buffer
	data      []string
	outputDir string
}

func buildListNetworksTestCase(
	t *testing.T,
	name string,
	data []string,
	outputDir string,
) listNetworksTableTestCase {
	tc := listNetworksTableTestCase{}
	tc.name = name
	tc.data = data
	tc.outputDir = filepath.Join("testdata", "table_tests", "network_tables", outputDir, "output")
	tc.fdOut = new(bytes.Buffer)
	return tc
}

func TestListNetworksTable(t *testing.T) {
	tcs := []listNetworksTableTestCase{
		buildListNetworksTestCase(
			t,
			"Empty",
			[]string{},
			"case_1",
		),
		buildListNetworksTestCase(
			t,
			"OK, one element",
			[]string{"A"},
			"case_2",
		),
		buildListNetworksTestCase(
			t,
			"OK, several elements",
			[]string{"A", "B", "C"},
			"case_3",
		),
	}

	for _, tc := range tcs {
		WriteListNetworksTable(tc.fdOut, tc.data)
		ensureResult(t, tc.fdOut, tc.outputDir)
	}
}
