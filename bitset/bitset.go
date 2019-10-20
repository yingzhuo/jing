/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package bitset

import "fmt"

type BitSet struct {
	bit   uint64
	count int
}

func (bs *BitSet) Add(i uint64) {
	bs.bit |= 1 << i
	bs.count += 1
}

func (bs *BitSet) Delete(i uint64) {
	bs.bit &= ^(1 << i)
}

func (bs *BitSet) Has(i uint64) bool {
	return bs.bit&(1<<i) != 0
}

func (bs *BitSet) Size() int {
	return bs.count
}

func (bs *BitSet) String() string {
	return fmt.Sprintf("%v", bs.bit)
}
