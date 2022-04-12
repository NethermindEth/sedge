package cmd

import (
	"bytes"
	"regexp"
	"testing"
)

func TestVersionCmdExecute(t *testing.T) {
	fdOut := new(bytes.Buffer)

	versionRegexp := regexp.MustCompile(`1click [0-9][0-9|\.]+`)

	rootCmd.SetArgs([]string{"version"})
	rootCmd.SetOut(fdOut)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("versionCmd.Execute() failed: %v", err)
	}

	if !versionRegexp.Match(fdOut.Bytes()) {
		t.Errorf("versionCmd.Execute() has an invalid output: %s", fdOut.String())
	}
}
