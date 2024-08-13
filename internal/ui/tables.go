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
	"fmt"
	"io"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/charmbracelet/glamour"
)

/*
WriteListClientsTable :
Prints the supported clients table

params :-
a. w io.Writer
Where the table is to be printed
b. data ListClientsTable
Table data

returns :-
None
*/
func WriteListClientsTable(w io.Writer, data *ListClientsTable) {
	var headers []*simpletable.Cell
	for _, clientType := range data.ClientTypes {
		headers = append(headers, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  fmt.Sprintf("%s Clients", clientType),
		})
	}

	var columns [][]*simpletable.Cell
	for _, clients := range data.Clients {
		column := []*simpletable.Cell{}
		for _, client := range clients {
			column = append(column, &simpletable.Cell{
				Align: simpletable.AlignLeft,
				Text:  client,
			})
		}
		columns = append(columns, column)
	}

	WriteSimpleTable(w, &SimpleTableData{
		Headers:      headers,
		Columns:      columns,
		DefaultAlign: simpletable.AlignLeft,
		Enumerate:    true,
	})
}

/*
WriteRandomizedClientsTable :
Prints the randomized clients table

params :-
a. w io.Writer
Where the table is to be printed
b. data [][]string
Table data

returns :-
None
*/
func WriteRandomizedClientsTable(w io.Writer, data RandomizedClientsTable) {
	headers := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: "Type of Client"},
		{Align: simpletable.AlignCenter, Text: "Randomized Client"},
	}

	columns := [][]*simpletable.Cell{}
	clientTypesColumn := []*simpletable.Cell{}
	for _, clientTypes := range data.ClientTypes {
		clientTypesColumn = append(clientTypesColumn, &simpletable.Cell{
			Align: simpletable.AlignLeft,
			Text:  clientTypes,
		})
	}
	clientColumn := []*simpletable.Cell{}
	for _, client := range data.Clients {
		clientColumn = append(clientColumn, &simpletable.Cell{
			Align: simpletable.AlignLeft,
			Text:  client,
		})
	}
	columns = append(columns, clientTypesColumn, clientColumn)

	WriteSimpleTable(w, &SimpleTableData{
		Headers:      headers,
		Columns:      columns,
		DefaultAlign: simpletable.AlignLeft,
		Enumerate:    true,
	})
}

/*
WriteSimpleTable :
Prints a simple table from given data

params :-
a. w io.Writer
Where the table is to be printed
b. data SimpleTableData
Table data

returns :-
None
*/
func WriteSimpleTable(w io.Writer, data *SimpleTableData) {
	// Initialize table
	table := simpletable.New()
	// Get maxium dimensions
	n := 0
	for _, column := range data.Columns {
		if n < len(column) {
			n = len(column)
		}
	}
	m := len(data.Headers)

	if len(data.Headers) == 0 && !data.Enumerate {
		return
	}

	// Add headers to table
	table.Header = &simpletable.Header{
		Cells: data.Headers,
	}

	if data.Enumerate { // Add number header
		table.Header.Cells = append([]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
		}, table.Header.Cells...)
	}

	if len(data.Headers) > 0 { // Don't write rows if no headers are provided
		for x := 0; x < n; x++ {
			// Initialize new row
			row := []*simpletable.Cell{}
			if data.Enumerate { // Add row number cell
				row = append(row, &simpletable.Cell{
					Align: simpletable.AlignCenter,
					Text:  fmt.Sprint(x + 1),
				})
			}
			for y := 0; y < m; y++ {
				if y < len(data.Columns) && x < len(data.Columns[y]) { // Add existing cell to row
					row = append(row, data.Columns[y][x])
				} else { // Add empty cell to row
					row = append(row, &simpletable.Cell{
						Align: data.DefaultAlign,
						Text:  "-",
					})
				}
			}
			// Add new row to table
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	// Print table
	table.SetStyle(simpletable.StyleCompact)
	fmt.Fprint(w, table.String())
	fmt.Fprintln(w)
}

/*
WriteListNetworksTable :
Prints the supported networks table

params :-
a. w io.Writer
Where the table is to be printed
b. data []string
Table data

returns :-
None
*/
func WriteListNetworksTable(w io.Writer, data []string) {
	headers := []*simpletable.Cell{
		{
			Align: simpletable.AlignCenter,
			Text:  "Supported Networks",
		},
	}

	var columns [][]*simpletable.Cell
	column := []*simpletable.Cell{}
	for _, network := range data {
		column = append(column, &simpletable.Cell{
			Align: simpletable.AlignLeft,
			Text:  network,
		})
	}
	columns = append(columns, column)

	WriteSimpleTable(w, &SimpleTableData{
		Headers:      headers,
		Columns:      columns,
		DefaultAlign: simpletable.AlignLeft,
		Enumerate:    true,
	})
}

/*
WriteLidoStatusTable :
Prints the Lido Node Operator Information

params :-
a. w io.Writer
Where the data is to be printed
b. data []string
Node Operator data
c. string
Data Header
returns :-
None
*/
func WriteLidoStatusTable(w io.Writer, data []string, header string) {
	var allData []string

	allData = append(allData, header)
	allData = append(allData, data...)
	info := strings.Join(allData, "\n")
	renderedInfo, err := glamour.Render(info, "dark")

	if err != nil {
		fmt.Fprint(w, info) // Fallback to plain text if rendering fails
	} else {
		fmt.Fprint(w, renderedInfo)
	}
}
