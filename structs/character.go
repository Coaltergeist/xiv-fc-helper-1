package structs

// An XIVDBCharacterSearch struct represents the response of a search for a Character
type XIVDBCharacterSearch struct {
	Characters struct {
		Results []struct {
			Name        string `json:"name"`
			Server      string `json:"server"`
			Icon        string `json:"icon"`
			LastUpdated string `json:"last_updated"`
			ID          int    `json:"id"`
			URL         string `json:"url"`
			URLType     string `json:"url_type"`
			URLAPI      string `json:"url_api"`
			URLXivdb    string `json:"url_xivdb"`
		} `json:"results"`
		Total  int `json:"total"`
		Paging struct {
			Page  int   `json:"page"`
			Total int   `json:"total"`
			Pages []int `json:"pages"`
			Next  int   `json:"next"`
			Prev  int   `json:"prev"`
		} `json:"paging"`
	} `json:"characters"`
}

// XIVDBCharacter represents a queried character
type XIVDBCharacter struct {
	LodestoneID                  int    `json:"lodestone_id"`
	Name                         string `json:"name"`
	Server                       string `json:"server"`
	Avatar                       string `json:"avatar"`
	Added                        string `json:"added"`
	LastUpdated                  string `json:"last_updated"`
	LastSynced                   string `json:"last_synced"`
	DataLastChanged              string `json:"data_last_changed"`
	DataHash                     string `json:"data_hash"`
	UpdateCount                  int    `json:"update_count"`
	AchievementsLastUpdated      string `json:"achievements_last_updated"`
	AchievementsLastChanged      string `json:"achievements_last_changed"`
	AchievementsPublic           int    `json:"achievements_public"`
	AchievementsScoreReborn      int    `json:"achievements_score_reborn"`
	AchievementsScoreLegacy      int    `json:"achievements_score_legacy"`
	AchievementsScoreRebornTotal int    `json:"achievements_score_reborn_total"`
	AchievementsScoreLegacyTotal int    `json:"achievements_score_legacy_total"`
	Deleted                      string `json:"deleted"`
	Priority                     int    `json:"priority"`
	Patch                        int    `json:"patch"`
	Data                         struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Server      string `json:"server"`
		Title       string `json:"title"`
		Avatar      string `json:"avatar"`
		Portrait    string `json:"portrait"`
		Biography   string `json:"biography"`
		Race        string `json:"race"`
		Clan        string `json:"clan"`
		Gender      string `json:"gender"`
		Nameday     string `json:"nameday"`
		ActiveClass struct {
			Role struct {
				ID             int    `json:"id"`
				Name           string `json:"name"`
				Abbr           string `json:"abbr"`
				IsJob          int    `json:"is_job"`
				ClassjobParent int    `json:"classjob_parent"`
				Icon           string `json:"icon"`
				Patch          int    `json:"patch"`
			} `json:"role"`
			Progress struct {
				Name  string `json:"name"`
				Level int    `json:"level"`
				Exp   struct {
					Current      int     `json:"current"`
					Max          int     `json:"max"`
					AtCap        bool    `json:"at_cap"`
					TotalGained  int     `json:"total_gained"`
					TotalMax     int     `json:"total_max"`
					TotalTogo    int     `json:"total_togo"`
					Percent      float64 `json:"percent"`
					TotalPercent float64 `json:"total_percent"`
				} `json:"exp"`
				ID   int `json:"id"`
				Data struct {
					ID             int    `json:"id"`
					Name           string `json:"name"`
					Abbr           string `json:"abbr"`
					IsJob          int    `json:"is_job"`
					ClassjobParent int    `json:"classjob_parent"`
					Icon           string `json:"icon"`
					Patch          int    `json:"patch"`
				} `json:"data"`
				LevelTogo    int     `json:"level_togo"`
				LevelPercent float64 `json:"level_percent"`
			} `json:"progress"`
		} `json:"active_class"`
		Guardian struct {
			Icon string `json:"icon"`
			Name string `json:"name"`
		} `json:"guardian"`
		City struct {
			Icon string `json:"icon"`
			Name string `json:"name"`
		} `json:"city"`
		GrandCompany *struct {
			Icon string `json:"icon"`
			Name string `json:"name"`
			Rank string `json:"rank"`
		} `json:"grand_company"`
	} `json:"data"`
	Portrait                       string  `json:"portrait"`
	LastActive                     string  `json:"last_active"`
	URL                            string  `json:"url"`
	URLAPI                         string  `json:"url_api"`
	URLXivdb                       string  `json:"url_xivdb"`
	URLLodestone                   string  `json:"url_lodestone"`
	URLType                        string  `json:"url_type"`
	AchievementsScoreRebornPercent float64 `json:"achievements_score_reborn_percent"`
	AchievementsScoreLegacyPercent int     `json:"achievements_score_legacy_percent"`
	Extras                         struct {
		Mounts struct {
			Obtained int     `json:"obtained"`
			Total    int     `json:"total"`
			Percent  float64 `json:"percent"`
		} `json:"mounts"`
		Minions struct {
			Obtained int     `json:"obtained"`
			Total    int     `json:"total"`
			Percent  float64 `json:"percent"`
		} `json:"minions"`
	} `json:"extras"`
}
