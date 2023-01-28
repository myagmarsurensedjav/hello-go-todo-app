package user

import (
	"hello-go-todo-app/db"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func GetUserById(userId int) (User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE id = $1", userId).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func InsertUser(email string, password string) (User, error) {
	var u User
	err := db.GetDB().QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password", email, password).Scan(&u.ID, &u.Email, &u.Password)
	return u, err
}
