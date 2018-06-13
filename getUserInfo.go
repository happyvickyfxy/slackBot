package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func getAllUsers() ([]slack.User, error) {
	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return nil, err
	}

	for _, user := range users {
		fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
	}

	return users, nil
}

func getOneUser() {
	user, err := api.GetUserInfo("xiaoyu")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
