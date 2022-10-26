package tools

import (
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
