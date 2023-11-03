package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
get team statistic: hard code tournament and season for now: EPL 23/24

@param id team_id, can found in struct `Team`
*/
func GetTeamStatistic(id int) (lineup types.TeamStatistics, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/team/%d/unique-tournament/17/season/52186/statistics/overall", id)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&lineup)
	if err != nil {
		fmt.Println(err)
	}
	return lineup, err
}
