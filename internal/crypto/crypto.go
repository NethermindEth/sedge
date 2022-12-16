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
