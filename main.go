package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	SlackBotToken := os.Getenv("SLACK_BOT_OAUTH_TOKEN")
	SlackChannelId := os.Getenv("SLACK_CHANNEL_ID")
	DeployedProjectName := os.Getenv("DEPLOYED_PROJ_NAME")
	DeployedEnvironmentName := os.Getenv("DEPLOYED_ENV_NAME")
	api := slack.New(SlackBotToken)

	channelID, timestamp, err := api.PostMessage(SlackChannelId,
		slack.MsgOptionText(fmt.Sprintf("Successfully deployed & updated %v environment %v!",
			DeployedProjectName,
			DeployedEnvironmentName), false))

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
