package cmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/embed"
	"github.com/paul-io/xiv-fc-helper/freecompany"
	"github.com/paul-io/xiv-fc-helper/lodestone"
)

func whoisUser(s *discordgo.Session, m *discordgo.Message) {
	mentions := m.Mentions
	if len(mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Please mention a user to lookup!")
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
	world := fc.World
	if len(world) == 0 {
		s.ChannelMessageSend(m.ChannelID, "FC not configured! Please mention me with the word `configure` to configure your server if you have the Manage Server permission")
		return
	}

	char, ok := fc.Characters[m.Mentions[0].ID]
	if !ok {
		s.ChannelMessageSend(m.ChannelID, "Discord user has not registered their character with this FC!")
		return
	}

	searchMsg, _ := s.ChannelMessageSend(m.ChannelID, "Searching...")
	defer func() {
		if searchMsg != nil {
			s.ChannelMessageDelete(m.ChannelID, searchMsg.ID)
		}
	}()

	character, err := lodestone.GetCharacterFromLodestone(fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/%d/", char.ID))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %s", err))
		return
	}

	em := embed.New()
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
		execute: whoisUser,
		trigger: "whois",
		aliases: []string{"who"},
		desc:    "Lookup someone else's character!",
	})
}
