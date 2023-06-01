package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/joehann-9s/cli-tasks/models"
)

func VerifyFile(file *os.File, tasks []models.Task) ([]models.Task, error) {

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []models.Task{}
	}

	return tasks, nil
}
