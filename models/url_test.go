package models

import (
	"github.com/stretchr/testify/assert" //一个断言包
	"gopkg.in/mgo.v2"
	"testing"
)

func init() {
	mongo_url := "127.0.0.1:27017"
	mongo_database := "short_url_test"
	session, err := mgo.Dial(mongo_url)
	if err != nil {
		panic(err)
	}
	_db = session.DB(mongo_database)
}

func TestMgo(t *testing.T) {
	var url = Url{
		Id:        1,
		SourceUrl: "http://www.qq.com",
	}

	//insert
	err := url.Insert()
	assert.Nil(t, err)

	//updata
	url.SourceUrl = "http://www.weixin.com"
	err = url.Update()
	assert.Nil(t, err) //判断t与err是否为空

	//find
	err = url.FindById()
	assert.Nil(t, err)
	assert.Equal(t, url.SourceUrl, "http://www.weixin.com")

	//delete
	err = url.DeleteById()
	assert.Nil(t, err)
}

func TestGenId(t *testing.T) {
	url := &Url{}
	url.SourceUrl = "http://www.facebook2.com"
	err := url.GenId() //自增id
	assert.Nil(t, err)

	err = url.Save() //保存更新（只用于保存到Url节点内）
	assert.Nil(t, err)
}
