package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var (
	_db *mgo.Database
)

func GetDB() *mgo.Database {
	if _db != nil {
		return _db
	}

	//首先根据ip与端口连接到MongoDB数据库
	session, err := mgo.Dial(beego.AppConfig.String("mongo_url"))
	if err != nil {
		panic(err)
	}

	//再获取到某个库
	_db = session.DB(beego.AppConfig.String("mongo_database"))
	return _db
}
