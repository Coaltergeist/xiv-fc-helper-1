package cmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/xiv-fc-helper/freecompany"
	"github.com/paul-io/xiv-fc-helper/lodestone"
)

func configureUser(s *discordgo.Session, m *discordgo.Message) {
	split := strings.Split(m.Content, " ")
	if len(split) != 3 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Usage: `%sadduser first-name last-name", commandPrefix))
		return
	}
	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		ch, err = s.Channel(m.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
			l.Println(err)
			return
		}
	}
	fc, err := freecompany.GetFreeCompany(ch.GuildID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
		l.Println(err)
		return
	}
	first := split[1]
	last := split[2]
	world := fc.World
	if len(world) == 0 {
		s.ChannelMessageSend(m.ChannelID, "FC not configured! Please mention me with the word `configure` to configure your server if you have the Manager Server permission")
		return
	}

	searchMsg, _ := s.ChannelMessageSend(m.ChannelID, "Searching...")
	defer func() {
		if searchMsg != nil {
			s.ChannelMessageDelete(m.ChannelID, searchMsg.ID)
		}
	}()
	url, err := lodestone.GetCharacterURL(first, last, world)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %s", err))
		return
	}

	character, err := lodestone.GetCharacterFromLodestone(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %s", err))
		return
	}
	fc.Characters[m.Author.ID] = &freecompany.Character{
		ID:        character.ID,
		FirstName: character.FirstName,
		LastName:  character.LastName,
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Registered **%s %s** of **%s** to your account!", character.FirstName, character.LastName, strings.Title(fc.World)))
	freecompany.Serialize(ch.GuildID)
}

func init() {
	add(&command{
		execute: configureUser,
		trigger: "adduser",
		aliases: []string{"iam"},
		desc:    "Configure your own user for your FC/discord!",
	})
}
