package web

import (
	"hello-go-todo-app/internal/auth"
	"hello-go-todo-app/internal/task"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.FormValue("name")

	err := task.AddTask(auth.GetUserId(r), taskName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(r)["task"])

	err := task.DeleteTask(auth.GetUserId(r), taskId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(r)["task"])
	isDone := r.FormValue("is_done") == "1"

	err := task.UpdateTaskStatus(auth.GetUserId(r), taskId, isDone)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func ClearDoneTasks(w http.ResponseWriter, r *http.Request) {
	err := task.ClearDoneTasks(auth.GetUserId(r))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
