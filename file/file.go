/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package file

import (
	"os"
	"time"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func IsDirExists(filename string) bool {
	if info, err := os.Stat(filename); err != nil {
		return false
	} else {
		return info.IsDir()
	}
}

func IsDirNotExists(filename string) bool {
	return !IsDirExists(filename)
}

func IsFileExists(filename string) bool {
	if info, err := os.Stat(filename); err != nil {
		return false
	} else {
		return !info.IsDir()
	}
}

func IsFileNotExists(filename string) bool {
	return !IsFileExists(filename)
}

func Size(filename string) int64 {
	if info, err := os.Stat(filename); err != nil {
		panic(err)
	} else {
		return info.Size()
	}
}

func Touch(filename string) error {

	var err error

	if IsFileExists(filename) {
		now := time.Now().Local()
		err = os.Chtimes(filename, now, now)
		return err
	}

	fg := os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(filename, fg, 0644)

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	return err
}

func MakeDir(dirname string) error {
	if IsDirExists(dirname) {
		return nil
	}
	return os.MkdirAll(dirname, 0777)
}
