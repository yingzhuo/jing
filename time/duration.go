/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package time

import "time"

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
