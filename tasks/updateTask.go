package tasks

import (
	"github.com/joehann-9s/cli-tasks/models"
)

func CompleteTask(tasks []models.Task, inputID int) []models.Task {
	for i, task := range tasks {
		if task.ID == inputID {
			tasks[i].Complete = !task.Complete
			break
		}
	}
	return tasks
}
