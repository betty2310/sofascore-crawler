package types

type Player struct {
	Name      string `json:"name"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Slug      string `json:"slug"`
	ShortName string `json:"shortName"`
	Position  string `json:"position"`
	UserCount int    `json:"userCount"`
	Id        int    `json:"id"`
	Country   struct {
		Alpha2 string `json:"alpha2"`
		Name   string `json:"name"`
	} `json:"country"`
	MarketValueCurrency  string `json:"marketValueCurrency"`
	DateOfBirthTimestamp int    `json:"dateOfBirthTimestamp"`
}

type Statistics struct {
	TotalPass          int     `json:"totalPass,omitempty"`
	AccuratePass       int     `json:"accuratePass,omitempty"`
	TotalLongBalls     int     `json:"totalLongBalls,omitempty"`
	AccurateLongBalls  int     `json:"accurateLongBalls,omitempty"`
	Saves              int     `json:"saves,omitempty"`
	MinutesPlayed      int     `json:"minutesPlayed,omitempty"`
	Touches            int     `json:"touches,omitempty"`
	Rating             float64 `json:"rating,omitempty"`
	PossessionLostCtrl int     `json:"possessionLostCtrl,omitempty"`
	RatingVersions     struct {
		Original    float64 `json:"original"`
		Alternative float64 `json:"alternative"`
	} `json:"ratingVersions,omitempty"`
	GoalsPrevented         float64 `json:"goalsPrevented,omitempty"`
	AerialLost             int     `json:"aerialLost,omitempty"`
	DuelLost               int     `json:"duelLost,omitempty"`
	TotalClearance         int     `json:"totalClearance,omitempty"`
	InterceptionWon        int     `json:"interceptionWon,omitempty"`
	ExpectedAssists        float64 `json:"expectedAssists,omitempty"`
	AerialWon              int     `json:"aerialWon,omitempty"`
	DuelWon                int     `json:"duelWon,omitempty"`
	ShotOffTarget          int     `json:"shotOffTarget,omitempty"`
	HitWoodwork            int     `json:"hitWoodwork,omitempty"`
	TotalTackle            int     `json:"totalTackle,omitempty"`
	Fouls                  int     `json:"fouls,omitempty"`
	ExpectedGoals          float64 `json:"expectedGoals,omitempty"`
	TotalContest           int     `json:"totalContest,omitempty"`
	WonContest             int     `json:"wonContest,omitempty"`
	BigChanceMissed        int     `json:"bigChanceMissed,omitempty"`
	OnTargetScoringAttempt int     `json:"onTargetScoringAttempt,omitempty"`
	ChallengeLost          int     `json:"challengeLost,omitempty"`
	Dispossessed           int     `json:"dispossessed,omitempty"`
	TotalCross             int     `json:"totalCross,omitempty"`
	KeyPass                int     `json:"keyPass,omitempty"`
	GoalAssist             int     `json:"goalAssist,omitempty"`
	AccurateCross          int     `json:"accurateCross,omitempty"`
	BigChanceCreated       int     `json:"bigChanceCreated,omitempty"`
	WasFouled              int     `json:"wasFouled,omitempty"`
	Goals                  int     `json:"goals,omitempty"`
}

type TeamLineup struct {
	Players []struct {
		Player       Player     `json:"player"`
		AvgRating    float64    `json:"avgRating"`
		ShirtNumber  int        `json:"shirtNumber"`
		JerseyNumber string     `json:"jerseyNumber"`
		Position     string     `json:"position"`
		Substitute   bool       `json:"substitute"`
		Statistics   Statistics `json:"statistics"`
		Captain      bool       `json:"captain,omitempty"`
	} `json:"players"`
	Formation      string `json:"formation"`
	MissingPlayers []struct {
		Player Player `json:"player"`
		Type   string `json:"type"`
		Reason int    `json:"reason"`
	} `json:"missingPlayers"`
}

type Lineup struct {
	Confirmed bool       `json:"confirmed"`
	Home      TeamLineup `json:"home"`
	Away      TeamLineup `json:"away"`
}
