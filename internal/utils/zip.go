package utils

import "fmt"

func ZipString(lists ...[]string) ([][]string, error) {
	size := len(lists[0])
	for _, list := range lists {
		if len(list) != size {
			return nil, fmt.Errorf("all lists must have the same size")
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
