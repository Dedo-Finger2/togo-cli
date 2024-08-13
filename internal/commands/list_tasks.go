package commands

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func ListTasks() {
	if len(flag.Args()) > 1 {
		switch flag.Arg(1) {
		case "--completed":
			listCompletedTasks()
			return
		case "--all":
			listAllTasks()
			return
		default:
			slog.Warn("Invalid argument. Listing uncompleted tasks instead.")
		}
	}

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

	defer file.Close()

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		slog.Error("Error trying to read CSV file content.", "error", err)
		os.Exit(1)
	}

	fmt.Println("-------------------------------")
	fmt.Println("# - " + userToGoList)
	fmt.Println("-------------------------------")
	for _, line := range content {
		var (
			taskID        = line[0]
			taskName      = line[1]
			taskCreatedAt = strings.Split(line[2], " ")[0] // Removes time
			taskCompleted = line[3]
			formattedLine = [3]string{taskID, taskName, taskCreatedAt}
		)

		if taskCompleted == "true" {
			continue
		}

		fmt.Println(strings.Join(formattedLine[:], "\t"))
	}

}
