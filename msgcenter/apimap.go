package msgcenter

import (
	"zjh/pb"
)

type HandleFunc func(int32, []byte) ([]byte, pb.MsgId)

var ApiMap map[pb.MsgId]HandleFunc

func init() {
	ApiMap = make(map[pb.MsgId]HandleFunc, 0)
}

func Register(msgId pb.MsgId, handler HandleFunc) {
	ApiMap[msgId] = handler
}

func GetHandler(msgId pb.MsgId) HandleFunc {
	return ApiMap[msgId]
}
