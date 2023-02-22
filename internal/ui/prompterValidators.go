package ui

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	"github.com/NethermindEth/sedge/internal/utils"
)

var (
	ErrInvalidEthereumAddress = errors.New("invalid ethereum address")
	ErrInvalidInt64String     = errors.New("invalid int64 string")
)

func EthAddressValidator(ans interface{}) error {
	if str, ok := ans.(string); ok && !utils.IsAddress(str) {
		return ErrInvalidEthereumAddress
	}
	return nil
}

func FilePathValidator(ans interface{}) error {
	if str, ok := ans.(string); ok {
		fileInfo, err := os.Stat(str)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return errors.New("is a directory not a file")
		}
	}
	return nil
}

var int64Regex = regexp.MustCompile("^[0-9]{1,15}$")

func Int64Validator(ans interface{}) error {
	if str, ok := ans.(string); ok {
		if !int64Regex.MatchString(str) {
			return ErrInvalidInt64String
		}
		if _, err := strconv.ParseInt(str, 10, 64); err != nil {
			return ErrInvalidInt64String
		}
	}
	return nil
}
