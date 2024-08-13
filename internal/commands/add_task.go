package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Dedo-Finger2/todo-list-cli/internal/types"
	"github.com/Dedo-Finger2/todo-list-cli/internal/utils"
)

func AddTask() {
	var taskName string

	// Get togolist name
	if strings.Contains(flag.Arg(1), "=") {
		taskName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		taskName = flag.Arg(2)
	}

	// Validation
	if taskName == "" {
		log.Println("[ERROR]: Task name cannot be empty.")
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
		log.Println("[WARN]: To go list not created yet. Try 'togo create --name Todo'")
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
		panic(err)
	}

	defer file.Close()

	// Appends task to the to do list file
	_, err = file.WriteString(fmt.Sprintf("\n%d,%s,%s,%t", task.ID, task.Name, task.CreatedAt, task.Completed))
	if err != nil {
		panic(err)
	}

	log.Println("[SUCCESS]: New task added to your to-go list.")
}
