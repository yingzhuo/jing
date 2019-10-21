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

type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) Sort() {
	sort.Sort(is)
}

func (is IntSlice) SortReverse() {
	sort.Sort(sort.Reverse(is))
}

func (is IntSlice) IsSorted() bool {
	return sort.IsSorted(is)
}

func (is IntSlice) Size() int {
	return len(is)
}

func (is IntSlice) Map(fn func(s int) int) IntSlice {
	ret := make(IntSlice, len(is))
	for _, s := range is {
		ret = append(ret, fn(s))
	}
	return ret
}

func (is IntSlice) Filter(fn func(s int) bool) IntSlice {
	ret := make(IntSlice, 0)
	for _, s := range is {
		if fn(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (is IntSlice) NotFilter(fn func(s int) bool) IntSlice {
	ret := make(IntSlice, 0)
	for _, s := range is {
		if !fn(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (is IntSlice) AnyMatch(fn func(s int) bool) bool {
	for _, s := range is {
		if fn(s) {
			return true
		}
	}
	return false
}

func (is IntSlice) AllMatch(fn func(s int) bool) bool {
	for _, s := range is {
		if !fn(s) {
			return false
		}
	}
	return true
}

func (is IntSlice) NoneMatch(fn func(s int) bool) bool {
	for _, s := range is {
		if fn(s) {
			return false
		}
	}
	return true
}

func (is IntSlice) Reverse() {
	if length := len(is); length == 0 || length == 1 {
		return
	}

	for i, j := 0, len(is)-1; i < j; i, j = i+1, j-1 {
		is[i], is[j] = is[j], is[i]
	}
}
