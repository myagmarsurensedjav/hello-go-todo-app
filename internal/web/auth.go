package web

import (
	"hello-go-todo-app/internal/auth"
	error2 "hello-go-todo-app/internal/error"
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
	d := auth.LoginFormData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	u, err := auth.Login(d)
	if err != nil {
		error2.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	auth.SetUserAuthSession(w, r, u.ID)

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

func Register(w http.ResponseWriter, r *http.Request) {
	data := auth.RegisterFormData{
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

	_, err := auth.Register(data)

	if err != nil {
		error2.SetErrorMessage(w, err.Error())
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
