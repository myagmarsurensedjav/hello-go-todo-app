package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"hello-go-todo-app/internal/auth"
	"hello-go-todo-app/internal/user"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var d auth.LoginFormData

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		respondWithError(w, Error{"Valid JSON data required.", http.StatusUnprocessableEntity})
		return
	}

	if _, err := govalidator.ValidateStruct(d); err != nil {
		respondWithError(w, Error{err.Error(), http.StatusUnprocessableEntity})
		return
	}

	// Attempt to login the user
	u, err := auth.Login(d)
	if err != nil {
		respondWithError(w, Error{err.Error(), http.StatusUnprocessableEntity})
		return
	}

	// Create a new access token for the user
	token, err := auth.CreateAccessToken(u, d.DeviceName)
	if err != nil {
		respondWithError(w, Error{err.Error(), http.StatusUnprocessableEntity})
		return
	}

	respondWithToken(w, token, u)
}

func respondWithToken(w http.ResponseWriter, token auth.PersonalAccessToken, u user.User) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth.LoginResponseWithToken{token.Token, token.ExpiresAt, u})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var d auth.RegisterFormData

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		respondWithError(w, Error{"Valid JSON data required.", http.StatusUnprocessableEntity})
		return
	}

	if _, err := govalidator.ValidateStruct(d); err != nil {
		respondWithError(w, Error{err.Error(), http.StatusUnprocessableEntity})
		return
	}

	u, err := auth.Register(d)
	if err != nil {
		respondWithError(w, Error{err.Error(), http.StatusUnprocessableEntity})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
