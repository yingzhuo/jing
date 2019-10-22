/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package flag

import "strings"

type ArrayFlags []string // flag.Value interface

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *ArrayFlags) String() string {
	return strings.Join(*i, ",")
}
