package tools

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	DatetimeFormat = "2006-01-02 15:04:05"
)

func DateNow() *Date {
	now := Date(time.Now())
	return &now
}

type Date time.Time

func (t *Date) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DateFormat))), nil
}

func (t *Date) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = Date(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Date) Value() (driver.Value, error) {
	tm := time.Time(t)
	var zero time.Time
	if tm.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return tm, nil
}

func (t Date) ToTime() time.Time {
	return time.Time(t)
}

func DateTimeNow() *DateTime {
	now := DateTime(time.Now())
	return &now
}

type DateTime time.Time

func (t *DateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DatetimeFormat))), nil
}

func (t *DateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t DateTime) Value() (driver.Value, error) {
	tm := time.Time(t)
	var zero time.Time
	if tm.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return tm, nil
}

func (t DateTime) ToTime() time.Time {
	return time.Time(t)
}
