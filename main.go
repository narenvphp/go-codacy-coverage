package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	controller "github.com/narenvphp/go-codacy-coverage/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.AllUsers).Methods("GET")
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
