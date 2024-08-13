package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dedo-Finger2/todo-list-cli/internal/types"
)

func WriteTaskInToGoListFile(task *types.Task, toGoListFolderPath string, userToGoListFile string) error {
	file, err := os.OpenFile(filepath.Join(toGoListFolderPath, userToGoListFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n%d,%s,%s,%t", task.ID, task.Name, task.CreatedAt, task.Completed))
	if err != nil {
		return err
	}

	return nil
}
