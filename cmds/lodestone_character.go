package cmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/embed"
	"github.com/paul-io/xiv-bot/lodestone"
)

func lodestoneCharacterSearchCommand(s *discordgo.Session, m *discordgo.Message) {
	split := strings.Split(m.Content, " ")
	if len(split) != 4 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Use %shelp character for correct usage", commandPrefix))
		return
	}
	em := embed.New()
	em.Color = s.State.UserColor(m.Author.ID, m.ChannelID)
	em.SetDesc("Searching...")
	searchMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
	defer func() {
		if searchMsg != nil {
			s.ChannelMessageDelete(m.ChannelID, searchMsg.ID)
		}
	}()

	world := split[1]
	first := split[2]
	last := split[3]

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
	em = embed.New()
	em.SetAuthor(fmt.Sprintf("%s %s of %s", character.FirstName, character.LastName, character.Server), character.LodestoneURL, "").
		SetThumbnail(character.JobImageURL)
	em.MessageEmbed.Color = s.State.UserColor(m.Author.ID, m.ChannelID)

	description := make([]string, 0, 5)
	description = append(description, fmt.Sprintf("**%s %s** | **%s**", character.Gender, character.Race, character.Faction))
	description = append(description, fmt.Sprintf("**Nameday**: %s | **Guardian**: %s", character.NameDay, character.Guardian))
	description = append(description, fmt.Sprintf("**Grand Company**: %s", character.GrandCompany))
	description = append(description, fmt.Sprintf("**Free Company**: %s", character.FreeCompany))
	em.SetDesc(strings.Join(description, "\n")).
		SetImage(character.ImageURL)

	s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
}

func init() {
	add(&command{
		execute: lodestoneCharacterSearchCommand,
		trigger: "character",
		aliases: []string{"char"},
		desc:    "Search for a character on the lodestone",
		usage:   "character first-name last-name",
	})
}
