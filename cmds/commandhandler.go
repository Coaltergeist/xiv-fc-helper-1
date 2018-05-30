package cmds

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	commandMap = make(map[string]*command)
	aliasMap   = make(map[string]string)
	l          = log.New(os.Stderr, "cmds: ", log.LstdFlags|log.Lshortfile)
)

// CommandPrefix is a global prefix for commands
const commandPrefix = "."

type command struct {
	execute      func(*discordgo.Session, *discordgo.Message)
	trigger      string
	aliases      []string
	desc         string
	usage        string
	commandCount int
	deleteAfter  bool
}

func add(c *command) {
	commandMap[c.trigger] = c
	for _, alias := range c.aliases {
		aliasMap[alias] = c.trigger
	}
	l.Printf("Added command %s | %d aliases", c.trigger, len(c.aliases))
}

// HandleCommand handles the command message
func handleCommand(s *discordgo.Session, m *discordgo.Message) {
	defer func() {
		if r := recover(); r != nil {
			l.Println("recovered in handleCommand")
		}
	}()
	cmdTrigger := strings.Split(m.Content, " ")[0][len(commandPrefix):]
	cmd, ok := commandMap[cmdTrigger]
	if !ok {
		cmd, ok = commandMap[aliasMap[cmdTrigger]]
		if !ok {
			return
		}
	}
	go func() {
		cmd.execute(s, m)
		cmd.commandCount++
		if cmd.deleteAfter {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		}
	}()
}

// OnMessage handles events on a user message
func OnMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || len(e.Message.Content) < len(commandPrefix) {
		return
	}
	if e.Message.Content[0:len(commandPrefix)] == commandPrefix {
		ch, err := s.State.Channel(e.ChannelID)
		if err != nil {
			ch, err = s.Channel(e.ChannelID)
			if err != nil {
				l.Println(err.Error())
				return
			}
		}
		mem, err := s.State.Member(ch.GuildID, e.Author.ID)
		if err != nil {
			mem, err = s.GuildMember(ch.GuildID, e.Author.ID)
			if err != nil {
				l.Println(err.Error())
				return
			}
		}

		if hasPermissions(mem) {
			handleCommand(s, e.Message)
		}
	}
}

// TODO: permission system
func hasPermissions(m *discordgo.Member) bool {
	return true
}
