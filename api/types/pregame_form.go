package types

type TeamForm struct {
	AvgRating string   `json:"avgRating"`
	Position  int      `json:"position"`
	Value     string   `json:"value"`
	Form      []string `json:"form"`
}

type PregameFormData struct {
	HomeTeam TeamForm `json:"homeTeam"`
	AwayTeam TeamForm `json:"awayTeam"`
	Label    string   `json:"label"`
}
