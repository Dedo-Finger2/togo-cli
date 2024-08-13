package utils

import (
	"io/fs"
	"os"
)

func GetAllFilesInToGoListFile() ([]fs.DirEntry, error) {
	outputPath, err := GetToGoListFolderPath()
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(outputPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}
