package csrf

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func Configure() mux.MiddlewareFunc {
	return csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.Secure(false),
		csrf.CookieName("csrf_token"),
	)
}
