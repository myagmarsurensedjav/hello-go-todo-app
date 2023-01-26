package app

import (
	"github.com/gorilla/mux"
	"hello-go-todo-app/internal/auth"
	"hello-go-todo-app/internal/web"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/", web.ShowLandingPage).Methods("Get")

	r.HandleFunc("/login", web.ShowLoginForm).Methods("Get")
	r.HandleFunc("/login", web.Login).Methods("Post")

	r.HandleFunc("/logout", web.Logout).Methods("Post")

	r.HandleFunc("/register", web.ShowRegisterForm).Methods("Get")
	r.HandleFunc("/register", web.Register).Methods("Post")

	r.HandleFunc("/dashboard", auth.AuthMiddleware(web.ShowDashboard)).Methods("Get")

	r.HandleFunc("/tasks", auth.AuthMiddleware(web.AddTask)).Methods("Post")
	r.HandleFunc("/tasks/{task}/remove", auth.AuthMiddleware(web.RemoveTask)).Methods("Post")
	r.HandleFunc("/tasks/{task}/done", auth.AuthMiddleware(web.MarkTaskAsDone)).Methods("Post")
	r.HandleFunc("/tasks/clear", auth.AuthMiddleware(web.ClearDoneTasks)).Methods("Post")

	r.HandleFunc("/profile", auth.AuthMiddleware(web.ShowProfile)).Methods("Get")

	r.HandleFunc("/db/migrate", auth.AdminMiddleware(web.Migrate)).Methods("Get")
}
