package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func handleMessage() {
	for msg := range rtm.IncomingEvents {
		fmt.Println("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "test"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			if ev.Text == "hello" {
				sendMessage("WAR1H0U3W", "hello, doggie")
			}

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
