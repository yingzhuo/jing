/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package str

import (
	"sort"
	"strings"
)

type StringSlice []string

func NewStringSlice() *StringSlice {
	return &StringSlice{}
}

func (ss *StringSlice) String() string {
	return strings.Join(*ss, ",")
}

func (ss *StringSlice) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

func (ss *StringSlice) Add(value string) {
	_ = ss.Set(value)
}

func (ss *StringSlice) Len() int {
	return len(*ss)
}

func (ss *StringSlice) Less(i, j int) bool {
	return (*ss)[i] < (*ss)[j]
}

func (ss *StringSlice) Swap(i, j int) {
	(*ss)[i], (*ss)[j] = (*ss)[j], (*ss)[i]
}

func (ss *StringSlice) Sort() {
	sort.Sort(ss)
}

func (ss *StringSlice) SortReverse() {
	sort.Sort(sort.Reverse(ss))
}

func (ss *StringSlice) IsSorted() bool {
	return sort.IsSorted(ss)
}

func (ss *StringSlice) Map(fn func(s string) string) *StringSlice {
	ret := make(StringSlice, 0)
	for _, s := range *ss {
		ret = append(ret, fn(s))
	}
	return &ret
}

func (ss *StringSlice) TrimEach() *StringSlice {
	return ss.Map(strings.TrimSpace)
}

func (ss *StringSlice) Unique() *StringSlice {
	dict := make(map[string]bool)

	nss := &StringSlice{}
	for _, it := range *ss {
		if !dict[it] {
			nss.Add(it)
			dict[it] = true
		}
	}
	return nss
}

func (ss *StringSlice) Reverse() {
	if length := len(*ss); length == 0 || length == 1 {
		return
	}
	for i, j := 0, len(*ss)-1; i < j; i, j = i+1, j-1 {
		(*ss)[i], (*ss)[j] = (*ss)[j], (*ss)[i]
	}
}

func (ss *StringSlice) Filter(fn func(s string) bool) *StringSlice {
	ret := make(StringSlice, 0)
	for _, s := range *ss {
		if fn(s) {
			ret = append(ret, s)
		}
	}
	return &ret
}

func (ss *StringSlice) NotFilter(fn func(s string) bool) *StringSlice {
	ret := make(StringSlice, 0)
	for _, s := range *ss {
		if !fn(s) {
			ret = append(ret, s)
		}
	}
	return &ret
}

func (ss *StringSlice) AnyMatch(fn func(s string) bool) bool {
	for _, s := range *ss {
		if fn(s) {
			return true
		}
	}
	return false
}

func (ss *StringSlice) AllMatch(fn func(s string) bool) bool {
	for _, s := range *ss {
		if !fn(s) {
			return false
		}
	}
	return true
}

func (ss *StringSlice) NoneMatch(fn func(s string) bool) bool {
	for _, s := range *ss {
		if fn(s) {
			return false
		}
	}
	return true
}
