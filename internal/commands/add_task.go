package commands

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Dedo-Finger2/todo-list-cli/internal/errors"
	"github.com/Dedo-Finger2/todo-list-cli/internal/types"
	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func AddTask() {
	var taskName string
	utils.DefineFlagValue(&taskName)

	// Validation
	if taskName == "" {
		errors.ResourceCannotBeEmpty("Task Name")
		os.Exit(1)
	}

	// Get current user
	user, err := utils.GetCurrentUser()
	if err != nil {
		slog.Error("Error trying to get current user.", "error", err)
		os.Exit(1)
	}

	// Create output
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	// Read all files in the to do list dir
	files, err := os.ReadDir(outputPath)
	if err != nil {
		slog.Error("Error trying to get files in ToGoList DIR", "error", err)
		os.Exit(1)
	}

	if len(files) < 2 {
		errors.ResourceNotFound("To-go list")
		os.Exit(1)
	}

	// Get the first file name.extension
	userToGoListFile := files[0].Name()

	// Get current time
	currentDateTime := time.Now().Local().Format(time.DateTime)

	// Read the content of the ids.json
	jsonData, err := os.ReadFile(filepath.Join(outputPath, "ids.json"))
	if err != nil {
		slog.Error("Error trying to read ids.json file in ToGoList DIR.", "error", err)
		os.Exit(1)
	}

	// Payload structure
	type JsonPayload struct {
		ID int64 `json:"id"`
	}

	// Payload
	var payload JsonPayload

	// Parse json into the variable
	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		slog.Error("Error trying to parse json content from ids.json into a variable.", "error", err)
		os.Exit(1)
	}

	// Create a new task
	task := types.Task{
		ID:        payload.ID + 1,
		Name:      taskName,
		CreatedAt: currentDateTime,
		Completed: false,
	}

	// Update the ID
	utils.UpdateLastIdStored(payload.ID + 1)

	// Open to do list file
	file, err := os.OpenFile(filepath.Join(outputPath, userToGoListFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("Error trying to open to-go list file.", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	// Appends task to the to do list file
	_, err = file.WriteString(fmt.Sprintf("\n%d,%s,%s,%t", task.ID, task.Name, task.CreatedAt, task.Completed))
	if err != nil {
		slog.Error("Error trying to update to-do list file.", "error", err)
		os.Exit(1)
	}

	slog.Info("New task added to your to-go list.", "Task-Name", taskName)
}
