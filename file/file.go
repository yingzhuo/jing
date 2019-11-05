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
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func IsExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func IsSymlink(filename string) bool {
	info, err := os.Lstat(filename)
	return err == nil && (info.Mode()&os.ModeSymlink != 0)
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
	if IsFileExists(dirname) {
		return fmt.Errorf("%s is exists, but it is not a dir", dirname)
	}
	return os.MkdirAll(dirname, os.ModePerm)
}

func RemoveFileOrDir(name string) error {

	if info, err := os.Stat(name); err != nil {
		return nil
	} else if info.IsDir() {
		return os.RemoveAll(name)
	} else {
		return os.Remove(name)
	}
}

func CleanDir(dirname string) error {
	if IsFileExists(dirname) {
		return fmt.Errorf("%s is exists, and it is not a dir", dirname)
	}

	if !strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}

	fs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, f := range fs {
		err := RemoveFileOrDir(dirname + f.Name())
		if err != nil {
			return err
		}
	}

	return nil
}
