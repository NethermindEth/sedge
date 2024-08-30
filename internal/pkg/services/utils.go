package services

import "github.com/docker/docker/api/types"

func convertPorts(ports []types.Port) []Port {
	res := make([]Port, len(ports))
	for i, p := range ports {
		res[i].IP = p.IP
		res[i].PrivatePort = p.PrivatePort
		res[i].PublicPort = p.PublicPort
	}
	return res
}

func StringPtr(s string) *string {
	return &s
}

// Checks if a string slice contains a string
func Contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func ContainerNameWithTag(containerName, tag string) string {
	if tag == "" {
		return containerName
	}
	return containerName + "-" + tag
}
