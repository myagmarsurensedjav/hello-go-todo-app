package auth

import (
	"errors"
	"hello-go-todo-app/db"
	"hello-go-todo-app/internal/hash"
	"hello-go-todo-app/internal/user"
	"time"
)

type LoginFormData struct {
	Email      string `json:"email" valid:"email,required"`
	Password   string `json:"password" valid:"length(6|20),required"`
	DeviceName string `json:"device_name" valid:"length(1|20),required"`
}

type LoginResponseWithToken struct {
	Token     string    `json:"token"`
	ExpiresAt string    `json:"expires_at"`
	User      user.User `json:"user"`
}

type RegisterFormData struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"length(6|20),required"`
}

func Register(data RegisterFormData) (user.User, error) {
	var u user.User
	if _, err := user.GetUserByEmail(data.Email); err == nil {
		return u, errors.New("User already exists")
	}

	return user.InsertUser(data.Email, hash.HashPassword(data.Password))
}

func Login(data LoginFormData) (user.User, error) {
	u, err := user.GetUserByEmail(data.Email)
	if err != nil {
		return u, errors.New("Could not find user")
	}

	if !hash.CheckPasswordHash(data.Password, u.Password) {
		return u, errors.New("Invalid email or password")
	}

	return u, nil
}

type PersonalAccessToken struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	DeviceName string `json:"device_name"`
	Token      string `json:"token"`
	ExpiresAt  string `json:"expires_at"`
}

func CreateAccessToken(u user.User, deviceName string) (PersonalAccessToken, error) {
	newToken := hash.GenerateRandomString(60)
	expiresAt := time.Now().AddDate(0, 0, 365)
	var token PersonalAccessToken
	row := db.GetDB().QueryRow("INSERT INTO personal_access_tokens (user_id, device_name, token, expires_at) VALUES ($1, $2, $3, $4) RETURNING id, user_id, device_name, token, expires_at", u.ID, deviceName, newToken, expiresAt)
	err := row.Scan(&token.ID, &token.UserID, &token.DeviceName, &token.Token, &token.ExpiresAt)
	return token, err
}
