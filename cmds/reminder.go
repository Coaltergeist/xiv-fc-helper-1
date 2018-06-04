package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/xiv-fc-helper/freecompany"
)

func toggleReminder(s *discordgo.Session, m *discordgo.Message) {
	ch, err := s.State.Channel(m.ChannelID)
	if err != nil {
		ch, err = s.Channel(m.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "error, please try again")
			l.Print(err)
			return
		}
	}
	if freecompany.GuildReceivingReminders(ch.GuildID) {
		err := freecompany.DeregisterReminder(ch.GuildID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
			l.Println(err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, "No longer receiving notifications")
	} else {
		err := freecompany.RegisterReminder(ch.GuildID, m.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
			l.Println(err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Now receiving notifications to <#%s>", m.ChannelID))
	}
}

func init() {
	add(&command{
		execute:     toggleReminder,
		trigger:     "reminder",
		aliases:     []string{"dailyreminders"},
		desc:        "Toggle on/off daily reminders!",
		usage:       "reminder",
		permissions: discordgo.PermissionManageServer,
	})
}
