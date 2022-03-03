package main

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"zjh/db"
	"zjh/log"
	"zjh/orm"
)

//import (
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"zjh/log"
//)
//
//type User struct {
//	username string
//	password string
//}
//
//func main() {
//	db, err := gorm.Open("mysql", "root:XX1516305754@/zjh?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		log.Debug("连接失败")
//	} else {
//		log.Debug("连接成功")
//	}
//	users := make([]User, 0)
//	//db.Table("accounts").Where("username = ?", "lvqiang").Find(&users)
//	db.Table("accounts").Find(&users)
//	log.Debug("users len: %v", len(users))
//	log.Debug("username: %v", users[0].username)
//	log.Debug("password: %v", users[0].password)
//	defer db.Close()
//}

func main() {
	account := &orm.Account{}
	account.Username = "df"
	account.Password = "dsjnsdkl"
	account.Lastlogin = time.Now()
	account.Createtime = time.Now()
	_, err := db.GetDBEngine().Table("account").Where("username = ?", "df").Update(account)
	if err != nil {
		log.Debug("err : %v", err.Error())
	}
	log.Debug("%v", account.Password)
}
