package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

// Function to print command events for analytics or debugging
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Event:")
		fmt.Println("Timestamp:", event.Timestamp)
		fmt.Println("Command:", event.Command)
		fmt.Println("Parameters:", event.Parameters)
		fmt.Println("Event:", event.Event)
		fmt.Println()
	}
}

func main() {
	// Retrieve Slack tokens from environment variables
	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")

	// Check if the tokens are properly configured
	if slackBotToken == "" || slackAppToken == "" {
		log.Fatal("SLACK_BOT_TOKEN and SLACK_APP_TOKEN must be set in the environment")
	}

	// Initialize the Slack bot client
	bot := slacker.NewClient(slackBotToken, slackAppToken)

	// Start a goroutine to monitor and log command events
	go printCommandEvents(bot.CommandEvents())

	// Define a command to calculate age based on the year of birth
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Calculate age based on year of birth",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Retrieve the 'year' parameter from the command
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year. Please enter a valid number.")
				return
			}

			// Calculate the current year and derive the age
			currentYear := time.Now().Year()
			age := currentYear - yob
			if age < 0 {
				response.Reply("You entered a future year. Please check your input.")
				return
			}

			// Respond with the calculated age
			response.Reply(fmt.Sprintf("Your age is approximately %d years.", age))
		},
	})

	// Create a context to manage the bot's lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the bot and handle any errors that occur
	log.Println("Slack bot is running...")
	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
