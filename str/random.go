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
	"math/rand"
)

var (
	alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetic   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric      = "0123456789"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// 生成随机n位长度英文字母字符串
//  n: 长度
//
// 例子:
//  RandAlphabetic(10)
func RandAlphabetic(n int) string {
	return RandString(n, alphabetic)
}

// 生成随机n位长度数字字符串
//  n: 长度
//
// 例子:
//  RandNumeric(10)
func RandNumeric(n int) string {
	return RandString(n, numeric)
}

// 生成随机n位长度数字或英文字母字符串
//  n: 长度
//
// 例子:
//  RandAlphanumeric(10)
func RandAlphanumeric(n int) string {
	return RandString(n, alphanumeric)
}

func RandString(n int, charset string) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(charset) {
			b[i] = charset[idx]
			i++
		}
	}
	return string(b)
}
