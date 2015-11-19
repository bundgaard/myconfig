package myconfig

import (
	"os/user"
)

/*
CheckUsername takes a string and compare it with the output of os/user/Current.Username return true or false based on output
*/

func CheckUsername(expected string) bool {
	myuser, err := user.Current()
	if err != nil {
		return false
	}

	if myuser.Username != expected {
		return false
	}
	return true
}
