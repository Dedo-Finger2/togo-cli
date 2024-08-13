package commands

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/Dedo-Finger2/todo-list-cli/internal/errors"
	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func IncompleteTask() {
	var taskID string
	utils.DefineFlagValue(&taskID)

	utils.Validator("TaskID", &taskID, []string{"not-null", "string-to-integer"})

	// Get current user
	user, err := utils.GetCurrentUser()
	if err != nil {
		slog.Error("Error trying to get current user.", "error", err)
		os.Exit(1)
	}

	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	files, err := os.ReadDir(outputPath)
	if err != nil {
		slog.Error("Error trying to read ToGoLists DIR.", "error", err)
		os.Exit(1)
	}

	userToGoList := files[0].Name()

	file, err := os.Open(filepath.Join(outputPath, userToGoList))
	if err != nil {
		slog.Error("Error trying to open userToGoList file.", "error", err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		slog.Error("Error trying to read CSV file content.", "error", err)
		os.Exit(1)
	}

	file.Close()

	writeFile, err := os.OpenFile(filepath.Join(outputPath, userToGoList), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		slog.Error("Error trying to open userToGoList file.", "error", err)
		os.Exit(1)
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
			slog.Warn("This task is not completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			fileTaskCompleted = "false"
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		errors.ResourceNotFound("TaskID")
	}

	slog.Info("Un-completed task!", "TaskID", taskID)
}
