package profile

import "strings"

type InvalidProfileError struct {
	message       string
	invalidFields []string
	missingFields []string
}

func (e InvalidProfileError) Error() string {
	// Nil error
	if e.message == "" {
		return ""
	}

	if len(e.invalidFields) == 0 && len(e.missingFields) == 0 {
		return e.message
	}

	msg := e.message + " -> "
	if len(e.invalidFields) > 0 {
		msg += "invalid fields: " + strings.Join(e.invalidFields, ", ")
	}
	if len(e.missingFields) > 0 {
		msg += "missing fields: " + strings.Join(e.missingFields, ", ")
	}
	return msg
}
