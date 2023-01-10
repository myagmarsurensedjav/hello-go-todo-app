package handler

import (
	"hello-go-todo-app/db"
	"html/template"
	"net/http"
)

type Task struct {
	ID     int
	Name   string
	IsDone bool
}

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard/index.html", "templates/layouts/base.html"))

	var user User
	db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE id = ?", r.Context().Value("user_id")).Scan(&user.ID, &user.Email, &user.Password)

	var tasks []Task

	rows, err := db.GetDB().Query("SELECT id, name, is_done FROM tasks WHERE user_id = ?", r.Context().Value("user_id"))

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Name, &task.IsDone)
		tasks = append(tasks, task)
	}

	tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"tasks": tasks,
		"user":  user,
	})
}
