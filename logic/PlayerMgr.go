package logic

import (
	"sync"
)

type PlayerManager struct {
	PlayerIdPlayerMap map[int32]*Player
	m_Locker          *sync.RWMutex
}

var playerMgr *PlayerManager

func init() {
	playerMgr = &PlayerManager{}
	playerMgr.PlayerIdPlayerMap = make(map[int32]*Player, 0)
	playerMgr.m_Locker = &sync.RWMutex{}
}

func (this *PlayerManager) GetPlayerNum() int {
	return len(this.PlayerIdPlayerMap)
}

func GetPlayerMgr() *PlayerManager {
	return playerMgr
}

func (this *PlayerManager) AddOrSetPlayerById(playerId int32, player *Player) {
	this.m_Locker.Lock()
	this.PlayerIdPlayerMap[playerId] = player
	this.m_Locker.Unlock()
}

func (this *PlayerManager) GetPlayerById(playerId int32) *Player {
	this.m_Locker.RLocker().Lock()
	player := this.PlayerIdPlayerMap[playerId]
	this.m_Locker.RLocker().Unlock()
	return player
}

func (this *PlayerManager) DeletePlayerById(playerId int32) bool {
	this.m_Locker.Lock()
	delete(this.PlayerIdPlayerMap, playerId)
	this.m_Locker.Unlock()
	return true
}
