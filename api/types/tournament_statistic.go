package types

type AvgRatingStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		AvgRating      float64 `json:"avgRating"`
		ID             int     `json:"id"`
		Matches        int     `json:"matches"`
		AwardedMatches int     `json:"awardedMatches"`
	} `json:"statistics"`
}

type GoalsScoredStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		GoalsScored    int `json:"goalsScored"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type GoalsConcededStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		GoalsConceded  int `json:"goalsConceded"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type BigChancesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		BigChances     int `json:"bigChances"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type BigChancesMissedStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		BigChancesMissed int `json:"bigChancesMissed"`
		ID               int `json:"id"`
		Matches          int `json:"matches"`
		AwardedMatches   int `json:"awardedMatches"`
	} `json:"statistics"`
}
type HitWoodworkStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		HitWoodwork    int `json:"hitWoodwork"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type YellowCardsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		YellowCards    int `json:"yellowCards"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type RedCardsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		RedCards       int `json:"redCards"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type AverageBallPossessionStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		AverageBallPossession float64 `json:"averageBallPossession"`
		ID                    int     `json:"id"`
		Matches               int     `json:"matches"`
		AwardedMatches        int     `json:"awardedMatches"`
	} `json:"statistics"`
}

type AccuratePassesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		AccuratePasses int `json:"accuratePasses"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type AccurateLongBallsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		AccurateLongBalls int `json:"accurateLongBalls"`
		ID                int `json:"id"`
		Matches           int `json:"matches"`
		AwardedMatches    int `json:"awardedMatches"`
	} `json:"statistics"`
}

type AccurateCrossesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		AccurateCrosses int `json:"accurateCrosses"`
		ID              int `json:"id"`
		Matches         int `json:"matches"`
		AwardedMatches  int `json:"awardedMatches"`
	} `json:"statistics"`
}

type ShotsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Shots          int `json:"shots"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}
type ShotsOnTargetStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		ShotsOnTarget  int `json:"shotsOnTarget"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type SuccessfulDribblesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		SuccessfulDribbles int `json:"successfulDribbles"`
		ID                 int `json:"id"`
		Matches            int `json:"matches"`
		AwardedMatches     int `json:"awardedMatches"`
	} `json:"statistics"`
}

type TacklesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Tackles        int `json:"tackles"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	} `json:"statistics"`
}

type InterceptionsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Interceptions  int `json:"interceptions"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type ClearancesStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Clearances     int `json:"clearances"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type CornersStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Corners        int `json:"corners"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type FoulsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		Fouls          int `json:"fouls"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type PenaltyGoalsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		PenaltyGoals   int `json:"penaltyGoals"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type PenaltyGoalsConcededStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		PenaltyGoalsConceded int `json:"penaltyGoalsConceded"`
		ID                   int `json:"id"`
		Matches              int `json:"matches"`
		AwardedMatches       int `json:"awardedMatches"`
	}
}

type CleanSheetsStatistics struct {
	Team       TeamInfo `json:"team"`
	Statistics struct {
		CleanSheets    int `json:"cleanSheets"`
		ID             int `json:"id"`
		Matches        int `json:"matches"`
		AwardedMatches int `json:"awardedMatches"`
	}
}

type TournamentStatistics struct {
	AvgRating                       []AvgRatingStatistics             `json:"avgRating"`
	GoalsScored                     []GoalsScoredStatistics           `json:"goalsScored"`
	GoalsConceded                   []GoalsConcededStatistics         `json:"goalsConceded"`
	BigChances                      []BigChancesStatistics            `json:"bigChances"`
	BigChancesMissed                []BigChancesMissedStatistics      `json:"bigChancesMissed"`
	HitWoodwork                     []HitWoodworkStatistics           `json:"hitWoodwork"`
	YellowCardsStatistics           []YellowCardsStatistics           `json:"yellowCards"`
	RedCardsStatistics              []RedCardsStatistics              `json:"redCards"`
	AverageBallPossessionStatistics []AverageBallPossessionStatistics `json:"averageBallPossession"`
	AccuratePassesStatistics        []AccuratePassesStatistics        `json:"accuratePasses"`
	AccurateLongBallsStatistics     []AccurateLongBallsStatistics     `json:"accurateLongBalls"`
	AccurateCrossesStatistics       []AccurateCrossesStatistics       `json:"accurateCrosses"`
	ShotsStatistics                 []ShotsStatistics                 `json:"shots"`
	ShotsOnTargetStatistics         []ShotsOnTargetStatistics         `json:"shotsOnTarget"`
	SuccessfulDribblesStatistics    []SuccessfulDribblesStatistics    `json:"successfulDribbles"`
	TacklesStatistics               []TacklesStatistics               `json:"tackles"`
	InterceptionsStatistics         []InterceptionsStatistics         `json:"interceptions"`
	ClearancesStatistics            []ClearancesStatistics            `json:"clearances"`
	CornersStatistics               []CornersStatistics               `json:"corners"`
	FoulsStatistics                 []FoulsStatistics                 `json:"fouls"`
	PenaltyGoalsStatistics          []PenaltyGoalsStatistics          `json:"penaltyGoals"`
	PenaltyGoalsConcededStatistics  []PenaltyGoalsConcededStatistics  `json:"penaltyGoalsConceded"`
	CleanSheetsStatistics           []CleanSheetsStatistics           `json:"cleanSheets"`
}

type TournamentStatisticsData struct {
	TopTeams TournamentStatistics `json:"topTeams"`
}
