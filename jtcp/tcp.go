package jtcp

import (
	"github.com/yingzhuo/jing/jstr"
	"net"
)

func IsReachable(address string) bool {
	if jstr.IsBlank(address) {
		return false
	}

	conn, err := net.Dial("tcp", address)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return err == nil
}
