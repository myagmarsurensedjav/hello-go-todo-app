package session

import (
	"github.com/gorilla/sessions"
	"hello-go-todo-app/internal/config"
	"net/http"
)

var store *sessions.CookieStore

func Init() {
	store = sessions.NewCookieStore([]byte(config.GetSessionKey()))
}

func Get(r *http.Request, name string) (*sessions.Session, error) {
	return store.Get(r, name)
}
