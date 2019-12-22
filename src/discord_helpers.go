package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
)

/*
 * Under the use of the MIT license, these implementations and modifications
 * are made from works originating from Allen Lydiard:
 *
 * https://github.com/FactoKit/FactoCord
 */

var logFile string

func setupChat(lFile string) {
	logFile = lFile
}

// Pipe chat to discord.
func Chat(s *discordgo.Session, m *discordgo.MessageCreate) {
	for {
		t, err := tail.TailFile(filepath.Join(config.FactorioDir, logFile), tail.Config{Follow: true})
		if err != nil {
			log.Println("Could not tail factorio.log", err)
		}
		for line := range t.Lines {
			if strings.Contains(line.Text, "[CHAT]") || strings.Contains(line.Text, "[JOIN]") || strings.Contains(line.Text, "[LEAVE]") {
				if !strings.Contains(line.Text, "<server>") {

					if strings.Contains(line.Text, "[JOIN]") ||
						strings.Contains(line.Text, "[LEAVE]") {
						TmpList := strings.Split(line.Text, " ")
						// Don't hard code the channelID! }:<
						s.ChannelMessageSend(config.DiscordChannelId, fmt.Sprintf("%s", strings.Join(TmpList[3:], " ")))
					} else {

						TmpList := strings.Split(line.Text, " ")
						TmpList[3] = strings.Replace(TmpList[3], ":", "", -1)
						if strings.Contains(strings.Join(TmpList, " "), "@") {
							index := LocateMentionPosition(TmpList)

							for _, position := range index {
								User := SearchForUser(TmpList[position])

								if User == nil {
									continue
								}
								TmpList[position] = User.Mention()
							}

						}

						s.ChannelMessageSend(config.DiscordChannelId, fmt.Sprintf("<%s>: %s", TmpList[3], strings.Join(TmpList[4:], " ")))
					}
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// SearchForUser searches for the user to be mentioned.
func SearchForUser(name string) *discordgo.User {
	name = strings.Replace(name, "@", "", -1)
	for _, user := range Users {
		if strings.ToLower(user.Nick) == strings.ToLower(name) ||
			strings.ToLower(user.User.Username) == strings.ToLower(name) {
			return user.User
		}
	}
	return nil
}

// LocateMentionPosition locates the position in a string list for the discord mention.
func LocateMentionPosition(List []string) []int {
	positionlist := []int{}
	for i, String := range List {
		if strings.Contains(String, "@") {
			positionlist = append(positionlist, i)
		}
	}
	return positionlist
}

// UserList is a struct for member info.
type UserList struct {
	UserID string
	Nick   string
	User   *discordgo.User
}

// Users is a slice of UserList.
var Users []UserList

// CacheDiscordMembers caches the users list to be searched.
func CacheDiscordMembers(s *discordgo.Session) {
	// Clear the users list
	Users = nil

	GuildChannel, err := s.Channel(config.DiscordChannelId)
	if err != nil {
		log.Println("An error occurred when attempting to read the Discord Guild Details: ", err)
	}
	GuildID := GuildChannel.GuildID
	members, err := s.State.Guild(GuildID)
	if err != nil {
		log.Println("An error occurred when attempting to read the Discord Guild Members Details: ", err)
	}
	for _, member := range members.Members {
		Users = append(Users, UserList{UserID: member.User.ID, Nick: member.Nick,
			User: member.User})
	}
}
