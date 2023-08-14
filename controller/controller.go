package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/narenvphp/go-codacy-coverage/config"
	model "github.com/narenvphp/go-codacy-coverage/model"
)

// AllUser = Select User API
func AllUsers(w http.ResponseWriter, r *http.Request) {
	var User model.User
	var response model.Response
	var arrUser []model.User

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, status FROM users")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&User.Id, &User.Name, &User.Status)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrUser = append(arrUser, User)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// InsertUser = Insert User API
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")

	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", name)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
