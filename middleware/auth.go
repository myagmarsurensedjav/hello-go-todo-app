package middleware

import (
	"context"
	"net/http"
	"strconv"
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

func GetUserId(r *http.Request) int {
	userId, _ := strconv.Atoi(r.Context().Value("user_id").(string))
	return userId
}
