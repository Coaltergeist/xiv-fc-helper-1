// Free Companies are FFXIV guilds
// By convention, a free company only uses 1 discord server at a time
// Thus, linking a 1 to 1 relationship between a FC and a discord server
package main

// A FreeCompany struct represents configuration data on a per-FC system
type FreeCompany struct {
	World string `json:"world"`

	Characters map[string]int `json:"characters"` // Map discord user ID -> ffxiv character id
}

// A Character is a relationship between a user's character and their discord account
// Since FFXIV isn't conducive to alternate characters, it's generally safe to assume
// you can link the relationship of FC -> Discord Account/XIV Character
type Character struct {
}
