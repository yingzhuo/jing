package jtcp

import (
	"github.com/yingzhuo/jing/jstr"
	"net"
)

func IsReachable(addr string) bool {
	if jstr.IsBlank(addr) {
		panic("addr is blank")
	}

	conn, err := net.Dial("tcp", addr)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	return err == nil
}

func IsNotReachable(addr string) bool {
	return !IsReachable(addr)
}
