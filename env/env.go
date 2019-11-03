/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package env

import (
	"os"
	"strconv"
	"time"
)

func LookupString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		return value
	}
}

func LookupInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		if ret, err := strconv.Atoi(value); err != nil {
			return defaultValue
		} else {
			return ret
		}
	}
}

func LookupInt64(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		if ret, err := strconv.ParseInt(value, 10, 64); err != nil {
			return defaultValue
		} else {
			return ret
		}
	}
}

func LookupFloat64(key string, defaultValue float64) float64 {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		if ret, err := strconv.ParseFloat(value, 64); err != nil {
			return defaultValue
		} else {
			return ret
		}
	}
}

func LookupDuration(key string, defaultValue time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		if ret, err := time.ParseDuration(value); err != nil {
			return defaultValue
		} else {
			return ret
		}
	}
}
