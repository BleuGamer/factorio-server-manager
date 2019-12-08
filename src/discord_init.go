package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func initDiscord(token string, channel string) {
	// Boolean to enable/disable discord integration.
	var shouldRun bool

	if token != "" && channel != "" {
		shouldRun = true
	} else {
		shouldRun = false
	}

	if shouldRun != true {
		log.Printf("Discord disabled.")
		return
	} else {
		log.Printf("Discord details found, spooling...")
	}

	// flog, err := os.OpenFile(config.FactorioLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Println("[Discord]: Cannot open factorio log, exiting: ", err)
	//	return
	//}

	//var Session *discordgo.Session
	bot, err := discordgo.New("Bot " + token)
	//Session = bot
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}
	time.Sleep(3 * time.Second)
	log.Println("Discord launched successfully!")

	bot.UpdateStatus(0, "Factorio")
}
