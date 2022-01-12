package utils

import (
	"math/rand"
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

/*
RandomizeClients :
Select a random execution client and a random consensus client

params :-
a. executionClients []string
List of execution clients
b. consensusClients []string
List of consensus clients

returns :-
a. string
Random execution client
b. string
Random consensus client
*/
func RandomizeClients(executionClients []string, consensusClients []string) (string, string) {
	executionClient := executionClients[rand.Intn(len(executionClients))]
	consensusClient := consensusClients[rand.Intn(len(consensusClients))]
	return executionClient, consensusClient
}
