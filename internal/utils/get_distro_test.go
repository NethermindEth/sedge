package utils_test

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetOSInfo(t *testing.T) {
	var (
		reID        = regexp.MustCompile(`^ID=(.*)$`)
		reVersionID = regexp.MustCompile(`^VERSION_ID=(.*)$`)
	)
	var distro utils.DistroInfo
	distro.Architecture = runtime.GOARCH
	file := "/etc/os-release"

	f, v := os.Open(file)

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if v == nil && cerr != nil {
			log.Errorf(configs.ClosingFileError, file)
			v = cerr
		}

	}()

	// Get the distro name
	s := bufio.NewScanner(f)
	for s.Scan() {
		if m := reID.FindStringSubmatch(s.Text()); m != nil {
			distro.Name = strings.Trim(m[1], `"`)
		} else if m := reVersionID.FindStringSubmatch(s.Text()); m != nil {
			distro.Version = strings.Trim(m[1], `"`)
		}

	}
	testingService, err := utils.GetOSInfo()
	assert.Nil(t, err, nil)
	assert.Equal(t, err, v)
	fmt.Println(distro)
	assert.Equal(t, testingService, distro)
}
