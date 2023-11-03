package types

type PlayerOverall struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"shortName"`
	Position  string `json:"position"`
	UserCount int    `json:"userCount"`
	ID        int    `json:"id"`
}

type PlayerRating struct {
	Statistics struct {
		Appearances int     `json:"appearances"`
		Rating      float64 `json:"rating"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerGoal struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		Goals       int `json:"goals"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerExpectedGoals struct {
	Statistics struct {
		Appearances   int     `json:"appearances"`
		ExpectedGoals float64 `json:"expectedGoals"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}
type PlayerAssists struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		Assists     int `json:"assists"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerExpectedAssists struct {
	Statistics struct {
		Appearances     int     `json:"appearances"`
		ExpectedAssists float64 `json:"expectedAssists"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerPenaltyGoals struct {
	Statistics struct {
		Appearances    int `json:"appearances"`
		PenaltyGoals   int `json:"penaltyGoals"`
		PenaltiesTaken int `json:"penaltiesTaken"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerFreeKickGoals struct {
	Statistics struct {
		Appearances      int `json:"appearances"`
		FreeKickGoal     int `json:"freeKickGoal"`
		ShotFromSetPiece int `json:"shotFromSetPiece"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerScoringFrequency struct {
	Statistics struct {
		Appearances      int     `json:"appearances"`
		ScoringFrequency float64 `json:"scoringFrequency"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerTotalShots struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		TotalShots  int `json:"totalShots"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerShotsOnTarget struct {
	Statistics struct {
		Appearances   int `json:"appearances"`
		ShotsOnTarget int `json:"shotsOnTarget"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerBigChancesMissed struct {
	Statistics struct {
		Appearances      int `json:"appearances"`
		BigChancesMissed int `json:"bigChancesMissed"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerBigChancesCreated struct {
	Statistics struct {
		Appearances       int `json:"appearances"`
		BigChancesCreated int `json:"bigChancesCreated"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerAccuratePasses struct {
	Statistics struct {
		Appearances              int     `json:"appearances"`
		AccuratePasses           int     `json:"accuratePasses"`
		AccuratePassesPercentage float64 `json:"accuratePassesPercentage"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerKeyPasses struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		KeyPasses   int `json:"keyPasses"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}
type PlayerAccurateLongBalls struct {
	Statistics struct {
		Apperances        int `json:"apperances"`
		AccurateLongBalls int `json:"accurateLongBalls"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}
type PlayerSuccessfulDribbles struct {
	Statistics struct {
		Appearances                  int     `json:"appearances"`
		SuccessfulDribbles           int     `json:"successfulDribbles"`
		SuccessfulDribblesPercentage float64 `json:"successfulDribblesPercentage"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerTackles struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		Tackles     int `json:"tackles"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerInterceptions struct {
	Statistics struct {
		Interceptions int `json:"interceptions"`
		Appearances   int `json:"appearances"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerClearances struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		Clearances  int `json:"clearances"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerPossessionLost struct {
	Statistics struct {
		Appearances    int `json:"appearances"`
		PossessionLost int `json:"possessionLost"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type PlayerYellowCards struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		YellowCards int `json:"yellowCards"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}
type PlayerRedCards struct {
	Statistics struct {
		Appearances int `json:"appearances"`
		RedCards    int `json:"redCards"`
	} `json:"statistics"`
	PlayedEnough bool          `json:"playedEnough"`
	Player       PlayerOverall `json:"player"`
}

type TeamPlayerData struct {
	TopPlayers struct {
		Rating             []PlayerRating             `json:"rating"`
		Goals              []PlayerGoal               `json:"goals"`
		ExpectedGoals      []PlayerExpectedGoals      `json:"expectedGoals"`
		Assists            []PlayerAssists            `json:"assists"`
		ExpectedAssists    []PlayerExpectedAssists    `json:"expectedAssists"`
		FreeKickGoal       []PlayerFreeKickGoals      `json:"freeKickGoal"`
		PenaltyGoals       []PlayerPenaltyGoals       `json:"penaltyGoals"`
		ScoringFrequency   []PlayerScoringFrequency   `json:"scoringFrequency"`
		TotalShots         []PlayerTotalShots         `json:"totalShots"`
		ShotsOnTarget      []PlayerShotsOnTarget      `json:"shotsOnTarget"`
		BigChancesMissed   []PlayerBigChancesMissed   `json:"bigChancesMissed"`
		BigChancesCreated  []PlayerBigChancesCreated  `json:"bigChancesCreated"`
		AccuratePasses     []PlayerAccuratePasses     `json:"accuratePasses"`
		KeyPasses          []PlayerKeyPasses          `json:"keyPasses"`
		AccurateLongBalls  []PlayerAccurateLongBalls  `json:"accurateLongBalls"`
		SuccessfulDribbles []PlayerSuccessfulDribbles `json:"successfulDribbles"`
		Tackles            []PlayerTackles            `json:"tackles"`
		Interceptions      []PlayerInterceptions      `json:"interceptions"`
		Clearances         []PlayerClearances         `json:"clearances"`
		PossessionLost     []PlayerPossessionLost     `json:"possessionLost"`
		YellowCards        []PlayerYellowCards        `json:"yellowCards"`
		RedCards           []PlayerRedCards           `json:"redCards"`
	} `json:"topPlayers"`
}
