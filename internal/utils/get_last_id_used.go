package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Dedo-Finger2/todo-list-cli/internal/types"
)

func GetLastIdUsed() (int64, error) {
	toGoListFolderPath, err := GetToGoListFolderPath()
	if err != nil {
		return 0, err
	}

	jsonData, err := os.ReadFile(filepath.Join(toGoListFolderPath, "ids.json"))
	if err != nil {
		return 0, err
	}

	var payload types.JsonPayload

	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		return 0, nil
	}

	return payload.ID, nil
}
