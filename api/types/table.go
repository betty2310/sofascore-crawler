package types

type TeamInfo struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"shortName"`
	UserCount int    `json:"userCount"`
	NameCode  string `json:"nameCode"`
	Type      int    `json:"type"`
	ID        int    `json:"id"`
}

type PromotionInfo struct {
	Text string `json:"text"`
	ID   int    `json:"id"`
}

type RowData struct {
	Team          TeamInfo      `json:"team"`
	Descriptions  []string      `json:"descriptions"`
	Promotion     PromotionInfo `json:"promotion"`
	Position      int           `json:"position"`
	Matches       int           `json:"matches"`
	Wins          int           `json:"wins"`
	ScoresFor     int           `json:"scoresFor"`
	ScoresAgainst int           `json:"scoresAgainst"`
	ID            int           `json:"id"`
	Losses        int           `json:"losses"`
	Draws         int           `json:"draws"`
	Points        int           `json:"points"`
}

type TournamentInfo struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TableData struct {
	Standings []struct {
		Tournament         TournamentInfo `json:"tournament"`
		Type               string         `json:"type"`
		Name               string         `json:"name"`
		Rows               []RowData      `json:"rows"`
		ID                 int            `json:"id"`
		UpdatedAtTimestamp int64          `json:"updatedAtTimestamp"`
	} `json:"standings"`
}
