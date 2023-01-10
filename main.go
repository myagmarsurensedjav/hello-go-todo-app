package main

import (
	"net/http"

	"hello-go-todo-app/handler"
	"hello-go-todo-app/middleware"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/", handler.ShowLandingPage).Methods("Get")

	r.HandleFunc("/login", handler.ShowLoginForm).Methods("Get")
	r.HandleFunc("/login", handler.Login).Methods("Post")

	r.HandleFunc("/logout", handler.Logout).Methods("Post")

	r.HandleFunc("/register", handler.ShowRegisterForm).Methods("Get")
	r.HandleFunc("/register", handler.Register).Methods("Post")

	r.HandleFunc("/dashboard", middleware.AuthMiddleware(handler.ShowDashboard)).Methods("Get")
}

func main() {
	r := mux.NewRouter()

	// Register middleware
	r.Use(middleware.ErrorMessageMiddleware)

	// Register routes
	registerRoutes(r)

	http.ListenAndServe(":8080", r)
}
