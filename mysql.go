package tools

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var mysql_ *gorm.DB

type MysqlConfig struct {
	Host         string
	Port         int
	Database     string
	Username     string
	Password     string
	Charset      string
	Timezone     string
	MaxIdleConns int
	MaxOpenConns int
}

func InitMysql(config *MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=true&loc=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset, config.Timezone)
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connect db error:\n" + err.Error())
	}
	pool, _ := instance.DB()
	pool.SetMaxIdleConns(config.MaxIdleConns)
	pool.SetMaxOpenConns(config.MaxOpenConns)
	pool.SetConnMaxLifetime(time.Hour)
	mysql_ = instance
}

func GetMysql() *gorm.DB {
	return mysql_
}
