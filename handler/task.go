package handler

import (
	"fmt"
	"hello-go-todo-app/db"
	"net/http"

	"github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	_, err := db.GetDB().Exec("INSERT INTO tasks (user_id, name) VALUES (?, ?)", r.Context().Value("user_id"), r.FormValue("name"))

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["task"]
	_, err := db.GetDB().Exec("DELETE FROM tasks WHERE id = ?", taskId)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["task"]
	IsDone := r.FormValue("is_done")

	fmt.Println(IsDone)

	if IsDone == "1" {
		IsDone = "1"
	} else {
		IsDone = "0"
	}

	_, err := db.GetDB().Exec("UPDATE tasks SET is_done = ? WHERE id = ?", IsDone, taskId)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
