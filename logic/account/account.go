package account

import (
	"google.golang.org/protobuf/proto"
	"zjh/log"
	"zjh/logic"
	"zjh/msgcenter"
	"zjh/orm"
	"zjh/pb"
)

func Init() {
	msgcenter.Register(pb.MsgId_LoginRequest, LoginReqHandler)
	msgcenter.Register(pb.MsgId_RegisterRequest, RegisterHandler)
}

func LoginReqHandler(playerId int32, data []byte) ([]byte, pb.MsgId) {
	ackErr := pb.ErrorType_Success

	player := logic.NewPlayer()

	var loginReq pb.LoginReq
	proto.Unmarshal(data, &loginReq)

	exist, _ := player.GetAccount().Exist("username = ? and password = ?", loginReq.Name, loginReq.PassWord)

	if exist == true {
		player.GetAccount().Get("username = ? and password = ?", loginReq.Name, loginReq.PassWord)
		logic.GetPlayerMgr().AddOrSetPlayerById(player.GetAccount().Id, player)

		if err := player.GetAccount().Update(); err != nil {
			log.Debug(err.Error())
		}

		dat, err := proto.Marshal(&pb.LoginAck{Error: &ackErr, PlayerId: &player.GetAccount().Id})
		if err != nil {
			log.Debug(err.Error())
		}
		return dat, pb.MsgId_LoginResponse
	}
	ackErr = pb.ErrorType_Fail
	dat, err := proto.Marshal(&pb.LoginAck{Error: &ackErr})
	if err != nil {
		log.Debug(err.Error())
	}
	return dat, pb.MsgId_LoginResponse
}

func RegisterHandler(playerId int32, data []byte) ([]byte, pb.MsgId) {
	ackErr := pb.ErrorType_Success

	account := &orm.Account{}

	var registerReq pb.RegisterReq
	proto.Unmarshal(data, &registerReq)

	exist, _ := account.Exist("username = ? and password = ?", registerReq.Name, registerReq.PassWord)

	if exist == true {
		ackErr = pb.ErrorType_Fail
		dat, err := proto.Marshal(&pb.RegisterAck{Error: &ackErr})
		if err != nil {
			log.Debug(err.Error())
		}
		return dat, pb.MsgId_RegisterAckResponse
	}
	account.Username = *registerReq.Name
	account.Password = *registerReq.PassWord
	if _, err := account.Insert(); err != nil {
		log.Debug("err : %v", err.Error())
	}

	ackErr = pb.ErrorType_Success
	dat, err := proto.Marshal(&pb.RegisterAck{Error: &ackErr})
	if err != nil {
		log.Debug(err.Error())
	}
	return dat, pb.MsgId_RegisterAckResponse
}
