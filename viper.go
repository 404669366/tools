package tools

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"os"
)

func InitViper(fileName string) {
	path, _ := os.Getwd()
	viper.SetConfigName(fileName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config error:\n" + err.Error())
	}
}

func InitViperFs(fileName string, configs embed.FS) {
	viper.SetConfigType("yml")
	config, err := configs.ReadFile(fileName)
	if err != nil {
		panic("read config error:\n" + err.Error())
	}
	if err := viper.ReadConfig(bytes.NewBuffer(config)); err != nil {
		panic("read config error:\n" + err.Error())
	}
}
