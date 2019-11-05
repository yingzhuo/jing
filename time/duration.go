/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package time

import (
	"fmt"
	"time"
)

func SubTime(t1, t2 time.Time) time.Duration {
	return t1.Sub(t2)
}

func SubTimeAbs(t1, t2 time.Time) time.Duration {
	d := SubTime(t1, t2)
	if d < 0 {
		return -d
	}
	return d
}

func ParseDurationWildly(value string) (time.Duration, error) {

	if d, err := time.ParseDuration(value); err == nil {
		return d, nil
	} else {
		if f, err := ParseTimeWildly(value); err == nil {
			sub := f.Local().UnixNano() - LocalNow().UnixNano()

			if sub <= 0 {
				return 0, nil
			} else {
				return time.Duration(sub), nil
			}
		} else {
			return 0, fmt.Errorf("can not parse duration (%v)", value)
		}
	}
}
