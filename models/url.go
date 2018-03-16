package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Url struct {
	Id        int    `bson:"_id"`
	SourceUrl string `bson:"SourceUrl"`

	ShortUrl string `bson:"-"` //暂时不存
}

var (
	URL_COLLECTION = "Url" //一个节点名
)

//自增id
func (url *Url) GenId() error {
	sourceUrl := url.SourceUrl
	err := GetDB().C(URL_COLLECTION).Find(bson.M{"SourceUrl": sourceUrl}).One(url) //在Url节点内查找，找到之后通过one返回id与url（路由链接）到url对象结构内
	if err != nil {
		//没有找到生成新的id
		url.Id, err = IncrMaxId("url")
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *Url) Save() error {
	return u.Upsert()
}

func (url *Url) Insert() error {
	return GetDB().C(URL_COLLECTION).Insert(url)
}

func (url *Url) FindById() error {
	return GetDB().C(URL_COLLECTION).FindId(url.Id).One(url)
}

func (url *Url) Update() error {
	return GetDB().C(URL_COLLECTION).Update(bson.M{"_id": url.Id}, url)
}

//若没有找到此条更新的数据则重新插入一条数据
func (url *Url) Upsert() error {
	_, err := GetDB().C(URL_COLLECTION).Upsert(bson.M{"_id": url.Id}, url)
	return err
}

func (url *Url) DeleteById() error {
	return GetDB().C(URL_COLLECTION).Remove(bson.M{"_id": url.Id})
}
