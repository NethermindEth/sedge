package env

import "regexp"

var (
	ReTTD    = regexp.MustCompile(`TTD=(.*)`)
	ReCONFIG = regexp.MustCompile(`CONFIG_URL=(.*)`)
	ReMEV    = regexp.MustCompile(`MEV=(.*)`)
	ReXEEV   = regexp.MustCompile(`XEE_VERSION=(.*)`)
)
