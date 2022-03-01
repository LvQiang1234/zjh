package apitest

import (
	"google.golang.org/protobuf/proto"
	"zjh/errors"
	"zjh/log"
	proto2 "zjh/proto"
)

func TestFunc(playerId uint32, data []byte) interface{} {
	log.Debug("playerId:%v", playerId)
	loginReq := proto2.LoginReq{}
	proto.Unmarshal(data, &loginReq)
	log.Debug("username: %v", loginReq.Name)
	log.Debug("password: %v", loginReq.PassWord)
	return errors.OK
}
