package main

import (
	"log"
	"strings"
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
	bot.AddHandlerOnce(Chat)
	time.Sleep(3 * time.Second)
	log.Println("Discord launched successfully!")
	bot.UpdateStatus(0, "Factorio")

	/*
		embedme := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xb87333, // Copper
			Description: "Factorio Message Test Embed",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Field",
					Value:  "Value",
					Inline: false,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://wiki.factorio.com/images/Rocket_silo_entity.png",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://mods-data.factorio.com/assets/832490ba5d54b75061d9f5c959c75dfa883613ab.thumb.png",
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "Factorio Announcement: ",
		}

		Session.ChannelMessageSendEmbed(config.DiscordChannelId, embedme)
	*/
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
			// commands.RunCommand(name, s, m)
			return
		}
		// Pipes normal chat allowing it to be seen ingame
		//_, err := io.WriteString(Pipe, fmt.Sprintf("[Discord] <%s>: %s\r\n", m.Author.Username, m.ContentWithMentionsReplaced()))
		//if err != nil {
		//	log.Println("An error occurred when attempting to pass Discord chat to in-game Details: ", err)
		//}
		return
	}
}
