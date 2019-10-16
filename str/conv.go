/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package str

import (
	"strconv"
	"time"
)

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsNotNumber(s string) bool {
	return !IsNumber(s)
}

func IsBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

func IsNotBool(s string) bool {
	return !IsBool(s)
}

func ToInt(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return n
	}
}

func ToInt64(s string) int64 {
	return int64(ToInt(s))
}

func ToDuration(s string) time.Duration {
	return time.Duration(ToInt64(s))
}

func ToBool(s string) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		panic(err)
	} else {
		return b
	}
}

func ToFloat64(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err != nil {
		panic(err)
	} else {
		return f
	}
}

func ToRuneSlice(s string) []rune {
	return []rune(s)
}
