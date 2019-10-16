package jing

import (
	"net"
)

func IsTcpServerReachable(addr string) bool {
	conn, err := net.Dial("tcp", addr)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return err == nil
}
