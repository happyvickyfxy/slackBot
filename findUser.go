package main

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"fmt"
)

//return a string of user id in a list of user
func findUserInList(users []slack.User, name string) (string, error) {
	for _, user := range users {
		if name == user.Profile.RealName {
			fmt.Println("user id: ", user.ID)
			return user.ID, nil
		}
	}
	return "", errors.New("user not found")
}
