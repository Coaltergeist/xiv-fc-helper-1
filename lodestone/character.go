package lodestone

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/paul-io/xiv-fc-helper/structs"
)

// GetCharacterURL resolves the URL to the lodestone of a character based on their
// name and server
func GetCharacterURL(firstName, lastName, server string) (string, error) {
	server = strings.Title(strings.ToLower(server))
	url := fmt.Sprintf("http://na.finalfantasyxiv.com/lodestone/character/?q=%s+%s&worldname=%s&classjob=&race_tribe=&order=",
		firstName, lastName, server)

	req := NewLodestoneRequest()
	req.SetURL(url)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return "", err
	}

	var foundHref string
	doc.Find(".entry__link").EachWithBreak(func(i int, s *goquery.Selection) bool {
		searchResultsName := s.Find("div.entry__box--world").First().Find("p.entry__name").First().Text()
		if strings.EqualFold(searchResultsName, fmt.Sprintf("%s %s", firstName, lastName)) {
			foundHref, _ = s.Attr("href")
			return false
		}
		return true
	})
	if len(foundHref) == 0 {
		return "", fmt.Errorf("character not found in search: %s", url)
	}

	return fmt.Sprintf("http://na.finalfantasyxiv.com%s", foundHref), nil
}

// GetCharacterFromLodestone uses the given valid lodestone
func GetCharacterFromLodestone(lodestoneURL string) (*structs.LodestoneCharacter, error) {
	req := NewLodestoneRequest()
	req.SetURL(lodestoneURL)
	body := req.queue().Consume()
	defer body.Body.Close()
	doc, err := goquery.NewDocumentFromReader(body.Body)
	if err != nil {
		return nil, err
	}

	character := &structs.LodestoneCharacter{}
	character.ImageURL, _ = doc.Find("div.character__detail__image").Find("img").First().Attr("src")

	// Count # of profile boxes - 5 has both grand/free, 4 has one or the other, 3 has neither
	profileElements := doc.Find("div.character__profile__data__detail").First().Find("div.character-block").Size()
	doc.Find("div.character-block__box").Each(func(i int, s *goquery.Selection) {
		switch i {
		// Race/gender case
		case 0:
			raceGender, _ := s.Find("p.character-block__name").First().Html()
			character.Race = html.UnescapeString(raceGender[0:strings.Index(raceGender, "<")])
			character.Faction = raceGender[strings.Index(raceGender, ">")+1 : strings.LastIndex(raceGender, "/")-1]
			if strings.Contains(raceGender, "♀") {
				character.Gender = "Female"
			} else {
				character.Gender = "Male"
			}
		// Birth/name case
		case 1:
			character.NameDay = html.UnescapeString(s.Find("p.character-block__birth").First().Text())
			character.Guardian = html.UnescapeString(s.Find("p.character-block__name").First().Text())
		// City state case
		case 2:
			character.CityState = s.Find("p.character-block__name").First().Text()
		case 3:
			// Grand company (or free if profile elements == 4)
			if profileElements == 5 || doc.Find("div.character__freecompany__crest").Size() == 0 {
				character.GrandCompany = html.UnescapeString(s.Find("p.character-block__name").First().Text())
				if doc.Find("div.character__freecompany__crest").Size() == 0 {
					character.FreeCompany = "none"
				}
			} else if profileElements == 4 && doc.Find("div.character__freecompany__crest").Size() > 0 {
				character.FreeCompany = html.UnescapeString(doc.Find("div.character__freecompany__name").First().Find("a").First().Text())
				character.GrandCompany = "none"
			}
		case 4:
			// Free company
			character.FreeCompany = html.UnescapeString(s.Find("div.character__freecompany__name").First().Find("a").First().Text())
		}
	})

	if profileElements == 3 {
		character.FreeCompany = "none"
		character.GrandCompany = "none"
	}

	character.JobImageURL, _ = doc.Find("div.character__class_icon").First().Find("img").First().Attr("src")
	character.LodestoneURL = lodestoneURL

	name := doc.Find(".frame__chara__name")
	character.FirstName = html.UnescapeString(strings.Split(name.Text(), " ")[0])
	character.LastName = html.UnescapeString(strings.Split(name.Text(), " ")[1])
	var (
		title  *goquery.Selection
		server *goquery.Selection
	)
	if name.Parent().Children().Size() == 2 {
		// No title
		character.Server = name.Next().Text()
	} else {
		if name.Prev().HasClass("frame__chara__title") {
			// Title is before
			title = name.Prev()
			server = name.Next()
			character.TitleBefore = true
		} else {
			// Title is next
			title = name.Next()
			server = title.Next()
		}
		character.Title = title.Text()
		character.Server = server.Text()
	}

	// This doesn't work if the URL doesn't end in a '/' but... yikes
	character.ID, _ = strconv.Atoi(string(lodestoneURL[strings.LastIndex(lodestoneURL[:len(lodestoneURL)-1], "/")+1 : len(lodestoneURL)-1]))

	return character, nil
}
