package cmds

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/embed"
	"github.com/paul-io/xiv-fc-helper/ffxivmb"
	"github.com/paul-io/xiv-fc-helper/structs"
	"github.com/paul-io/xiv-fc-helper/xivdb"
)

func marketboardSearchCommand(s *discordgo.Session, m *discordgo.Message) {
	//s.ChannelMessageSend(m.ChannelID, "Hellooooo there!")
	var (
		data        []byte
		queryID     int
		itemResults structs.XIVDBItemSearch
		request     *xivdb.SearchRequest
	)

	split := strings.Split(m.Content, " ")
	if len(split) < 2 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Use %shelp item for correct usage", commandPrefix))
		return
	}
	serverName := split[1]
	itemName := split[2:]

	if len(itemName) == 1 {
		if id, err := strconv.Atoi(itemName[0]); err == nil {
			// Multi result search
			queryID, err = getIDFromHistoricalSearch(m.ChannelID, xivdb.ITEM, id)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}
			goto finalQuery
		}
	}

	request = xivdb.NewSearchRequest()
	request.SetType(xivdb.ITEM)
	request.SetSearch(fmt.Sprintf("%s", strings.Join(itemName, "+")))
	data = request.Queue().Consume()

	if err := json.Unmarshal(data, &itemResults); err != nil {
		panic(err)
	}
	if itemResults.Items.Total == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No results found for %s", itemName))
		return
	}
	if itemResults.Items.Total > 1 {
		count := min(multiResultLimit, itemResults.Items.Total)
		multiResults := make([]int, count)
		for i := 0; i < count; i++ {
			// Create the historical results
			multiResults[i] = itemResults.Items.Results[i].ID
		}
		saveHistoricalSearch(m.ChannelID, xivdb.ITEM, multiResults)
		msg := make([]string, count)
		for i := 0; i < count; i++ {
			msg[i] = fmt.Sprintf("**%d)** %s", i+1, itemResults.Items.Results[i].Name)
		}
		em := embed.New().SetTitle("Multiple results found").SetDesc(strings.Join(msg, "\n")).
			SetFooter(fmt.Sprintf("use %sitem # to get results on an item on this list", commandPrefix))
		em.MessageEmbed.Color = s.State.UserColor(m.Author.ID, m.ChannelID)
		s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		return
	}
	queryID = itemResults.Items.Results[0].ID

finalQuery:
	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.ITEM)
	queryRequest.SetID(queryID)
	data = queryRequest.Queue().Consume()

	var item structs.XIVDBItem
	if err := json.Unmarshal(data, &item); err != nil {
		panic(err)
	}

	if item.IsUntradable == 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Sorry, %s is not able to be put on the marketboard", itemName))
		return
	}

	url := fmt.Sprintf("https://www.ffxivmb.com/items/%s/%d", serverName, item.ID)
	listing, err := ffxivmb.GetMBData(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	fmt.Println(listing)
	return
}

func init() {
	add(&command{
		execute: marketboardSearchCommand,
		trigger: "marketboard",
		aliases: []string{"pricecheck"},
		desc:    "Am I alive?",
		usage:   "ping",
	})
}
