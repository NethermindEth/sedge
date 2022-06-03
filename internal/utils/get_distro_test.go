package utils

import (
	"runtime"
	"testing"
)

func TestGetDistro(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}
	_, err := getOSInfo()
	if err != nil {
		t.Errorf("getOsInfo() failed: %v", err)
	}

	//TODO: validate distro info
}

func TestGetDistroName(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}
	_, err := GetDistroName()
	if err != nil {
		t.Errorf("GetDistroName() failed: %v", err)
	}

	//TODO: validate distro name
}
