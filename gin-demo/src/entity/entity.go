package entity

import (
	"gopkg.in/mgo.v2/bson"
)

//Mysql信息结构，需要注意的是：结构名称(Users)和字段名、字段类型需要和Mysql表一一对应。
type Users struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Remember_token string `json:"remember_token"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
}

//Mongo信息结构，需要注意的是：字段名、字段类型需要和Mongo表一一对应。
type UsersMongo struct {
	Id             bson.ObjectId `bson:"_id" json:"id"`
	Uid            int64         `json:"uid"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	Remember_token string        `json:"remember_token"`
	Created_at     int64         `json:"created_at"`
	Updated_at     int64         `json:"updated_at"`
	Stat           int64         `json:"stat"`
}
