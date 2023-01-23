package task

import "hello-go-todo-app/db"

type Task struct {
	ID     int
	Name   string
	IsDone bool
}

func GetTasks(userId int) ([]Task, error) {
	var tasks []Task
	rows, err := db.GetDB().Query("SELECT id, name, is_done FROM tasks WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Name, &task.IsDone)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func AddTask(userId int, name string) error {
	_, err := db.GetDB().Exec("INSERT INTO tasks (user_id, name) VALUES (?, ?)", userId, name)
	return err
}

func DeleteTask(userId int, taskId int) error {
	_, err := db.GetDB().Exec("DELETE FROM tasks WHERE user_id = ? and id = ?", userId, taskId)
	return err
}

func UpdateTaskStatus(userId int, taskId int, isDone bool) error {
	_, err := db.GetDB().Exec("UPDATE tasks SET is_done = ? WHERE user_id = ? and id = ?", isDone, userId, taskId)
	return err
}
