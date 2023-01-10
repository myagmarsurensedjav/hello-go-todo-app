package handler

import (
	"hello-go-todo-app/db"
	"html/template"
	"net/http"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard/index.html", "templates/layouts/base.html"))

	var user User
	db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE id = ?", r.Context().Value("user_id")).Scan(&user.ID, &user.Email, &user.Password)

	tmpl.ExecuteTemplate(w, "base", user)
}
