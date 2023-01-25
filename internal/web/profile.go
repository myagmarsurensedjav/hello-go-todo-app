package web

import (
	"encoding/json"
	"hello-go-todo-app/internal/auth"
	"hello-go-todo-app/internal/task"
	"hello-go-todo-app/internal/user"
	template "html/template"
	"net/http"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/profile/index.html", "templates/layouts/base.html"))

	u, err := user.GetUserById(auth.GetUserId(r))

	if err != nil {
		panic(err)
	}

	uJson, _ := json.Marshal(u)

	data, err := task.GetTasksCountByLast15Days(auth.GetUserId(r))

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"User":     u,
		"UserJson": string(uJson),
		"Data":     data,
	})
}
