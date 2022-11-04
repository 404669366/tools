package tools

import (
	"fmt"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	DatetimeFormat = "2006-01-02 15:04:05"
)

type Date time.Time

func (lt *Date) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*lt)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DateFormat))), nil
}

type DateTime time.Time

func (lt *DateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*lt)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DatetimeFormat))), nil
}
