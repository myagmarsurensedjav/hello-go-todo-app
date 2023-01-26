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
	rows, err := db.GetDB().Query("SELECT id, name, is_done FROM tasks WHERE user_id = $1", userId)

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
	rows, err := db.GetDB().Query("SELECT id, name, is_done FROM tasks WHERE user_id = $1 and project = $2", userId, project)

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
	_, err := db.GetDB().Exec("INSERT INTO tasks (user_id, name, project) VALUES ($1, $2, $3)", userId, name, project)
	return err
}

func DeleteTask(userId int, taskId int) error {
	_, err := db.GetDB().Exec("DELETE FROM tasks WHERE user_id = $1 and id = $2", userId, taskId)
	return err
}

func UpdateTaskStatus(userId int, taskId int, isDone bool) error {
	_, err := db.GetDB().Exec("UPDATE tasks SET is_done = $1 WHERE user_id = $2 and id = $3", isDone, userId, taskId)
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

func ClearDoneTasks(userId int) error {
	_, err := db.GetDB().Exec("DELETE FROM tasks WHERE user_id = $1 and is_done = true", userId)
	return err
}

func GetTasksCountByLast15Days(userId int) (map[string]int, error) {
	var tasksCountByLast15Days map[string]int
	rows, err := db.GetDB().Query("SELECT DATE(created_at) as date, COUNT(*) as count FROM tasks WHERE user_id = $1 GROUP BY date ORDER BY date DESC LIMIT 15", userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var date string
		var count int
		rows.Scan(&date, &count)
		tasksCountByLast15Days[date] = count
	}

	return tasksCountByLast15Days, nil
}
