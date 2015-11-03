package gocat

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlack(t *testing.T) {
	token := os.Getenv("SLACK_TOKEN")
	if "" != token {
		fmt.Printf("%s\n", token)
		api := slack.New(token)
		// fmt.Printf("%s\n", api)
		// If you set debugging, it will log all requests to the console
		// Useful when encountering issues
		// api.SetDebug(true)
		channels, err := api.GetChannels(true) //false)
		if err != nil {
			//fmt.Printf("%s\n", err)
			return
		}
		for _, channel := range channels {
			fmt.Printf("%s ", channel.Name)
			//fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
		}
	}
}
