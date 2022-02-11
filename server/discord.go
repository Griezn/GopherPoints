package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const auth_key = "token"

var (
	BotID string
)

func main() {
	discord, err := discordgo.New("Bot " + auth_key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := discord.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID

	discord.AddHandler(messageHandler)
	err = discord.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running...")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
	if m.Content == "allo" {
		_, err := s.ChannelMessageSend(m.ChannelID, "yeet")
		if err != nil {
			return
		}
	}
}
