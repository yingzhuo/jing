/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package tcp

import (
	"net"

	"github.com/yingzhuo/jing/str"
)

// tcp可连通时返回true
// 否则返回false
//
// 例子:
//  IsReachable("localhost:8080")
func IsReachable(addr string) bool {
	if str.IsBlank(addr) {
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

// tcp不可连通时返回true
// 否则返回false
//
// 例子:
//  IsNotReachable("localhost:8080")
func IsNotReachable(addr string) bool {
	return !IsReachable(addr)
}
