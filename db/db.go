package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"zjh/log"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:XX1516305754@/zjh?charset=utf8")
	if err != nil {
		log.Debug("连接数据库失败")
	} else {
		log.Debug("连接数据库成功")
	}
}

func GetDBEngine() *xorm.Engine {
	return engine
}
