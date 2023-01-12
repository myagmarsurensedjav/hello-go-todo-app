package main

import (
	"log"
	"net/http"

	"hello-go-todo-app/db"
	"hello-go-todo-app/handler"
	"hello-go-todo-app/middleware"

	"github.com/gorilla/csrf"
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

	r.HandleFunc("/tasks", middleware.AuthMiddleware(handler.AddTask)).Methods("Post")
	r.HandleFunc("/tasks/{task}/remove", middleware.AuthMiddleware(handler.RemoveTask)).Methods("Post")
	r.HandleFunc("/tasks/{task}/done", middleware.AuthMiddleware(handler.MarkTaskAsDone)).Methods("Post")
}

func configureCSRF() mux.MiddlewareFunc {
	return csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.Secure(false),
		csrf.CookieName("csrf_token"),
	)
}

func main() {
	r := mux.NewRouter()

	// Register middleware
	r.Use(middleware.ErrorMessageMiddleware)

	// Register routes
	registerRoutes(r)

	r.Use(configureCSRF())

	// Init DB
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.GetDB().Close()

	http.ListenAndServe(":8080", r)
}
