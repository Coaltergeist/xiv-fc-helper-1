package cmds

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/embed"
)

func helpCommand(s *discordgo.Session, m *discordgo.Message) {
	split := strings.Split(m.Content, " ")
	if len(split) == 2 {
		// Specific command
		var (
			givenName string
			cmdName   string
			cmd       *command
			ok        bool
		)
		givenName = split[1]
		if cmdName, ok = aliasMap[givenName]; !ok {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No command found for `%s`", givenName))
			return
		}
		if cmd, ok = commandMap[cmdName]; !ok {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No command found for valid command name `%s`", givenName))
			return
		}
		em := embed.New()
		em.Color = s.State.UserColor(m.Author.ID, m.ChannelID)
		em.SetTitle(strings.Title(cmd.trigger))
		em.SetDesc(fmt.Sprintf("**Description**: %s\n**Usage**: %s%s", cmd.desc, commandPrefix, cmd.usage))
		s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
	} else {
		msg := make([]string, 0, len(commandMap))
		for _, cmd := range commandMap {
			msg = append(msg, fmt.Sprintf("**%s** - %s", cmd.trigger, cmd.desc))
		}
		sort.Strings(msg)
		em := embed.New()
		em.Color = s.State.UserColor(m.Author.ID, m.ChannelID)
		em.SetDesc(strings.Join(msg, "\n"))
		var sendChannelID string
		privChannel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			sendChannelID = m.ChannelID
		} else {
			sendChannelID = privChannel.ID
			s.ChannelMessageSend(m.ChannelID, "PM'd you the command list!")
		}
		em.SetFooter("Use .help command-name to get detailed info!")
		s.ChannelMessageSendEmbed(sendChannelID, em.MessageEmbed)
	}
}

func init() {
	add(&command{
		execute: helpCommand,
		trigger: "help",
		aliases: []string{},
		desc:    "Get help on my commands!",
		usage:   "help [command-name]",
	})
}
