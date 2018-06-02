package cmds

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/paul-io/discordgo-embeds/colors"
	"github.com/paul-io/discordgo-embeds/embed"
	"github.com/paul-io/xiv-fc-helper/structs"
	"github.com/paul-io/xiv-fc-helper/xivdb"
)

func recipeCommand(s *discordgo.Session, m *discordgo.Message) {
	split := strings.Split(m.Content, " ")
	if len(split) <= 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Incorrect usage. Use %shelp item for correct usage", commandPrefix))
		return
	}
	recipeName := split[1:]

	request := xivdb.NewSearchRequest()
	request.SetType(xivdb.RECIPE)
	request.SetSearch(fmt.Sprintf("%s", strings.Join(recipeName, "+")))
	data := request.Queue().Consume()

	var recipeResults structs.XIVDBRecipeSearch
	if err := json.Unmarshal(data, &recipeResults); err != nil {
		panic(err)
	}
	if recipeResults.Recipes.Total == 0 || recipeResults.Recipes.Total > 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Non-1 result count: %d\n<%s>", recipeResults.Recipes.Total, request.GetEndpoint()))
		return
	}

	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.RECIPE)
	queryRequest.SetID(recipeResults.Recipes.Results[0].ID)
	data = queryRequest.Queue().Consume()

	var recipe structs.XIVDBRecipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		panic(err)
	}

	message, _ := s.ChannelMessageSend(m.ChannelID, "Hold on Kupo, I'm fetching that for you!")

	matMap := make(map[string]int)

	matsNeeded := ""

	tree := recipe.Tree
	for i := range tree {
		matsNeeded += fmt.Sprintf("\n%d %s", tree[i].Quantity, tree[i].Name)
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
	job := ""
	level := fmt.Sprintf("%d", recipe.Level)
	job += "Level " + level + " "
	for i := 0; i < recipe.Stars; i++ {
		job += "★"
	}

	title := fmt.Sprintf("%s", recipe.ClassName)

	em.SetAuthor(author, recipe.URLXivdb, "").
		SetThumbnail(recipe.Icon).
		SetColor(colors.Cyan()).
		AddField(title, job, false).
		AddField("Recipe", matsNeeded, true)

	description := ""
	for k, v := range matMap {
		description += fmt.Sprintf("%d %s\n", v, k)
	}
	em.AddField("All Mats Needed", description, true)
	s.ChannelMessageDelete(m.ChannelID, message.ID)
	s.ChannelMessageSendEmbed(m.ChannelID, em.MessageEmbed)
}

func materialRecursion(node structs.XIVDBRecipeTree, matMap map[string]int, multiplier int) map[string]int {
	//fmt.Println(node.Name)
	recipe, isRecipe := idLookup(node.ID)
	if !isRecipe {
		matMap[recipe.Name] += multiplier
	} else {
		tree := recipe.Tree
		for i := range tree {
			if len(tree[i].Synths) != 0 {
				multiplier *= tree[i].Quantity
				synths := tree[i].Synths
				for k := range synths {
					matMap = materialRecursion(synths[k], matMap, multiplier)
				}
			} else {
				matMap[tree[i].Name] += tree[i].Quantity
			}
		}
	}
	return matMap
}

func idLookup(id int) (structs.XIVDBRecipe, bool) {
	request := xivdb.NewSearchRequest()
	request.SetType(xivdb.RECIPE)
	request.SetSearch(fmt.Sprintf("%d", id))
	data := request.Queue().Consume()

	var recipeResults structs.XIVDBRecipeSearch
	var emptyRecipe structs.XIVDBRecipe
	if err := json.Unmarshal(data, &recipeResults); err != nil {
		panic(err)
	}
	if recipeResults.Recipes.Total == 0 || recipeResults.Recipes.Total > 1 {
		return emptyRecipe, false
	}

	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.RECIPE)
	queryRequest.SetID(recipeResults.Recipes.Results[0].ID)
	data = queryRequest.Queue().Consume()

	var recipe structs.XIVDBRecipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		panic(err)
	}

	return recipe, true
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
