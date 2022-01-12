package ui

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

/*
WriteListClientsTable :
Prints the supported clients table

params :-
a. data [][]string
Table data

returns :-
None
*/
func WriteListClientsTable(data [][]string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Execution Client"},
			{Align: simpletable.AlignCenter, Text: "Consensus Client"},
		},
	}

	for i, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignCenter, Text: interface{}(row[0]).(string)},
			{Align: simpletable.AlignCenter, Text: interface{}(row[1]).(string)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
	fmt.Println()
}

/*
WriteRandomizedClientsTable :
Prints the randomized clients table

params :-
a. data [][]string
Table data

returns :-
None
*/
func WriteRandomizedClientsTable(data [][]string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Type of Client"},
			{Align: simpletable.AlignCenter, Text: "Randomized Client"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: interface{}(row[0]).(string)},
			{Align: simpletable.AlignCenter, Text: interface{}(row[1]).(string)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
	fmt.Println()
}
