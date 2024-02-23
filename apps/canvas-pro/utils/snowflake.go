package utils

import (
	"github.com/bwmarrin/snowflake"
)

var snowFlakeNode, _ = snowflake.NewNode(1)

func GetSnowFlakeNode() *snowflake.Node {
	return snowFlakeNode
}

func GenerateId() snowflake.ID {
	return snowFlakeNode.Generate()
}
