package utils

import "path"

func GetToGoListFolderPath() (string, error) {
	user, err := GetCurrentUser()
	if err != nil {
		return "", err
	}
	outputPath := path.Join(user.HomeDir, "Documents", "ToGoLists")
	return outputPath, nil
}
