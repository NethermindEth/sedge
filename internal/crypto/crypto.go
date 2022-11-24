package crypto

import (
	"crypto/rand"
	"errors"
	"fmt"
)

func GenerateJWTSecret() (string, error) {
	secret := make([]byte, 32)
	n, err := rand.Read(secret)
	if err != nil {
		return "", err
	}
	if n != 32 {
		return "", errors.New("cannot generate 32 bytes long secret")
	}

	return fmt.Sprintf("%x", secret), nil
}
