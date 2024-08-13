package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Dedo-Finger2/todo-list-cli/internal/errors"
	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func CompleteTask() {
	var taskID string
	utils.DefineFlagValue(&taskID)

	utils.Validator("TaskID", &taskID, []string{"not-null", "string-to-integer"})

	files, err := utils.GetAllFilesInToGoListFile()
	if err != nil {
		slog.Error("Failed to get all files from ToGoList folder.", "error", err)
		os.Exit(1)
	}

	toGoListFolderPath, err := utils.GetToGoListFolderPath()
	if err != nil {
		slog.Error("Failed to get ToGoList folder path.", "error", err)
		os.Exit(1)
	}

	userToGoListName := files[0].Name()

	userToGoListFileContent, err := utils.GetUserToGoListContent(toGoListFolderPath, userToGoListName)
	if err != nil {
		slog.Error("Failed on trying to get User's ToGo List file content", "error", err)
		os.Exit(1)
	}

	writeFile, err := os.OpenFile(filepath.Join(toGoListFolderPath, userToGoListName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		slog.Error("Error trying to open userToGoList file.", "error", err)
		os.Exit(1)
	}

	defer writeFile.Close()

	var taskFound = false

	for _, line := range userToGoListFileContent {
		var (
			fileTaskID        = line[0]
			fileTaskName      = line[1]
			fileTaskCreatedAt = line[2]
			fileTaskCompleted = line[3]
		)

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			slog.Warn("This task is already completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted != "true" {
			fileTaskCompleted = "true"
			slog.Info("Task completed!", "TaskID", taskID)
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		errors.ResourceNotFound("TaskID")
	}
}
