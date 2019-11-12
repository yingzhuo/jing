/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package os

import (
	"os"
	"runtime"
	"strings"
)

func GetPWD() string {
	dir, _ := os.Getwd()
	if dir == "" {
		dir = "."
	}
	return dir
}

func IsWindows() bool {
	return strings.EqualFold("windows", runtime.GOOS)
}

func IsMacOS() bool {
	return strings.EqualFold("darwin", runtime.GOOS)
}

func IsLinux() bool {
	return strings.EqualFold("linux", runtime.GOOS)
}
