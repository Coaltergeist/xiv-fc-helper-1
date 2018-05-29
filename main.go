package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
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

	if err = d.Open(); err != nil {
		panic(err)
	}

	l.Println("Bot is now running")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill)
	<-shutdown
}

func init() {
	l = log.New(os.Stderr, "main: ", log.LstdFlags|log.Lshortfile)
	fileContents, err := ioutil.ReadFile("configuration/MainConfig.json")
	if err != nil {
		l.Panicln(err.Error())
	}
	err = json.Unmarshal(fileContents, &conf)
	if err != nil {
		l.Panicln(err.Error())
	}
}
