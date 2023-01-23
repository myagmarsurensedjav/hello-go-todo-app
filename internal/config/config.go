package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	App struct {
		Key string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
}

var config Config

func GetConfig() Config {
	return config
}

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// App config
	config.App.Key = os.Getenv("APP_KEY")

	// Database config
	config.Db.Host = os.Getenv("DB_HOST")
	config.Db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	config.Db.User = os.Getenv("DB_USER")
	config.Db.Password = os.Getenv("DB_PASSWORD")
	config.Db.Dbname = os.Getenv("DB_NAME")

	return nil
}

func GetSessionKey() string {
	return config.App.Key
}
