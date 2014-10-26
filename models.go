package hearthscience

type Card struct {
	Name         string   `json:"name"`
	Cost         int      `json:"cost"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Faction      string   `json:"faction"`
	Race         string   `json:"race"`
	PlayerClass  string   `json:"playerClass"`
	Text         string   `json:"text"`
	InPlayText   string   `json:"inPlayText"`
	Mechanics    []string `json:"mechanics"`
	Flavor       string   `json:"flavor"`
	Artist       string   `json:"artist"`
	Attack       int      `json:"attack"`
	Health       int      `json:"health"`
	Durability   int      `json:"durability"`
	Id           string   `json:"id"`
	Collectable  bool     `json:"collectable"`
	Elite        bool     `json:"elite"`
	HowToGet     string   `json:"howToGet"`
	HowToGetGold string   `json:"howToGetGold"`
}

type SetList struct {
	List []string
}

type HearthstoneJSONVersion struct {
	Version string `json:"version"`
}
