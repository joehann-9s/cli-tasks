package tasks

import (
	"fmt"

	"github.com/joehann-9s/cli-tasks/models"
)

func GetTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
		return
	}

	fmt.Println("Tasks:")
	fmt.Println(" State | ID | Name")
	for _, task := range tasks {
		status := "🟩"
		if task.Complete {
			status = "✅"
		}

		fmt.Printf("  %v   | %v  | %v\n", status, task.ID, task.Name)
	}
}
