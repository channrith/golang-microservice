package domain

import (
	"fmt"
	"net/http"

	"github.com/channrith/golang-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Channrith", LastName: "Moul", Email: "channrith.moul@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	// user := users[userId]
	// if user == nil {
	// 	return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
	// }

	// return user, nil

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
