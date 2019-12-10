package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func discordEmbedServerLaunch(
	serverLocation string,
	fserverLocation string,
	s *discordgo.Session) {

	embedme := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xb87333, // Copper
		Description: "Factorio server details:",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "IP:PORT",
				Value:  fserverLocation,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name: "Lobby",
				// TODO: Variable this.
				Value:  "Seablock Testing Initiative",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://camo.githubusercontent.com/d5f508f9a43def47f889777e95b9d6f972648b10/687474703a2f2f692e696d6775722e636f6d2f713774627a64482e706e67",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://mods-data.factorio.com/assets/832490ba5d54b75061d9f5c959c75dfa883613ab.thumb.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Factorio Server Manager is now live at: " + serverLocation,
	}

	s.ChannelMessageSendEmbed(config.DiscordChannelId, embedme)

}
