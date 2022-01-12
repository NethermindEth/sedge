package utils

import (
	"strings"
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

