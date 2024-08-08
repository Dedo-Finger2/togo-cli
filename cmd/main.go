package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

// Structs
type Task struct {
	ID        int64
	Name      string
	CreatedAt string
	Completed bool
}

// Constants
const (
	TOGOLISTFILEEXTENSION string = ".txt"
)

// Flags
var toGoListName string
var taskName string

// Initializing the flags
func init() {
	flag.StringVar(&toGoListName, "name", "", "sets a name for your todo list.")
	flag.StringVar(&taskName, "task", "", "sets a name for a task.")
	flag.Parse()
}

func createToGoList() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		toGoListName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		toGoListName = flag.Arg(2)
	}

	// Validation
	if toGoListName == "" {
		fmt.Println("name cannot be empty.")
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
	file, err := os.Create(filepath.Join(outputPath, filepath.Base(toGoListName+TOGOLISTFILEEXTENSION)))
	if err != nil {
		panic(err)
	}

	// Closes the file at the end of the func
	defer file.Close()

	// Write down the string into the file
	_, err = file.WriteString("ID,NAME,CREATED_AT,COMPLETED\n")
	if err != nil {
		panic(err)
	}
}

func main() {
	var command string = flag.Arg(0)

	if len(flag.Args()) > 1 {
		switch command {
		case "add":
			fmt.Println("Wanna a break from the adds?" + taskName)
			return
		case "create":
			createToGoList()
			return
		default:
			fmt.Println(fmt.Sprintf("command '%s' not found.", command))
			return
		}
	} else {
		fmt.Println("use --help to see all avaliables commands.")
		return
	}
}
