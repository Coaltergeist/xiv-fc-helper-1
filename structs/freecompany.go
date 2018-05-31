package structs

// LodestoneFreeCompany is a free company profile from the lodestone
type LodestoneFreeCompany struct {
	ID     int    `json:"id"`
	Server string `json:"server"`
	Name   string `json:"name"`

	Tag     string `json:"tag"`
	Slogan  string `json:"slogan"`
	Faction string `json:"faction"`

	Reputation string `json:"reputation"`
}
