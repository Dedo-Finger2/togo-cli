package utils

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func CreateJsonIdStorageFile() {
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
