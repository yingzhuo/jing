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
	"strings"
	"unicode"
)

const (
	EmptyString = ""
)

func Length(s string) int {
	return len(ToRuneSlice(s))
}

func CountRune(s string) int {
	return len(ToRuneSlice(s))
}

func CountRuneFunc(s string, fn func(r rune) bool) int {
	sum := 0
	for _, ch := range []rune(s) {
		if fn(ch) {
			sum++
		}
	}
	return sum
}

func AnyMatches(s string, fn func(r rune) bool) bool {
	for _, ch := range []rune(s) {
		if fn(ch) {
			return true
		}
	}
	return false
}

func AllMatches(s string, fn func(r rune) bool) bool {
	for _, ch := range []rune(s) {
		if !fn(ch) {
			return false
		}
	}
	return true
}

func NoneMatches(s string, fn func(r rune) bool) bool {
	for _, ch := range ToRuneSlice(s) {
		if fn(ch) {
			return false
		}
	}
	return true
}

func IsEmpty(s string) bool {
	return s == EmptyString
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsBlank(s string) bool {
	return IsEmpty(strings.TrimSpace(s))
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func TrimWhitespace(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

func Reverse(s string) string {
	if s == EmptyString {
		return EmptyString
	}
	r := strings.Split(s, EmptyString)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return strings.Join(r, EmptyString)
}
