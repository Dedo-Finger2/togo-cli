package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func GetUserToGoListContent(toGoListFolderPath string, userToGoList string) ([][]string, error) {
	file, err := os.Open(filepath.Join(toGoListFolderPath, userToGoList))
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)

	content, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	file.Close()

	return content, nil
}
