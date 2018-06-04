package cmds

import (
	"testing"

	"github.com/paul-io/xiv-fc-helper/xivdb"
	"github.com/stretchr/testify/assert"
)

func TestHistoricalSearch(t *testing.T) {
	channelID := "1234"
	entityIDs := []int{10, 20, 30, 40}
	saveHistoricalSearch(channelID, xivdb.RECIPE, entityIDs)

	id, err := getIDFromHistoricalSearch(channelID, xivdb.RECIPE, 1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, id, 10)
	_, err = getIDFromHistoricalSearch(channelID, xivdb.ITEM, 1)
	assert.Error(t, err)
	_, err = getIDFromHistoricalSearch(channelID, xivdb.RECIPE, 5)
	assert.Error(t, err)
}
