package main

import (
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type discordConfig struct {
	HasDiscordDetails bool
	ActiveSession     *discordgo.Session
}

var DiscordConfig discordConfig

func initDiscord(token string, channel string, adminChannel string) {

	if token != "" && channel != "" {
		DiscordConfig.HasDiscordDetails = true
	} else {
		DiscordConfig.HasDiscordDetails = false
	}

	if DiscordConfig.HasDiscordDetails != true {
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
	Session := bot
	DiscordConfig.ActiveSession = Session
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
	bot.AddHandler(messageCreate)
	bot.AddHandlerOnce(Chat)
	time.Sleep(3 * time.Second)
	log.Println("Discord launched successfully!")
	Session.UpdateStatus(0, "Factorio")

	discordEmbedServerLaunch()

}

// Send discord messages to game.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("[" + m.Author.Username + "] " + m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == config.DiscordChannelId {
		if strings.HasPrefix(m.Content, config.DiscordPrefix) {
			//command := strings.Split(m.Content[1:len(m.Content)], " ")
			//name := strings.ToLower(command[0])
			// TODO: commands
			//commands.RunCommand(name, s, m)
			return
		}
		// Pipes normal chat allowing it to be seen ingame
		//_, err := io.WriteString(Pipe, fmt.Sprintf("[Discord] <%s>: %s\r\n", m.Author.Username, m.ContentWithMentionsReplaced()))
		//if err != nil {
		//	log.Println("An error occurred when attempting to pass Discord chat to in-game Details: ", err)
		//}
		// --
		if FactorioServ.Running {
			data := "[Discord] <" + m.Author.Username + ">: " + m.ContentWithMentionsReplaced() + "\r\n"
			reqId, err := FactorioServ.SendCommand(data)
			if err != nil {
				log.Printf("Error sending rcon command: %s", err)
				return
			}

			log.Printf("Command send to Factorio: %s, with rcon request id: %v", data, reqId)

		}
		// --
		return
	}
}
