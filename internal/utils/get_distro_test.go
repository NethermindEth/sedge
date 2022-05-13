package utils

import "testing"

func TestGetDistro(t *testing.T) {
	_, err := getOSInfo()
	if err != nil {
		t.Errorf("getOsInfo() failed: %v", err)
	}

	//TODO: validate distro info
}

func TestGetDistroName(t *testing.T) {
	_, err := GetDistroName()
	if err != nil {
		t.Errorf("GetDistroName() failed: %v", err)
	}

	//TODO: validate distro name
}
