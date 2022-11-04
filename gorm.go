package tools

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

var mysql_ *gorm.DB

type MysqlConfig struct {
	Host            string
	Port            int
	Database        string
	Username        string
	Password        string
	Charset         string
	Timezone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	LogWriters      io.Writer
	LogConfigs      logger.Config
}

func InitMysql(config *MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=true&loc=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset, config.Timezone)
	mysql_ = InitGorm(mysql.Open(dsn), config.MaxIdleConns, config.MaxOpenConns, config.ConnMaxLifetime, config.LogWriters, config.LogConfigs)
}

func GetMysql() *gorm.DB {
	return mysql_
}

// InitGorm
// LogConfigs
//	SlowThreshold 慢SQL阈值
//	LogLevel 日志级别
//	IgnoreRecordNotFoundError 忽略ErrRecordNotFound(记录未找到)错误
//	Colorful 禁用彩色打印
func InitGorm(dialector gorm.Dialector, MaxIdleConns, MaxOpenConns int, ConnMaxLifetime time.Duration, LogWriters io.Writer, LogConfigs logger.Config) *gorm.DB {
	instance, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(LogWriters, "\r\n", log.LstdFlags),
			LogConfigs,
		),
	})
	if err != nil {
		panic("Connect db error:\n" + err.Error())
	}
	pool, _ := instance.DB()
	pool.SetMaxIdleConns(MaxIdleConns)
	pool.SetMaxOpenConns(MaxOpenConns)
	pool.SetConnMaxLifetime(ConnMaxLifetime)
	return instance
}

func GormDateNow() *GormDate {
	now := GormDate(time.Now())
	return &now
}

type GormDate time.Time

func (t *GormDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DateFormat))), nil
}

func (t *GormDate) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = GormDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t GormDate) Value() (driver.Value, error) {
	tm := time.Time(t)
	var zero time.Time
	if tm.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return tm, nil
}

func (t GormDate) ToTime() time.Time {
	return time.Time(t)
}

func GormDateTimeNow() *GormDateTime {
	now := GormDateTime(time.Now())
	return &now
}

type GormDateTime time.Time

func (t *GormDateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DatetimeFormat))), nil
}

func (t *GormDateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = GormDateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t GormDateTime) Value() (driver.Value, error) {
	tm := time.Time(t)
	var zero time.Time
	if tm.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return tm, nil
}

func (t GormDateTime) ToTime() time.Time {
	return time.Time(t)
}

type GormSliceInt []int

func (t *GormSliceInt) Scan(val interface{}) error {
	if value := string(val.([]byte)); value != "" {
		temps := strings.Split(value, ",")
		*t = make([]int, 0, len(temps))
		for _, temp := range temps {
			v, _ := strconv.Atoi(temp)
			*t = append(*t, v)
		}
	}
	return nil
}

func (t GormSliceInt) Value() (value driver.Value, err error) {
	temp := make([]string, 0, len(t))
	for _, v := range t {
		temp = append(temp, fmt.Sprintf("%v", v))
	}
	return strings.Join(temp, ","), nil
}

type GormSliceString []string

func (t *GormSliceString) Scan(val interface{}) error {
	if value := string(val.([]byte)); value != "" {
		*t = strings.Split(value, ",")
	}
	return nil
}

func (t GormSliceString) Value() (value driver.Value, err error) {
	return strings.Join(t, ","), nil
}
