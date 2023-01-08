package main

import (
	"net/http"

	"hello-go-todo-app/handlers"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.ShowLandingPage).Methods("Get")

	r.HandleFunc("/login", handlers.ShowLoginForm).Methods("Get")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/register", handlers.ShowRegisterForm).Methods("Get")
}

func main() {
	r := mux.NewRouter()

	registerRoutes(r)

	http.ListenAndServe(":8080", r)
}
