/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package flag

import (
	"strings"
	"unicode"
)

type ArrayFlags []string // flag.Value interface

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *ArrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *ArrayFlags) Size() int {
	return len(*i)
}

func (i *ArrayFlags) IsEmpty() bool {
	return i.Size() == 0
}

func (i *ArrayFlags) Map(fn func(string) string) ArrayFlags {
	ret := ArrayFlags{}
	for _, v := range *i {
		ret = append(ret, fn(v))
	}
	return ret
}

func (i *ArrayFlags) MapTrim() ArrayFlags {
	return i.Map(func(s string) string {
		s = strings.TrimLeftFunc(s, unicode.IsSpace)
		s = strings.TrimRightFunc(s, unicode.IsSpace)
		return s
	})
}
