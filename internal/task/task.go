package task

import (
	"hello-go-todo-app/db"
	"regexp"
	"strings"
)

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

func GetTasksByProject(userId int, project string) ([]Task, error) {
	var tasks []Task
	rows, err := db.GetDB().Query("SELECT id, name, is_done FROM tasks WHERE user_id = ? and project = ?", userId, project)

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

/**
 * It finds #projectName from the task name by regex. For example:
 * #projectName task name -> ProjectName
 */
func findProjectNameFromTask(taskName string) string {
	re := regexp.MustCompile(`#(\w+)`)
	matches := re.FindStringSubmatch(taskName)

	if len(matches) == 0 {
		return "Inbox"
	}

	return strings.ToUpper(string(matches[1][0])) + matches[1][1:]
}

func AddTask(userId int, name string) error {
	project := findProjectNameFromTask(name)
	_, err := db.GetDB().Exec("INSERT INTO tasks (user_id, name, project) VALUES (?, ?, ?)", userId, name, project)
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

func FilterTasksByStatus(tasks *[]Task, isDone bool) []Task {
	var filteredTasks []Task

	for _, task := range *tasks {
		if task.IsDone == isDone {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}