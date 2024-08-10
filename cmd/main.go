package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	cliPkg "github.com/Dedo-Finger2/todo-list-cli/internal/cli"
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
	TOGOLISTFILEEXTENSION string = ".csv"
)

var (
	toDoListName       string
	taskName           string
	taskID             string
	listCompletedTasks string
)

var cli = cliPkg.Cli{}

func init() {
	flag.Usage = cli.Help

	cli.AddFlag("name", "sets a name for your todo list.", &toDoListName)
	cli.AddFlag("task", "sets a name for a task.", &taskName)
	cli.AddFlag("id", "chooses a task id", &taskID)
	cli.AddFlag("completed", "list only completed tasks", &listCompletedTasks)

	cli.ParseFlagsAndCommands()
}

func createToGoList() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		toDoListName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		toDoListName = flag.Arg(2)
	}

	// Validation
	if toDoListName == "" {
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
	file, err := os.Create(filepath.Join(outputPath, filepath.Base(toDoListName+TOGOLISTFILEEXTENSION)))
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

	file, err := os.OpenFile(filepath.Join(outputPath, "ids.json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

	if len(files) < 2 {
		fmt.Println("to go list not created. try 'togo create --name Todo'")
		return
	}

	// Get the first file name.extension
	userToGoListFile := files[0].Name()

	// Get current time
	currentDateTime := time.Now().Local().Format(time.DateTime)

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

func listTasks() {
	if len(flag.Args()) > 1 {
		switch flag.Arg(1) {
		case "--completed":
			listCompletedOnly()
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
			formatedLine  = [3]string{taskID, taskName, taskCreatedAt}
		)

		if taskCompleted == "true" {
			continue
		}

		fmt.Println(strings.Join(formatedLine[:], "\t"))
	}
}

func listAllTasks() {
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
			formatedLine  = [4]string{taskID, taskName, taskCreatedAt, taskCompleted}
		)

		fmt.Println(strings.Join(formatedLine[:], "\t"))
	}
}

func listCompletedOnly() {
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
	for idx, line := range content {
		var (
			taskID        = line[0]
			taskName      = line[1]
			taskCreatedAt = strings.Split(line[2], " ")[0] // Removes time
			taskCompleted = line[3]
			formatedLine  = [3]string{taskID, taskName, taskCreatedAt}
		)

		if taskCompleted != "true" && idx != 0 {
			continue
		}

		fmt.Println(strings.Join(formatedLine[:], "\t"))
	}
}

func completeTask() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskID = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskID = flag.Arg(2)
	}

	if taskID == "" {
		fmt.Println("invalid task id.")
		return
	}

	if convertedValue, err := strconv.Atoi(taskID); err != nil || convertedValue == 0 {
		fmt.Println("invalid task id.")
		return
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

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	file.Close()

	writeFile, err := os.OpenFile(filepath.Join(outputPath, userToGoList), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	defer writeFile.Close()

	var taskFound = false

	for _, line := range content {
		var (
			fileTaskID        = line[0]
			fileTaskName      = line[1]
			fileTaskCreatedAt = line[2]
			fileTaskCompleted = line[3]
		)

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			fmt.Println("task already completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted != "true" {
			fileTaskCompleted = "true"
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		fmt.Println("task with id '" + taskID + "' was not found.")
	}
}

func unCompleteTask() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskID = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskID = flag.Arg(2)
	}

	if taskID == "" {
		fmt.Println("invalid task id.")
		return
	}

	if convertedValue, err := strconv.Atoi(taskID); err != nil || convertedValue == 0 {
		fmt.Println("invalid task id.")
		return
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

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	file.Close()

	writeFile, err := os.OpenFile(filepath.Join(outputPath, userToGoList), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	defer writeFile.Close()

	var taskFound = false

	for _, line := range content {
		var (
			fileTaskID        = line[0]
			fileTaskName      = line[1]
			fileTaskCreatedAt = line[2]
			fileTaskCompleted = line[3]
		)

		if taskID == fileTaskID && fileTaskCompleted == "false" {
			fmt.Println("task is not completed.")
			taskFound = true
		}

		if taskID == fileTaskID && fileTaskCompleted == "true" {
			fileTaskCompleted = "false"
			taskFound = true
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		fmt.Println("task with id '" + taskID + "' was not found.")
	}
}

func deleteTask() {
	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskID = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskID = flag.Arg(2)
	}

	if taskID == "" {
		fmt.Println("invalid task id.")
		return
	}

	if convertedValue, err := strconv.Atoi(taskID); err != nil || convertedValue == 0 {
		fmt.Println("invalid task id.")
		return
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

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	file.Close()

	writeFile, err := os.OpenFile(filepath.Join(outputPath, userToGoList), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	defer writeFile.Close()

	var taskFound = false

	for _, line := range content {
		var (
			fileTaskID        = line[0]
			fileTaskName      = line[1]
			fileTaskCreatedAt = line[2]
			fileTaskCompleted = line[3]
		)

		if taskID == fileTaskID {
			taskFound = true
			continue
		}

		writeFile.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", fileTaskID, fileTaskName, fileTaskCreatedAt, fileTaskCompleted))
	}

	if !taskFound {
		fmt.Println("task with id '" + taskID + "' was not found.")
	}
}

func main() {
	cli.Start()
}
