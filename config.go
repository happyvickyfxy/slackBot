package main

import "github.com/nlopes/slack"

const botToken = "xoxb-223299515715-380375831267-eah8mDJHR4Wi7ta58QxbQSQx"

var (
	api = slack.New(botToken)
	rtm = api.NewRTM()
)
