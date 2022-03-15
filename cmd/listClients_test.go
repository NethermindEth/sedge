package cmd

import (
	"testing"
)

func TestBuildData(t *testing.T) {
	testingData, err := buildData()

	if (testingData == nil) && (err != nil) {
		t.Fail()
	} else if testingData != nil && err == nil {
		t.Skip()
	} else {
		t.Error(err)
	}
}

func TestListClientsCmd(t *testing.T) {
	testingService := listClientsCmd
	if testingService == nil {
		t.Fail()
	} else {
		t.Skip()
	}
}
