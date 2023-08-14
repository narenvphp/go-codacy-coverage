package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/go-codacy-coverage")
	if err != nil {
		panic(err.Error())
	}
	return db
}
