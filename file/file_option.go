/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package file

func FileOption(firstFilename string, filenames ...string) (string, bool) {
	firstFilename = BetterFilename(firstFilename)

	if IsFileExists(firstFilename) {
		return firstFilename, true
	}

	for _, filename := range filenames {
		filename = BetterFilename(filename)
		if IsFileExists(filename) {
			return filename, true
		}
	}
	return "", false
}

func DicOption(firstDirname string, dirnames ...string) (string, bool) {
	firstDirname = BetterFilename(firstDirname)

	if IsDirExists(firstDirname) {
		return firstDirname, true
	}

	for _, dirname := range dirnames {
		dirname = BetterFilename(dirname)
		if IsDirExists(dirname) {
			return dirname, true
		}
	}
	return "", false
}
