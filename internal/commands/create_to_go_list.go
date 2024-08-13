package commands

import (
	"flag"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func CreateToGoList() {
	var toDoListName string
	const TOGOLISTFILEEXTENSION string = ".csv"

	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		toDoListName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		toDoListName = flag.Arg(2)
	}

	// Validation
	if toDoListName == "" {
		log.Println("[ERROR]: To-go list's name cannot be empty.")
		return
	}

	// Get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Creates output path
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	// Make a new DIR
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Create the file
	file, err := os.Create(filepath.Join(outputPath, filepath.Base(toDoListName+TOGOLISTFILEEXTENSION)))
	if err != nil {
		panic(err)
	}

	// Create json id storage file
	utils.CreateJsonIdStorageFile()

	// Closes the file at the end of the func
	defer file.Close()

	// Write down the string into the file
	_, err = file.WriteString("ID,NAME,CREATED_AT,COMPLETED")
	if err != nil {
		panic(err)
	}
}
