package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var snowFlakeNode, _ = snowflake.NewNode(1)

func GetSnowFlakeNode() *snowflake.Node {
	return snowFlakeNode
}	

func GetRandomId() string {
	return fmt.Sprintf("%v", snowFlakeNode.Generate())
}
