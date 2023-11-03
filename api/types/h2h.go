package types

type Duel struct {
	HomeWins int `json:"homeWins"`
	AwayWins int `json:"awayWins"`
	Draws    int `json:"draws"`
}

type H2H struct {
	TeamDuel    Duel `json:"teamDuel"`
	ManagerDuel Duel `json:"managerDuel"`
}
