package tools

import (
	"net"
)

func HasPing(host string) bool {
	conn, err := net.Dial("ip4", host)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
