//go:build linux || darwin
// +build linux darwin

package utils

import "syscall"

func SetUmask(mask int) {
	syscall.Umask(mask)
}
