package tasks

import (
	"github.com/joehann-9s/cli-tasks/models"
)

func CreateTask(tasks []models.Task, input string) []models.Task {
	newTask := models.Task{
		ID:       getNextID(tasks),
		Name:     input,
		Complete: false,
	}

	tasks = append(tasks, newTask)
	return tasks
}

func getNextID(tasks []models.Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1

}
