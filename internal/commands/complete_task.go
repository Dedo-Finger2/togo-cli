package commands

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func CompleteTask() {
	var taskID string

	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskID = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskID = flag.Arg(2)
	}

	if taskID == "" {
		log.Println("[WARN]: TaskID cannot be empty.")
		return
	}

	if convertedValue, err := strconv.Atoi(taskID); err != nil || convertedValue == 0 {
		log.Println("[WARN]: Invalid task id, it must be a valid integer.")
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

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			log.Println("[WARN]: This task is already completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted != "true" {
			fileTaskCompleted = "true"
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		log.Println("[ERROR]: Task with id '" + taskID + "' was not found in your to-go list.")
	}

}
