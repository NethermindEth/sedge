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
package crypto

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/NethermindEth/sedge/configs"
)

/*
GenerateJWTSecret:
Generate a 32 bytes secret and return it hex encoded

returns :-
a. string
The generated hex encode secret
b. error
Generation error if any, nil otherwise
*/
func GenerateJWTSecret() (string, error) {
	secret := make([]byte, 32)
	n, err := rand.Read(secret)
	if err != nil {
		return "", err
	}
	if n != 32 {
		return "", errors.New(configs.CannotGenerateSecret)
	}

	return fmt.Sprintf("%x", secret), nil
}
