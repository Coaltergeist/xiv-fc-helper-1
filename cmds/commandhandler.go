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

	permissions int
}

func add(c *command) {
	commandMap[c.trigger] = c
	for _, alias := range c.aliases {
		aliasMap[alias] = c.trigger
	}
	l.Printf("Added command %s | %d aliases", c.trigger, len(c.aliases))
}

// HandleCommand handles the command message
func handleCommand(s *discordgo.Session, e *discordgo.MessageCreate) {
	defer func() {
		if r := recover(); r != nil {
			l.Println("recovered in handleCommand")
		}
	}()
	cmdTrigger := strings.Split(e.Message.Content, " ")[0][len(commandPrefix):]
	cmd, ok := commandMap[cmdTrigger]
	if !ok {
		cmd, ok = commandMap[aliasMap[cmdTrigger]]
		if !ok {
			return
		}
	}
	go func() {
		if !hasPermissions(s, e, cmd) {
			return
		}
		cmd.execute(s, e.Message)
		cmd.commandCount++
		if cmd.deleteAfter {
			s.ChannelMessageDelete(e.Message.ChannelID, e.Message.ID)
		}
	}()
}

// OnMessage handles events on a user message
func OnMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || len(e.Message.Content) < len(commandPrefix) {
		return
	}
	if e.Message.Content[0:len(commandPrefix)] == commandPrefix {
		handleCommand(s, e)
	}
}

func hasPermissions(s *discordgo.Session, e *discordgo.MessageCreate, cmd *command) bool {
	perms, err := s.State.UserChannelPermissions(e.Author.ID, e.ChannelID)
	if err != nil {
		panic(err)
	}
	if cmd.permissions == 0 || cmd.permissions&perms > 0 {
		return true
	}
	return false
}
