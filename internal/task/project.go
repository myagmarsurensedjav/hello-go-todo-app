package task

import "hello-go-todo-app/db"

type Project struct {
	Name       string
	TasksCount int
}

func GetProjects(userId int) ([]Project, error) {
	var projects []Project
	rows, err := db.GetDB().Query("SELECT DISTINCT project as Name, count(*) as TasksCount FROM tasks WHERE user_id = $1 group by project", userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var project Project
		rows.Scan(&project.Name, &project.TasksCount)
		projects = append(projects, project)
	}

	return projects, nil
}
