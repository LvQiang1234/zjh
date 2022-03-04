package account

import (
	"google.golang.org/protobuf/proto"
	"zjh/log"
	"zjh/logic"
	"zjh/msgcenter"
	"zjh/orm"
	"zjh/pb"
	"zjh/tool"
)

func Init() {
	msgcenter.Register(pb.MsgId_LoginRequest, LoginReqHandler)
	msgcenter.Register(pb.MsgId_RegisterRequest, RegisterHandler)
	msgcenter.Register(pb.MsgId_GetPlayerInfoRequest, GetPlayerInfoHandler)
	msgcenter.Register(pb.MsgId_UpdateCoinRequest, UpdateCoinHandler)
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
	account.Imageid = int32(tool.RAND.RandI(0, 18))
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

func GetPlayerInfoHandler(playerId int32, data []byte) ([]byte, pb.MsgId) {
	ackErr := pb.ErrorType_Success
	info := logic.GetPlayerMgr().GetPlayerById(playerId).GetAccount()
	if info == nil {
		log.Debug("玩家中没有信息")
		ackErr = pb.ErrorType_Fail
		dat, err := proto.Marshal(&pb.PlayerInfoAck{Error: &ackErr})
		if err != nil {
			log.Debug(err.Error())
		}
		return dat, pb.MsgId_GetPlayerInfoResponse
	}
	dat, err := proto.Marshal(&pb.PlayerInfoAck{PlayerName: &info.Username, Coins: &info.Coin, ImageId: &info.Imageid, Error: &ackErr})
	if err != nil {
		log.Debug(err.Error())
	}
	return dat, pb.MsgId_GetPlayerInfoResponse
}

func UpdateCoinHandler(playerId int32, data []byte) ([]byte, pb.MsgId) {
	ackErr := pb.ErrorType_Success
	account := logic.GetPlayerMgr().GetPlayerById(playerId).GetAccount()
	if account == nil {
		log.Debug("获取玩家账户信息失败!")
		dat, err := proto.Marshal(&pb.UpdateCoinAck{Error: &ackErr, CoinNum: &account.Coin})
		if err != nil {
			log.Debug(err.Error())
		}
		return dat, pb.MsgId_UpdateCoinResponse
	}

	var coinReq pb.UpdateCoinReq
	proto.Unmarshal(data, &coinReq)
	account.Coin += *coinReq.Num
	account.Update()

	dat, err := proto.Marshal(&pb.UpdateCoinAck{Error: &ackErr, CoinNum: &account.Coin})
	if err != nil {
		log.Debug(err.Error())
	}
	return dat, pb.MsgId_UpdateCoinResponse
}
