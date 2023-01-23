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
	currentProject := r.FormValue("project")

	user, err := user.GetUserById(userId)

	if err != nil {
		panic(err)
	}

	var tasks []task.Task

	if currentProject == "" {
		tasks, err = task.GetTasks(userId)
	} else {
		tasks, err = task.GetTasksByProject(userId, currentProject)
	}

	undoneTasks := task.FilterTasksByStatus(&tasks, false)
	doneTasks := task.FilterTasksByStatus(&tasks, true)

	if err != nil {
		panic(err)
	}

	projects, _ := task.GetProjects(auth.GetUserId(r))

	var allTasksCount int

	for _, project := range projects {
		allTasksCount += project.TasksCount
	}

	tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"undoneTasks":    undoneTasks,
		"doneTasks":      doneTasks,
		"user":           user,
		"projects":       projects,
		"currentProject": currentProject,
		"allTasksCount":  allTasksCount,
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}
