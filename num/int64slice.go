/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package num

import "sort"

type Int64Slice []int64

func (is Int64Slice) Len() int {
	return len(is)
}

func (is Int64Slice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is Int64Slice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is Int64Slice) Sort() {
	sort.Sort(is)
}

func (is Int64Slice) SortReverse() {
	sort.Sort(sort.Reverse(is))
}

func (is Int64Slice) IsSorted() bool {
	return sort.IsSorted(is)
}

func (is Int64Slice) Size() int {
	return len(is)
}

func (is Int64Slice) Map(fn func(s int64) int64) Int64Slice {
	ret := make(Int64Slice, len(is))
	for _, s := range is {
		ret = append(ret, fn(s))
	}
	return ret
}

func (is Int64Slice) Filter(fn func(s int64) bool) Int64Slice {
	ret := make(Int64Slice, 0)
	for _, s := range is {
		if fn(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (is Int64Slice) NotFilter(fn func(s int64) bool) Int64Slice {
	ret := make(Int64Slice, 0)
	for _, s := range is {
		if !fn(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (is Int64Slice) AnyMatch(fn func(s int64) bool) bool {
	for _, s := range is {
		if fn(s) {
			return true
		}
	}
	return false
}

func (is Int64Slice) AllMatch(fn func(s int64) bool) bool {
	for _, s := range is {
		if !fn(s) {
			return false
		}
	}
	return true
}

func (is Int64Slice) NoneMatch(fn func(s int64) bool) bool {
	for _, s := range is {
		if fn(s) {
			return false
		}
	}
	return true
}

func (is Int64Slice) Reverse() {
	if length := len(is); length == 0 || length == 1 {
		return
	}

	for i, j := 0, len(is)-1; i < j; i, j = i+1, j-1 {
		is[i], is[j] = is[j], is[i]
	}
}
