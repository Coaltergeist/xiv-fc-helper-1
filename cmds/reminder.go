package cmds

import "github.com/bwmarrin/discordgo"

func toggleReminder(s *discordgo.Session, m *discordgo.Message) {

}

func init() {
	add(&command{
		execute: toggleReminder,
		trigger: "reminder",
		aliases: []string{"dailyreminders"},
		desc:    "Toggle on/off daily reminders!",
	})
}
