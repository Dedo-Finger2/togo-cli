package utils

import "os/user"

func GetCurrentUser() (*user.User, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	return user, nil
}
