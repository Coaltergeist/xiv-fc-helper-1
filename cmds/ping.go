package cmds

import (
	"github.com/bwmarrin/discordgo"
)

func pingCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "Hellooooo there!")
}

func init() {
	add(&command{
		execute: pingCommand,
		trigger: "ping",
		aliases: []string{"pingme"},
		desc:    "Am I alive?",
	})
}
