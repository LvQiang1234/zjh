package apitest

import (
	"zjh/errors"
	"zjh/log"
)

func TestFunc(playerId uint64, dat []byte) interface{} {
	log.Debug("playerId:%v", playerId)
	return errors.OK
}
