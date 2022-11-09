package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {

	// Setting up the Slack env
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4313735501776-4302714749857-8YBAeSPI12H6eKaEKxLwmhES")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A048J1760GJ-4275501196311-0e75d1b037832e88542f8f84ab05a73102d8681a0bc1bb263ae0a6068fc9ef09")

	// Creating Bot

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Creating a goroutine

	go printCmdEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func printCmdEvents(analyticsChannels <-chan *slacker.CommandEvent) {

	for event := range analyticsChannels {
		fmt.Println("Command Plattee")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("\n")

	}

}
