package lodestone

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/paul-io/xiv-fc-helper/structs"

	"github.com/PuerkitoBio/goquery"
)

// GetFreeCompanyURL resolves the URL to the lodesteone of an FC
func GetFreeCompanyURL(name, server string) (string, error) {
	server = strings.Title(strings.ToLower(server))
	url := fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/freecompany/?q=%s&worldname=%s&character_count=&activetime=&join=&house=&order=",
		url.QueryEscape(name), server)

	req := NewLodestoneRequest()
	req.SetURL(url)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return "", err
	}

	var foundHref string
	doc.Find(".entry__block").EachWithBreak(func(i int, s *goquery.Selection) bool {
		searchResultsName := s.Find("div.entry__freecompany__box").First().Find("p.entry__name").First().Text()
		if strings.EqualFold(strings.ToLower(searchResultsName), strings.ToLower(name)) {
			foundHref, _ = s.Attr("href")
			return false
		}
		return true
	})
	if len(foundHref) == 0 {
		return "", fmt.Errorf("fc not found in search: %s", url)
	}

	return fmt.Sprintf("http://na.finalfantasyxiv.com%s", foundHref), nil
}

// GetFreeCompanyFromLodestone gets an FC from a valid FC URL
func GetFreeCompanyFromLodestone(lodestoneURL string) (*structs.LodestoneFreeCompany, error) {
	req := NewLodestoneRequest()
	req.SetURL(lodestoneURL)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return nil, err
	}

	fc := &structs.LodestoneFreeCompany{}
	infoBox := doc.Find(".entry__freecompany__box").First()
	infoBox.Children().Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			// Allied faction
			fc.Faction = strings.TrimSpace(s.Text()[:strings.Index(s.Text(), "<")-1])
			fc.Reputation = s.Text()[strings.Index(s.Text(), "<")+1 : strings.Index(s.Text(), ">")]
		case 1:
			fc.Name = s.Text()
		case 2:
			fc.Server = strings.TrimSpace(s.Text())
		}
	})

	fc.Slogan = doc.Find(".freecompany__text.freecompany__text__message").Text()
	fc.Tag = doc.Find(".freecompany__text.freecompany__text__tag").Text()

	fc.ID, _ = strconv.Atoi(string(lodestoneURL[strings.LastIndex(lodestoneURL[:len(lodestoneURL)-1], "/")+1 : len(lodestoneURL)-1]))
	return fc, nil
}
