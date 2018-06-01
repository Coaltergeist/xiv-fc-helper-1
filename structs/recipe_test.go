package structs

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/paul-io/xiv-bot/xivdb"
)

func TestRecipe(t *testing.T) {
	id := 32225

	queryRequest := xivdb.NewQueryRequest()
	queryRequest.SetType(xivdb.RECIPE)
	queryRequest.SetID(id)
	data := queryRequest.Queue().Consume()

	var item XIVDBRecipe
	if err := json.Unmarshal(data, &item); err != nil {
		panic(err)
	}
	spew.Dump(item.Tree)
}
