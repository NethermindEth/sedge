package env

import "regexp"

var (
	ReTTD    = regexp.MustCompile(`(.*)_TTD=(.*)`)
	ReCONFIG = regexp.MustCompile(`CONFIG_URL=(.*)`)
)
