package logic

import (
	"zjh/orm"
)

type Player struct {
	account *orm.Account
}

func NewPlayer() *Player {
	player := &Player{}
	player.account = &orm.Account{}
	return player
}

func (this *Player) GetAccount() *orm.Account {
	return this.account
}
