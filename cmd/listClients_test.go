package cmd

import (
	"bytes"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestListClientsCmd(t *testing.T) {
	fdOut := new(bytes.Buffer)

	rootCmd.SetArgs([]string{"clients"})
	rootCmd.SetOut(fdOut)
	log.SetOutput(fdOut)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("listClientsCmd.Execute() failed: %v", err)
	}
}
