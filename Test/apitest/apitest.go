package apitest

import (
	"google.golang.org/protobuf/proto"
	"zjh/errors"
	"zjh/log"
	"zjh/pb"
)

func TestFunc(playerId uint, data []byte) interface{} {
	log.Debug("playerId:%v", playerId)
	loginReq := pb.LoginReq{}
	proto.Unmarshal(data, &loginReq)
	log.Debug("username: %v", *loginReq.Name)
	log.Debug("password: %v", *loginReq.PassWord)
	return errors.OK
}
