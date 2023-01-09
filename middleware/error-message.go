package middleware

import (
	"context"
	"net/http"
)

const errorMessageKey = "error_message"

func ErrorMessageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessageCookie, err := r.Cookie(errorMessageKey)

		// If error message cookie is not found, continue to next handler
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// Set error message to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "error_message", errorMessageCookie.Value)

		// Delete error message cookie
		http.SetCookie(w, &http.Cookie{
			Name:   errorMessageKey,
			Value:  "",
			MaxAge: -1,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
