package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"hello-go-todo-app/internal/config"
)

var db *sql.DB

func getDSN() string {
	dbConfig := config.GetConfig().Db

	if dbConfig.Driver == "mysql" {
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
	}

	if dbConfig.Driver == "postgres" {
		return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Dbname)
	}

	return ""
}

func InitDB() error {
	fmt.Println(getDSN())

	var err error
	db, err = sql.Open(config.GetConfig().Db.Driver, getDSN())

	if err != nil {
		return err
	}

	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
