package myconfig

import (
	"os/user"
)

/*
CheckUsername takes a string and compare it with the output of os/user/Current.Username return true or false based on output
*/

func CheckUsername(expected string) (bool, error) {
	myuser, err := user.Current()
	if err != nil {
		return false, err
	}

	if myuser.Username != expected {
		return false, nil
	}
	return true, nil
}
