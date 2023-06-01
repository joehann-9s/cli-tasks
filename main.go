package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joehann-9s/cli-tasks/models"
	funcs_task "github.com/joehann-9s/cli-tasks/tasks"
	"github.com/joehann-9s/cli-tasks/utils"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []models.Task

	tasks, err = utils.VerifyFile(file, tasks)
	utils.Welcome()

	inputID := 0
	switch os.Args[1] {
	case "list":
		funcs_task.GetTasks(tasks)

	case "create":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Write your task:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		tasks = funcs_task.CreateTask(tasks, input)
		utils.SaveTasks(file, tasks)
		funcs_task.GetTasks(tasks)

	case "complete":
		inputID = utils.GetIDFromArgs()
		tasks = funcs_task.CompleteTask(tasks, inputID)
		utils.SaveTasks(file, tasks)
		funcs_task.GetTasks(tasks)

	case "delete":
		inputID = utils.GetIDFromArgs()
		tasks = funcs_task.DeleteTask(tasks, inputID)
		utils.SaveTasks(file, tasks)

		funcs_task.GetTasks(tasks)
	default:
		utils.PrintUsage()

	}

}
