package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/channrith/golang-microservice/mvc/entity"
	"github.com/channrith/golang-microservice/mvc/service"
	"github.com/channrith/golang-microservice/mvc/util"
)

func GetUserById(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr := &util.ApplicationError{
			Message:    "user_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(userErr)
		res.WriteHeader(userErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, userErr := service.GetUserById(userId)
	if userErr != nil {
		jsonValue, _ := json.Marshal(userErr)
		res.WriteHeader(userErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}

func AddUser(res http.ResponseWriter, req *http.Request) {
	var user = entity.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	// fmt.Printf("Person: %s", userData.Email)

	// user := entity.User{}
	// user.SetFirstName("Pich Chy")
	// user.SetLastName("Moul")
	// user.SetEmail("pichchy.moul@yahoo.com")

	newUser, userErr := service.AddUser(user)
	if userErr != nil {
		jsonValue, _ := json.Marshal(userErr)
		res.WriteHeader(userErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(newUser)
	res.Write(jsonValue)
}
