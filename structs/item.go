package structs

// XIVDBItemSearch represents an item search on XIVDB
type XIVDBItemSearch struct {
	Items struct {
		Results []struct {
			AttributesParamsSpecial []interface{} `json:"attributes_params_special"`
			CategoryName            string        `json:"category_name"`
			Color                   string        `json:"color"`
			Icon                    string        `json:"icon"`
			ID                      int           `json:"id"`
			KindName                string        `json:"kind_name"`
			LevelEquip              int           `json:"level_equip"`
			LevelItem               int           `json:"level_item"`
			Name                    string        `json:"name"`
			Rarity                  int           `json:"rarity"`
			URL                     string        `json:"url"`
			URLAPI                  string        `json:"url_api"`
			URLType                 string        `json:"url_type"`
			URLXivdb                string        `json:"url_xivdb"`
		} `json:"results"`
		Total  int `json:"total"`
		Paging struct {
			Page  int   `json:"page"`
			Total int   `json:"total"`
			Pages []int `json:"pages"`
			Next  int   `json:"next"`
			Prev  int   `json:"prev"`
		} `json:"paging"`
	} `json:"items"`
}

// XIVDBItem represents a queried item on XIVDB
type XIVDBItem struct {
	Achievements    interface{} `json:"achievements"`
	Action          int         `json:"action"`
	AetherialReduce int         `json:"aetherial_reduce"`
	// AttributesBase  struct {
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
	// 	Patch           int `json:"patch"`
	// } `json:"attributes_base"`
	AttributesParams           []interface{} `json:"attributes_params"`
	AttributesParamsSpecial    []interface{} `json:"attributes_params_special"`
	BaseParamModifier          int           `json:"base_param_modifier"`
	BonusName                  interface{}   `json:"bonus_name"`
	CanBeHq                    int           `json:"can_be_hq"`
	CategoryName               string        `json:"category_name"`
	ClassjobCategory           interface{}   `json:"classjob_category"`
	ClassjobRepair             interface{}   `json:"classjob_repair"`
	ClassjobUse                int           `json:"classjob_use"`
	Classjobs                  []interface{} `json:"classjobs"`
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
	CooldownSeconds            int           `json:"cooldown_seconds"`
	Craftable                  interface{}   `json:"craftable"`
	Desynthesize               int           `json:"desynthesize"`
	Enemies                    interface{}   `json:"enemies"`
	EquipSlotCategory          int           `json:"equip_slot_category"`
	EquippableByGenderF        int           `json:"equippable_by_gender_f"`
	EquippableByGenderM        int           `json:"equippable_by_gender_m"`
	EquippableByRaceAura       int           `json:"equippable_by_race_aura"`
	EquippableByRaceElezen     int           `json:"equippable_by_race_elezen"`
	EquippableByRaceHyur       int           `json:"equippable_by_race_hyur"`
	EquippableByRaceLalafell   int           `json:"equippable_by_race_lalafell"`
	EquippableByRaceMiqote     int           `json:"equippable_by_race_miqote"`
	EquippableByRaceRoegadyn   int           `json:"equippable_by_race_roegadyn"`
	Gathering                  interface{}   `json:"gathering"`
	GcTurnIn                   int           `json:"gc_turn_in"`
	GrandCompany               int           `json:"grand_company"`
	Help                       string        `json:"help"`
	HelpCns                    string        `json:"help_cns"`
	HelpDe                     string        `json:"help_de"`
	HelpEn                     string        `json:"help_en"`
	HelpFr                     string        `json:"help_fr"`
	HelpJa                     string        `json:"help_ja"`
	Icon                       string        `json:"icon"`
	IconHq                     string        `json:"icon_hq"`
	IconLodestone              string        `json:"icon_lodestone"`
	IconLodestoneHq            string        `json:"icon_lodestone_hq"`
	ID                         int           `json:"id"`
	Instances                  interface{}   `json:"instances"`
	IsCollectable              int           `json:"is_collectable"`
	IsConvertible              int           `json:"is_convertible"`
	IsCrestWorthy              int           `json:"is_crest_worthy"`
	IsDated                    int           `json:"is_dated"`
	IsDesynthesizable          int           `json:"is_desynthesizable"`
	IsDyeable                  int           `json:"is_dyeable"`
	IsIndisposable             int           `json:"is_indisposable"`
	IsLegacy                   int           `json:"is_legacy"`
	IsProjectable              int           `json:"is_projectable"`
	IsPvp                      int           `json:"is_pvp"`
	IsReducible                int           `json:"is_reducible"`
	IsUnique                   int           `json:"is_unique"`
	IsUntradable               int           `json:"is_untradable"`
	ItemAction                 int           `json:"item_action"`
	ItemDuration               int           `json:"item_duration"`
	ItemGlamour                interface{}   `json:"item_glamour"`
	ItemRepair                 interface{}   `json:"item_repair"`
	ItemSearchCategory         int           `json:"item_search_category"`
	ItemSeries                 int           `json:"item_series"`
	ItemSpecialBonus           int           `json:"item_special_bonus"`
	ItemUICategory             int           `json:"item_ui_category"`
	ItemUIKind                 int           `json:"item_ui_kind"`
	KindName                   string        `json:"kind_name"`
	LevelEquip                 int           `json:"level_equip"`
	LevelItem                  int           `json:"level_item"`
	Leves                      interface{}   `json:"leves"`
	LodestoneID                string        `json:"lodestone_id"`
	LodestoneType              string        `json:"lodestone_type"`
	MateriaSlotCount           int           `json:"materia_slot_count"`
	MaterializeType            int           `json:"materialize_type"`
	ModelMain                  string        `json:"model_main"`
	ModelSub                   string        `json:"model_sub"`
	Name                       string        `json:"name"`
	NameCns                    string        `json:"name_cns"`
	NameDe                     string        `json:"name_de"`
	NameEn                     string        `json:"name_en"`
	NameFr                     string        `json:"name_fr"`
	NameJa                     string        `json:"name_ja"`
	Noun                       interface{}   `json:"noun"`
	NounCns                    string        `json:"noun_cns"`
	NounDe                     interface{}   `json:"noun_de"`
	NounEn                     interface{}   `json:"noun_en"`
	NounFr                     interface{}   `json:"noun_fr"`
	NounJa                     interface{}   `json:"noun_ja"`
	ParsedLodestone            int           `json:"parsed_lodestone"`
	ParsedLodestoneTime        interface{}   `json:"parsed_lodestone_time"`
	// Patch                      struct {
	// 	Name   string `json:"name"`
	// 	Number string `json:"number"`
	// 	Patch  int    `json:"patch"`
	// 	URL    string `json:"url"`
	// } `json:"patch"`
	Plural      interface{} `json:"plural"`
	PluralCns   string      `json:"plural_cns"`
	PluralDe    interface{} `json:"plural_de"`
	PluralEn    interface{} `json:"plural_en"`
	PluralFr    interface{} `json:"plural_fr"`
	PluralJa    interface{} `json:"plural_ja"`
	PriceHigh   int         `json:"price_high"`
	PriceLow    int         `json:"price_low"`
	PriceMid    int         `json:"price_mid"`
	PriceSell   int         `json:"price_sell"`
	PriceSellHq int         `json:"price_sell_hq"`
	PvpRank     int         `json:"pvp_rank"`
	Quests      interface{} `json:"quests"`
	Rarity      int         `json:"rarity"`
	Recipes     []struct {
		ClassName string `json:"class_name"`
		Classjob  struct {
			Abbr           string `json:"abbr"`
			ClassjobParent int    `json:"classjob_parent"`
			Icon           string `json:"icon"`
			ID             int    `json:"id"`
			IsJob          int    `json:"is_job"`
			Name           string `json:"name"`
			//Patch          int    `json:"patch"`
		} `json:"classjob"`
		Color         string `json:"color"`
		CraftQuantity int    `json:"craft_quantity"`
		ElementName   string `json:"element_name"`
		Icon          string `json:"icon"`
		IconHq        string `json:"icon_hq"`
		ID            int    `json:"id"`
		Item          struct {
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
			// 	Patch           int `json:"patch"`
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
			PriceMid     int         `json:"price_mid"`
			Rarity       int         `json:"rarity"`
			SeriesName   interface{} `json:"series_name"`
			SlotEquip    int         `json:"slot_equip"`
			SlotName     interface{} `json:"slot_name"`
			StackSize    int         `json:"stack_size"`
			Updated      string      `json:"updated"`
			URL          string      `json:"url"`
			URLAPI       string      `json:"url_api"`
			URLLodestone string      `json:"url_lodestone"`
			URLType      string      `json:"url_type"`
			URLXivdb     string      `json:"url_xivdb"`
			URLXivdbDe   string      `json:"url_xivdb_de"`
			URLXivdbFr   string      `json:"url_xivdb_fr"`
			URLXivdbJa   string      `json:"url_xivdb_ja"`
		} `json:"item"`
		ItemIcon          int    `json:"item_icon"`
		ItemIconLodestone string `json:"item_icon_lodestone"`
		ItemName          string `json:"item_name"`
		ItemRarity        int    `json:"item_rarity"`
		ItemURL           string `json:"item_url"`
		Level             int    `json:"level"`
		LevelDiff         int    `json:"level_diff"`
		LevelView         int    `json:"level_view"`
		LodestoneID       string `json:"lodestone_id"`
		LodestoneType     string `json:"lodestone_type"`
		Name              string `json:"name"`
		// Patch             struct {
		// 	Name   string `json:"name"`
		// 	Number string `json:"number"`
		// 	Patch  int    `json:"patch"`
		// 	URL    string `json:"url"`
		// } `json:"patch"`
		Stars     int    `json:"stars"`
		StarsHTML string `json:"stars_html"`
		Tree      []struct {
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
			// 	Patch           int `json:"patch"`
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
			PriceMid     int         `json:"price_mid"`
			Quantity     int         `json:"quantity"`
			Rarity       int         `json:"rarity"`
			SeriesName   interface{} `json:"series_name"`
			SlotEquip    int         `json:"slot_equip"`
			SlotName     interface{} `json:"slot_name"`
			StackSize    int         `json:"stack_size"`
			Updated      string      `json:"updated"`
			URL          string      `json:"url"`
			URLAPI       string      `json:"url_api"`
			URLLodestone string      `json:"url_lodestone"`
			URLType      string      `json:"url_type"`
			URLXivdb     string      `json:"url_xivdb"`
			URLXivdbDe   string      `json:"url_xivdb_de"`
			URLXivdbFr   string      `json:"url_xivdb_fr"`
			URLXivdbJa   string      `json:"url_xivdb_ja"`
		} `json:"tree"`
		URL          string `json:"url"`
		URLAPI       string `json:"url_api"`
		URLLodestone string `json:"url_lodestone"`
		URLType      string `json:"url_type"`
		URLXivdb     string `json:"url_xivdb"`
		URLXivdbDe   string `json:"url_xivdb_de"`
		URLXivdbFr   string `json:"url_xivdb_fr"`
		URLXivdbJa   string `json:"url_xivdb_ja"`
	} `json:"recipes"`
	ReducibleClassjob    interface{} `json:"reducible_classjob"`
	ReducibleLevel       int         `json:"reducible_level"`
	Salvage              int         `json:"salvage"`
	SeriesName           interface{} `json:"series_name"`
	Shops                interface{} `json:"shops"`
	SlotEquip            int         `json:"slot_equip"`
	SlotName             interface{} `json:"slot_name"`
	SortKey              int         `json:"sort_key"`
	SpecialShopsCurrency interface{} `json:"special_shops_currency"`
	SpecialShopsObtain   []struct {
		Icon  string `json:"icon"`
		ID    int    `json:"id"`
		Items struct {
			CostCollectabilityRating1 int `json:"cost_collectability_rating_1"`
			CostCollectabilityRating2 int `json:"cost_collectability_rating_2"`
			CostCollectabilityRating3 int `json:"cost_collectability_rating_3"`
			CostCount1                int `json:"cost_count_1"`
			CostCount2                int `json:"cost_count_2"`
			CostCount3                int `json:"cost_count_3"`
			CostH12                   int `json:"cost_h1_2"`
			CostHq1                   int `json:"cost_hq_1"`
			CostHq3                   int `json:"cost_hq_3"`
			CostItem1                 struct {
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
				// 	Patch           int `json:"patch"`
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
				PriceMid     int         `json:"price_mid"`
				Rarity       int         `json:"rarity"`
				SeriesName   interface{} `json:"series_name"`
				SlotEquip    int         `json:"slot_equip"`
				SlotName     interface{} `json:"slot_name"`
				StackSize    int         `json:"stack_size"`
				Updated      string      `json:"updated"`
				URL          string      `json:"url"`
				URLAPI       string      `json:"url_api"`
				URLLodestone string      `json:"url_lodestone"`
				URLType      string      `json:"url_type"`
				URLXivdb     string      `json:"url_xivdb"`
				URLXivdbDe   string      `json:"url_xivdb_de"`
				URLXivdbFr   string      `json:"url_xivdb_fr"`
				URLXivdbJa   string      `json:"url_xivdb_ja"`
			} `json:"cost_item_1"`
			CostItem2     interface{} `json:"cost_item_2"`
			CostItem3     interface{} `json:"cost_item_3"`
			QuestItem1    interface{} `json:"quest_item_1"`
			ReceiveCount1 int         `json:"receive_count_1"`
			ReceiveCount2 int         `json:"receive_count_2"`
			ReceiveHq1    int         `json:"receive_hq_1"`
			ReceiveHq2    int         `json:"receive_hq_2"`
			ReceiveItem1  struct {
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
				// 	Patch           int `json:"patch"`
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
				PriceMid     int         `json:"price_mid"`
				Rarity       int         `json:"rarity"`
				SeriesName   interface{} `json:"series_name"`
				SlotEquip    int         `json:"slot_equip"`
				SlotName     interface{} `json:"slot_name"`
				StackSize    int         `json:"stack_size"`
				Updated      string      `json:"updated"`
				URL          string      `json:"url"`
				URLAPI       string      `json:"url_api"`
				URLLodestone string      `json:"url_lodestone"`
				URLType      string      `json:"url_type"`
				URLXivdb     string      `json:"url_xivdb"`
				URLXivdbDe   string      `json:"url_xivdb_de"`
				URLXivdbFr   string      `json:"url_xivdb_fr"`
				URLXivdbJa   string      `json:"url_xivdb_ja"`
			} `json:"receive_item_1"`
			ReceiveItem2                    interface{} `json:"receive_item_2"`
			ReceiveSpecialShopItemCategory1 int         `json:"receive_special_shop_item_category_1"`
			ReceiveSpecialShopItemCategory2 int         `json:"receive_special_shop_item_category_2"`
			SpecialShop                     int         `json:"special_shop"`
		} `json:"items"`
		Name         string `json:"name"`
		SpecialUsage string `json:"special_usage"`
		URL          string `json:"url"`
		URLAPI       string `json:"url_api"`
		URLType      string `json:"url_type"`
		URLXivdb     string `json:"url_xivdb"`
		URLXivdbDe   string `json:"url_xivdb_de"`
		URLXivdbFr   string `json:"url_xivdb_fr"`
		URLXivdbJa   string `json:"url_xivdb_ja"`
	} `json:"special_shops_obtain"`
	StackSize       int    `json:"stack_size"`
	Stain           int    `json:"stain"`
	StartsWithVowel int    `json:"starts_with_vowel"`
	Updated         string `json:"updated"`
	URL             string `json:"url"`
	URLAPI          string `json:"url_api"`
	URLLodestone    string `json:"url_lodestone"`
	URLType         string `json:"url_type"`
	URLXivdb        string `json:"url_xivdb"`
	URLXivdbDe      string `json:"url_xivdb_de"`
	URLXivdbFr      string `json:"url_xivdb_fr"`
	URLXivdbJa      string `json:"url_xivdb_ja"`
	Cid             int    `json:"_cid"`
	Type            string `json:"_type"`
}
