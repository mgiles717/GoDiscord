package main

import (
	"GoDiscord/bot"
	"encoding/json"
	"log"
	"os"
)

func main() {
	config, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Discord token not found/invalid!")
	}

	var jsonContent map[string]string
	err = json.Unmarshal(config, &jsonContent)
	if err != nil {
		log.Fatal("Error unmarshalling config.json!")
	}

	botToken := jsonContent["BOT_TOKEN"]
	bot.BotToken = botToken
	bot.Run()
}
