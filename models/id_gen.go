package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IdGen struct {
	Id       string `bson:"_id"`      //此id为MongoDB系统生成的id
	MaxValue int    `bson:"MaxValue"` //有新记录过来时此值加一
}

var (
	IdGen_COLLECTION = "IdGen" //MongoDB的id节点名（保存id名字与此id的对应的链接数）
)

func IncrMaxId(id string) (maxId int, err error) {
	idGen := &IdGen{}
	//自增MaxValue
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"MaxValue": 1}}, //如何更新（把inc对应的MaxValue字段加一）
		ReturnNew: true,                                  //是否更新最新的记录
		Upsert:    true,                                  //如果没有找到对应记录，则插入此条记录
	}

	//apply方法来自增MaxValue值,通过change的设置规则，
	//到MongoDB内IdGen节点查找到此id的记录数有记录就把idGen对象的MaxValue值与MongoDB内IdGen节点上的值加一
	//无就把idGen对象的MaxValue值与MongoDB内IdGen节点上的值从1开始
	_, err = GetDB().C(IdGen_COLLECTION).Find(bson.M{"_id": id}).Apply(change, idGen)
	if err != nil {
		return
	}

	//把查找后的结果值返回出去
	maxId = idGen.MaxValue
	return

}
