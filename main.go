package main

import (
	"flag"
	"log"
	"os"
	"simple_command/commands"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openInBrowser := flag.String("openInBrowser", "", "Example value: https://www.youtube.com/")
	notifyDiscord := flag.String("notifyDiscord", "", "Example usage: --notifyDiscord 'Hi, world my man'")

	flag.Parse()

	if *openInBrowser == "" && *notifyDiscord == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *openInBrowser != "" {
		commands.Openbrowser(*openInBrowser)
	}

	if *notifyDiscord != "" {
		commands.NotifyDiscord(*notifyDiscord)
	}

}
