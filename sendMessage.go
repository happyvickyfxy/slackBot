package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"io/ioutil"
	"encoding/json"
)

//send the message to the user
//return the status
func sendMessage(user string, msg string) error {
	chanID, err := findChannel(user)
	if err != nil {
		return err
	}
	params := slack.PostMessageParameters{}
	content, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("Failed to open file.")
	}
	json.Unmarshal(content, &params)
	_, _, err = api.PostMessage(chanID, msg, params)
	return err
}

//find a direct message channel to the user
//returns some status and the channel ID
func findChannel(user string) (string, error) {
	_, _, chanID, err := api.OpenIMChannel(user)
	if err != nil {
		fmt.Println("find channel failed")
		return "", err
	}
	return chanID, nil
}