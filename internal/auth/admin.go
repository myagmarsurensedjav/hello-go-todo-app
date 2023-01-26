package auth

import (
	"hello-go-todo-app/internal/config"
	"net/http"
)

func AdminMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationCode := r.Header.Get("Authorization")

		if authorizationCode != config.GetConfig().Auth.AdminKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}
