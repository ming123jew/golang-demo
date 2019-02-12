package service

import (
	"rpc-comment/entity"
	"rpc-comment/model"
)

type UserService struct {
	userModel model.Users
}

func (s *UserService) GetUser(uid int64) (entity.Users, error) {

	//userModel := model.Users{}

	var userObj entity.Users
	var err error

	userObj,err = s.userModel.Get(uid)
	if err != nil {
		return userObj,err
	}

	return userObj,nil
}
