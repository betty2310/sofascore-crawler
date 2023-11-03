package types

type Tournament struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type RoundInfo struct {
	Round int `json:"round"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Team struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"shortName"`
	Type      int    `json:"type"`
	ID        int    `json:"id"`
	Country   struct {
		Alpha2 string `json:"alpha2"`
		Name   string `json:"name"`
	} `json:"country"`
}

type Score struct {
	Current    int `json:"current"`
	Display    int `json:"display"`
	Period1    int `json:"period1"`
	Period2    int `json:"period2"`
	Normaltime int `json:"normaltime"`
}

type Event struct {
	Tournament               Tournament `json:"tournament"`
	RoundInfo                RoundInfo  `json:"roundInfo"`
	CustomId                 string     `json:"customId"`
	Status                   Status     `json:"status"`
	WinnerCode               int        `json:"winnerCode"`
	HomeTeam                 Team       `json:"homeTeam"`
	AwayTeam                 Team       `json:"awayTeam"`
	HomeScore                Score      `json:"homeScore"`
	AwayScore                Score      `json:"awayScore"`
	HasGlobalHighlights      bool       `json:"hasGlobalHighlights"`
	HasEventPlayerStatistics bool       `json:"hasEventPlayerStatistics"`
	HasEventPlayerHeatMap    bool       `json:"hasEventPlayerHeatMap"`
	ID                       int        `json:"id"`
	StartTimestamp           int        `json:"startTimestamp"`
	EndTimestamp             int        `json:"endTimestamp"`
	Slug                     string     `json:"slug"`
}

type Events struct {
	Events []Event `json:"events"`
}
