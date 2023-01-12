package handler

import (
	"hello-go-todo-app/hash"
	"hello-go-todo-app/middleware"
	"hello-go-todo-app/model"
	"html/template"
	"net/http"

	"github.com/asaskevich/govalidator"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/login.html"))

	tmpl.Execute(w, map[string]interface{}{
		"errorMessage": middleware.GetErrorMessage(r),
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	passwordInputValue := r.FormValue("password")

	user, err := model.GetUserByEmail(r.FormValue("email"))

	// Check if user exists
	if err != nil {
		middleware.SetErrorMessage(w, "Could not find user with that email")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if password is correct
	if !hash.CheckPasswordHash(passwordInputValue, user.Password) {
		// Redirect back with error message
		middleware.SetErrorMessage(w, "Invalid email or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	middleware.SetUserAuthSession(w, r, user.ID)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	middleware.UnsetUserAuthSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	err := model.InsertUser(data.Email, hash.HashPassword(data.Password))

	if err != nil {
		middleware.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
