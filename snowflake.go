package tools

import "github.com/bwmarrin/snowflake"

var snow *snowflake.Node

func InitSnowflake(node int) {
	snow, _ = snowflake.NewNode(int64(node))
}

func GenerateId() int64 {
	return snow.Generate().Int64()
}
