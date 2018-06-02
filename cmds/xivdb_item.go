package cmds

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/coaltergeist/xiv-fc-helper-1/structs"
	"github.com/coaltergeist/xiv-fc-helper-1/xivdb"
	"github.com/paul-io/discordgo-embeds/embed"
)

func xivdbItemSearchCommand(s *discordgo.Session, m *discordgo.Message) {
	split := strings.Split(m.Content, " ")
	if len(split) <= 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Use %shelp item for correct usage", commandPrefix))
		return
	}
	itemName := split[1:]

	request := xivdb.NewSearchRequest()
	request.SetType(xivdb.ITEM)
	request.SetSearch(fmt.Sprintf("%s", strings.Join(itemName, "+")))
	data := request.Queue().Consume()

	var itemResults structs.XIVDBItemSearch
	if err := json.Unmarshal(data, &itemResults); err != nil {
		panic(err)
	}
	if itemResults.Items.Total == 0 || itemResults.Items.Total > 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Non-1 result count: %d\n<%s>", itemResults.Items.Total, request.GetEndpoint()))
		return
	}

	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.ITEM)
	queryRequest.SetID(itemResults.Items.Results[0].ID)
	data = queryRequest.Queue().Consume()

	var item structs.XIVDBItem
	if err := json.Unmarshal(data, &item); err != nil {
		panic(err)
	}

	em := embed.New()
	em.SetAuthor(item.Name, item.URLXivdb, "").
		SetThumbnail(item.Icon).
		SetDesc(item.Help)
	em.Color = s.State.UserColor(m.Author.ID, m.ChannelID)

	s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)

}

func init() {
	add(&command{
		execute: xivdbItemSearchCommand,
		trigger: "item",
		aliases: []string{"xivdbitem"},
		desc:    "Search for an item on xivdb",
		usage:   "item item name",
	})
}
