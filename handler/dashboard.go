package handler

import (
	"html/template"
	"net/http"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard/index.html", "templates/layouts/base.html"))

	db := openDB()
	defer db.Close()

	var user User
	db.QueryRow("SELECT id, email, password FROM users WHERE id = ?", r.Context().Value("user_id")).Scan(&user.ID, &user.Email, &user.Password)

	tmpl.ExecuteTemplate(w, "base", user)
}
