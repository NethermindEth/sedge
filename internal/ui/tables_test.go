package ui

import (
	"bytes"
	"testing"

	"github.com/alexeyco/simpletable"
)

func TestWriteListClientsTable(t *testing.T) {
	fdOut := new(bytes.Buffer)

	inputs := []ListClientsTable{
		{},
		{Clients: [][]string{{}, {}, {}}},
		{Clients: [][]string{{"a"}, {"b"}, {}}},
		{ClientTypes: []string{"A", "C"}, Clients: [][]string{{"a", "b"}, {"c"}}},
	}

	for _, input := range inputs {
		WriteListClientsTable(fdOut, &input)
	}
}

func TestWriteRandomizedClientsTable(t *testing.T) {
	fdOut := new(bytes.Buffer)

	input := []RandomizedClientsTable{
		{},
		{ClientTypes: []string{}},
		{Clients: []string{}},
		{ClientTypes: []string{}, Clients: []string{}},
		{ClientTypes: []string{"A", "B"}, Clients: []string{"a", "b", "c"}},
		{ClientTypes: []string{"A", "B", "C"}, Clients: []string{"a", "b"}},
	}

	for _, input := range input {
		WriteRandomizedClientsTable(fdOut, input)
	}
}

func TestWriteSimpleTable(t *testing.T) {
	fdOut := new(bytes.Buffer)

	inputs := []SimpleTableData{
		{
			Headers: []*simpletable.Cell{},
			Columns: [][]*simpletable.Cell{
				{
					{Text: "A"},
				},
			},
		},
		{
			Headers: []*simpletable.Cell{},
			Columns: [][]*simpletable.Cell{
				{
					{Text: "A"},
				},
			},
			Enumerate: true,
		},
		{
			Headers: []*simpletable.Cell{
				{Text: "A"},
				{Text: "B"},
			},
			Columns: [][]*simpletable.Cell{},
		},
		{
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
		{
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
	}

	for _, input := range inputs {
		WriteSimpleTable(fdOut, &input)
	}
}
