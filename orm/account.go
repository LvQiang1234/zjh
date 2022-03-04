package orm

import (
	"github.com/go-xorm/xorm"
	"time"
	"zjh/db"
)

type Account struct {
	Id         int32
	Username   string
	Password   string
	Imageid    int32
	Coin       int32
	Createtime time.Time `xorm:"created"`
	Lastlogin  time.Time `xorm:"updated"`
}

const tableName string = "account"

var accountSession *xorm.Session

func init() {
	accountSession = db.GetDBEngine().Table(tableName)
}

func (this *Account) Get(query interface{}, args ...*string) error {
	var interAgrs []interface{}
	for _, arg := range args {
		interAgrs = append(interAgrs, *arg)
	}
	_, err := accountSession.Where(query, interAgrs...).Get(this)
	return err
}

func (this *Account) Insert() (int64, error) {
	return accountSession.InsertOne(this)
}

//func (this *Account) Find(query string, args ...interface{}) error{
//	return accountSession.Where(query, args...).Find(this)
//}

func (this *Account) Exist(query interface{}, args ...*string) (bool, error) {
	var interAgrs []interface{}
	for _, arg := range args {
		interAgrs = append(interAgrs, *arg)
	}
	exist, err := accountSession.Where(query, interAgrs...).Exist(this)
	return exist, err
}

func (this *Account) Update() error {
	_, err := accountSession.Where("id = ?", this.Id).Update(this)
	return err
}

func (this *Account) Delete() {
	accountSession.Delete(this)
}
