package csrf

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"hello-go-todo-app/internal/config"
)

func Configure() mux.MiddlewareFunc {
	return csrf.Protect(
		[]byte(config.GetConfig().App.Key),
		csrf.Secure(false),
		csrf.CookieName("csrf_token"),
	)
}
