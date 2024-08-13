package utils

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
)

func CheckToGoListFolderContent(files *[]fs.DirEntry) error {
	missingFiles := []string{}
	ExtensionCount := map[string]uint8{
		".csv":  0,
		".json": 0,
	}

	for _, value := range *files {
		if filepath.Ext(value.Name()) == ".csv" {
			ExtensionCount[".csv"]++
		}

		if filepath.Ext(value.Name()) == ".json" {
			ExtensionCount[".json"]++
		}
	}

	if ExtensionCount[".csv"] == 0 {
		missingFiles = append(missingFiles, "csv")
	}

	if ExtensionCount[".json"] == 0 {
		missingFiles = append(missingFiles, "json")
	}

	if len(missingFiles) > 0 {
		return errors.New("missing file(s) in ToGoList folder" + strings.Join(missingFiles, ", "))
	}

	return nil
}
