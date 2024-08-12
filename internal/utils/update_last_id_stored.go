package utils

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func UpdateLastIdStored(id int64) {
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
