package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	dotenv "github.com/joho/godotenv"
)

func main() {

	// Create a new Discord session using the provided bot token.
	token := os.Getenv("TOKEN")

	if token == "" {
		err := dotenv.Load(".env")
		if err != nil {
			panic("Cannot start server, no token available")
		}
		token = os.Getenv("TOKEN")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "no" {
		s.ChannelMessageSend(m.ChannelID, "Ha Get Rekt noob!")
	}

	if m.Content == "bruh" {
		imgEmbed := embed.NewEmbed().SetImage("https://biographyhub.com/wp-content/uploads/2021/04/Arsenal-RL.jpg").MessageEmbed
		s.ChannelMessageSendEmbed(m.ChannelID, imgEmbed)
	}
}
