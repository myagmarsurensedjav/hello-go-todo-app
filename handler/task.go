package handler

import (
	"hello-go-todo-app/middleware"
	"hello-go-todo-app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.FormValue("name")

	err := model.AddTask(middleware.GetUserId(r), taskName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(r)["task"])

	err := model.DeleteTask(middleware.GetUserId(r), taskId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(mux.Vars(r)["task"])
	isDone := r.FormValue("is_done") == "1"

	err := model.UpdateTaskStatus(middleware.GetUserId(r), taskId, isDone)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
