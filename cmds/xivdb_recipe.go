package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/colors"
	"github.com/paul-io/discordgo-embeds/embed"
	"github.com/paul-io/xiv-fc-helper/structs"
	"github.com/paul-io/xiv-fc-helper/xivdb"
)

func recipeCommand(s *discordgo.Session, m *discordgo.Message) {
	var (
		data          []byte
		queryID       int
		recipeResults structs.XIVDBRecipeSearch
		request       *xivdb.SearchRequest
	)
	split := strings.Split(m.Content, " ")
	if len(split) < 2 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Use %shelp recijpe for correct usage", commandPrefix))
		return
	}
	recipeName := split[1:]
	if len(recipeName) == 1 {
		if id, err := strconv.Atoi(recipeName[0]); err == nil {
			// Multi result search
			queryID, err = getIDFromHistoricalSearch(m.ChannelID, xivdb.RECIPE, id)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}
			goto finalQuery
		}
	}

	request = xivdb.NewSearchRequest()
	request.SetType(xivdb.RECIPE)
	request.SetSearch(fmt.Sprintf("%s", strings.Join(recipeName, "+")))
	data = request.Queue().Consume()

	if err := json.Unmarshal(data, &recipeResults); err != nil {
		panic(err)
	}
	if recipeResults.Recipes.Total == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No results found for %s", recipeName))
		return
	}
	if recipeResults.Recipes.Total > 1 {
		count := min(multiResultLimit, recipeResults.Recipes.Total)
		multiResults := make([]int, count)
		for i := 0; i < count; i++ {
			// Create the historical results
			multiResults[i] = recipeResults.Recipes.Results[i].ID
		}
		saveHistoricalSearch(m.ChannelID, xivdb.RECIPE, multiResults)
		msg := make([]string, count)
		for i := 0; i < count; i++ {
			msg[i] = fmt.Sprintf("**%d)** %s", i+1, recipeResults.Recipes.Results[i].Name)
		}
		em := embed.New().SetTitle("Multiple results found").SetDesc(strings.Join(msg, "\n")).
			SetFooter(fmt.Sprintf("use %srecipe # to get results on a recipe on this list", commandPrefix))
		em.MessageEmbed.Color = s.State.UserColor(m.Author.ID, m.ChannelID)
		s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
		return
	}
	queryID = recipeResults.Recipes.Results[0].ID

finalQuery:
	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.RECIPE)
	queryRequest.SetID(queryID)
	data = queryRequest.Queue().Consume()

	var recipe structs.XIVDBRecipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		panic(err)
	}

	message, _ := s.ChannelMessageSend(m.ChannelID, "Hold on Kupo, I'm fetching that for you!")

	matMap := make(map[string]int)

	var matsNeeded bytes.Buffer

	tree := recipe.Tree
	for i := range tree {
		matsNeeded.WriteString(fmt.Sprintf("\n%d %s", tree[i].Quantity, tree[i].Name))
		if len(tree[i].Synths) != 0 {
			quantity := tree[i].Quantity
			synths := tree[i].Synths
			for k := range synths {
				matMap = materialRecursion(synths[k], matMap, quantity)
			}
		} else {
			matMap[tree[i].Name] += tree[i].Quantity
		}
	}

	em := embed.New()

	author := recipe.Name
	var job bytes.Buffer
	job.WriteString(fmt.Sprintf("Level %d ", recipe.Level))
	for i := 0; i < recipe.Stars; i++ {
		job.WriteString("★")
	}

	em.SetAuthor(author, recipe.URLXivdb, "").
		SetThumbnail(recipe.Icon).
		SetColor(colors.Cyan()).
		AddField(recipe.ClassName, job.String(), false).
		AddField("Recipe", matsNeeded.String(), true)

	var (
		realMats  bytes.Buffer
		catalysts bytes.Buffer
	)
	for k, v := range matMap {
		if strings.Contains(k, " Crystal") || strings.Contains(k, " Shard") || strings.Contains(k, " Cluster") {
			catalysts.WriteString(fmt.Sprintf("%d %s\n", v, k))
		} else {
			realMats.WriteString(fmt.Sprintf("%d %s\n", v, k))
		}
	}
	em.AddField("All Mats", realMats.String(), true)
	em.AddField("All Catalysts", catalysts.String(), true)
	if message != nil {
		s.ChannelMessageDelete(m.ChannelID, message.ID)
	}
	s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
}

func materialRecursion(node structs.XIVDBRecipeTree, matMap map[string]int, multiplier int) map[string]int {
	recipe, isRecipe := idLookup(node.ID)
	//fmt.Println(recipe.Name)
	if !isRecipe {
		matMap[recipe.Name] += multiplier
	} else {
		yield := recipe.CraftQuantity
		multiplier := int(math.Ceil(float64(multiplier) / float64(yield)))
		tree := recipe.Tree
		for i := range tree {
			if len(tree[i].Synths) != 0 {
				synths := tree[i].Synths
				for k := range synths {
					matMap = materialRecursion(synths[k], matMap, (multiplier * tree[i].Quantity))
				}
			} else {
				matMap[tree[i].Name] += (multiplier * tree[i].Quantity)
			}
		}
	}
	return matMap
}

func idLookup(id int) (*structs.XIVDBRecipe, bool) {
	request := xivdb.NewSearchRequest()
	request.SetType(xivdb.RECIPE)
	request.SetSearch(fmt.Sprintf("%d", id))
	data := request.Queue().Consume()

	var recipeResults structs.XIVDBRecipeSearch
	if err := json.Unmarshal(data, &recipeResults); err != nil {
		panic(err)
	}
	if recipeResults.Recipes.Total == 0 || recipeResults.Recipes.Total > 1 {
		return &structs.XIVDBRecipe{}, false
	}

	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.RECIPE)
	queryRequest.SetID(recipeResults.Recipes.Results[0].ID)
	data = queryRequest.Queue().Consume()

	var recipe structs.XIVDBRecipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		panic(err)
	}

	return &recipe, true
}

func init() {
	add(&command{
		execute: recipeCommand,
		trigger: "recipe",
		aliases: []string{"craft"},
		desc:    "Gives the relevant info of a crafting recipe",
		usage:   "recipe recipe-name",
	})
}
