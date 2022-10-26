package tools

import "github.com/robfig/cron/v3"

func InitCrontab(register func(*cron.Cron)) {
	register(cron.New())
}
