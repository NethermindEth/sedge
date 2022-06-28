package env

import "regexp"

var (
	ReTTD    = regexp.MustCompile(`TTD=(.*)`)
	ReCONFIG = regexp.MustCompile(`CONFIG_URL=(.*)`)
)
