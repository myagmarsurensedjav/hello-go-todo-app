package auth

import (
	"context"
	"hello-go-todo-app/internal/session"
	"net/http"
)

const userSessionName = "session"

func AuthMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Get(r, userSessionName)

		userId, ok := s.Values["user_id"].(int)

		if !ok || userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		next(w, r.WithContext(ctx))
	})
}

func SetUserAuthSession(w http.ResponseWriter, r *http.Request, userId int) {
	s, _ := session.Get(r, userSessionName)
	s.Options.HttpOnly = true
	s.Options.Secure = true
	s.Options.SameSite = http.SameSiteStrictMode
	s.Options.Path = "/"
	s.Options.MaxAge = 60 * 60 * 24 * 7 // 7 days

	s.Values["user_id"] = userId
	s.Save(r, w)
}

func UnsetUserAuthSession(w http.ResponseWriter, r *http.Request) {
	s, _ := session.Get(r, userSessionName)
	s.Options.MaxAge = -1
	s.Save(r, w)
}

func GetUserId(r *http.Request) int {
	return r.Context().Value("user_id").(int)
}
