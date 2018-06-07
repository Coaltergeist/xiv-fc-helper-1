package ffxivmb

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/coaltergeist/xiv-fc-helper-1/structs"
)

// GetMBData returns some shit lol
func GetMBData(ffxivmbURL string) (*structs.MBListing, error) {
	req := NewMBRequest()
	req.SetURL(ffxivmbURL)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return nil, err
	}

	listing := &structs.MBListing{}
	fmt.Println("ok")

	currentPrices := doc.Find("table.table2")
	//history := doc.Find("table.table3")

	currentPrices = currentPrices.Find("tbody")
	//history = history.Find("tbody")

	lowestPrice := currentPrices.Find("tr.odd")

	for i := range lowestPrice.Nodes {
		if lowestPrice.Nodes[i].Data == "No data available in table" {
			return nil, errors.New("That item has no listing")
		}
		if i == 1 {
			(listing.Sale).Price = lowestPrice.Nodes[i].Data
			fmt.Println((listing.Sale).Price)
		} else if i == 2 {
			(listing.Sale).Quantity = lowestPrice.Nodes[i].Data
			fmt.Println((listing.Sale).Quantity)
		} else if i == 3 {
			(listing.Sale).Total = lowestPrice.Nodes[i].Data
			fmt.Println((listing.Sale).Total)
		}
	}

	// for i := range currentPrices.Nodes {
	// }

	return listing, nil
}
