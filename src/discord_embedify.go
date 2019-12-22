package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var toSession *discordgo.Session

func setupSession(s *discordgo.Session) {
	toSession = s
}

func discordEmbedServerLaunch() {

	embedme := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Factorio Server Manager",
		Color:       0xb87333, // Copper
		Description: "Server manager details:",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Server Location:",
				Value:  config.ServerLocation,
				Inline: false,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://camo.githubusercontent.com/d5f508f9a43def47f889777e95b9d6f972648b10/687474703a2f2f692e696d6775722e636f6d2f713774627a64482e706e67",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://mods-data.factorio.com/assets/832490ba5d54b75061d9f5c959c75dfa883613ab.thumb.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}

	toSession.ChannelMessageSendEmbed(config.DiscordAdminChannelId, embedme)

}

func discordEmbedFactorioLaunch() {

	embedme := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Factorio Server Launched!",
		Color:       0x00ff00, // Green
		Description: "Server details:",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Lobby Name:",
				Value:  config.LobbyName,
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	toSession.ChannelMessageSendEmbed(config.DiscordChannelId, embedme)
}

func discordEmbedFactorioClose() {

	embedme := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Factorio Server Shutdown!",
		Color:       0xff0000, // Red
		Description: "The server has stopped.",
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	toSession.ChannelMessageSendEmbed(config.DiscordChannelId, embedme)
}
