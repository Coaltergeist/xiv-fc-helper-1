package cmds

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paul-io/xiv-fc-helper/xivdb"

	"github.com/bwmarrin/discordgo"
)

var (
	commandMap = make(map[string]*command)
	aliasMap   = make(map[string]string)
	l          = log.New(os.Stderr, "cmds: ", log.LstdFlags|log.Lshortfile)
)

const (
	// CommandPrefix is a global prefix for commands
	commandPrefix = "."

	multiResultLimit = 15
)

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
	aliasMap[c.trigger] = c.trigger
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
	if perms, err := s.UserChannelPermissions(s.State.User.ID, e.ChannelID); err == nil {
		if perms&discordgo.PermissionSendMessages == 0 {
			// Can't speak in the channel? Don't bother
			return
		}
	}
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

var (
	// Map a channelID+RequestType (concatenated with "+") to a mapping of int (result line) to its ID and request type
	// This allows multiple requests of different types in the same channel to exist simultaneously
	// in a worst case scenario
	multiResultMap = make(map[string]*historicalSearch) // Some java level of declarations :^)
)

type historicalSearch struct {
	requestType xivdb.RequestType
	mapping     map[int]int // Map line # -> entity ID
}

// Save a historical search that can be called upon later
// Results need to be ordered by result number
func saveHistoricalSearch(channelID string, requestType xivdb.RequestType, entityIDs []int) {
	key := fmt.Sprintf("%s+%d", channelID, requestType)
	delete(multiResultMap, key)
	multiResultMap[key] = &historicalSearch{
		requestType: requestType,
		mapping:     make(map[int]int, len(entityIDs)),
	}
	// 1 indexed map
	for idx, entityID := range entityIDs {
		multiResultMap[key].mapping[idx+1] = entityID
	}
}

// Retrieve historical search results based on channel ID and request type
// Result ID is the 1 indexed line # of the search result
func getIDFromHistoricalSearch(channelID string, requestType xivdb.RequestType, resultID int) (int, error) {
	key := fmt.Sprintf("%s+%d", channelID, requestType)
	historicalSearch, ok := multiResultMap[key]
	if !ok {
		return -1, errors.New("prior multi-search result not found in this channel")
	}
	entityID, ok := historicalSearch.mapping[resultID]
	if !ok {
		return -1, fmt.Errorf("no results found for multi-result #%d", resultID)
	}
	return entityID, nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
