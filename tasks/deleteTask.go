package tasks

import (
	"github.com/joehann-9s/cli-tasks/models"
)

func DeleteTask(tasks []models.Task, inputID int) []models.Task {
	for i, task := range tasks {
		if task.ID == inputID {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func DeleteAllTasks() {

}
