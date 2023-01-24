package web

import (
	"hello-go-todo-app/internal/auth"
	error2 "hello-go-todo-app/internal/error"
	"hello-go-todo-app/internal/hash"
	"hello-go-todo-app/internal/user"
	"html/template"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/csrf"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/login.html"))

	tmpl.Execute(w, map[string]interface{}{
		"errorMessage":   error2.GetErrorMessage(r),
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	passwordInputValue := r.FormValue("password")

	user, err := user.GetUserByEmail(r.FormValue("email"))

	// Check if user exists
	if err != nil {
		error2.SetErrorMessage(w, "Could not find user with that email")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if password is correct
	if !hash.CheckPasswordHash(passwordInputValue, user.Password) {
		error2.SetErrorMessage(w, "Invalid email or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	auth.SetUserAuthSession(w, r, user.ID)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	auth.UnsetUserAuthSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth/register.html"))
	tmpl.Execute(w, map[string]interface{}{
		"errorMessage":   error2.GetErrorMessage(r),
		csrf.TemplateTag: csrf.TemplateField(r),
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
		error2.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Password should be the same as password confirmation
	if data.Password != r.FormValue("password_confirmation") {
		error2.SetErrorMessage(w, "Password confirmation does not match")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Check if user already exists
	{
		_, err := user.GetUserByEmail(data.Email)

		if err == nil {
			error2.SetErrorMessage(w, "User with that email already exists")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
	}

	// Create user
	err := user.InsertUser(data.Email, hash.HashPassword(data.Password))

	if err != nil {
		error2.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
