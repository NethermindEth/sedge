package utils

import (
	"fmt"

	"github.com/NethermindEth/1Click/configs"
)

/*
ZipString :
Zip string slices

params :-
a. lists ...[]string
String slices to be zipped

returns :-
a. [][]string
Zipped string slices
b. error
Error if any
*/
func ZipString(lists ...[]string) ([][]string, error) {
	size := len(lists[0])
	for _, list := range lists {
		if len(list) != size {
			return nil, fmt.Errorf(configs.ZipError)
		}
	}

	zippedList := make([][]string, size)
	for _, list := range lists {
		for i, item := range list {
			zippedList[i] = append(zippedList[i], item)
		}
	}

	return zippedList, nil
}
