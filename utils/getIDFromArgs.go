package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetIDFromArgs() int {
	if len(os.Args) < 3 {
		fmt.Println("Write a task's ID")
		return -1
	}

	inputID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("ID have to be a number, try again")
		return -1
	}

	return inputID
}
