package main

import "github.com/nlopes/slack"

//send the message to the user
//return the status
func sendMessage(user string, msg string) error {
	chanID, err := findChannel(user)
	if err != nil {
		return err
	}
	_, _, err = api.PostMessage(chanID, msg, slack.PostMessageParameters{})
	return err
}

//find a direct message channel to the user
//returns some status and the channel ID
func findChannel(user string) (string, error) {
	_, _, chanID, err := api.OpenIMChannel(user)
	if err != nil {
		return "", err
	}
	return chanID, nil
}