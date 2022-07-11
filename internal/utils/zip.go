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
package utils

import (
	"fmt"

	"github.com/NethermindEth/sedge/configs"
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
	if len(lists) == 0 {
		return [][]string{}, nil
	}

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
