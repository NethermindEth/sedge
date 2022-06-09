package utils

import (
	"fmt"
)

/*
CheckErr :
Helper function for tests. Check if a given error was expected.

params :-
a. desc string
Description of the error in case of failure
b. isError bool
Whether the error was expected or not
c. err error
Error to check

returns :-
a. error
Error if any
*/
func CheckErr(descr string, isErr bool, err error) error {
	l := err == nil && isErr
	r := err != nil && !isErr

	if l || r {
		return fmt.Errorf("%s failed, unexpected error value: %v", descr, err)
	}
	return nil
}
