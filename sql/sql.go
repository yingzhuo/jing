/**********************************************************************************************************************
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 **********************************************************************************************************************/
package sql

import (
	"database/sql"
	"time"
)

func NewNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	} else {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
}

func NewNullTime(s *time.Time) sql.NullTime {
	if s == nil {
		return sql.NullTime{}
	} else {
		return sql.NullTime{
			Time:  *s,
			Valid: true,
		}
	}
}

func NewNullInt32(s *int) sql.NullInt32 {
	if s == nil {
		return sql.NullInt32{}
	} else {
		return sql.NullInt32{
			Int32: int32(*s),
			Valid: true,
		}
	}
}

func NewNullInt64(s *int64) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{}
	} else {
		return sql.NullInt64{
			Int64: *s,
			Valid: true,
		}
	}
}

func NewNullBool(s *bool) sql.NullBool {
	if s == nil {
		return sql.NullBool{}
	} else {
		return sql.NullBool{
			Bool:  *s,
			Valid: true,
		}
	}
}

func NewNullFloat64(s *float64) sql.NullFloat64 {
	if s == nil {
		return sql.NullFloat64{}
	} else {
		return sql.NullFloat64{
			Float64: *s,
			Valid:   true,
		}
	}
}
