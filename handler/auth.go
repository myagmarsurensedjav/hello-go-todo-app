package handler

import (
	"database/sql"
	"fmt"
	"hello-go-todo-app/middleware"
	"html/template"
	"net/http"

	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@(localhost:3306)/go-todo?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/login.html"))

	tmpl.Execute(w, map[string]interface{}{
		"errorMessage": middleware.GetErrorMessage(r),
	})
}

type User struct {
	ID       int
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := openDB()
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", r.FormValue("email")).Scan(&user.ID, &user.Email, &user.Password)

	// Check if user exists
	if err != nil {
		middleware.SetErrorMessage(w, "Could not find user with that email")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if password is correct
	if user.Password != r.FormValue("password") {
		// Redirect back with error message
		middleware.SetErrorMessage(w, "Invalid email or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: fmt.Sprintf("%d", user.ID),
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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

	db := openDB()
	defer db.Close()

	// Insert user to database
	db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", data.Email, data.Password)

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
