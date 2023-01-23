package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello-go-todo-app/internal/config"
)

var db *sql.DB

func getDSN() string {
	dbConfig := config.GetConfig().Db
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
}

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", getDSN())

	if err != nil {
		return err
	}

	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
