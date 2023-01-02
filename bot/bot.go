package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func Run() {
	fmt.Println("Starting bot...")
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	fmt.Println("The bot is now live!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if message.Content == "ping" {
		_, err := discord.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			log.Fatal(err)
		}
	}
}
