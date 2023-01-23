package web

import (
	"hello-go-todo-app/internal/auth"
	"hello-go-todo-app/internal/task"
	"hello-go-todo-app/internal/user"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard/index.html", "templates/layouts/base.html"))

	userId := auth.GetUserId(r)

	user, err := user.GetUserById(userId)

	if err != nil {
		panic(err)
	}

	tasks, err := task.GetTasks(userId)

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"tasks":          tasks,
		"user":           user,
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}
