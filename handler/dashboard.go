package handler

import (
	"hello-go-todo-app/middleware"
	"hello-go-todo-app/model"
	"html/template"
	"net/http"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard/index.html", "templates/layouts/base.html"))

	userId := middleware.GetUserId(r)

	user, err := model.GetUserById(userId)

	if err != nil {
		panic(err)
	}

	tasks, err := model.GetTasks(userId)

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"tasks": tasks,
		"user":  user,
	})
}
