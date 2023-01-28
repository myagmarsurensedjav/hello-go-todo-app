package app

import (
	"fmt"
	"hello-go-todo-app/db"
	"hello-go-todo-app/internal/config"
	"hello-go-todo-app/internal/csrf"
	error2 "hello-go-todo-app/internal/error"
	"hello-go-todo-app/internal/session"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	// Web routes
	web := a.Router.PathPrefix("/").Subrouter()
	web.Use(csrf.Configure())
	web.Use(error2.ErrorMessageMiddleware)
	registerWebRoutes(web)

	// API routes
	api := a.Router.PathPrefix("/api/").Subrouter()
	api.Use(apiMiddleware)
	registerApiRoutes(api)

	return a
}

func (a *App) Start() {
	err := config.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	session.Init()

	// Init DB
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	defer db.GetDB().Close()

	go fmt.Println("App started on http://0.0.0.0:8080")

	http.ListenAndServe(":8080", a.Router)
}
