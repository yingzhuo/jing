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
	"path/filepath"
)

type FileWalker struct {
	Start   string
	OnDir   func(dirname string)
	OnFile  func(filename string)
	OnError func(path string, err error)
}

func (fw *FileWalker) Run() {

	if IsDirNotExists(fw.Start) {
		return
	}

	filepath.Walk(fw.Start, func(path string, info os.FileInfo, err error) error {

		if err != nil && fw.OnError != nil {
			fw.OnError(path, err)
			return nil
		}

		if info.IsDir() && fw.OnDir != nil {
			fw.OnDir(path)
		}

		if !info.IsDir() && fw.OnFile != nil {
			fw.OnFile(path)
		}

		return nil
	})
}
