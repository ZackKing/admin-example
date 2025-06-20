package uuid

import (
	"github.com/bwmarrin/snowflake"
)

const (
	LogNodeID  = 1
	UserNodeID = 2
)

var LogNode *snowflake.Node
var UserNode *snowflake.Node

var NodeMap = map[int]*snowflake.Node{
	LogNodeID:  LogNode,
	UserNodeID: UserNode,
}

func Init() {
	LogNode, _ = snowflake.NewNode(LogNodeID)
	UserNode, _ = snowflake.NewNode(UserNodeID)
}

func GenSnowflakeID(id int) int64 {
	node := NodeMap[id]
	if node == nil {
		return 0
	}
	return LogNode.Generate().Int64()
}
