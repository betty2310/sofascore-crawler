package server

import (
	"betty/fb88/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLineups(id int) (lineup types.Lineup, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/event/%d/lineups", id)
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
