package main

import "github.com/nlopes/slack"

const botToken = "xxxx"

var (
	api = slack.New(botToken)
	rtm = api.NewRTM()
)
