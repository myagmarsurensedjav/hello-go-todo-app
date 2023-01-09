package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdCookie, err := r.Cookie("user_id")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		userId := userIdCookie.Value

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		next(w, r.WithContext(ctx))
	})
}
