package configs

import "fmt"

var loggingDrivers = map[string]string{
	"none": "",
	"json": "json-file",
}

func ValidateLoggingFlag(loggingFlag string) error {
	if _, ok := loggingDrivers[loggingFlag]; !ok {
		return fmt.Errorf(InvalidLoggingFlag, loggingFlag)
	}
	return nil
}

func GetLoggingDriver(loggingFlag string) string {
	return loggingDrivers[loggingFlag]
}

func ValidLoggingFlags() []string {
	flags := make([]string, 0, len(loggingDrivers))
	for flag := range loggingDrivers {
		flags = append(flags, flag)
	}
	return flags
}
