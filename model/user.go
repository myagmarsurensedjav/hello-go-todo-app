package model

import "hello-go-todo-app/db"

type User struct {
	ID       int
	Email    string
	Password string
}

func GetUserById(userId int) (User, error) {
	var user User
	err := db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE id = ?", userId).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
