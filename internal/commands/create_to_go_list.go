package commands

import (
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func CreateToGoList() {
	var toDoListName string
	utils.DefineFlagValue(&toDoListName)
	const TO_GO_LIST_FILE_EXTENSION string = ".csv"

	utils.Validator("toDoListName", &toDoListName, []string{"not-null"})

	user, err := utils.GetCurrentUser()
	if err != nil {
		slog.Error("Error trying to get current user.", "error", err)
		os.Exit(1)
	}

	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	if err = os.MkdirAll(outputPath, os.ModePerm); err != nil {
		slog.Error("Error trying to create ToGoList folder.", "error", err)
		os.Exit(1)
	}

	file, err := os.Create(filepath.Join(outputPath, filepath.Base(toDoListName+TO_GO_LIST_FILE_EXTENSION)))
	if err != nil {
		slog.Error("Error trying to create ToGoList.csv file.", "error", err)
		os.Exit(1)
	}

	utils.CreateJsonIdStorageFile()

	defer file.Close()

	_, err = file.WriteString("ID,NAME,CREATED_AT,COMPLETED")
	if err != nil {
		slog.Error("Error trying to add CSV headers", "error", err)
		os.Exit(1)
	}
}
