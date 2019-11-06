/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package bin

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetBinaryDir() string {
	app := os.Args[0]
	if !strings.ContainsAny(app, `/\`) {
		app, _ = exec.LookPath(app)
	}
	dir, _ := filepath.Abs(filepath.Dir(app))
	return dir
}
