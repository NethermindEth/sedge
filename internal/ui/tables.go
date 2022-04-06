package ui

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

/*
WriteListClientsTable :
Prints the supported clients table

params :-
a. data ListClientsTable
Table data

returns :-
None
*/
func WriteListClientsTable(data *ListClientsTable) {
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

	WriteSimpleTable(&SimpleTableData{
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
a. data [][]string
Table data

returns :-
None
*/
func WriteRandomizedClientsTable(data RandomizedClientsTable) {

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

	WriteSimpleTable(&SimpleTableData{
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
a. data SimpleTableData
Table data

returns :-
None
*/
func WriteSimpleTable(data *SimpleTableData) {
	//Initialize table
	table := simpletable.New()
	//Get maxium dimensions
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

	//Add headers to table
	table.Header = &simpletable.Header{
		Cells: data.Headers,
	}

	if data.Enumerate { // Add number header
		table.Header.Cells = append([]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
		}, table.Header.Cells...)
	}

	for x := 0; x < n; x++ {
		//Initialize new row
		row := []*simpletable.Cell{}
		if data.Enumerate { //Add row number cell
			row = append(row, &simpletable.Cell{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprint(x + 1),
			})
		}
		for y := 0; y < m; y++ {
			if y < len(data.Columns) && x < len(data.Columns[y]) { //Add existing cell to row
				row = append(row, data.Columns[y][x])
			} else { //Add empty cell to row
				row = append(row, &simpletable.Cell{
					Align: data.DefaultAlign,
					Text:  "-",
				})
			}
		}
		//Add new row to table
		table.Body.Cells = append(table.Body.Cells, row)
	}
	//Print table
	table.SetStyle(simpletable.StyleCompact)
	table.Println()
	fmt.Println()
}
