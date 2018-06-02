package structs

// XIVDBRecipeSearch represents a search on xivdb
type XIVDBRecipeSearch struct {
	Recipes struct {
		Results []struct {
			ClassName     string      `json:"class_name"`
			Color         string      `json:"color"`
			Icon          string      `json:"icon"`
			IconLodestone string      `json:"icon_lodestone"`
			ID            int         `json:"id"`
			ItemName      string      `json:"item_name"`
			Level         int         `json:"level"`
			LevelDiff     int         `json:"level_diff"`
			LevelView     int         `json:"level_view"`
			Masterbook    interface{} `json:"masterbook"`
			Name          string      `json:"name"`
			Stars         int         `json:"stars"`
			StarsHTML     string      `json:"stars_html"`
			URL           string      `json:"url"`
			URLAPI        string      `json:"url_api"`
			URLType       string      `json:"url_type"`
			URLXivdb      string      `json:"url_xivdb"`
		} `json:"results"`
		Total  int `json:"total"`
		Paging struct {
			Page  int   `json:"page"`
			Total int   `json:"total"`
			Pages []int `json:"pages"`
			Next  int   `json:"next"`
			Prev  int   `json:"prev"`
		} `json:"paging"`
	} `json:"recipes"`
}

// XIVDBRecipe represents a full recipe for a craftable item
type XIVDBRecipe struct {
	CanHq         int    `json:"can_hq"`
	CanQuickSynth int    `json:"can_quick_synth"`
	ClassName     string `json:"class_name"`
	Classjob      struct {
		Abbr           string `json:"abbr"`
		ClassjobParent int    `json:"classjob_parent"`
		Icon           string `json:"icon"`
		ID             int    `json:"id"`
		IsJob          int    `json:"is_job"`
		Name           string `json:"name"`
		//Patch          int    `json:"patch"`
	} `json:"classjob"`
	Color                    string          `json:"color"`
	CraftQuantity            int             `json:"craft_quantity"`
	CraftType                int             `json:"craft_type"`
	Difficulty               int             `json:"difficulty"`
	DifficultyFactor         float64         `json:"difficulty_factor"`
	Durability               int             `json:"durability"`
	DurabilityFactor         float64         `json:"durability_factor"`
	ElementName              string          `json:"element_name"`
	Icon                     string          `json:"icon"`
	IconHq                   string          `json:"icon_hq"`
	ID                       int             `json:"id"`
	IsSecondary              int             `json:"is_secondary"`
	IsSpecializationRequired int             `json:"is_specialization_required"`
	Item                     XIVDBRecipeTree `json:"item"`
	ItemIcon                 int             `json:"item_icon"`
	ItemIconLodestone        string          `json:"item_icon_lodestone"`
	ItemName                 string          `json:"item_name"`
	ItemRarity               int             `json:"item_rarity"`
	ItemRequired             int             `json:"item_required"`
	ItemURL                  string          `json:"item_url"`
	Level                    int             `json:"level"`
	LevelDiff                int             `json:"level_diff"`
	LevelView                int             `json:"level_view"`
	LodestoneID              string          `json:"lodestone_id"`
	LodestoneType            string          `json:"lodestone_type"`
	Masterbook               interface{}     `json:"masterbook"`
	MaterialPoint            int             `json:"material_point"`
	Name                     string          `json:"name"`
	NameCns                  string          `json:"name_cns"`
	NameDe                   string          `json:"name_de"`
	NameEn                   string          `json:"name_en"`
	NameFr                   string          `json:"name_fr"`
	NameJa                   string          `json:"name_ja"`
	Number                   int             `json:"number"`
	// Patch                    struct {
	// 	Name   string `json:"name"`
	// 	Number string `json:"number"`
	// 	Patch  int    `json:"patch"`
	// 	URL    string `json:"url"`
	// } `json:"patch"`
	Quality                 int               `json:"quality"`
	QualityFactor           float64           `json:"quality_factor"`
	QuickSynthControl       int               `json:"quick_synth_control"`
	QuickSynthCraftsmanship int               `json:"quick_synth_craftsmanship"`
	RecipeElement           int               `json:"recipe_element"`
	RecipeNotebook          int               `json:"recipe_notebook"`
	RequiredControl         int               `json:"required_control"`
	RequiredCraftsmanship   int               `json:"required_craftsmanship"`
	Stars                   int               `json:"stars"`
	StarsHTML               string            `json:"stars_html"`
	StatusRequired          int               `json:"status_required"`
	Tree                    []XIVDBRecipeTree `json:"tree"`
	UnlockKey               int               `json:"unlock_key"`
	URL                     string            `json:"url"`
	URLAPI                  string            `json:"url_api"`
	URLLodestone            string            `json:"url_lodestone"`
	URLType                 string            `json:"url_type"`
	URLXivdb                string            `json:"url_xivdb"`
	URLXivdbDe              string            `json:"url_xivdb_de"`
	URLXivdbFr              string            `json:"url_xivdb_fr"`
	URLXivdbJa              string            `json:"url_xivdb_ja"`
	WorkMax                 int               `json:"work_max"`
	Cid                     int               `json:"_cid"`
	Type                    string            `json:"_type"`
}

// XIVDBRecipeTree represents the tree of recipes to create a specific recipe
type XIVDBRecipeTree struct {
	// AttributesBase struct {
	// 	AutoAttack      int `json:"auto_attack"`
	// 	AutoAttackHq    int `json:"auto_attack_hq"`
	// 	BlockRate       int `json:"block_rate"`
	// 	BlockRateHq     int `json:"block_rate_hq"`
	// 	BlockStrength   int `json:"block_strength"`
	// 	BlockStrengthHq int `json:"block_strength_hq"`
	// 	Damage          int `json:"damage"`
	// 	DamageHq        int `json:"damage_hq"`
	// 	Defense         int `json:"defense"`
	// 	DefenseHq       int `json:"defense_hq"`
	// 	Delay           int `json:"delay"`
	// 	DelayHq         int `json:"delay_hq"`
	// 	Dps             int `json:"dps"`
	// 	DpsHq           int `json:"dps_hq"`
	// 	ID              int `json:"id"`
	// 	MagicDamage     int `json:"magic_damage"`
	// 	MagicDamageHq   int `json:"magic_damage_hq"`
	// 	MagicDefense    int `json:"magic_defense"`
	// 	MagicDefenseHq  int `json:"magic_defense_hq"`
	// 	//Patch           int `json:"patch"`
	// } `json:"attributes_base"`
	AttributesParams           []interface{} `json:"attributes_params"`
	AttributesParamsSpecial    []interface{} `json:"attributes_params_special"`
	BonusName                  interface{}   `json:"bonus_name"`
	CategoryName               string        `json:"category_name"`
	ClassjobCategory           interface{}   `json:"classjob_category"`
	Color                      string        `json:"color"`
	ConnectAchievement         int           `json:"connect_achievement"`
	ConnectCraftable           int           `json:"connect_craftable"`
	ConnectEnemyDrop           int           `json:"connect_enemy_drop"`
	ConnectGathering           int           `json:"connect_gathering"`
	ConnectInstance            int           `json:"connect_instance"`
	ConnectInstanceChest       int           `json:"connect_instance_chest"`
	ConnectInstanceReward      int           `json:"connect_instance_reward"`
	ConnectLeve                int           `json:"connect_leve"`
	ConnectQuestReward         int           `json:"connect_quest_reward"`
	ConnectRecipe              int           `json:"connect_recipe"`
	ConnectShop                int           `json:"connect_shop"`
	ConnectSpecialshopCost1    int           `json:"connect_specialshop_cost_1"`
	ConnectSpecialshopCost2    int           `json:"connect_specialshop_cost_2"`
	ConnectSpecialshopCost3    int           `json:"connect_specialshop_cost_3"`
	ConnectSpecialshopReceive1 int           `json:"connect_specialshop_receive_1"`
	ConnectSpecialshopReceive2 int           `json:"connect_specialshop_receive_2"`
	Connected                  bool          `json:"connected"`
	Help                       string        `json:"help"`
	Icon                       string        `json:"icon"`
	IconHq                     string        `json:"icon_hq"`
	IconLodestone              string        `json:"icon_lodestone"`
	ID                         int           `json:"id"`
	ItemUICategory             int           `json:"item_ui_category"`
	KindName                   string        `json:"kind_name"`
	LevelEquip                 int           `json:"level_equip"`
	LevelItem                  int           `json:"level_item"`
	LodestoneID                string        `json:"lodestone_id"`
	LodestoneType              string        `json:"lodestone_type"`
	Name                       string        `json:"name"`
	// Patch                      struct {
	// 	Name   string `json:"name"`
	// 	Number string `json:"number"`
	// 	Patch  int    `json:"patch"`
	// 	URL    string `json:"url"`
	// } `json:"patch"`
	PriceMid     int                        `json:"price_mid"`
	Quantity     int                        `json:"quantity"`
	Rarity       int                        `json:"rarity"`
	SeriesName   interface{}                `json:"series_name"`
	SlotEquip    int                        `json:"slot_equip"`
	SlotName     interface{}                `json:"slot_name"`
	StackSize    int                        `json:"stack_size"`
	Updated      string                     `json:"updated"`
	URL          string                     `json:"url"`
	URLAPI       string                     `json:"url_api"`
	URLLodestone string                     `json:"url_lodestone"`
	URLType      string                     `json:"url_type"`
	URLXivdb     string                     `json:"url_xivdb"`
	URLXivdbDe   string                     `json:"url_xivdb_de"`
	URLXivdbFr   string                     `json:"url_xivdb_fr"`
	URLXivdbJa   string                     `json:"url_xivdb_ja"`
	Synths       map[string]XIVDBRecipeTree `json:"synths,omitempty"`
}
