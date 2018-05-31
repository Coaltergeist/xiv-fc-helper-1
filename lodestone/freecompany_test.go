package lodestone

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetFCURL(t *testing.T) {
	server := "Sargatanas"
	fc := "Sanguinem Tempestas"

	url, err := GetFreeCompanyURL(fc, server)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}

func TestGetFC(t *testing.T) {
	url := "https://na.finalfantasyxiv.com/lodestone/freecompany/9237305048202013337/"
	fc, err := GetFreeCompanyFromLodestone(url)
	if err != nil {
		panic(err)
	}
	spew.Dump(fc)
}
