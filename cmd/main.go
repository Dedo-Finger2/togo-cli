package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"
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
	_, err = file.WriteString("ID,NAME,CREATED_AT,COMPLETED")
	if err != nil {
		panic(err)
	}
}

func addTask() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskName = flag.Arg(2)
	}

	// Validation
	if taskName == "" {
		fmt.Println("name cannot be empty.")
		return
	}

	// Get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	files, err := os.ReadDir(outputPath)
	if err != nil {
		panic(err)
	}

	userToGoListFile := files[0].Name()

	file, err := os.OpenFile(filepath.Join(outputPath, userToGoListFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	currentDateTime := time.Now().Local().Format(time.RFC1123)

	task := Task{
		ID:        1,
		Name:      taskName,
		CreatedAt: currentDateTime,
		Completed: false,
	}

	_, err = file.WriteString(fmt.Sprintf("\n%d,%s,%s,%t", task.ID, task.Name, task.CreatedAt, task.Completed))
	if err != nil {
		panic(err)
	}

	fmt.Println("new task added.")
}

func main() {
	var command string = flag.Arg(0)

	if len(flag.Args()) > 1 {
		switch command {
		case "add":
			addTask()
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
