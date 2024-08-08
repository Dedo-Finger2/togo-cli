package main

import (
	"encoding/json"
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

	// Create json id storage file
	createJsonIdStorageFile()

	// Closes the file at the end of the func
	defer file.Close()

	// Write down the string into the file
	_, err = file.WriteString("ID,NAME,CREATED_AT,COMPLETED")
	if err != nil {
		panic(err)
	}
}

func createJsonIdStorageFile() {
	// Get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Creates output path
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	// Create json file
	file, err := os.Create(filepath.Join(outputPath, "ids.json"))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Insert id 0
	file.WriteString(`{ "id": 0 }`)
}

func updateJsonStoreLastId(id int64) {
	// Get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	file, err := os.OpenFile(filepath.Join(outputPath, "ids.json"), os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(fmt.Sprintf(`{ "id":%d }`, id))
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

	// Create output
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")

	// Read all files in the to do list dir
	files, err := os.ReadDir(outputPath)
	if err != nil {
		panic(err)
	}

	// Get the first file name.extension
	userToGoListFile := files[0].Name()

	// Get current time
	currentDateTime := time.Now().Local().Format(time.RFC1123)

	// Read the content of the ids.json
	jsonData, err := os.ReadFile(filepath.Join(outputPath, "ids.json"))
	if err != nil {
		panic(err)
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
		panic(err)
	}

	// Create a new task
	task := Task{
		ID:        payload.ID + 1,
		Name:      taskName,
		CreatedAt: currentDateTime,
		Completed: false,
	}

	// Update the ID
	updateJsonStoreLastId(payload.ID + 1)

	// Open to do list file
	file, err := os.OpenFile(filepath.Join(outputPath, userToGoListFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Appends task to the to do list file
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
