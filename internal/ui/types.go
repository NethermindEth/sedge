package ui

import "github.com/alexeyco/simpletable"

type SimpleTableData struct {
	Headers      []*simpletable.Cell
	Columns      [][]*simpletable.Cell
	DefaultAlign int
	Enumerate    bool
}

type ListClientsTable struct {
	ClientTypes []string
	Clients     [][]string
}

type RandomizedClientsTable struct {
	ClientTypes []string
	Clients     []string
}
