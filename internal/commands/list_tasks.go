package commands

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
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
			fmt.Println("invalid argument. listing uncompleted tasks instead.")
		}
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

	defer file.Close()

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
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
