package commands

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func IncompleteTask() {
	var taskID string

	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskID = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskID = flag.Arg(2)
	}

	if taskID == "" {
		fmt.Println("invalid task id.")
		return
	}

	if convertedValue, err := strconv.Atoi(taskID); err != nil || convertedValue == 0 {
		fmt.Println("invalid task id.")
		return
	}

	// Get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Create output
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	files, err := os.ReadDir(outputPath)
	if err != nil {
		panic(err)
	}

	userToGoList := files[0].Name()

	file, err := os.Open(filepath.Join(outputPath, userToGoList))
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	file.Close()

	writeFile, err := os.OpenFile(filepath.Join(outputPath, userToGoList), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	defer writeFile.Close()

	var taskFound = false

	for _, line := range content {
		var (
			fileTaskID        = line[0]
			fileTaskName      = line[1]
			fileTaskCreatedAt = line[2]
			fileTaskCompleted = line[3]
		)

		if taskID == fileTaskID && fileTaskCompleted == "false" {
			fmt.Println("task is not completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			fileTaskCompleted = "false"
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		fmt.Println("task with id '" + taskID + "' was not found.")
	}
}
