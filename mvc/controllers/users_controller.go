package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/channrith/golang-microservice/mvc/services"
	"github.com/channrith/golang-microservice/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(userErr)
		res.WriteHeader(userErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, userErr := services.GetUser(userId)
	if userErr != nil {
		jsonValue, _ := json.Marshal(userErr)
		res.WriteHeader(userErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
