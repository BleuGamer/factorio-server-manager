package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func discordEmbedServerLaunch(
	serverLocation string,
	targetChannel string,
	s *discordgo.Session) {

	embedme := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:     "Factorio Server Manager",
		Color:       0xb87333, // Copper
		Description: "Server manager details:",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Server Location:",
				Value: "https://" + serverLocation,
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

	s.ChannelMessageSendEmbed(targetChannel, embedme)

}
