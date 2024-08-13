package commands

import (
	"log/slog"
	"os"
	"time"

	"github.com/Dedo-Finger2/todo-list-cli/internal/types"
	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func AddTask() {
	var taskName string
	utils.DefineFlagValue(&taskName)

	utils.Validator("Task-Name", &taskName, []string{"not-null"})

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

	if err = utils.CheckToGoListFolderContent(&files); err != nil {
		slog.Error("Failed on validating required files in ToGoList folder.", "error", err)
		os.Exit(1)
	}

	userToGoListFile := files[0].Name()
	currentDateTime := time.Now().Local().Format(time.DateTime)

	lastInsertedTaskID, err := utils.GetLastIdUsed()
	if err != nil {
		slog.Error("Failed on trying to get last used ID.", "error", err)
		os.Exit(1)
	}

	task := types.Task{
		ID:        lastInsertedTaskID + 1,
		Name:      taskName,
		CreatedAt: currentDateTime,
		Completed: false,
	}

	utils.UpdateLastIdStored(lastInsertedTaskID + 1)

	if err = utils.WriteTaskInToGoListFile(&task, toGoListFolderPath, userToGoListFile); err != nil {
		slog.Error("Failed on trying to write task in ToGo List file.", "error", err)
		os.Exit(1)
	}

	slog.Info("New task added to your to-go list.", "Task-Name", taskName)
}
