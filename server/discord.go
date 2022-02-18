package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type points struct {
	Vak  string `json:"vak"`
	Test toets  `json:"toets"`
}

type toets struct {
	Onderwerp string `json:"onderwerp"`
	Points    string `json:"points"`
	Pointsmax string `json:"pointsmax"`
	Procent   string `json:"procent"`
}

const authKey = "token"

var (
	BotID string
)

func main() {
	discord, err := discordgo.New("Bot " + authKey)
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
	readJSON()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//close discord and clean it up
	_ = discord.Close()
}

//This function will be called when a message is created in a channel
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
	if m.Content == "allo" {
		punt := readJSON()
		embed := &discordgo.MessageEmbed{
			Title:       punt[0].Test.Onderwerp,
			Description: punt[0].Test.Points + punt[0].Test.Pointsmax,
			Color:       chooseColor(punt[0].Test.Procent),
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
	}
}

//read the JSON file
func readJSON() []points {
	var puntenlijst []points

	//open file
	data, err := ioutil.ReadFile("./server/yeet.json")
	if err != nil {
		log.Fatalln("Could not read file!")
		return nil
	}

	//Change json to array
	err = json.Unmarshal(data, &puntenlijst)
	if err != nil {
		log.Fatalln("Could not unmarshal JSON!")
		return nil
	}

	return puntenlijst
}

//change the color of the embed according to the percentage
func chooseColor(punt string) int {
	var color int
	if punt < "50%" {
		color = 0xff0000
	} else if punt < "70%" {
		color = 0xff7700
	} else if punt >= "70%" {
		color = 0x60ff00
	}
	return color
}
