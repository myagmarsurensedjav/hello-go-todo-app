package handlers

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	email    string
	password string
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	// Get error message from cookie
	errorMessageCookie, err := r.Cookie("error_message")

	// Delete error message cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "error_message",
		MaxAge: -1,
	})

	var errorMessage string

	if err != nil {
		errorMessage = ""
	} else {
		errorMessage = errorMessageCookie.Value
	}

	tmpl := template.Must(template.ParseFiles("views/auth/login.html"))

	tmpl.Execute(w, map[string]interface{}{
		"errorMessage": errorMessage,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "admin@example.com" && password == "secret1234" {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: "logged_in",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

		return
	}

	// Error message should be used in the next request
	http.SetCookie(w, &http.Cookie{
		Name:  "error_message",
		Value: "Invalid email or password",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/auth/register.html"))
	tmpl.Execute(w, nil)
}
