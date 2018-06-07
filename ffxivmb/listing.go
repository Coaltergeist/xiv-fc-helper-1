package ffxivmb

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/coaltergeist/xiv-fc-helper-1/structs"
)

// GetMBData returns some shit lol
func GetMBData(ffxivmbURL string) (*structs.MBListing, error) {
	fmt.Println(ffxivmbURL)
	req := NewMBRequest()
	req.SetURL(ffxivmbURL)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return nil, err
	}

	listing := &structs.MBListing{}

	// currentPrices := doc.Find("table.table2")
	// //history := doc.Find("table.table3")
	// fmt.Println(currentPrices.Text())

	// currentPrices = currentPrices.Find("tbody")
	// //history = history.Find("tbody")
	// fmt.Println(currentPrices.Text())

	// lowestPrice := currentPrices.Find("tr.odd")
	// fmt.Println(lowestPrice.Text())

	// for i := range lowestPrice.Nodes {
	// 	if lowestPrice.Nodes[i].Data == "No data available in table" {
	// 		return nil, errors.New("That item has no listing")
	// 	}
	// 	if i == 1 {
	// 		(listing.Sale).Price = lowestPrice.Nodes[i].Data
	// 		fmt.Println((listing.Sale).Price)
	// 	} else if i == 2 {
	// 		(listing.Sale).Quantity = lowestPrice.Nodes[i].Data
	// 		fmt.Println((listing.Sale).Quantity)
	// 	} else if i == 3 {
	// 		(listing.Sale).Total = lowestPrice.Nodes[i].Data
	// 		fmt.Println((listing.Sale).Total)
	// 	}
	// }

	a := make([]string, 0)

	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		a = append(a, s.Text())
	})

	for i := 0; i < 100; i++ {
		if a[i] == "" {
			fmt.Println("ok")
		}
	}

	// for i := range currentPrices.Nodes {
	// }

	return listing, nil
}
