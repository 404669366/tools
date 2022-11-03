package tools

import "github.com/bwmarrin/snowflake"

var snow *snowflake.Node

func InitSnowflake(node int) {
	var err error
	snow, err = snowflake.NewNode(int64(node))
	if err != nil {
		panic("init snowflake error : " + err.Error())
	}
}

func GetSnowflakeId() int64 {
	return snow.Generate().Int64()
}
