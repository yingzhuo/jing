/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package file

import "os"

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func IsDirExists(name string) bool {
	if info, err := os.Stat(name); err != nil {
		panic(err)
	} else {
		return info.IsDir()
	}
}

func IsFileExists(name string) bool {
	if info, err := os.Stat(name); err != nil {
		panic(err)
	} else {
		return !info.IsDir()
	}
}

func Size(name string) int64 {
	if info, err := os.Stat(name); err != nil {
		panic(err)
	} else {
		return info.Size()
	}
}
