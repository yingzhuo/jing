/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package conv

import (
	"fmt"
	"strconv"
	"strings"
)

func ToString(a interface{}) string {
	switch s := a.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	case fmt.GoStringer:
		return s.GoString()
	case []rune:
		return string(s)
	default:
		return fmt.Sprintf("%v", s)
	}
}

func ToBool(a interface{}) (bool, error) {
	switch b := a.(type) {
	case bool:
		return b, nil
	case string:
		return stringToBool(b)
	case uint, uint8, uint16, uint32, uint64:
		return b != 0, nil
	case int, int8, int16, int32, int64:
		return b != 0, nil
	case float32, float64:
		return b != 0, nil
	default:
		return false, fmt.Errorf("can not conv to bool")
	}
}

func ToInt(a interface{}) (int, error) {
	switch i := a.(type) {
	case uint:
		return int(i), nil
	case uint8:
		return int(i), nil
	case uint16:
		return int(i), nil
	case uint32:
		return int(i), nil
	case uint64:
		return int(i), nil
	case int:
		return int(i), nil
	case int8:
		return int(i), nil
	case int16:
		return int(i), nil
	case int32:
		return int(i), nil
	case int64:
		return int(i), nil
	case float32:
		return int(i), nil
	case float64:
		return int(i), nil
	case bool:
		if i {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		n, err := stringToInt64(i)
		if err != nil {
			return 0, err
		} else {
			return int(n), nil
		}
	default:
		return 0, fmt.Errorf("can not conv to int")
	}
}

func ToInt64(a interface{}) (int64, error) {
	switch i := a.(type) {
	case uint:
		return int64(i), nil
	case uint8:
		return int64(i), nil
	case uint16:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case uint64:
		return int64(i), nil
	case int:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case int64:
		return int64(i), nil
	case float32:
		return int64(i), nil
	case float64:
		return int64(i), nil
	case bool:
		if i {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		n, err := stringToInt64(i)
		if err != nil {
			return 0, err
		} else {
			return int64(n), nil
		}
	default:
		return 0, fmt.Errorf("can not conv to int64")
	}
}

func ToFloat32(a interface{}) (float32, error) {
	switch i := a.(type) {
	case uint:
		return float32(i), nil
	case uint8:
		return float32(i), nil
	case uint16:
		return float32(i), nil
	case uint32:
		return float32(i), nil
	case uint64:
		return float32(i), nil
	case int:
		return float32(i), nil
	case int8:
		return float32(i), nil
	case int16:
		return float32(i), nil
	case int32:
		return float32(i), nil
	case int64:
		return float32(i), nil
	case float32:
		return i, nil
	case float64:
		return float32(i), nil
	case bool:
		if i {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		n, err := stringToFloat64(i)
		if err != nil {
			return 0, err
		} else {
			return float32(n), nil
		}
	default:
		return 0, fmt.Errorf("can not conv to float32")
	}
}

func ToFloat64(a interface{}) (float64, error) {
	switch i := a.(type) {
	case uint:
		return float64(i), nil
	case uint8:
		return float64(i), nil
	case uint16:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case int:
		return float64(i), nil
	case int8:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case float32:
		return float64(i), nil
	case float64:
		return i, nil
	case bool:
		if i {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		n, err := stringToFloat64(i)
		if err != nil {
			return 0, err
		} else {
			return n, nil
		}
	default:
		return 0, fmt.Errorf("can not conv to float64")
	}
}

// -------------------------------------------------------------------------------------------------------------------

func stringToBool(s string) (bool, error) {
	s = strings.ToLower(s)

	switch s {
	case "t", "true", "on", "1", "y", "yes":
		return true, nil
	case "f", "false", "off", "0", "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("can not conv to bool")
	}
}

func stringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func stringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
