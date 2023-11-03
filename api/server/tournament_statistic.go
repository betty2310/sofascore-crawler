package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTournamentStatistic() (statistic types.TournamentStatisticsData, error error) {
	url := "https://api.sofascore.com/api/v1/unique-tournament/17/season/52186/top-teams/overall"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&statistic)
	if err != nil {
		fmt.Println(err)
	}
	return statistic, err
}
