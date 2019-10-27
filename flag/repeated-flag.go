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
	"sort"
	"strings"

	"github.com/yingzhuo/jing/str"
)

type RepeatedFlag []string

func (r *RepeatedFlag) Len() int {
	return len(*r)
}

func (r *RepeatedFlag) Less(i, j int) bool {
	return (*r)[i] < (*r)[j]
}

func (r *RepeatedFlag) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *RepeatedFlag) String() string {
	return strings.Join(*r, ",")
}

func (r *RepeatedFlag) Set(s string) error {
	*r = append(*r, s)
	return nil
}

func (r *RepeatedFlag) Map(f func(string) string) *RepeatedFlag {
	rf := &RepeatedFlag{}
	for _, it := range *r {
		*rf = append(*rf, f(it))
	}
	return rf
}

func (r *RepeatedFlag) Filter(f func(string) bool) *RepeatedFlag {
	rf := &RepeatedFlag{}
	for _, it := range *r {
		if f(it) {
			*rf = append(*rf, it)
		}
	}
	return rf
}

func (r *RepeatedFlag) Unique() *RepeatedFlag {
	dict := make(map[string]bool)

	rf := &RepeatedFlag{}
	for _, it := range *r {
		if found := dict[it]; !found {
			dict[it] = true
			*rf = append(*rf, it)
		}
	}
	return rf
}

func (r *RepeatedFlag) SortAsc() *RepeatedFlag {
	sort.Sort(r)
	return r
}

func (r *RepeatedFlag) SortDesc() *RepeatedFlag {
	sort.Sort(sort.Reverse(r))
	return r
}

func (r *RepeatedFlag) RemoveEmpty() *RepeatedFlag {
	return r.Filter(func(s string) bool {
		return str.IsNotEmpty(s)
	})
}

func (r *RepeatedFlag) RemoveBlank() *RepeatedFlag {
	return r.Filter(func(s string) bool {
		return str.IsNotBlank(s)
	})
}
