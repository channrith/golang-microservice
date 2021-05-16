package entity

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/channrith/golang-microservice/mvc/config"
	"github.com/channrith/golang-microservice/mvc/util"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (user *User) SetFirstName(firstName string) { user.FirstName = firstName }
func (user *User) SetLastName(lastName string)   { user.LastName = lastName }
func (user *User) SetEmail(email string)         { user.Email = email }

func GetUserById(userId int64) (
	*User,
	*util.ApplicationError,
) {
	db, err := sql.Open("mysql", config.DB_URI)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM users WHERE id = ?", userId)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		user := &User{}
		err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			log.Fatal(err)
		}

		return user, nil
	}

	return nil, &util.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

func AddUser(
	firstName string,
	lastName string,
	email string,
) (
	*User,
	*util.ApplicationError,
) {
	db, err := sql.Open("mysql", config.DB_URI)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO users(first_name, last_name, email) VALUES ('" + firstName + "', '" + lastName + "', '" + email + "')"
	res, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user := &User{uint64(lastId), firstName, lastName, email}
	return user, nil
}
