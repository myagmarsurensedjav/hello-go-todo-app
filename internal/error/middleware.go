package error

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

		UnsetErrorMessage(w)

		// Set error message to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "error_message", errorMessageCookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetErrorMessage(w http.ResponseWriter, message string) {
	http.SetCookie(w, &http.Cookie{
		Name:  errorMessageKey,
		Value: message,
	})
}

func GetErrorMessage(r *http.Request) string {
	errorMessage, ok := r.Context().Value(errorMessageKey).(string)

	if !ok {
		return ""
	}

	return errorMessage
}

func UnsetErrorMessage(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   errorMessageKey,
		Value:  "",
		MaxAge: -1,
	})
}
