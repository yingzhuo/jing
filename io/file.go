/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package io

import "io/ioutil"

func LoadFileString(filename string) string {
	return string(LoadFileBytes(filename))
}

func LoadFileBytes(filename string) []byte {
	if bytes, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		return bytes
	}
}
