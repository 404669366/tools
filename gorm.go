package tools

import (
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
	LogWriters      []io.Writer
	LogConfigs      logger.Config
}

func InitMysql(config *MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=true&loc=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset, config.Timezone)
	mysql_ = InitGorm(mysql.Open(dsn), config.MaxIdleConns, config.MaxOpenConns, config.ConnMaxLifetime, config.LogWriters, config.LogConfigs)
}

func GetMysql() *gorm.DB {
	return mysql_
}

// InitGorm LogConfigs=>SlowThreshold 慢SQL阈值,LogLevel 日志级别,IgnoreRecordNotFoundError忽略ErrRecordNotFound(记录未找到)错误,Colorful 禁用彩色打印
func InitGorm(dialector gorm.Dialector, MaxIdleConns, MaxOpenConns int, ConnMaxLifetime time.Duration, LogWriters []io.Writer, LogConfigs logger.Config) *gorm.DB {
	instance, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(io.MultiWriter(LogWriters...), "\r\n", log.LstdFlags),
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
