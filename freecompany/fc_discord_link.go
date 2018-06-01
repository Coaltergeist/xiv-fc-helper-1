// Package freecompany relates things about the model and storage of FCs
// Free Companies are FFXIV guilds
// By convention, a free company only uses 1 discord server at a time
// Thus, linking a 1 to 1 relationship between a FC and a discord server
package freecompany

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/paul-io/xiv-fc-helper/lodestone"
	"github.com/paul-io/xiv-fc-helper/reminders"

	"github.com/bwmarrin/discordgo"
)

var (
	discordLink map[string]*FreeCompany // Map discord ID -> FC
	serverNames []string

	l            *log.Logger
	channelRegex = regexp.MustCompile(channelRegexPhrase)
)

type configState int

const (
	// PRECONFIG represents being in a state prior to configuration
	PRECONFIG = iota
	// SERVER represents configuring which server/world the FC is in
	SERVER
	// SERVERCONFIRM represents confirming the server name
	SERVERCONFIRM
	// FCNAME represents configuring the FC itself
	FCNAME
	// FCCONFIRM represents confirming the FC name
	FCCONFIRM
	// RESETTIMERS represents reset timer reminders
	RESETTIMERS
	// RESETCONFIRM represents confirmation of reset timers
	RESETCONFIRM
	// FINISHED represents the final, finished state
	FINISHED
)

const (
	channelRegexPhrase = "<#(\\d+)>"
)

// A FreeCompany struct represents configuration data on a per-FC system
type FreeCompany struct {
	World string `json:"world"`
	ID    int    `json:"id"`
	Name  string `json:"name"`

	ConfigurationState configState `json:"configState"`
	ConfiguringUser    string      `json:"configuringUser"`

	ResetReminderChannel string `json:"resetReminderChannel"`

	Characters map[string]*Character `json:"characters"` // Map discord user ID -> Character
}

// A Character is a relationship between a user's character and their discord account
// Since FFXIV isn't conducive to alternate characters, it's generally safe to assume
// you can link the relationship of FC -> Discord Account/XIV Character (1 to 1)
type Character struct {
	ID int `json:"id"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// GetFreeCompany retrieves a free company for a given discord guild ID
func GetFreeCompany(discordGuildID string) (*FreeCompany, error) {
	fc, ok := discordLink[discordGuildID]
	if !ok {
		return nil, errors.New("free company not registered yet")
	}
	return fc, nil
}

// ConfigureOnMessage is a state machine to configure a discord server's FC
func ConfigureOnMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || e.Author.ID == s.State.User.ID {
		return
	}
	ch, err := s.State.Channel(e.ChannelID)
	if err != nil {
		ch, err = s.Channel(e.ChannelID)
		if err != nil {
			l.Panic(err)
		}
	}
	fc, ok := discordLink[ch.GuildID]
	if !ok {
		discordLink[ch.GuildID] = &FreeCompany{
			ConfigurationState: PRECONFIG,
		}
		fc = discordLink[ch.GuildID]
	}
	if fc.ConfigurationState == FINISHED {
		return
	}
	if strings.Contains(e.Message.Content, "stop") {
		fc.ConfigurationState = PRECONFIG
		fc.ConfiguringUser = ""
		s.ChannelMessageSend(e.ChannelID, "If you want to start again, just mention me with the word \"configure\" if you have the Manage Server permission!")
		return
	}
	switch fc.ConfigurationState {
	case PRECONFIG:
		break
	default:
		if e.Author.ID != fc.ConfiguringUser {
			return
		}
	}
	switch fc.ConfigurationState {
	case PRECONFIG:
		perms, err := s.State.UserChannelPermissions(e.Author.ID, e.ChannelID)
		if err != nil {
			l.Println(err)
			return
		}
		if perms&discordgo.PermissionManageServer > 0 && messageMentions(e.Message.Mentions, s.State.User.ID) {
			if strings.Contains(e.Message.ContentWithMentionsReplaced(), "configure") {
				// Start configuration and set this user as configuring
				msg := []string{
					"Let's start configuration! Please enter your FFXIV server name (i.e. Balmung)",
					"If you ever want to stop and restart, just type \"stop\"",
				}
				s.ChannelMessageSend(e.ChannelID, strings.Join(msg, "\n"))
				fc.ConfiguringUser = e.Author.ID
				fc.ConfigurationState = SERVER
			}
		}
	case SERVER:
		givenName := e.Message.Content
		if !containsIgnoreCase(serverNames, givenName) {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Server `%s` doesn't exist! Please try again", givenName))
			return
		}
		fc.World = strings.ToLower(givenName)
		s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Set your world as `%s`. Please respond with `yes` to confirm, or `no` to retry", fc.World))
		fc.ConfigurationState = SERVERCONFIRM
	case SERVERCONFIRM:
		if strings.EqualFold(e.Message.Content, "yes") {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("`%s` confirmed. Please enter your Free Company name", fc.World))
			fc.ConfigurationState = FCNAME
		} else if strings.EqualFold(e.Message.Content, "no") {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Please enter your server name"))
			fc.ConfigurationState = SERVER
			fc.World = ""
		} else {
			s.ChannelMessageSend(e.ChannelID, "Unknown input. Please respond with `yes` or `no`")
		}
	case FCNAME:
		givenName := e.Message.Content
		url, err := lodestone.GetFreeCompanyURL(givenName, fc.World)
		if err != nil {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Error: %s\nPlease try again", err))
			l.Println(err)
			return
		}
		retrievedFC, err := lodestone.GetFreeCompanyFromLodestone(url)
		if err != nil {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Error: %s\nPlease try again", err))
			l.Println(err)
			return
		}
		msg := []string{
			fmt.Sprintf("Retrieved this guild: %s %s on server %s", retrievedFC.Name, retrievedFC.Tag, retrievedFC.Server),
			fmt.Sprintf("Please respond with `yes` to confirm, or `no` to try another search"),
		}
		fc.ID = retrievedFC.ID
		fc.Name = retrievedFC.Name
		s.ChannelMessageSend(e.ChannelID, strings.Join(msg, "\n"))
		fc.ConfigurationState = FCCONFIRM
	case FCCONFIRM:
		if strings.EqualFold(e.Message.Content, "yes") {
			msg := []string{
				fmt.Sprintf("`%s` confirmed. Do you want reset timer reminders in a channel? This can be turned on/off at any time.", fc.Name),
				"Mention the channel you want to receive reminders, or `no` to decline",
			}
			s.ChannelMessageSend(e.ChannelID, strings.Join(msg, "\n"))
			fc.ConfigurationState = RESETTIMERS
		} else if strings.EqualFold(e.Message.Content, "no") {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Please enter your Free Company name"))
			fc.ConfigurationState = FCNAME
			fc.ID = 0
			fc.Name = ""
		} else {
			s.ChannelMessageSend(e.ChannelID, "Unknown input. Please respond with `yes` or `no`")
		}
	case RESETTIMERS:
		if strings.EqualFold(e.Message.Content, "no") {
			s.ChannelMessageSend(e.ChannelID, "Finished configuring! You can enable this feature later with .resettimers. Use .help to see what I can do!")
			fc.ConfigurationState = FINISHED
			return
		}
		channels := channelRegex.FindStringSubmatch(e.Message.Content)
		if len(channels) != 2 {
			s.ChannelMessageSend(e.ChannelID, "No channel mention found. If you don't want to configure this, just type `no`")
			return
		}
		channelID := channels[1]
		desiredChannel, err := s.State.Channel(channelID)
		if err != nil {
			if desiredChannel, err = s.Channel(channelID); err != nil {
				s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Error: %s\nPlease try again, or say `stop` to stop", err))
				l.Println(err)
				return
			}
		}
		perms, err := s.UserChannelPermissions(s.State.User.ID, desiredChannel.ID)
		if err != nil {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Error: %s\nPlease try again, or say `stop` to stop", err))
			l.Println(err)
			return
		}
		if perms&discordgo.PermissionSendMessages == 0 {
			s.ChannelMessageSend(e.ChannelID, "Error: I don't have permissions to talk in that channel\nPlease try again, or say `stop` to stop")
			l.Println(err)
			return
		}
		s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("I'll send daily/weekly reminders to the channel <#%s>. Is that ok? Respond with `yes` or `no`", channelID))
		fc.ResetReminderChannel = channelID
		fc.ConfigurationState = RESETCONFIRM
	case RESETCONFIRM:
		if strings.EqualFold(e.Message.Content, "yes") {
			msg := []string{
				fmt.Sprintf("<#%s> confirmed. Configuration finished! Type .help if you want to see what else I can do", fc.ResetReminderChannel),
			}
			s.ChannelMessageSend(e.ChannelID, strings.Join(msg, "\n"))
			fc.ConfigurationState = FINISHED
			reminders.RegisterReminders(fc.ResetReminderChannel, ch.GuildID)
		} else if strings.EqualFold(e.Message.Content, "no") {
			s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Please enter your Free Company name"))
			fc.ConfigurationState = RESETTIMERS
			fc.ResetReminderChannel = ""
		} else {
			s.ChannelMessageSend(e.ChannelID, "Unknown input. Please respond with `yes` or `no`")
		}
	}
	Serialize(ch.GuildID)

}

func messageMentions(arr []*discordgo.User, id string) bool {
	for _, user := range arr {
		if user.ID == id {
			return true
		}
	}
	return false
}

func containsIgnoreCase(arr []string, toFind string) bool {
	for _, found := range arr {
		if strings.EqualFold(found, toFind) {
			return true
		}
	}
	return false
}

func Serialize(discordServerID string) error {
	path := fmt.Sprintf("resources/guilds/%s/", discordServerID)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}
	file, err := os.Create(fmt.Sprintf("%s%s", path, "FCConfiguration.json"))
	if err != nil {
		return err
	}
	defer file.Close()
	fc, ok := discordLink[discordServerID]
	if !ok {
		return fmt.Errorf("no configuration in memory for guildID %s", discordServerID)
	}
	serialized, err := json.MarshalIndent(fc, "", "	")
	if err != nil {
		return err
	}
	file.Write(serialized)
	return nil
}

func deserialize() error {
	path := "resources/guilds"
	folders, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range folders {
		if !f.IsDir() {
			continue
		}
		info, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/FCConfiguration.json", path, f.Name()))
		if err != nil {
			test := &FreeCompany{
				ConfigurationState: PRECONFIG,
			}
			info, _ = json.Marshal(test)
			ioutil.WriteFile(fmt.Sprintf("%s/%s/FCConfiguration.json", path, f.Name()), info, 0644)
		}
		var fc FreeCompany
		json.Unmarshal(info, &fc)
		discordLink[f.Name()] = &fc
		if fc.Characters == nil {
			fc.Characters = make(map[string]*Character)
		}
		if len(fc.ResetReminderChannel) > 0 {
			if err = reminders.RegisterReminders(fc.ResetReminderChannel, f.Name()); err != nil {
				l.Println(err)
			}
		}
	}
	return nil
}

func init() {
	l = log.New(os.Stderr, "main: ", log.LstdFlags|log.Lshortfile)
	discordLink = make(map[string]*FreeCompany)
	deserialize()

	fileContents, err := ioutil.ReadFile("resources/config/ServerList.json")
	if err != nil {
		l.Panic(err.Error())
	}
	err = json.Unmarshal(fileContents, &serverNames)
	if err != nil {
		l.Panic(err.Error())
	}
}
