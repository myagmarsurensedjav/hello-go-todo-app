package app

import (
	"fmt"
	"hello-go-todo-app/db"
	"hello-go-todo-app/internal/csrf"
	error2 "hello-go-todo-app/internal/error"
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

	registerRoutes(a.Router)

	a.Router.Use(error2.ErrorMessageMiddleware)
	a.Router.Use(csrf.Configure())

	return a
}

func (a *App) Start() {
	// Init DB
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	defer db.GetDB().Close()

	go fmt.Println("App started on http://0.0.0.0:8080")

	http.ListenAndServe(":8080", a.Router)
}
