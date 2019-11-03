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

var timePatterns = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02T15:04:05", // iso8601 without timezone
	time.RFC1123Z,
	time.RFC1123,
	"2006-01-02 15:04:05.999999999 -0700 MST",
	time.RFC822Z,
	time.RFC822,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05",
	"02 Jan 06 15:04 MST",
	"2006-01-02",
	"02 Jan 2006",
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05 -0700",
	"2006/01/02 15:04:05 -07:00",
	"2006/01/02 15:04:05 -0700",
	"2006/01/02",
	"2006/01/02 15:04:05",
}

func ParseTimeWildly(s string) (*time.Time, error) {
	for _, pattern := range timePatterns {
		if d, err := time.Parse(pattern, s); err == nil {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("unable to parse date: %s", s)
}
