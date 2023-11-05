package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"betty/fb88/api/types"

	"github.com/fatih/color"
)

func GetMatches(date string) (matches types.Events, error error) {
	url := fmt.Sprintf("https://api.sofascore.com/api/v1/sport/football/scheduled-events/%s", date)
	response, err := http.Get(url)
	if err != nil {
		color.Red("Error fetching: %v\n\nMaybe you get banned :v", err)
	}
	err = json.NewDecoder(response.Body).Decode(&matches)
	if err != nil {
		color.Red("Error fetching: %v\n\nMaybe you get banned :v", err)
	}
	return matches, err
}
