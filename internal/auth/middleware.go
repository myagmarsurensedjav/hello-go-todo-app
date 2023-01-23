package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func AuthMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		userId, ok := session.Values["user_id"].(int)

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
	session, _ := store.Get(r, "session")
	session.Options.HttpOnly = true
	session.Options.Secure = true
	session.Options.SameSite = http.SameSiteStrictMode
	session.Options.Path = "/"
	session.Options.MaxAge = 60 * 60 * 24 * 7 // 7 days

	session.Values["user_id"] = userId
	session.Save(r, w)
}

func UnsetUserAuthSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
}

func GetUserId(r *http.Request) int {
	return r.Context().Value("user_id").(int)
}
