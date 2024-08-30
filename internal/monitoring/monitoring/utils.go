package monitoring

import (
	"fmt"
	"net"
	"sort"

	log "github.com/sirupsen/logrus"
)

var ErrDefaultPortInvalid = fmt.Errorf("default port invalid")

// Checks if port is occupied in a given host
func assignPorts(host string, defaults map[string]uint16) (ports map[string]uint16, err error) {
	ports = make(map[string]uint16)
	mask := make(map[uint16]bool)

	keys := make([]string, 0, len(defaults))
	for k := range defaults {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := defaults[k]
		if v == 0 {
			return ports, fmt.Errorf("%w: %s", ErrDefaultPortInvalid, k)
		}
		for !portAvailable(host, v) || mask[v] {
			v = v + 1
		}
		ports[k] = v
		mask[v] = true
	}

	return
}

// Checks if port is occupied in a given host
func portAvailable(ip string, port uint16) bool {
	log.Debugf("checking occupation of %s:%d", ip, port)
	netIp := net.ParseIP("127.0.0.1")
	if ip != "localhost" && ip != "0.0.0.0" {
		netIp = net.ParseIP(ip)
		if netIp == nil {
			log.Debugf("invalid host ip address")
			return true
		}
	}
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", netIp.String(), port))
	if err != nil {
		log.Debugf("error checking  %s:%d occupation: %v", ip, port, err)
		return false
	}
	sock.Close()
	return true
}
