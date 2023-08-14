package model

type User struct {
	Id     string `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Status string `form:"city" json:"status"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}
