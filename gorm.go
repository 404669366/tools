package tools

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
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
