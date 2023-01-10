package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", "root:secret@(localhost:3306)/go-todo?parseTime=true")

	if err != nil {
		return err
	}

	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
