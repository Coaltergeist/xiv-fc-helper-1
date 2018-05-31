package lodestone

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetCharacterURL(t *testing.T) {
	first := "neko"
	last := "mcdoogal"
	server := "Sargatanas"

	url, err := GetCharacterURL(first, last, server)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}

func TestGetCharacter(t *testing.T) {
	// Has FC no GC https://na.finalfantasyxiv.com/lodestone/character/21175993/
	// Has GC no FC https://na.finalfantasyxiv.com/lodestone/character/10168111/
	// Has FC and GC https://na.finalfantasyxiv.com/lodestone/character/12285614/
	url := "https://na.finalfantasyxiv.com/lodestone/character/12285614/"
	character, err := GetCharacterFromLodestone(url)
	if err != nil {
		panic(err)
	}
	spew.Dump(character)
}
