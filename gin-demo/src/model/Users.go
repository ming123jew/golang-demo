package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"rpc-comment/config"

	_ "github.com/go-sql-driver/mysql" //注意这一行，必须不完全import

	"rpc-comment/entity"
)

type Users struct {
}

func (m *Users) Get(uid int64) (entity.Users, error) {

	client := newClient()

	var userObj entity.Users
	var err error
	ret, err := client.Where("id= ?", uid).Get(&userObj)
	if !ret {
		return userObj, err
	}

	return userObj, nil
}

func newClient() (*xorm.Engine) {

	mysql_conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.Mysql_username, config.Mysql_password, config.Mysql_host, config.Mysql_port, config.Mysql_dbname)
	engine, _ := xorm.NewEngine("mysql", mysql_conf)

	return engine
}