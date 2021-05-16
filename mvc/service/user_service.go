package service

import (
	"github.com/channrith/golang-microservice/mvc/entity"
	"github.com/channrith/golang-microservice/mvc/util"
)

func GetUserById(userId int64) (*entity.User, *util.ApplicationError) {
	return entity.GetUserById(userId)
}

func AddUser(userData entity.User) (*entity.User, *util.ApplicationError) {
	return entity.AddUser(userData.FirstName, userData.LastName, userData.Email)
}
