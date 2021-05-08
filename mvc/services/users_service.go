package services

import (
	"github.com/channrith/golang-microservice/mvc/domain"
	"github.com/channrith/golang-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
