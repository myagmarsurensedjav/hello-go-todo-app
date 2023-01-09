package handler

import (
	"database/sql"
	"hello-go-todo-app/middleware"
	"html/template"
	"net/http"

	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/login.html"))

	tmpl.Execute(w, map[string]interface{}{
		"errorMessage": middleware.GetErrorMessage(r),
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Attempt to login with email and password
	if email == "admin@example.com" && password == "secret1234" {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: "logged_in",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Redirect back with error message
	middleware.SetErrorMessage(w, "Invalid email or password")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/register.html"))
	tmpl.Execute(w, map[string]interface{}{
		"errorMessage": middleware.GetErrorMessage(r),
	})
}

type RegisterFormData struct {
	Email    string `valid:"email,required"`
	Password string `valid:"length(6|20),required"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	data := RegisterFormData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Validate form data
	if _, err := govalidator.ValidateStruct(&data); err != nil {
		middleware.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Password should be the same as password confirmation
	if data.Password != r.FormValue("password_confirmation") {
		middleware.SetErrorMessage(w, "Password confirmation does not match")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("mysql", "root:secret@(localhost:3306)/go-todo?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Insert user to database
	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", data.Email, data.Password)

	if err != nil {
		panic(err.Error())
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
