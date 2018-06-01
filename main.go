package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/paul-io/xiv-bot/freecompany"
	"github.com/paul-io/xiv-bot/reminders"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/xiv-bot/cmds"
)

type config struct {
	BotToken string `json:"token"`
}

var (
	l    *log.Logger
	conf config
)

func main() {
	d, err := discordgo.New("Bot " + conf.BotToken)
	defer d.Close()

	if err != nil {
		panic(err)
	}
	l.Println("Starting bot")

	d.AddHandler(cmds.OnMessage)

	d.AddHandler(onGuildJoin)

	d.AddHandler(freecompany.ConfigureOnMessage)

	if err = d.Open(); err != nil {
		l.Panic(err)
	}

	l.Println("Bot is now running")

	reminders.RegisterSession(d)
	l.Println("Registered singleton session for reminders")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill)
	<-shutdown
}

func onGuildJoin(s *discordgo.Session, gc *discordgo.GuildCreate) {
	// Fires everytime on startup. Check if config exists already
	if _, err := os.Stat(fmt.Sprintf("resources/guilds/%s", gc.Guild.ID)); os.IsNotExist(err) {
		for _, channel := range gc.Channels {
			if channel.Type == discordgo.ChannelTypeGuildText {
				perms, _ := s.State.UserChannelPermissions(s.State.User.ID, channel.ID)
				if perms&discordgo.PermissionSendMessages > 0 {
					msg := []string{
						"Hi, I'm here to help manage your FFXIV FC!",
						"To start basic configuration, have someone with the Manage Server permission mention me and say \"configure\"",
					}
					s.ChannelMessageSend(channel.ID, strings.Join(msg, "\n"))
					break
				}
			}
		}
	}
}

func init() {
	if _, err := os.Stat("resources/guilds"); os.IsNotExist(err) {
		os.Mkdir("resources/guilds", os.ModePerm)
	}
	l = log.New(os.Stderr, "main: ", log.LstdFlags|log.Lshortfile)
	fileContents, err := ioutil.ReadFile("resources/config/MainConfig.json")
	if err != nil {
		l.Panic(err.Error())
	}
	err = json.Unmarshal(fileContents, &conf)
	if err != nil {
		l.Panic(err.Error())
	}
}
