package tools

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var redis_ *redis.Client

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	Db          int
	IdleTimeout time.Duration
	PoolSize    int
}

func InitRedis(config *RedisConfig) {
	redis_ = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%v", config.Host, config.Port),
		Password:    config.Password,
		DB:          config.Db,
		IdleTimeout: config.IdleTimeout,
		PoolSize:    config.PoolSize,
	})
	_, err := redis_.Ping(context.Background()).Result()
	if err != nil {
		panic("Connect rdb error:\n" + err.Error())
	}
}

func GetRedis() *redis.Client {
	return redis_
}
